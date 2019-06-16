package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	ctypes "github.com/cosmos/cosmos-sdk/types"
)

type OrderStatus byte

type OrderHashes []string

const (
	StatusUnFilled  OrderStatus = 0x01
	StatusCancelled OrderStatus = 0x02
	StatusFilled    OrderStatus = 0x03
)

func (status OrderStatus) String() string {
	switch status {
	case StatusUnFilled:
		return "UnFilled"
	case StatusCancelled:
		return "Cancelled"
	case StatusFilled:
		return "Filled"
	default:
		return ""
	}
}

// Base MakeOrder
type BaseMakeOrder struct {
	BaseToken        ctypes.Coin
	QuoteToken       ctypes.Coin
	MakerAddress     ctypes.AccAddress
	TakerAddress     ctypes.AccAddress
	ExpirationHeight uint64
	OrderHash        string
	Signature        []byte
	Status           OrderStatus
}

func NewBaseMakeOrder(baseToken, quoteToken ctypes.Coin, makerAddress, takerAddress ctypes.AccAddress, height uint64, hash string, sign []byte, status OrderStatus) BaseMakeOrder {
	return BaseMakeOrder{
		BaseToken:        baseToken,
		QuoteToken:       quoteToken,
		MakerAddress:     makerAddress,
		TakerAddress:     takerAddress,
		ExpirationHeight: height,
		OrderHash:        hash,
		Signature:        sign,
		Status:           status,
	}
}

func MustMarshalMakeOrder(cdc *codec.Codec, baseMakeOrder BaseMakeOrder) []byte {
	return cdc.MustMarshalBinaryLengthPrefixed(baseMakeOrder)
}

func MustUnMarshalMakeOrder(cdc *codec.Codec, value []byte) (BaseMakeOrder, ctypes.Error) {
	order, err := unMarshalMakeOrder(cdc, value)
	if err != nil {
		return BaseMakeOrder{}, ctypes.ErrInternal("cannot unmarshal make order ")
	}
	
	return order, nil
}

func unMarshalMakeOrder(cdc *codec.Codec, value []byte) (order BaseMakeOrder, err error) {
	err = cdc.UnmarshalBinaryLengthPrefixed(value, &order)
	return order, err
}

func MustMarshalOrdersByAddress(cdc *codec.Codec, hashes OrderHashes) []byte {
	return cdc.MustMarshalBinaryLengthPrefixed(hashes)
}

func MustUnMarshalOrdersByAddress(cdc *codec.Codec, value []byte) (OrderHashes, ctypes.Error) {
	order, err := unMarshalOrderHashes(cdc, value)
	if err != nil {
		return OrderHashes{}, ctypes.ErrInternal("cannot unmarshal order  ")
	}
	
	return order, nil
}
func unMarshalOrderHashes(cdc *codec.Codec, value []byte) (order OrderHashes, err error) {
	err = cdc.UnmarshalBinaryLengthPrefixed(value, &order)
	return order, err
}

func SignBytesForMakeOrder(makerAddress ctypes.AccAddress, baseToken, quoteToken ctypes.Coin, height uint64) []byte {
	bz, err := MsgCdc.MarshalJSON(struct {
		BaseToken        ctypes.Coin
		QuoteToken       ctypes.Coin
		MakerAddress     ctypes.AccAddress
		ExpirationHeight uint64
	}{
		BaseToken:        baseToken,
		QuoteToken:       quoteToken,
		MakerAddress:     makerAddress,
		ExpirationHeight: height,
	})
	
	if err != nil {
		panic(err)
	}
	return ctypes.MustSortJSON(bz)
}
