package client

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
	"github.com/tendermint/go-amino"
	
	"ibc-marketplace/modules/orders/client/cli"
)

// ModuleClient exports all client functionality from this module
type ModuleClient struct {
	makerStoreKey string
	takerStoreKey string
	cdc           *amino.Codec
}

func NewModuleClient(makerStoreKey, takerStoreKey string, cdc *amino.Codec) ModuleClient {
	return ModuleClient{makerStoreKey, takerStoreKey, cdc}
}

// GetQueryCmd returns the cli query commands for this module
func (mc ModuleClient) GetQueryCmd() *cobra.Command {
	// Group orders queries under a subcommand
	relayerQueryCmd := &cobra.Command{
		Use:   "orders",
		Short: "Querying commands for the orders module",
	}
	
	relayerQueryCmd.AddCommand(client.GetCommands()...)
	
	return relayerQueryCmd
}

// GetTxCmd returns the transaction commands for this module
func (mc ModuleClient) GetTxCmd() *cobra.Command {
	relayerTxCmd := &cobra.Command{
		Use:   "orders",
		Short: "orders transactions subcommands",
	}
	
	relayerTxCmd.AddCommand(client.PostCommands(
		cli.MakeOrderTxCmd(mc.cdc),
		// cli.TakeOrderTxCmd(mc.cdc),
	)...)
	
	return relayerTxCmd
}
