package cli

import (
	flag "github.com/spf13/pflag"
)

const (
	FlagMakerAddress       = "makerAddress"
	FlagTakerAddress       = "takerAddress"
	FlagTakerAssetAmount   = "takerAssetAmount"
	FlagMakerAssetAmount   = "makerAssetAmount"
	FlagExpirationInHeight = "expirationInHeight"
	FlagTakerFillAmount    = "takerFillAmount"
	FlagVDFProof           = "vdfProof"
	FlagVDFIterations      = "vdfIterations"
	FlagPage               = "page"
	FlagPerPage            = "perPage"
	FlagOrderHash          = "orderHash"
)

var (
	fsMakerAddress       = flag.NewFlagSet("", flag.ContinueOnError)
	fsTakerAddress       = flag.NewFlagSet("", flag.ContinueOnError)
	fsTakerAssetAmount   = flag.NewFlagSet("", flag.ContinueOnError)
	fsMakerAssetAmount   = flag.NewFlagSet("", flag.ContinueOnError)
	fsExpirationInHeight = flag.NewFlagSet("", flag.ContinueOnError)
	fsVDFProof           = flag.NewFlagSet("", flag.ContinueOnError)
	fsVDFIterations      = flag.NewFlagSet("", flag.ContinueOnError)
	fsPage               = flag.NewFlagSet("", flag.ContinueOnError)
	fsOrderHash          = flag.NewFlagSet("", flag.ContinueOnError)
	
	fsPerPage         = flag.NewFlagSet("", flag.ContinueOnError)
	fsTakerFillAmount = flag.NewFlagSet("", flag.ContinueOnError)
)

func init() {
	fsMakerAddress.String(FlagMakerAddress, "", "maker address of an order")
	fsTakerAddress.String(FlagTakerAddress, "", "taker address")
	fsTakerAssetAmount.String(FlagTakerAssetAmount, "", "taker asset amount")
	fsMakerAssetAmount.String(FlagMakerAssetAmount, "", "maker asset amount")
	fsExpirationInHeight.String(FlagExpirationInHeight, "", "expiration height")
	fsVDFProof.String(FlagVDFProof, "", "vdf proof")
	fsVDFIterations.String(FlagVDFIterations, "", "vdf iterations")
	fsPage.String(FlagPage, "", "which page you want to access")
	fsPerPage.String(FlagPerPage, "", "per page how many you want to get")
	fsTakerFillAmount.String(FlagTakerFillAmount, "", "fill amount of taker of given order")
	fsOrderHash.String(FlagOrderHash, "", "order hash")
	
}
