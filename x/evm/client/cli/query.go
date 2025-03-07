package cli

import (
	"context"
	"fmt"
	"math/big"

	"cosmossdk.io/core/address"
	"cosmossdk.io/math"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/initia-labs/minievm/x/evm/types"
)

func GetQueryCmd(ac address.Codec) *cobra.Command {
	queryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the evm module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	queryCmd.AddCommand(
		GetCmdCall(ac),
	)
	return queryCmd
}

func GetCmdCall(ac address.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "call [sender] [contract-addr] [input-hex-string]",
		Short: "Call deployed evm contract",
		Long:  "Call deployed evm contract",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			_, err = ac.StringToBytes(args[0])
			if err != nil {
				return err
			}

			_, err = types.ContractAddressFromString(ac, args[1])
			if err != nil {
				return err
			}

			_, err = hexutil.Decode(args[2])
			if err != nil {
				return err
			}

			value, err := cmd.Flags().GetString(FlagValue)
			if err != nil {
				return errors.Wrap(err, "failed to get value")
			}
			val, ok := new(big.Int).SetString(value, 10)
			if !ok {
				return fmt.Errorf("invalid value: %s", value)
			}

			trace, err := cmd.Flags().GetBool(FlagTrace)
			if err != nil {
				return err
			}

			var traceOption *types.TraceOptions
			if trace {
				withMemory, err := cmd.Flags().GetBool(FlagWithMemory)
				if err != nil {
					return err
				}
				withStack, err := cmd.Flags().GetBool(FlagWithStack)
				if err != nil {
					return err
				}
				withStorage, err := cmd.Flags().GetBool(FlagWithStorage)
				if err != nil {
					return err
				}
				withReturnData, err := cmd.Flags().GetBool(FlagWithReturnData)
				if err != nil {
					return err
				}
				traceOption = &types.TraceOptions{
					WithMemory:     withMemory,
					WithStack:      withStack,
					WithStorage:    withStorage,
					WithReturnData: withReturnData,
				}
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.Call(
				context.Background(),
				&types.QueryCallRequest{
					Sender:       args[0],
					ContractAddr: args[1],
					Input:        args[2],
					Value:        math.NewIntFromBigInt(val),
					TraceOptions: traceOption,
				},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
	cmd.Flags().AddFlagSet(FlagTraceOptions())
	cmd.Flags().String(FlagValue, "0", "value")
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
