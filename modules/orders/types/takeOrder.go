package types

import (
	ctypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/codec"
)

type BaseTakeOrder struct {
	Order           BaseMakeOrder
	TakerFillAmount ctypes.Coin
	VdfProof        string
	VdtIterations   ctypes.Uint
	FromAddress     ctypes.AccAddress
	OrderHash       string
}

func MustMarshalTakeOrder(cdc *codec.Codec, baseTakeOrders []BaseTakeOrder) []byte {
	return cdc.MustMarshalBinaryLengthPrefixed(baseTakeOrders)
}

func MustUnMarshalTakeOrder(cdc *codec.Codec, value []byte) ([]BaseTakeOrder, ctypes.Error) {
	orders, err := unMarshalTakeOrder(cdc, value)
	if err != nil {
		return []BaseTakeOrder{}, ctypes.ErrInternal("cannot unmarshal make order ")
	}
	
	return orders, nil
}

func unMarshalTakeOrder(cdc *codec.Codec, value []byte) (orders []BaseTakeOrder, err error) {
	err = cdc.UnmarshalBinaryLengthPrefixed(value, &orders)
	return orders, err
}
