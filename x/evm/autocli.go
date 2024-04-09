package evm

import (
	"fmt"

	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	"github.com/cosmos/cosmos-sdk/version"
	evmv1 "github.com/initia-labs/minievm/api/minievm/evm/v1"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: evmv1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Code",
					Use:       "code [contract_addr]",
					Short:     "Query contract code bytes",
					Example:   fmt.Sprintf("%s query evm code 0x1", version.AppName),
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "contract_addr"},
					},
				},
				{
					RpcMethod: "State",
					Use:       "state [contract_addr] [key]",
					Short:     "Query contract state",
					Example:   fmt.Sprintf("%s query evm state 0x1 0x123...", version.AppName),
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "contract_addr"},
						{ProtoField: "key"},
					},
				},
				{
					RpcMethod: "ContractAddrByDenom",
					Use:       "contract-addr-by-denom [denom]",
					Short:     "Query corresponding contract address by denom",
					Alias:     []string{"contract-addr"},
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "denom"},
					},
				},
				{
					RpcMethod: "Denom",
					Use:       "denom [contract_addr]",
					Short:     "Query corresponding denom by contract address",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "contract_addr"},
					},
				},
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Query the evm module params",
				},
				{
					RpcMethod: "Call",
					Skip:      true,
				},
			},
			EnhanceCustomCommand: true, // We still have manual commands in evm that we want to keep
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service: evmv1.Msg_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "MsgCreate",
					Skip:      true,
				},
				{
					RpcMethod: "MsgCreate2",
					Skip:      true,
				},
				{
					RpcMethod: "MsgCall",
					Skip:      true,
				},
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
			},
			EnhanceCustomCommand: false, // use custom commands only until v0.51
		},
	}
}
