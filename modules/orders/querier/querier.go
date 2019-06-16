package querier

import (
	"github.com/cosmos/cosmos-sdk/codec"
	ctypes "github.com/cosmos/cosmos-sdk/types"
	"ibc-marketplace/modules/orders/keeper"
	abci "github.com/tendermint/tendermint/abci/types"
)

// NewQuerier is the module level router for state queries
func NewQuerier(keeper keeper.Keeper, cdc *codec.Codec) ctypes.Querier {
	return func(ctx ctypes.Context, path []string, req abci.RequestQuery) ([]byte, ctypes.Error) {
		switch path[0] {
		
		default:
			return nil, ctypes.ErrUnknownRequest("unknown orders query endpoint")
		}
	}
}
