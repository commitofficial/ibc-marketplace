package main

import (
	"encoding/json"
	"io"
	
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/cli"
	dbm "github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"
	tmtypes "github.com/tendermint/tendermint/types"
	
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	
	"ibc-marketplace/app"
	relayerInit "ibc-marketplace/app/genesis"
	rserver "ibc-marketplace/server"
)

// relayerd custom flags
const flagInvCheckPeriod = "inv-check-period"

var invCheckPeriod uint

func main() {
	cdc := app.MakeCodec()
	
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(sdk.Bech32PrefixAccAddr, sdk.Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(sdk.Bech32PrefixValAddr, sdk.Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(sdk.Bech32PrefixConsAddr, sdk.Bech32PrefixConsPub)
	config.Seal()
	
	ctx := server.NewDefaultContext()
	cobra.EnableCommandSorting = false
	rootCmd := &cobra.Command{
		Use:               "relayerd",
		Short:             "Relayer Daemon (server)",
		PersistentPreRunE: server.PersistentPreRunEFn(ctx),
	}
	
	rootCmd.AddCommand(relayerInit.InitCmd(ctx, cdc))
	rootCmd.AddCommand(relayerInit.CollectGenTxsCmd(ctx, cdc))
	rootCmd.AddCommand(relayerInit.GenTxCmd(ctx, cdc))
	rootCmd.AddCommand(relayerInit.AddGenesisAccountCmd(ctx, cdc))
	rootCmd.AddCommand(relayerInit.ValidateGenesisCmd(ctx, cdc))
	rootCmd.AddCommand(client.NewCompletionCmd(rootCmd, true))
	
	rserver.AddCommands(ctx, cdc, rootCmd, newApp, exportAppStateAndTMValidators)
	
	// prepare and add flags
	executor := cli.PrepareBaseCmd(rootCmd, "RA", app.DefaultNodeHome)
	rootCmd.PersistentFlags().UintVar(&invCheckPeriod, flagInvCheckPeriod,
		0, "Assert registered invariants every N blocks")
	err := executor.Execute()
	if err != nil {
		// handle with #870
		panic(err)
	}
}

func newApp(logger log.Logger, db dbm.DB, traceStore io.Writer) abci.Application {
	return app.NewRelayerApp(
		logger, db, traceStore, true, invCheckPeriod,
		baseapp.SetPruning(store.NewPruningOptionsFromString(viper.GetString("pruning"))),
		baseapp.SetMinGasPrices(viper.GetString(server.FlagMinGasPrices)),
	)
}

func exportAppStateAndTMValidators(
	logger log.Logger, db dbm.DB, traceStore io.Writer, height int64, forZeroHeight bool, jailWhiteList []string,
) (json.RawMessage, []tmtypes.GenesisValidator, error) {
	
	if height != -1 {
		rApp := app.NewRelayerApp(logger, db, traceStore, false, uint(1))
		err := rApp.LoadHeight(height)
		if err != nil {
			return nil, nil, err
		}
		return rApp.ExportAppStateAndValidators(forZeroHeight, jailWhiteList)
	}
	rApp := app.NewRelayerApp(logger, db, traceStore, true, uint(1))
	return rApp.ExportAppStateAndValidators(forZeroHeight, jailWhiteList)
}
