package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgCreateMakeOrder{}, "orders/CreateMakeOrderMsg", nil)
	cdc.RegisterConcrete(MsgSubmitTakeOrder{}, "orders/CreateTakeOrderMsg", nil)
}

var MsgCdc *codec.Codec

func init() {
	MsgCdc = codec.New()
}
