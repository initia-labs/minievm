package keeper

import (
	"context"
	"math/big"

	"cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	erc721 "github.com/initia-labs/minievm/x/evm/contracts/ics721_erc721"
	"github.com/initia-labs/minievm/x/evm/types"
)

type ERC721Keeper struct {
	*Keeper
	*abi.ABI
	ERC721Bin []byte
}

func NewERC721Keeper(k *Keeper) (types.IERC721Keeper, error) {
	abi, err := erc721.Ics721Erc721MetaData.GetAbi()
	if err != nil {
		return ERC721Keeper{}, err
	}

	erc721Bin, err := hexutil.Decode(erc721.Ics721Erc721Bin)
	if err != nil {
		return ERC721Keeper{}, err
	}

	return &ERC721Keeper{k, abi, erc721Bin}, nil
}

func (k ERC721Keeper) isCollectionInitialized(ctx context.Context, classId string) (bool, error) {
	return k.ERC721ContractAddrsByClassId.Has(ctx, classId)
}

func (k ERC721Keeper) CreateOrUpdateClass(ctx context.Context, classId, classUri, classData string) error {
	if ok, err := k.isCollectionInitialized(ctx, classId); err != nil {
		return err
	} else if !ok {
		contractAddr, err := k.nextContractAddress(ctx, types.StdAddress)
		if err != nil {
			return err
		}

		inputBz, err := k.ABI.Pack("", classId, classId, classUri)
		if err != nil {
			return types.ErrFailedToPackABI.Wrap(err.Error())
		}

		ret, _, err := k.EVMCreate(ctx, types.StdAddress, append(k.ERC721Bin, inputBz...))
		if err != nil {
			return err
		}

		if err := k.ERC721ClassIdsByContractAddr.Set(ctx, contractAddr.Bytes(), classId); err != nil {
			return err
		}

		if err := k.ERC721ContractAddrsByClassId.Set(ctx, classId, contractAddr.Bytes()); err != nil {
			return err
		}

		sdkCtx := sdk.UnwrapSDKContext(ctx)
		// emit erc721 created event
		sdkCtx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeERC721Created,
				sdk.NewAttribute(types.AttributeKeyClassId, classId),
				sdk.NewAttribute(types.AttributeKeyContract, hexutil.Encode(ret)),
			),
		)
	} // update not supported; ignore

	return nil
}

func (k ERC721Keeper) Transfers(ctx context.Context, sender, receiver sdk.AccAddress, classId string, tokenIds []string) error {
	senderAddr, err := k.convertToEVMAddress(ctx, sender, true)
	if err != nil {
		return err
	}
	receiverAddr, err := k.convertToEVMAddress(ctx, receiver, false)
	if err != nil {
		return err
	}

	for _, tokenId := range tokenIds {
		contractAddr, err := k.GetContractAddrByClassId(ctx, classId)
		if err != nil {
			return err
		}

		intTokenId, ok := types.TokenIdToBigInt(classId, tokenId)
		if !ok {
			return types.ErrInvalidTokenId
		}
		inputBz, err := k.ABI.Pack("safeTransferFrom", senderAddr, receiverAddr, intTokenId)
		if err != nil {
			return types.ErrFailedToPackABI.Wrap(err.Error())
		}

		// ignore the return values
		_, _, err = k.EVMCall(ctx, senderAddr, contractAddr, inputBz)
		if err != nil {
			return err
		}
	}
	return nil
}

func (k ERC721Keeper) Burn(
	ctx context.Context, owner common.Address,
	tokenId *big.Int, contractAddr common.Address,
) error {
	inputBz, err := k.ABI.Pack("burn", tokenId)
	if err != nil {
		return types.ErrFailedToPackABI.Wrap(err.Error())
	}

	_, _, err = k.EVMCall(ctx, owner, contractAddr, inputBz)
	if err != nil {
		return err
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	// emit erc721 minted event
	sdkCtx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeERC721Burned,
			sdk.NewAttribute(types.AttributeKeyContract, contractAddr.String()),
			sdk.NewAttribute(types.AttributeKeyTokenId, tokenId.String()),
		),
	)
	return nil
}

func (k ERC721Keeper) Burns(ctx context.Context, owner sdk.AccAddress, classId string, tokenIds []string) error {
	contractAddr, err := k.ERC721ContractAddrsByClassId.Get(ctx, classId)
	if err != nil {
		return err
	}

	ownerAddr, err := k.convertToEVMAddress(ctx, owner, false)
	for _, tokenId := range tokenIds {
		intTokenId, ok := types.TokenIdToBigInt(classId, tokenId)
		if !ok {
			return types.ErrInvalidTokenId
		}
		err := k.Burn(ctx, ownerAddr, intTokenId, common.BytesToAddress(contractAddr))
		if err != nil {
			return err
		}
	}

	return nil
}

