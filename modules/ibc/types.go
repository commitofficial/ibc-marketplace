package ibc

import (
	ctypes "github.com/comdex-blockchain/types"
	
	"github.com/ibc-marketplace/modules/orders"
)

const (
	Exchange = "ibc/exchange"
	RouteKey = "ibc"
)

type IBCPacket struct {
	SourceAddress      string
	DestinationAddress string
	TakeOrder          orders.BaseTakeOrder
}

type MsgIBCExchangeOrder struct {
	IBCPacket
	Relayer ctypes.AccAddress
}

var _ ctypes.Msg = MsgIBCExchangeOrder{}

func (msg MsgIBCExchangeOrder) Type() string         { return Exchange }
func (msg MsgIBCExchangeOrder) Route() string        { return RouteKey }
func (msg MsgIBCExchangeOrder) GetSignBytes() []byte { return nil }
func (msg MsgIBCExchangeOrder) GetSigners() []ctypes.AccAddress {
	return []ctypes.AccAddress{msg.Relayer}
}
func (msg MsgIBCExchangeOrder) ValidateBasic() ctypes.Error { return nil }
