package types

import (
	ctypes "github.com/comdex-blockchain/types"
)

type BaseTakeOrder struct {
	Order           BaseMakeOrder
	TakerFillAmount ctypes.Coin
	VdfProof        string
	VdtIterations   ctypes.Uint
	FromAddress     ctypes.AccAddress
	OrderHash       string
}
