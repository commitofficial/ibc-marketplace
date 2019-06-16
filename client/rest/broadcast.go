package rest

import (
	"net/http"
	
	context2 "github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth"
)

func BroadcastRest(w http.ResponseWriter, cliCtx context2.CLIContext, cdc *codec.Codec, stdTx auth.StdTx, mode string) {
	
	txBytes, err := cdc.MarshalBinaryLengthPrefixed(stdTx)
	if err != nil {
		rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	cliCtx = cliCtx.WithBroadcastMode(mode)
	
	res, err := cliCtx.BroadcastTx(txBytes)
	if err != nil {
		rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	
	rest.PostProcessResponse(w, cdc, res, cliCtx.Indent)
}
