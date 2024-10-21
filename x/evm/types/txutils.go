package types

import (
	"github.com/ethereum/go-ethereum/common"
	coretypes "github.com/ethereum/go-ethereum/core/types"
)

func ConvertCosmosAccessListToEth(cosmosAccessList []AccessTuple) coretypes.AccessList {
	if len(cosmosAccessList) == 0 {
		return nil
	}
	coreAccessList := make(coretypes.AccessList, len(cosmosAccessList))
	for i, a := range cosmosAccessList {
		storageKeys := make([]common.Hash, len(a.StorageKeys))
		for j, s := range a.StorageKeys {
			storageKeys[j] = common.HexToHash(s)
		}
		coreAccessList[i] = coretypes.AccessTuple{
			Address:     common.HexToAddress(a.Address),
			StorageKeys: storageKeys,
		}
	}
	return coreAccessList
}

func ConvertEthAccessListToCosmos(ethAccessList coretypes.AccessList) []AccessTuple {
	if len(ethAccessList) == 0 {
		return nil
	}
	accessList := make([]AccessTuple, len(ethAccessList))
	for i, al := range ethAccessList {
		storageKeys := make([]string, len(al.StorageKeys))
		for j, s := range al.StorageKeys {
			storageKeys[j] = s.String()
		}
		accessList[i] = AccessTuple{
			Address:     al.Address.String(),
			StorageKeys: storageKeys,
		}
	}
	return accessList
}
