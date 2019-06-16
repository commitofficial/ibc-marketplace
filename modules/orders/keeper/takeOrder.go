package keeper

import (
	"sort"
	
	"github.com/ibc-marketplace/modules/orders/types"
	
	ctypes "github.com/cosmos/cosmos-sdk/types"
)

// set TakeOrderByHash => list of takeOrders by OrderHash
func (keeper Keeper) SetTakeOrdersByHash(ctx ctypes.Context, baseTakeOrder types.BaseTakeOrder) {
	store := ctx.KVStore(keeper.takerStoreKey)
	baseTakeOrders, _ := keeper.GetTakeOrdersByHash(ctx, []byte(baseTakeOrder.OrderHash))
	baseTakeOrders = append(baseTakeOrders, baseTakeOrder)
	store.Set(GetTakeOrderKey([]byte(baseTakeOrder.OrderHash)), types.MustMarshalTakeOrder(keeper.cdc, baseTakeOrders))
	
	// set OrderHash to Respective TakerAddress
	keeper.SetTakeOrdersByAddress(ctx, baseTakeOrder.FromAddress, baseTakeOrder.OrderHash)
}

func (keeper Keeper) GetTakeOrdersByHash(ctx ctypes.Context, orderHash []byte) ([]types.BaseTakeOrder, ctypes.Error) {
	store := ctx.KVStore(keeper.takerStoreKey)
	data := store.Get(GetTakeOrderKey(orderHash))
	if data == nil {
		return []types.BaseTakeOrder{}, ctypes.ErrInternal("take orders doesn't exist")
	}
	
	orders, err := types.MustUnMarshalTakeOrder(keeper.cdc, data)
	if err != nil {
		return []types.BaseTakeOrder{}, err
	}
	sort.Slice(orders, func(i, j int) bool {
		return orders[i].VdtIterations.GT(orders[j].VdtIterations)
	})
	
	return orders, nil
}

// getter and setter for the get the orders based on sender address
func (keeper Keeper) SetTakeOrdersByAddress(ctx ctypes.Context, address ctypes.AccAddress, orderHash string) {
	store := ctx.KVStore(keeper.takerStoreKey)
	hashs, _ := keeper.GetMakeOrdersByAddress(ctx, address)
	hashs = append(hashs, orderHash)
	
	store.Set(GetOrdersByAddressKey(address.Bytes()), types.MustMarshalOrdersByAddress(keeper.cdc, hashs))
}

func (keeper Keeper) GetTakeOrdersByAddress(ctx ctypes.Context, address ctypes.Address) (types.OrderHashes, ctypes.Error) {
	var orderHashes types.OrderHashes
	store := ctx.KVStore(keeper.takerStoreKey)
	data := store.Get(GetOrdersByAddressKey(address.Bytes()))
	if data == nil {
		return orderHashes, ctypes.ErrInternal("Orders not exist")
	}
	
	return types.MustUnMarshalOrdersByAddress(keeper.cdc, data)
}

// get Take Orders
func (keeper Keeper) GetTakerOrders(ctx ctypes.Context) ([][]types.BaseTakeOrder, ctypes.Error) {
	var order []types.BaseTakeOrder
	var orders [][]types.BaseTakeOrder
	store := ctx.KVStore(keeper.takerStoreKey)
	
	iterator := ctypes.KVStorePrefixIterator(store, TakeOrderKey)
	defer iterator.Close()
	
	for ; iterator.Valid(); iterator.Next() {
		
		keeper.cdc.MustUnmarshalBinaryLengthPrefixed(iterator.Value(), &order)
		orders = append(orders, order)
	}
	return orders, nil
}
