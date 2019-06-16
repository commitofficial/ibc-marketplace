package keeper

import (
	"github.com/ibc-marketplace/modules/orders/types"
	
	ctypes "github.com/cosmos/cosmos-sdk/types"
)

// Set MakeOrders
func (keeper Keeper) SetMakeOrderByOrderHash(ctx ctypes.Context, baseMakeOrder types.BaseMakeOrder) {
	store := ctx.KVStore(keeper.makerStoreKey)
	bz := types.MustMarshalMakeOrder(keeper.cdc, baseMakeOrder)
	store.Set(GetMakeOrderKey([]byte(baseMakeOrder.OrderHash)), bz)
	keeper.SetMakeOrdersByAddress(ctx, baseMakeOrder.MakerAddress, baseMakeOrder.OrderHash)
	
}

func (keeper Keeper) GetMakeOrderByHash(ctx ctypes.Context, orderHash []byte) (types.BaseMakeOrder, ctypes.Error) {
	store := ctx.KVStore(keeper.takerStoreKey)
	bz := store.Get(GetMakeOrderKey(orderHash))
	if bz == nil {
		return types.BaseMakeOrder{}, ctypes.ErrInternal("orders doesn't exist")
	}
	return types.MustUnMarshalMakeOrder(keeper.cdc, bz)
}

func (keeper Keeper) SetMakeOrdersByAddress(ctx ctypes.Context, address ctypes.AccAddress, orderHash string) {
	store := ctx.KVStore(keeper.makerStoreKey)
	hashs, _ := keeper.GetMakeOrdersByAddress(ctx, address)
	hashs = append(hashs, orderHash)
	
	store.Set(GetOrdersByAddressKey(address.Bytes()), types.MustMarshalOrdersByAddress(keeper.cdc, hashs))
}

func (keeper Keeper) GetMakeOrdersByAddress(ctx ctypes.Context, address ctypes.Address) (types.OrderHashes, ctypes.Error) {
	var orderHashes types.OrderHashes
	store := ctx.KVStore(keeper.makerStoreKey)
	data := store.Get(GetOrdersByAddressKey(address.Bytes()))
	if data == nil {
		return orderHashes, ctypes.ErrInternal("Orders not exist")
	}
	
	return types.MustUnMarshalOrdersByAddress(keeper.cdc, data)
}

// Get list of MakeOrders
func (keeper Keeper) GetOrderBook(ctx ctypes.Context) ([]types.BaseMakeOrder, ctypes.Error) {
	var orders []types.BaseMakeOrder
	
	store := ctx.KVStore(keeper.makerStoreKey)
	iterator := ctypes.KVStorePrefixIterator(store, MakeOrderKey)
	defer iterator.Close()
	
	for ; iterator.Valid(); iterator.Next() {
		order, err := types.MustUnMarshalMakeOrder(keeper.cdc, iterator.Value())
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}
