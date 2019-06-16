package cli

import (
	"encoding/hex"
	
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtxb "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	
	"ibc-marketplace/modules/orders"
)

const DefalultPassword = "rgukt123"

// GetCmdMakeOrder :
func MakeOrderTxCmd(cdc *codec.Codec) *cobra.Command {
	
	cmd := &cobra.Command{
		Use:   "makeOrder",
		Short: "set make orders",
		RunE: func(cmd *cobra.Command, args []string) error {
			
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)
			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			
			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}
			
			from := cliCtx.GetFromAddress()
			
			makerAssetAmount := viper.GetString(FlagMakerAssetAmount)
			takerAssetAmount := viper.GetString(FlagTakerAssetAmount)
			expirationHeight := viper.GetInt64(FlagExpirationInHeight)
			
			baseAssetData, err := sdk.ParseCoin(makerAssetAmount)
			if err != nil {
				return err
			}
			
			quoteAssetData, err := sdk.ParseCoin(takerAssetAmount)
			if err != nil {
				return err
			}
			
			keybase, err := keys.NewKeyBaseFromHomeFlag()
			if err != nil {
				return err
			}
			
			signBytes := orders.SignBytesForMakeOrder(from, baseAssetData, quoteAssetData, uint64(expirationHeight))
			signature, _, err := keybase.Sign(cliCtx.FromName, DefalultPassword, signBytes)
			if err != nil {
				return err
			}
			
			orderHash := hex.EncodeToString(signature)
			
			msg := orders.MsgCreateMakeOrder(baseAssetData, quoteAssetData, from, uint64(expirationHeight), signature, orderHash[0:64])
			
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
		},
	}
	
	cmd.Flags().AddFlagSet(fsTakerAssetAmount)
	cmd.Flags().AddFlagSet(fsMakerAssetAmount)
	cmd.Flags().AddFlagSet(fsExpirationInHeight)
	
	_ = cmd.MarkFlagRequired(FlagTakerAssetAmount)
	_ = cmd.MarkFlagRequired(FlagMakerAssetAmount)
	_ = cmd.MarkFlagRequired(FlagExpirationInHeight)
	return cmd
	
}
