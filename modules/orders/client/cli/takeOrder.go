package cli

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtxb "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	
	"ibc-marketplace/modules/orders"
)

func TakeOrderTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "takeOrder",
		Short: "take order",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)
			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			
			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}
			
			from := cliCtx.GetFromAddress()
			
			orderHash := viper.GetString(FlagOrderHash)
			takerFillAmount := viper.GetString(FlagTakerFillAmount)
			vdfProof := viper.GetString(FlagVDFProof)
			vdfIterations := sdk.NewUintFromString(viper.GetString(FlagVDFIterations))
			
			baseToken, err := sdk.ParseCoin(takerFillAmount)
			if err != nil {
				return err
			}
			
			msg := orders.MsgSubmitTakeOrder(baseToken, orderHash, vdfProof, vdfIterations, from)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}
			
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
		},
	}
	cmd.Flags().AddFlagSet(fsVDFIterations)
	cmd.Flags().AddFlagSet(fsVDFProof)
	cmd.Flags().AddFlagSet(fsTakerFillAmount)
	cmd.Flags().AddFlagSet(fsOrderHash)
	
	_ = cmd.MarkFlagRequired(FlagVDFIterations)
	_ = cmd.MarkFlagRequired(FlagVDFProof)
	_ = cmd.MarkFlagRequired(FlagOrderHash)
	_ = cmd.MarkFlagRequired(FlagTakerFillAmount)
	
	return cmd
}
