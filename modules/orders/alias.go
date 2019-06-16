package orders

import (
	"ibc-marketplace/modules/orders/keeper"
	"ibc-marketplace/modules/orders/querier"
	"ibc-marketplace/modules/orders/types"
)

type (
	BaseTakeOrder = types.BaseTakeOrder
	BaseMakeOrder = types.BaseMakeOrder
	Keeper = keeper.Keeper
)

var (
	NewKeeper             = keeper.NewKeeper
	NewQuerier            = querier.NewQuerier
	RegisterCodec         = types.RegisterCodec
	MsgCreateMakeOrder    = types.NewMsgCreateMakeOrder
	SignBytesForMakeOrder = types.SignBytesForMakeOrder
	MsgSubmitTakeOrder    = types.NewMsgSubmitTakeOrder
)

const (
	MakerStoreKey = types.MakerStoreKey
	TakerStoreKey = types.TakerStoreKey
	ModuleName    = types.ModuleName
	QuerierRoute  = types.QuerierRoute
	RouterKey     = types.RouterKey
)
