package main

import (
	"fmt"
	"path/filepath"
	"strconv"

	cmtcmd "github.com/cometbft/cometbft/cmd/cometbft/commands"
	"github.com/spf13/cobra"

	dbm "github.com/cosmos/cosmos-db"

	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/server/types"
)

// NewMultipleRollbackCmd creates a command to rollback CometBFT and multistore state by one height.
func NewMultipleRollbackCmd(appCreator types.AppCreator) *cobra.Command {
	removeBlock := false
	cmd := &cobra.Command{
		Use:   "mrollback [height]",
		Short: "rollback Cosmos SDK and CometBFT state to the given height",
		Long: `
A state rollback is performed to recover from an incorrect application state transition,
when CometBFT has persisted an incorrect app hash and is thus unable to make
progress. Rollback overwrites a state with the state at the given height. All
blocks after the given height are removed from the blockchain.
`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := server.GetServerContextFromCmd(cmd)

			height, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				return err
			}
			if height <= 0 {
				return fmt.Errorf("height must be greater than 0")
			}

			dataDir := filepath.Join(ctx.Config.RootDir, "data")
			db, err := dbm.NewDB("application", server.GetAppDBBackend(ctx.Viper), dataDir)
			if err != nil {
				return err
			}
			app := appCreator(ctx.Logger, db, nil, ctx.Viper)
			if curHeight := app.CommitMultiStore().LatestVersion(); height >= curHeight {
				return fmt.Errorf("height must be less than the current height %d", curHeight)
			}

			// rollback CometBFT state
			height, hash, err := cmtcmd.RollbackStateTo(ctx.Config, height, removeBlock)
			if err != nil {
				return fmt.Errorf("failed to rollback CometBFT state: %w", err)
			}
			// rollback the multistore

			if err := app.CommitMultiStore().RollbackToVersion(height); err != nil {
				return fmt.Errorf("failed to rollback to version: %w", err)
			}

			fmt.Printf("Rolled back state to height %d and hash %X\n", height, hash)
			return nil
		},
	}
	cmd.Flags().BoolVar(&removeBlock, "hard", false, "remove blocks as well as state")
	return cmd
}
