package types

import (
	ctypes "github.com/cosmos/cosmos-sdk/types"
)

const SubmitTakeOrderType = "submitTakeOrder"

type MsgSubmitTakeOrder struct {
	TakeFillAmount ctypes.Coin
	OrderHash      string
	VDFProof       string
		VDFIterations  ctypes.Uint
	FromAddress    ctypes.AccAddress
}

func NewMsgSubmitTakeOrder(takeAmount ctypes.Coin, hash string, vdfProof string, vdfIterations ctypes.Uint, from ctypes.AccAddress) MsgSubmitTakeOrder {
	return MsgSubmitTakeOrder{
		TakeFillAmount: takeAmount,
		OrderHash:      hash,
		VDFProof:       vdfProof,
		VDFIterations:  vdfIterations,
		FromAddress:    from,
	}
}

var _ ctypes.Msg = MsgSubmitTakeOrder{}

func (msg MsgSubmitTakeOrder) Type() string  { return SubmitTakeOrderType }
func (msg MsgSubmitTakeOrder) Route() string { return RouterKey }
func (msg MsgSubmitTakeOrder) GetSigners() []ctypes.AccAddress {
	return []ctypes.AccAddress{msg.FromAddress}
}
func (msg MsgSubmitTakeOrder) GetSignBytes() []byte        { return nil }
func (msg MsgSubmitTakeOrder) ValidateBasic() ctypes.Error { return nil }
