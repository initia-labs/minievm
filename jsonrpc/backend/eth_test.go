package backend_test

import (
	"bytes"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/holiman/uint256"
	rpctypes "github.com/initia-labs/minievm/jsonrpc/types"
	"github.com/initia-labs/minievm/tests"
	"github.com/initia-labs/minievm/x/evm/contracts/erc20"
	evmtypes "github.com/initia-labs/minievm/x/evm/types"
	"github.com/stretchr/testify/require"
)

func Test_GetBalance(t *testing.T) {
	input := setupBackend(t)
	app, _, backend, _, privKeys := input.app, input.addrs, input.backend, input.addrs, input.privKeys

	// send 1GAS to newly created account
	addr := common.Address{1, 2, 3}
	tx, _ := tests.GenerateTx(t, app, privKeys[0], &addr, nil, big.NewInt(1_000_000_000_000_000_000))
	_, finalizeRes := tests.ExecuteTxs(t, app, tx)
	tests.CheckTxResult(t, finalizeRes.TxResults[0], true)

	// wait feeFetcher interval
	time.Sleep(3 * time.Second)

	// 0 at genesis block
	balance, err := backend.GetBalance(addr, rpc.BlockNumberOrHashWithNumber(0))
	require.NoError(t, err)
	require.Equal(t, big.NewInt(0), balance.ToInt())

	// 1GAS at latest block
	balance, err = backend.GetBalance(addr, rpc.BlockNumberOrHashWithNumber(-1))
	require.NoError(t, err)
	require.Equal(t, big.NewInt(1_000_000_000_000_000_000), balance.ToInt())
}

func Test_Call(t *testing.T) {
	input := setupBackend(t)
	app, _, backend, addrs, privKeys := input.app, input.addrs, input.backend, input.addrs, input.privKeys

	tx, _ := tests.GenerateCreateERC20Tx(t, app, privKeys[0])
	_, finalizeRes := tests.ExecuteTxs(t, app, tx)
	tests.CheckTxResult(t, finalizeRes.TxResults[0], true)

	events := finalizeRes.TxResults[0].Events
	createEvent := events[len(events)-3]
	require.Equal(t, evmtypes.EventTypeContractCreated, createEvent.GetType())

	contractAddr, err := hexutil.Decode(createEvent.Attributes[0].Value)
	require.NoError(t, err)

	// mint 1_000_000 tokens to the first address
	contractEVMAddr := common.BytesToAddress(contractAddr)
	tx, _ = tests.GenerateMintERC20Tx(t, app, privKeys[0], contractEVMAddr, addrs[0], new(big.Int).SetUint64(1_000_000_000_000))
	tx2, _ := tests.GenerateTransferERC20Tx(t, app, privKeys[0], contractEVMAddr, addrs[1], new(big.Int).SetUint64(1_000_000), tests.SetNonce(2))
	_, finalizeRes = tests.ExecuteTxs(t, app, tx, tx2)
	tests.CheckTxResult(t, finalizeRes.TxResults[0], true)
	tests.CheckTxResult(t, finalizeRes.TxResults[1], true)

	// call balanceOf function
	abi, err := erc20.Erc20MetaData.GetAbi()
	require.NoError(t, err)

	inputBz, err := abi.Pack("balanceOf", addrs[1])
	require.NoError(t, err)

	// query call with genesis block
	genesisNumber := rpc.BlockNumberOrHashWithNumber(0)
	retBz, err := backend.Call(rpctypes.TransactionArgs{
		From:  &addrs[0],
		To:    &contractEVMAddr,
		Input: (*hexutil.Bytes)(&inputBz),
		Value: nil,
		Nonce: nil,
		AccessList: &types.AccessList{
			types.AccessTuple{
				Address: contractEVMAddr,
				StorageKeys: []common.Hash{
					common.HexToHash("0x00"),
				},
			},
		},
	}, &genesisNumber, nil, nil)
	require.NoError(t, err)
	require.Empty(t, retBz)

	// query to latest block
	retBz, err = backend.Call(rpctypes.TransactionArgs{
		From:  &addrs[0],
		To:    &contractEVMAddr,
		Input: (*hexutil.Bytes)(&inputBz),
		Value: nil,
		Nonce: nil,
		AccessList: &types.AccessList{
			types.AccessTuple{
				Address: contractEVMAddr,
				StorageKeys: []common.Hash{
					common.HexToHash("0x00"),
				},
			},
		},
	}, nil, nil, nil)
	require.NoError(t, err)

	res, err := abi.Unpack("balanceOf", retBz)
	require.NoError(t, err)

	require.Equal(t, new(big.Int).SetUint64(1_000_000), res[0].(*big.Int))
}

