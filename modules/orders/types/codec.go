package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)



func RegisterCodec(cdc *codec.Codec) {
	// cdc.RegisterConcrete(CreateTakeOrderMsg{}, "orders/CreateTakeOrderMsg", nil)
	// cdc.RegisterConcrete(CreateMakeOrderMsg{}, "orders/CreateMakeOrderMsg", nil)
	// cdc.RegisterConcrete(ZeroExOrder{}, "order", nil)
	// cdc.RegisterConcrete(BaseMakeOrder{}, "baseMakeOrder", nil)
	// cdc.RegisterConcrete(AddFeeRecipientMsg{}, "orders/FeeRecipientMsg", nil)
	// cdc.RegisterConcrete(DeleteFeeRecipientMsg{}, "orders/DeleteFeeRecipientMsg", nil)
	// cdc.RegisterConcrete(OrderConfigMsg{}, "orders/orderConfigMsg", nil)
}
var MsgCdc *codec.Codec

func init() {
	MsgCdc = codec.New()
}