func (k ERC721Keeper) Mint(
	ctx context.Context, receiver common.Address,
	tokenId *big.Int, tokenOriginId, tokenUri string, contractAddr common.Address,
) error {
	inputBz, err := k.ABI.Pack("mint", receiver, tokenId, tokenUri, tokenOriginId)
	if err != nil {
		return types.ErrFailedToPackABI.Wrap(err.Error())
	}

	_, _, err = k.EVMCall(ctx, types.StdAddress, contractAddr, inputBz)
	if err != nil {
		return err
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	// emit erc721 minted event
	sdkCtx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeERC721Minted,
			sdk.NewAttribute(types.AttributeKeyContract, contractAddr.String()),
			sdk.NewAttribute(types.AttributeKeyTokenId, tokenId.String()),
			sdk.NewAttribute(types.AttributeKeyTokenOriginId, tokenOriginId),
		),
	)
	return nil
}

func (k ERC721Keeper) Mints(
	ctx context.Context, receiver sdk.AccAddress,
	classId string, tokenIds, tokenUris, tokenData []string,
) error {
	contractAddr, err := k.ERC721ContractAddrsByClassId.Get(ctx, classId)
	if err != nil {
		return err
	}

	receiverAddr, err := k.convertToEVMAddress(ctx, receiver, false)

	for i, tokenId := range tokenIds {
		intTokenId, ok := types.TokenIdToBigInt(classId, tokenId)
		if !ok {
			return types.ErrInvalidTokenId
		}
		err := k.Mint(ctx, receiverAddr, intTokenId, tokenId, tokenUris[i], common.BytesToAddress(contractAddr))
		if err != nil {
			return err
		}
	}

	return nil
}

func (k ERC721Keeper) GetClassInfo(ctx context.Context, classId string) (classUri string, classDescs string, err error) {
	contractAddrBz, err := k.ERC721ContractAddrsByClassId.Get(ctx, classId)
	if err != nil {
		return "", "", err
	}
	contractAddr := common.BytesToAddress(contractAddrBz)

	classUri, err = k.classURI(ctx, contractAddr)
	if err != nil {
		return "", "", err
	}

	return classUri, "", err
}

func (k ERC721Keeper) GetTokenInfos(ctx context.Context, classId string, tokenIds []string) (tokenUris []string, tokenDescs []string, err error) {
	contractAddrBz, err := k.ERC721ContractAddrsByClassId.Get(ctx, classId)
	if err != nil {
		return nil, nil, err
	}
	contractAddr := common.BytesToAddress(contractAddrBz)

	tokenUris = make([]string, len(tokenIds))
	for i, tokenId := range tokenIds {
		intTokenId, ok := types.TokenIdToBigInt(classId, tokenId)
		if !ok {
			return nil, nil, types.ErrInvalidTokenId
		}
		tokenUri, err := k.tokenURI(ctx, intTokenId, contractAddr)
		if err != nil {
			return nil, nil, err
		}
		tokenUris[i] = tokenUri
	}
	return tokenUris, make([]string, len(tokenIds)), err
}

func (k ERC721Keeper) balanceOf(ctx context.Context, addr, contractAddr common.Address) (math.Int, error) {
	inputBz, err := k.ABI.Pack("balanceOf", addr)
	if err != nil {
		return math.ZeroInt(), types.ErrFailedToPackABI.Wrap(err.Error())
	}

	retBz, err := k.EVMStaticCall(ctx, types.NullAddress, contractAddr, inputBz)
	if err != nil {
		return math.ZeroInt(), err
	}

	res, err := k.ABI.Unpack("balanceOf", retBz)
	if err != nil {
		return math.ZeroInt(), types.ErrFailedToUnpackABI.Wrap(err.Error())
	}

	balance, ok := res[0].(*big.Int)
	if !ok {
		return math.ZeroInt(), types.ErrFailedToDecodeOutput
	}

	return math.NewIntFromBigInt(balance), nil
}

func (k ERC721Keeper) BalanceOf(ctx context.Context, addr sdk.AccAddress, classId string) (math.Int, error) {
	contractAddr, err := k.GetContractAddrByClassId(ctx, classId)
	if err != nil {
		return math.ZeroInt(), err
	}

	evmAddr, err := k.convertToEVMAddress(ctx, addr, false)
	if err != nil {
		return math.ZeroInt(), err
	}

	return k.balanceOf(ctx, evmAddr, contractAddr)
}

