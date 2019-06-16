package types

import (
	ctypes "github.com/cosmos/cosmos-sdk/types"
)

const (
	CreateMakeOrderType = "createMakeOrderMsg"
)

type MsgCreateMakeOrder struct {
	BaseToken        ctypes.Coin
	QuoteToken       ctypes.Coin
	FromAddress      ctypes.AccAddress
	ExpirationHeight uint64
	Signature        []byte
	OrderHash        string
}

func NewMsgCreateMakeOrder(baseToken, quoteToken ctypes.Coin, from ctypes.AccAddress, height uint64, sign []byte, hash string) MsgCreateMakeOrder {
	return MsgCreateMakeOrder{
		BaseToken:        baseToken,
		QuoteToken:       quoteToken,
		FromAddress:      from,
		ExpirationHeight: height,
		Signature:        sign,
		OrderHash:        hash,
	}
}

var _ ctypes.Msg = MsgCreateMakeOrder{}

func (msg MsgCreateMakeOrder) Type() string                { return CreateMakeOrderType }
func (msg MsgCreateMakeOrder) Route() string               { return RouterKey }
func (msg MsgCreateMakeOrder) ValidateBasic() ctypes.Error { return nil }
func (msg MsgCreateMakeOrder) GetSigners() []ctypes.AccAddress {
	return []ctypes.AccAddress{msg.FromAddress}
}
func (msg MsgCreateMakeOrder) GetSignBytes() []byte { return nil }
