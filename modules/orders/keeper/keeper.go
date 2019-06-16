package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	ctypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
)

type Keeper struct {
	cdc           *codec.Codec
	makerStoreKey ctypes.StoreKey
	takerStoreKey ctypes.StoreKey
	bankKeeper    bank.Keeper
	
	// codespace ctypes.CodespaceType
}

func NewKeeper(cdc *codec.Codec, makerStoreKey, takerStoreKey ctypes.StoreKey, bankKeeper bank.Keeper) Keeper {
	return Keeper{
		cdc:           cdc,
		makerStoreKey: makerStoreKey,
		takerStoreKey: takerStoreKey,
		bankKeeper:    bankKeeper,
	}
}
