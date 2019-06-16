package orders

import (
	ctypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"ibc-marketplace/modules/orders/tags"
	
	keeper "ibc-marketplace/modules/orders/keeper"
	"ibc-marketplace/modules/orders/types"
	
	"fmt"
)

// NewHandler returns a handler for "orders" type messages.
func NewHandler(keeper keeper.Keeper, bk bank.Keeper) ctypes.Handler {
	return func(ctx ctypes.Context, msg ctypes.Msg) ctypes.Result {
		switch msg := msg.(type) {
		case types.MsgCreateMakeOrder:
			return handleCreateMakeOrder(keeper, bk, ctx, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized order Msg type: %v", msg.Type())
			return ctypes.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleCreateMakeOrder(k keeper.Keeper, bk bank.Keeper, ctx ctypes.Context, order types.MsgCreateMakeOrder) ctypes.Result {
	if !bk.HasCoins(ctx, order.FromAddress, ctypes.Coins{order.BaseToken}) {
		return ctypes.ErrInsufficientCoins("coins not available in your account").Result()
	}
	
	baseMakeOrder := types.NewBaseMakeOrder(order.BaseToken, order.QuoteToken, order.FromAddress, nil, order.ExpirationHeight, order.OrderHash, order.Signature, types.StatusUnFilled)
	k.SetMakeOrderByOrderHash(ctx, baseMakeOrder)
	
	resTags := ctypes.NewTags(
		tags.FromAddress, order.FromAddress.String(),
		tags.OrderHash, order.OrderHash,
	)
	
	return ctypes.Result{
		Tags: resTags,
	}
}