func Test_StorageAt(t *testing.T) {
	input := setupBackend(t)
	app, _, backend, addrs, privKeys := input.app, input.addrs, input.backend, input.addrs, input.privKeys

	tx, _ := tests.GenerateCreateERC20Tx(t, app, privKeys[0])
	_, finalizeRes := tests.ExecuteTxs(t, app, tx)
	tests.CheckTxResult(t, finalizeRes.TxResults[0], true)

	events := finalizeRes.TxResults[0].Events
	createEvent := events[len(events)-3]
	require.Equal(t, evmtypes.EventTypeContractCreated, createEvent.GetType())

	contractAddr, err := hexutil.Decode(createEvent.Attributes[0].Value)
	require.NoError(t, err)

	// slot 0: 0x000000000000000000000000a7cdbd8580affc8dd17c2c28cb7ae891c711fd18
	// slot 1: 0x0000000000000000000000000000000000000000000000000000000000000000
	// slot 2: 0x0000000000000000000000000000000000000000000000000000000000000000
	// slot 3: 0x666f6f0000000000000000000000000000000000000000000000000000000006 (name)
	// slot 4: 0x666f6f0000000000000000000000000000000000000000000000000000000006 (symbol)
	// slot 5: 0x0000000000000000000000000000000000000000000000000000000000000006 (decimals)
	// slot 6: 0x0000000000000000000000000000000000000000000000000000000000000000 (total supply)
	slot := common.HexToHash("03")
	retBz, err := backend.GetStorageAt(common.BytesToAddress(contractAddr), slot, rpc.BlockNumberOrHashWithNumber(-1))
	require.NoError(t, err)
	require.Equal(t, common.Hex2Bytes("666f6f0000000000000000000000000000000000000000000000000000000006"), []byte(retBz))

	slot = common.HexToHash("04")
	retBz, err = backend.GetStorageAt(common.BytesToAddress(contractAddr), slot, rpc.BlockNumberOrHashWithNumber(-1))
	require.NoError(t, err)
	require.Equal(t, common.Hex2Bytes("666f6f0000000000000000000000000000000000000000000000000000000006"), []byte(retBz))

	slot = common.HexToHash("05")
	retBz, err = backend.GetStorageAt(common.BytesToAddress(contractAddr), slot, rpc.BlockNumberOrHashWithNumber(-1))
	require.NoError(t, err)
	require.Equal(t, common.Hex2Bytes("0000000000000000000000000000000000000000000000000000000000000006"), []byte(retBz))

	slot = common.HexToHash("06")
	retBz, err = backend.GetStorageAt(common.BytesToAddress(contractAddr), slot, rpc.BlockNumberOrHashWithNumber(-1))
	require.NoError(t, err)
	require.Equal(t, common.Hex2Bytes("0000000000000000000000000000000000000000000000000000000000000000"), []byte(retBz))

	// mint 1_000_000 tokens to the first address
	contractEVMAddr := common.BytesToAddress(contractAddr)
	tx, _ = tests.GenerateMintERC20Tx(t, app, privKeys[0], contractEVMAddr, addrs[0], new(big.Int).SetUint64(1_000_000_000_000))
	_, finalizeRes = tests.ExecuteTxs(t, app, tx)
	tests.CheckTxResult(t, finalizeRes.TxResults[0], true)

	slot = common.HexToHash("06")
	retBz, err = backend.GetStorageAt(common.BytesToAddress(contractAddr), slot, rpc.BlockNumberOrHashWithNumber(-1))
	require.NoError(t, err)
	expectedTotalSupply := uint256.NewInt(1_000_000_000_000).Bytes32()
	require.Equal(t, expectedTotalSupply[:], []byte(retBz))
}

func Test_Code(t *testing.T) {
	input := setupBackend(t)
	app, _, backend, _, privKeys := input.app, input.addrs, input.backend, input.addrs, input.privKeys

	tx, _ := tests.GenerateCreateERC20Tx(t, app, privKeys[0])
	_, finalizeRes := tests.ExecuteTxs(t, app, tx)
	tests.CheckTxResult(t, finalizeRes.TxResults[0], true)

	events := finalizeRes.TxResults[0].Events
	createEvent := events[len(events)-3]
	require.Equal(t, evmtypes.EventTypeContractCreated, createEvent.GetType())

	contractAddr, err := hexutil.Decode(createEvent.Attributes[0].Value)
	require.NoError(t, err)

	code, err := backend.GetCode(common.BytesToAddress(contractAddr), rpc.BlockNumberOrHashWithNumber(-1))
	require.NoError(t, err)
	require.NotEmpty(t, code)

	erc20Code := hexutil.MustDecode(erc20.Erc20MetaData.Bin)
	initCodeOP := common.Hex2Bytes("5ff3fe")
	initCodePos := bytes.Index(erc20Code, initCodeOP)
	erc20RuntimeCode := erc20Code[initCodePos+3:]
	require.True(t, bytes.Equal(erc20RuntimeCode, code))
}

func Test_ChainID(t *testing.T) {
	input := setupBackend(t)
	app, _, backend, _, _ := input.app, input.addrs, input.backend, input.addrs, input.privKeys

	chainID, err := backend.ChainId()
	require.NoError(t, err)

	evmChainID := evmtypes.ConvertCosmosChainIDToEthereumChainID(app.ChainID())
	require.Equal(t, evmChainID, chainID.ToInt())
}