func (k ERC721Keeper) ownerOf(ctx context.Context, tokenId *big.Int, contractAddr common.Address) (common.Address, error) {
	inputBz, err := k.ABI.Pack("ownerOf", tokenId)
	if err != nil {
		return types.NullAddress, types.ErrFailedToPackABI.Wrap(err.Error())
	}

	retBz, err := k.EVMStaticCall(ctx, types.NullAddress, contractAddr, inputBz)
	if err != nil {
		return types.NullAddress, err
	}

	res, err := k.ABI.Unpack("ownerOf", retBz)
	if err != nil {
		return types.NullAddress, types.ErrFailedToUnpackABI.Wrap(err.Error())
	}

	owner, ok := res[0].(common.Address)
	if !ok {
		return types.NullAddress, types.ErrFailedToDecodeOutput
	}

	return owner, nil
}

func (k ERC721Keeper) OwnerOf(ctx context.Context, tokenId string, classId string) (common.Address, error) {
	contractAddr, err := k.GetContractAddrByClassId(ctx, classId)
	if err != nil {
		return types.NullAddress, err
	}

	tokenIdInt, ok := types.TokenIdToBigInt(classId, tokenId)
	if !ok {
		return types.NullAddress, types.ErrInvalidTokenId
	}

	return k.ownerOf(ctx, tokenIdInt, contractAddr)
}

func (k ERC721Keeper) name(ctx context.Context, contractAddr common.Address) (string, error) {
	inputBz, err := k.ABI.Pack("name")
	if err != nil {
		return "", types.ErrFailedToPackABI.Wrap(err.Error())
	}

	retBz, err := k.EVMStaticCall(ctx, types.NullAddress, contractAddr, inputBz)
	if err != nil {
		return "", err
	}

	res, err := k.ABI.Unpack("name", retBz)
	if err != nil {
		return "", types.ErrFailedToUnpackABI.Wrap(err.Error())
	}

	name, ok := res[0].(string)
	if !ok {
		return name, types.ErrFailedToDecodeOutput
	}

	return name, nil
}

func (k ERC721Keeper) symbol(ctx context.Context, contractAddr common.Address) (string, error) {
	inputBz, err := k.ABI.Pack("symbol")
	if err != nil {
		return "", types.ErrFailedToPackABI.Wrap(err.Error())
	}

	retBz, err := k.EVMStaticCall(ctx, types.NullAddress, contractAddr, inputBz)
	if err != nil {
		return "", err
	}

	res, err := k.ABI.Unpack("symbol", retBz)
	if err != nil {
		return "", types.ErrFailedToUnpackABI.Wrap(err.Error())
	}

	symbol, ok := res[0].(string)
	if !ok {
		return symbol, types.ErrFailedToDecodeOutput
	}

	return symbol, nil
}

func (k ERC721Keeper) classURI(ctx context.Context, contractAddr common.Address) (string, error) {
	inputBz, err := k.ABI.Pack("classURI")
	if err != nil {
		return "", types.ErrFailedToPackABI.Wrap(err.Error())
	}

	retBz, err := k.EVMStaticCall(ctx, types.NullAddress, contractAddr, inputBz)
	if err != nil {
		return "", err
	}

	res, err := k.ABI.Unpack("classURI", retBz)
	if err != nil {
		return "", types.ErrFailedToUnpackABI.Wrap(err.Error())
	}

	classUri, ok := res[0].(string)
	if !ok {
		return classUri, types.ErrFailedToDecodeOutput
	}

	return classUri, nil
}

func (k ERC721Keeper) tokenURI(ctx context.Context, tokenId *big.Int, contractAddr common.Address) (string, error) {
	inputBz, err := k.ABI.Pack("tokenURI", tokenId)
	if err != nil {
		return "", types.ErrFailedToPackABI.Wrap(err.Error())
	}

	retBz, err := k.EVMStaticCall(ctx, types.NullAddress, contractAddr, inputBz)
	if err != nil {
		return "", err
	}

	res, err := k.ABI.Unpack("tokenURI", retBz)
	if err != nil {
		return "", types.ErrFailedToUnpackABI.Wrap(err.Error())
	}

	tokenUri, ok := res[0].(string)
	if !ok {
		return tokenUri, types.ErrFailedToDecodeOutput
	}

	return tokenUri, nil
}
