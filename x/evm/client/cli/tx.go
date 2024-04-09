package cli

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"cosmossdk.io/core/address"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"

	"github.com/initia-labs/minievm/x/evm/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(ac address.Codec) *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "EVM transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	txCmd.AddCommand(
		CreateCmd(ac),
		Create2Cmd(ac),
		CallCmd(ac),
	)
	return txCmd
}

func CreateCmd(ac address.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [bin file1] [bin file2] [...]",
		Short: "Deploy evm contracts with CREATE opcode",
		Long: strings.TrimSpace(
			fmt.Sprintf(`
Deploy evm contracts. allowed to upload up to 100 files at once.

Example:
$ %s tx evm create \
    ERC20.bin \
	CustomDex.bin --from mykey
`, version.AppName,
			),
		),
		Args:    cobra.RangeArgs(1, 100),
		Aliases: []string{"CREATE"},
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			contracts := make([][]byte, len(args))
			for i, arg := range args {
				hexStrBz, err := os.ReadFile(arg)
				if err != nil {
					return err
				}

				contracts[i], err = hexutil.Decode("0x" + string(hexStrBz))
				if err != nil {
					return err
				}
			}

			sender, err := ac.BytesToString(clientCtx.FromAddress)
			if err != nil {
				return err
			}

			msgs := make([]sdk.Msg, len(contracts))
			for i, contract := range contracts {
				msgs[i] = &types.MsgCreate{
					Sender: sender,
					Code:   hexutil.Encode(contract),
				}
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgs...)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func Create2Cmd(ac address.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create2 [slat]:[bin file1] [salt]:[bin file2] [...]",
		Short: "Deploy evm contracts with CREATE2 opcode",
		Long: strings.TrimSpace(
			fmt.Sprintf(`
Deploy evm contracts. allowed to upload up to 100 files at once.

Example:
$ %s tx evm create2 \
    1:ERC20.bin \
	CustomDex.bin --from mykey
`, version.AppName,
			),
		),
		Args:    cobra.RangeArgs(2, 100),
		Aliases: []string{"CREATE2"},
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			salts := make([]uint64, len(args))
			contracts := make([][]byte, len(args))
			for i, arg := range args {
				s := strings.Split(arg, ":")
				if len(s) != 2 {
					return fmt.Errorf("invalid argument: %s", arg)
				}

				salts[i], err = strconv.ParseUint(s[0], 10, 64)
				if err != nil {
					return err
				}

				hexStrBz, err := os.ReadFile(s[1])
				if err != nil {
					return err
				}

				contracts[i], err = hexutil.Decode("0x" + string(hexStrBz))
				if err != nil {
					return err
				}
			}

			sender, err := ac.BytesToString(clientCtx.FromAddress)
			if err != nil {
				return err
			}

			msgs := make([]sdk.Msg, len(contracts))
			for i, contract := range contracts {
				msgs[i] = &types.MsgCreate2{
					Sender: sender,
					Salt:   salts[i],
					Code:   hexutil.Encode(contract),
				}
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgs...)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func CallCmd(ac address.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "call [contract-addr] [input-hex-string]",
		Short: "Call a evm contract",
		Long: strings.TrimSpace(
			fmt.Sprintf(`
Call a contract with input bytes.

Example:
$ %s tx evm call 0x1 0x123456 --from mykey
`, version.AppName,
			),
		),
		Args:    cobra.ExactArgs(2),
		Aliases: []string{"CALL"},
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			_, err = types.ContractAddressFromString(ac, args[0])
			if err != nil {
				return err
			}

			_, err = hexutil.Decode(args[1])
			if err != nil {
				return err
			}

			sender, err := ac.BytesToString(clientCtx.FromAddress)
			if err != nil {
				return err
			}

			msg := types.MsgCall{
				Sender:       sender,
				ContractAddr: args[0],
				Input:        args[1],
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
