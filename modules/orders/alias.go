package orders

import (
	"github.com/ibc-marketplace/modules/orders/keeper"
	"github.com/ibc-marketplace/modules/orders/types"
	"github.com/ibc-marketplace/modules/orders/querier"
)

type (
	BaseTakeOrder = types.BaseTakeOrder
	BaseMakeOrder = types.BaseMakeOrder
	Keeper = keeper.Keeper
)

var (
	NewKeeper     = keeper.NewKeeper
	NewQuerier    = querier.NewQuerier
	RegisterCodec = types.RegisterCodec
)

const (
	MakerStoreKey = types.MakerStoreKey
	TakerStoreKey = types.TakerStoreKey
	ModuleName    = types.ModuleName
	QuerierRoute  = types.QuerierRoute
	RouterKey     = types.RouterKey
)
