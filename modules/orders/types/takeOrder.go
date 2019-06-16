package types

import (
	ctypes "github.com/comdex-blockchain/types"
)

type BaseTakeOrder struct {
	Order         BaseMakeOrder
	VdfProof      string            `json:"vdfProof"`
	VdtIterations ctypes.Uint       `json:"vdfIterations"`
	FromAddress   ctypes.AccAddress `json:"fromAddress"`
	OrderHash     string            `json:"orderHash"`
}
