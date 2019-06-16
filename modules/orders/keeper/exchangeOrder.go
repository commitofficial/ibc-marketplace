package keeper

import (
	"fmt"
	"strings"
	
	ctypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	
	"ibc-marketplace/modules/orders/tags"
	"ibc-marketplace/modules/orders/types"
)

func (keeper Keeper) ExchangeOrder(ctx ctypes.Context, bk bank.Keeper, takerAddress ctypes.AccAddress, takerFillAmount ctypes.Coin, orderHash []byte) (ctypes.Tags, ctypes.Error) {
	baseMakeOrder, err := keeper.GetMakeOrderByHash(ctx, orderHash)
	if err != nil {
		return nil, err
	}
	
	// atomic swap based on baseMakeOrder
	if !strings.EqualFold(takerFillAmount.Denom, baseMakeOrder.QuoteToken.Denom) {
		return nil, ctypes.ErrInternal("tokens are not same")
	}
	takerAmount := (takerFillAmount.Amount.Quo(baseMakeOrder.QuoteToken.Amount)).Mul(baseMakeOrder.BaseToken.Amount)
	
	takerAsset := ctypes.Coin{baseMakeOrder.BaseToken.Denom, takerAmount}
	if takerAsset.IsNegative() {
		return nil, ctypes.ErrInternal("negative amount")
	}
	
	// Deduct tokens amount from both accounts
	
	_, _, err = bk.SubtractCoins(ctx, takerAddress, ctypes.Coins{takerFillAmount})
	if err != nil {
		return nil, ctypes.ErrInsufficientCoins("insufficient coins to subtract from taker")
	}
	
	coins, _, err := bk.SubtractCoins(ctx, baseMakeOrder.MakerAddress, ctypes.Coins{takerAsset})
	fmt.Println("maker coins", coins, err)
	if err != nil {
		return nil, ctypes.ErrInsufficientCoins("insufficient coins to subtract from maker")
	}
	
	// Adding tokens amount
	_, _, err = bk.AddCoins(ctx, takerAddress, ctypes.Coins{takerAsset})
	if err != nil {
		return nil, ctypes.ErrInsufficientCoins("insufficient coins to add for taker")
	}
	
	_, _, err = bk.AddCoins(ctx, baseMakeOrder.MakerAddress, ctypes.Coins{takerFillAmount})
	if err != nil {
		return nil, ctypes.ErrInsufficientCoins("insufficient coins to add for maker")
	}
	
	baseMakeOrder.Status = types.StatusFilled
	baseMakeOrder.TakerAddress = takerAddress
	
	keeper.SetMakeOrderByOrderHash(ctx, baseMakeOrder)
	
	resTags := ctypes.NewTags(
		tags.MakerAddress, baseMakeOrder.MakerAddress.String(),
		tags.TakerAddress, takerAddress.String(),
		tags.TakerFillAmount, takerFillAmount.String(),
	)
	
	return resTags, nil
}
