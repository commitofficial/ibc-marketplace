package types

import (
	ctypes "github.com/comdex-blockchain/types"
)

type OrderStatus byte

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
	ExpirationHeight int64
	OrderHash        string
	Signature        string
	Status           OrderStatus
}
