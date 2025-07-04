package cli

import (
	"fmt"
	"slices"
	"strings"
	"time"

	"cosmossdk.io/core/address"
	"github.com/spf13/cobra"

	evmtypes "github.com/initia-labs/minievm/x/evm/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/authz"
	authzcli "github.com/cosmos/cosmos-sdk/x/authz/client/cli"
	bank "github.com/cosmos/cosmos-sdk/x/bank/types"
)

// Flag names and valuesi
const (
	FlagContracts = "contracts"
	evm           = "evm"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(ac, vc address.Codec) *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        authz.ModuleName,
		Short:                      "Authorization transactions subcommands",
		Long:                       "Authorize and revoke access to execute transactions on behalf of your address",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(
		NewCmdGrantAuthorization(ac, vc),
		authzcli.NewCmdRevokeAuthorization(ac),
		authzcli.NewCmdExecAuthorization(),
	)

	return txCmd
}

func NewCmdGrantAuthorization(ac, vc address.Codec) *cobra.Command {
	originCmd := authzcli.NewCmdGrantAuthorization(ac)
	cmd := &cobra.Command{
		Use:   "grant <grantee> <authorization_type=\"send\"|\"generic\"|\"delegate\"|\"unbond\"|\"redelegate\"|\"evm\"> --from <granter>",
		Short: "Grant authorization to an address",
		Long: strings.TrimSpace(
			fmt.Sprintf(`grant authorization to an address to execute a transaction on your behalf:

Examples:
 $ %s tx %s grant init1vrit.. send %s --spend-limit=1000uinit --from=init1vrit..
 $ %s tx %s grant init1vrit.. generic --msg-type=/cosmos.gov.v1beta1.MsgVote --from=init1vrit..
 $ %s tx %s grant init1vrit.. evm --contracts "0x1234567890123456789012345678901234567890,0xabcdefabcdefabcdefabcdefabcdefabcdefabcd" --from=init1vrit..
`, version.AppName, authz.ModuleName, bank.SendAuthorization{}.MsgTypeURL(), version.AppName, authz.ModuleName, version.AppName, authz.ModuleName),
		),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			// If not a EVM authorization, delegate to standard authz CLI handler
			if args[1] != evm {
				return originCmd.RunE(cmd, args)
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			grantee, err := ac.StringToBytes(args[0])
			if err != nil {
				return err
			}

			contractsStr, err := cmd.Flags().GetString(FlagContracts)
			if err != nil {
				return err
			}

			contracts := []string{}
			if contractsStr != "" {
				contracts = strings.Split(contractsStr, ",")
				for i, contract := range contracts {
					contracts[i] = strings.TrimSpace(contract)
				}

				// Remove empty strings
				contracts = slices.DeleteFunc(contracts, func(c string) bool {
					return c == ""
				})
			}

			authorization := evmtypes.NewCallAuthorization(contracts)
			if err := authorization.ValidateBasic(); err != nil {
				return err
			}

			expire, err := getExpireTime(cmd)
			if err != nil {
				return err
			}

			msg, err := authz.NewMsgGrant(clientCtx.GetFromAddress(), grantee, authorization, expire)
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().AddFlagSet(originCmd.Flags())
	cmd.Flags().String(FlagContracts, "", "The contracts of evm authorization, a comma-separated string of contract addresses.")
	return cmd
}

func getExpireTime(cmd *cobra.Command) (*time.Time, error) {
	exp, err := cmd.Flags().GetInt64(authzcli.FlagExpiration)
	if err != nil {
		return nil, err
	}
	if exp == 0 {
		return nil, nil
	}
	e := time.Unix(exp, 0)
	return &e, nil
}
