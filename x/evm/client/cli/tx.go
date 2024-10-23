package cli

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"

	"cosmossdk.io/core/address"
	"cosmossdk.io/math"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
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
		Use:   "create [bin-file] --input [input-hex-string] --value [value]",
		Short: "Deploy evm contracts with CREATE opcode",
		Long: strings.TrimSpace(
			fmt.Sprintf(`
Deploy evm contracts. allowed to upload up to 100 files at once.

Example:
$ %s tx evm create ERC20.bin --input 0x1234 --value 100 --from mykey
`, version.AppName,
			),
		),
		Args:    cobra.ExactArgs(1),
		Aliases: []string{"CREATE"},
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			sender, err := ac.BytesToString(clientCtx.FromAddress)
			if err != nil {
				return err
			}

			codeBz, val, err := prepareContractCreation(cmd, args[0])
			if err != nil {
				return err
			}

			msg := &types.MsgCreate{
				Sender: sender,
				Code:   hexutil.Encode(codeBz),
				Value:  math.NewIntFromBigInt(val),
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().String(FlagInput, "0x", "input hex string")
	cmd.Flags().String(FlagValue, "0", "value")
	return cmd
}

func Create2Cmd(ac address.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create2 [salt] [bin file] --input [input-hex-string] --value [value]",
		Short: "Deploy evm contracts with CREATE2 opcode",
		Long: strings.TrimSpace(
			fmt.Sprintf(`
Deploy evm contracts. allowed to upload up to 100 files at once.

Example:
$ %s tx evm create2 100 ERC20.bin --input 0x1234 --value 100 --from mykey
`, version.AppName,
			),
		),
		Args:    cobra.ExactArgs(2),
		Aliases: []string{"CREATE2"},
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			sender, err := ac.BytesToString(clientCtx.FromAddress)
			if err != nil {
				return err
			}

			salt, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return errors.Wrap(err, "failed to parse salt")
			}

			codeBz, val, err := prepareContractCreation(cmd, args[1])
			if err != nil {
				return err
			}

			msg := &types.MsgCreate2{
				Sender: sender,
				Salt:   salt,
				Code:   hexutil.Encode(codeBz),
				Value:  math.NewIntFromBigInt(val),
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().String(FlagInput, "0x", "input hex string")
	cmd.Flags().String(FlagValue, "0", "value")
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

			value, err := cmd.Flags().GetString(FlagValue)
			if err != nil {
				return errors.Wrap(err, "failed to get value")
			}
			val, ok := new(big.Int).SetString(value, 10)
			if !ok {
				return fmt.Errorf("invalid value: %s", value)
			}

			msg := types.MsgCall{
				Sender:       sender,
				ContractAddr: args[0],
				Input:        args[1],
				Value:        math.NewIntFromBigInt(val),
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().String(FlagValue, "0", "value")
	return cmd
}

func readContractBinFile(binFile string) ([]byte, error) {
	contractBz, err := os.ReadFile(binFile)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read contract file")
	}

	contractBz, err = hex.DecodeString(string(contractBz))
	if err != nil {
		return nil, errors.Wrap(err, "failed to read contract file: expect hex string")
	}

	return contractBz, nil
}

func prepareContractCreation(cmd *cobra.Command, contractFile string) ([]byte, *big.Int, error) {
	contractBz, err := readContractBinFile(contractFile)
	if err != nil {
		return nil, nil, err
	}

	input, err := cmd.Flags().GetString(FlagInput)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to get input")
	}
	inputBz, err := hexutil.Decode(input)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to decode input")
	}

	value, err := cmd.Flags().GetString(FlagValue)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to get value")
	}
	val, ok := new(big.Int).SetString(value, 10)
	if !ok {
		return nil, nil, fmt.Errorf("invalid value: %s", value)
	}

	return append(contractBz, inputBz...), val, nil
}
