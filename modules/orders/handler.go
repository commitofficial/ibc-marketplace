package orders

import (
	ctypes "github.com/cosmos/cosmos-sdk/types"
	keeper "github.com/ibc-marketplace/modules/orders/keeper"
	"fmt"
)

// NewHandler returns a handler for "orders" type messages.
func NewHandler(keeper keeper.Keeper) ctypes.Handler {
	return func(ctx ctypes.Context, msg ctypes.Msg) ctypes.Result {
		switch msg := msg.(type) {
		
		default:
			errMsg := fmt.Sprintf("Unrecognized order Msg type: %v", msg.Type())
			return ctypes.ErrUnknownRequest(errMsg).Result()
		}
	}
}
