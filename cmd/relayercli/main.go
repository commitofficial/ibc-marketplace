package main

import (
	"fmt"
	"net/http"
	"os"
	"path"
	
	client2 "ibc-marketplace/modules/orders/client"
	
	"github.com/cosmos/cosmos-sdk/x/mint"
	
	"github.com/rakyll/statik/fs"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	
	"github.com/tendermint/go-amino"
	"github.com/tendermint/tendermint/libs/cli"
	
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/client/lcd"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	
	"ibc-marketplace/app"
	"ibc-marketplace/modules/orders"
	"ibc-marketplace/version"
	
	at "github.com/cosmos/cosmos-sdk/x/auth"
	auth "github.com/cosmos/cosmos-sdk/x/auth/client/rest"
	bank "github.com/cosmos/cosmos-sdk/x/bank/client/rest"
	dist "github.com/cosmos/cosmos-sdk/x/distribution/client/rest"
	gv "github.com/cosmos/cosmos-sdk/x/gov"
	gov "github.com/cosmos/cosmos-sdk/x/gov/client/rest"
	mintrest "github.com/cosmos/cosmos-sdk/x/mint/client/rest"
	sl "github.com/cosmos/cosmos-sdk/x/slashing"
	slashing "github.com/cosmos/cosmos-sdk/x/slashing/client/rest"
	st "github.com/cosmos/cosmos-sdk/x/staking"
	staking "github.com/cosmos/cosmos-sdk/x/staking/client/rest"
	
	authcmd "github.com/cosmos/cosmos-sdk/x/auth/client/cli"
	bankcmd "github.com/cosmos/cosmos-sdk/x/bank/client/cli"
	crisisclient "github.com/cosmos/cosmos-sdk/x/crisis/client"
	distcmd "github.com/cosmos/cosmos-sdk/x/distribution"
	distClient "github.com/cosmos/cosmos-sdk/x/distribution/client"
	govClient "github.com/cosmos/cosmos-sdk/x/gov/client"
	mintclient "github.com/cosmos/cosmos-sdk/x/mint/client"
	slashingclient "github.com/cosmos/cosmos-sdk/x/slashing/client"
	stakingclient "github.com/cosmos/cosmos-sdk/x/staking/client"
	
	_ "github.com/cosmos/cosmos-sdk/client/lcd/statik"
)

func main() {
	cobra.EnableCommandSorting = false
	
	cdc := app.MakeCodec()
	
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(sdk.Bech32PrefixAccAddr, sdk.Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(sdk.Bech32PrefixValAddr, sdk.Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(sdk.Bech32PrefixConsAddr, sdk.Bech32PrefixConsPub)
	config.Seal()
	
	mc := []sdk.ModuleClients{
		govClient.NewModuleClient(gv.StoreKey, cdc),
		distClient.NewModuleClient(distcmd.StoreKey, cdc),
		stakingclient.NewModuleClient(st.StoreKey, cdc),
		mintclient.NewModuleClient(mint.StoreKey, cdc),
		slashingclient.NewModuleClient(sl.StoreKey, cdc),
		crisisclient.NewModuleClient(sl.StoreKey, cdc),
		client2.NewModuleClient(orders.MakerStoreKey, orders.TakerStoreKey, cdc),
	}
	
	rootCmd := &cobra.Command{
		Use:   "relayercli",
		Short: "Command line interface for interacting with relayerd",
	}
	
	// Add --chain-id to persistent flags and mark it required
	rootCmd.PersistentFlags().String(client.FlagChainID, "", "Chain ID of tendermint node")
	rootCmd.PersistentPreRunE = func(_ *cobra.Command, _ []string) error {
		return initConfig(rootCmd)
	}
	
	// Construct Root Command
	rootCmd.AddCommand(
		rpc.StatusCommand(),
		client.ConfigCmd(app.DefaultCLIHome),
		queryCmd(cdc, mc),
		txCmd(cdc, mc),
		client.LineBreak,
		lcd.ServeCommand(cdc, registerRoutes),
		client.LineBreak,
		keys.Commands(),
		client.LineBreak,
		version.VersionCmd,
		client.NewCompletionCmd(rootCmd, true),
	)
	
	// Add flags and prefix all env exposed with GA
	executor := cli.PrepareMainCmd(rootCmd, "RA", app.DefaultCLIHome)
	
	err := executor.Execute()
	if err != nil {
		fmt.Printf("Failed executing CLI command: %s, exiting...\n", err)
		os.Exit(1)
	}
}

func queryCmd(cdc *amino.Codec, mc []sdk.ModuleClients) *cobra.Command {
	queryCmd := &cobra.Command{
		Use:     "query",
		Aliases: []string{"q"},
		Short:   "Querying subcommands",
	}
	
	queryCmd.AddCommand(
		rpc.ValidatorCommand(cdc),
		rpc.BlockCommand(),
		tx.SearchTxCmd(cdc),
		tx.QueryTxCmd(cdc),
		client.LineBreak,
		authcmd.GetAccountCmd(at.StoreKey, cdc),
	)
	
	for _, m := range mc {
		mQueryCmd := m.GetQueryCmd()
		if mQueryCmd != nil {
			queryCmd.AddCommand(mQueryCmd)
		}
	}
	
	return queryCmd
}

func txCmd(cdc *amino.Codec, mc []sdk.ModuleClients) *cobra.Command {
	txCmd := &cobra.Command{
		Use:   "tx",
		Short: "Transactions subcommands",
	}
	
	txCmd.AddCommand(
		bankcmd.SendTxCmd(cdc),
		client.LineBreak,
		authcmd.GetSignCommand(cdc),
		authcmd.GetMultiSignCommand(cdc),
		tx.GetBroadcastCommand(cdc),
		tx.GetEncodeCommand(cdc),
		client.LineBreak,
	)
	
	for _, m := range mc {
		txCmd.AddCommand(m.GetTxCmd())
	}
	
	return txCmd
}

func registerRoutes(rs *lcd.RestServer) {
	// registerSwaggerUI(rs)
	rpc.RegisterRoutes(rs.CliCtx, rs.Mux)
	tx.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc)
	auth.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc, at.StoreKey)
	bank.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc, rs.KeyBase)
	dist.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc, distcmd.StoreKey)
	staking.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc, rs.KeyBase)
	slashing.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc, rs.KeyBase)
	gov.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc)
	mintrest.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc)
}

func registerSwaggerUI(rs *lcd.RestServer) {
	statikFS, err := fs.New()
	if err != nil {
		panic(err)
	}
	staticServer := http.FileServer(statikFS)
	rs.Mux.PathPrefix("/swagger-ui/").Handler(http.StripPrefix("/swagger-ui/", staticServer))
}

func initConfig(cmd *cobra.Command) error {
	home, err := cmd.PersistentFlags().GetString(cli.HomeFlag)
	if err != nil {
		return err
	}
	
	cfgFile := path.Join(home, "config", "config.toml")
	if _, err := os.Stat(cfgFile); err == nil {
		viper.SetConfigFile(cfgFile)
		
		if err := viper.ReadInConfig(); err != nil {
			return err
		}
	}
	if err := viper.BindPFlag(client.FlagChainID, cmd.PersistentFlags().Lookup(client.FlagChainID)); err != nil {
		return err
	}
	if err := viper.BindPFlag(cli.EncodingFlag, cmd.PersistentFlags().Lookup(cli.EncodingFlag)); err != nil {
		return err
	}
	return viper.BindPFlag(cli.OutputFlag, cmd.PersistentFlags().Lookup(cli.OutputFlag))
}
