package keeper

import (
	ctypes "github.com/comdex-blockchain/types"
)

var (
	MakeOrderKey = []byte{0x01}
	
	TakeOrderKey = []byte{0x02}
	
	AddressKey = []byte{0x03}
)

func GetMakeOrderKey(orderHash []byte) []byte {
	return append(MakeOrderKey, orderHash...)
}

func GetTakeOrderKey(orderHash []byte) []byte {
	return append(TakeOrderKey, orderHash...)
}

func GetOrdersByAddressKey(address ctypes.AccAddress) []byte {
	return append(AddressKey, address...)
}
