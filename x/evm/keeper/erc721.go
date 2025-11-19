package keeper

import (
	"context"
	"errors"
	"math/big"

	"cosmossdk.io/collections"
	"cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	erc721 "github.com/initia-labs/minievm/x/evm/contracts/ics721_erc721"
	"github.com/initia-labs/minievm/x/evm/types"

	nfttransfertypes "github.com/initia-labs/initia/x/ibc/nft-transfer/types"
)

var _ nfttransfertypes.NftKeeper = ERC721Keeper{}

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

// GetERC721ABI implements IERC721Keeper.
func (k ERC721Keeper) GetERC721ABI() *abi.ABI {
	return k.ABI
}

func (k ERC721Keeper) isCollectionInitialized(ctx context.Context, classId string) (bool, error) {
	return k.ERC721ContractAddrsByClassId.Has(ctx, classId)
}

func (k ERC721Keeper) CreateOrUpdateClass(ctx context.Context, classId, classUri, classData string) error {
	if ok, err := k.isCollectionInitialized(ctx, classId); err != nil {
		return err
	} else if !ok {
		inputBz, err := k.Pack("", classId, classId)
		if err != nil {
			return types.ErrFailedToPackABI.Wrap(err.Error())
		}

		ret, contractAddr, _, err := k.EVMCreate(ctx, types.StdAddress, append(k.ERC721Bin, inputBz...), nil, nil)
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
	}

	// update class uri
	contractAddr, err := types.ContractAddressFromClassId(ctx, k, classId)
	if err != nil {
		return err
	}

	if err := k.ERC721ClassURIs.Set(ctx, contractAddr.Bytes(), classUri); err != nil {
		return err
	}

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

	contractAddr, err := types.ContractAddressFromClassId(ctx, k, classId)
	if err != nil {
		return err
	}

	for _, tokenId := range tokenIds {
		intTokenId, ok := types.TokenIdToBigInt(classId, tokenId)
		if !ok {
			return types.ErrInvalidTokenId
		}
		inputBz, err := k.Pack("safeTransferFrom", senderAddr, receiverAddr, intTokenId)
		if err != nil {
			return types.ErrFailedToPackABI.Wrap(err.Error())
		}

		// ignore the return values
		_, _, err = k.EVMCall(ctx, senderAddr, contractAddr, inputBz, nil, nil, nil)
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
	inputBz, err := k.Pack("burn", tokenId)
	if err != nil {
		return types.ErrFailedToPackABI.Wrap(err.Error())
	}

	_, _, err = k.EVMCall(ctx, owner, contractAddr, inputBz, nil, nil, nil)
	if err != nil {
		return err
	}

	// emit erc721 burned event
	sdkCtx := sdk.UnwrapSDKContext(ctx)
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
	contractAddr, err := types.ContractAddressFromClassId(ctx, k, classId)
	if err != nil {
		return err
	}

	ownerAddr, err := k.convertToEVMAddress(ctx, owner, false)
	if err != nil {
		return err
	}

	for _, tokenId := range tokenIds {
		intTokenId, ok := types.TokenIdToBigInt(classId, tokenId)
		if !ok {
			return types.ErrInvalidTokenId
		}

		err := k.Burn(ctx, ownerAddr, intTokenId, contractAddr)
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
	inputBz, err := k.Pack("mint", receiver, tokenId, tokenUri, tokenOriginId)
	if err != nil {
		return types.ErrFailedToPackABI.Wrap(err.Error())
	}

	_, _, err = k.EVMCall(ctx, types.StdAddress, contractAddr, inputBz, nil, nil, nil)
	if err != nil {
		return err
	}

	// emit erc721 minted event
	sdkCtx := sdk.UnwrapSDKContext(ctx)
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
	contractAddr, err := types.ContractAddressFromClassId(ctx, k, classId)
	if err != nil {
		return err
	}

	receiverAddr, err := k.convertToEVMAddress(ctx, receiver, false)
	if err != nil {
		return err
	}

	for i, tokenId := range tokenIds {
		intTokenId, ok := types.TokenIdToBigInt(classId, tokenId)
		if !ok {
			return types.ErrInvalidTokenId
		}

		err := k.Mint(ctx, receiverAddr, intTokenId, tokenId, tokenUris[i], contractAddr)
		if err != nil {
			return err
		}
	}

	return nil
}

func (k ERC721Keeper) GetClassInfo(ctx context.Context, classId string) (className string, classUri string, classDescs string, err error) {
	contractAddr, err := types.ContractAddressFromClassId(ctx, k, classId)
	if err != nil {
		return "", "", "", err
	}

	// ERC721s relayed from the other chains via IBC have classUri else return empty str.
	classUri, err = k.ERC721ClassURIs.Get(ctx, contractAddr.Bytes())
	if err != nil && !errors.Is(err, collections.ErrNotFound) {
		return "", "", "", err
	} else if err != nil {
		classUri = ""
	}

	className, err = k.name(ctx, contractAddr)
	if err != nil {
		return "", "", "", err
	}

	return className, classUri, "", err
}

func (k ERC721Keeper) GetTokenInfos(ctx context.Context, classId string, tokenIds []string) (tokenUris []string, tokenDescs []string, err error) {
	contractAddr, err := types.ContractAddressFromClassId(ctx, k, classId)
	if err != nil {
		return nil, nil, err
	}

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

func (k ERC721Keeper) GetOriginTokenInfos(ctx context.Context, classId string, tokenIds []*big.Int) (tokenOriginIds, tokenUris []string, err error) {
	contractAddr, err := types.ContractAddressFromClassId(ctx, k, classId)
	if err != nil {
		return nil, nil, err
	}

	tokenOriginIds = make([]string, len(tokenIds))
	tokenUris = make([]string, len(tokenIds))
	for i, tokenId := range tokenIds {
		tokenUri, err := k.tokenURI(ctx, tokenId, contractAddr)
		if err != nil {
			return nil, nil, err
		}
		tokenUris[i] = tokenUri

		tokenOriginId, err := k.tokenOriginId(ctx, tokenId, contractAddr)
		if err != nil {
			return nil, nil, err
		}
		tokenOriginIds[i] = tokenOriginId
	}
	return tokenOriginIds, tokenUris, err
}

func (k ERC721Keeper) balanceOf(ctx context.Context, addr, contractAddr common.Address) (math.Int, error) {
	inputBz, err := k.Pack("balanceOf", addr)
	if err != nil {
		return math.ZeroInt(), types.ErrFailedToPackABI.Wrap(err.Error())
	}

	retBz, err := k.EVMStaticCall(ctx, types.NullAddress, contractAddr, inputBz, nil)
	if err != nil {
		return math.ZeroInt(), err
	}

	res, err := k.Unpack("balanceOf", retBz)
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
	contractAddr, err := types.ContractAddressFromClassId(ctx, k, classId)
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
	inputBz, err := k.Pack("ownerOf", tokenId)
	if err != nil {
		return types.NullAddress, types.ErrFailedToPackABI.Wrap(err.Error())
	}

	retBz, err := k.EVMStaticCall(ctx, types.NullAddress, contractAddr, inputBz, nil)
	if err != nil {
		return types.NullAddress, err
	}

	res, err := k.Unpack("ownerOf", retBz)
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
	contractAddr, err := types.ContractAddressFromClassId(ctx, k, classId)
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
	inputBz, err := k.Pack("name")
	if err != nil {
		return "", types.ErrFailedToPackABI.Wrap(err.Error())
	}

	retBz, err := k.EVMStaticCall(ctx, types.NullAddress, contractAddr, inputBz, nil)
	if err != nil {
		return "", err
	}

	res, err := k.Unpack("name", retBz)
	if err != nil {
		return "", types.ErrFailedToUnpackABI.Wrap(err.Error())
	}

	name, ok := res[0].(string)
	if !ok {
		return name, types.ErrFailedToDecodeOutput
	}

	return name, nil
}

func (k ERC721Keeper) tokenURI(ctx context.Context, tokenId *big.Int, contractAddr common.Address) (string, error) {
	inputBz, err := k.Pack("tokenURI", tokenId)
	if err != nil {
		return "", types.ErrFailedToPackABI.Wrap(err.Error())
	}

	retBz, err := k.EVMStaticCall(ctx, types.NullAddress, contractAddr, inputBz, nil)
	if err != nil {
		return "", err
	}

	res, err := k.Unpack("tokenURI", retBz)
	if err != nil {
		return "", types.ErrFailedToUnpackABI.Wrap(err.Error())
	}

	tokenUri, ok := res[0].(string)
	if !ok {
		return tokenUri, types.ErrFailedToDecodeOutput
	}

	return tokenUri, nil
}

func (k ERC721Keeper) tokenOriginId(ctx context.Context, tokenId *big.Int, contractAddr common.Address) (string, error) {
	inputBz, err := k.Pack("tokenOriginId", tokenId)
	if err != nil {
		return "", types.ErrFailedToPackABI.Wrap(err.Error())
	}

	retBz, err := k.EVMStaticCall(ctx, types.NullAddress, contractAddr, inputBz, nil)
	if err != nil {
		return "", err
	}

	res, err := k.Unpack("tokenOriginId", retBz)
	if err != nil {
		return "", types.ErrFailedToUnpackABI.Wrap(err.Error())
	}

	tokenOriginId, ok := res[0].(string)
	if !ok {
		return tokenOriginId, types.ErrFailedToDecodeOutput
	}

	return tokenOriginId, nil
}
