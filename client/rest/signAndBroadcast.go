package rest

import (
	"net/http"
	
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	
	"github.com/cosmos/cosmos-sdk/x/auth"
	authtxb "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
)

func SignAndBroadcast(w http.ResponseWriter, br rest.BaseReq, cliCtx context.CLIContext, cdc *codec.Codec, mode, password string, msgs []sdk.Msg) {
	
	gasAdj, ok := rest.ParseFloat64OrReturnBadRequest(w, br.GasAdjustment, client.DefaultGasAdjustment)
	if !ok {
		return
	}
	
	simAndExec, gas, err := client.ParseGas(br.Gas)
	if err != nil {
		rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	
	keybase, err := keys.NewKeyBaseFromHomeFlag()
	if err != nil {
		rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	
	txBldr := authtxb.NewTxBuilder(
		utils.GetTxEncoder(cdc), br.AccountNumber, br.Sequence, gas, gasAdj,
		br.Simulate, br.ChainID, br.Memo, br.Fees, br.GasPrices,
	)
	txBldr = txBldr.WithKeybase(keybase)
	
	if br.Simulate || simAndExec {
		if gasAdj < 0 {
			rest.WriteErrorResponse(w, http.StatusBadRequest, client.ErrInvalidGasAdjustment.Error())
			return
		}
		
		txBldr, err = utils.EnrichWithGas(txBldr, cliCtx, msgs)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		
		if br.Simulate {
			rest.WriteSimulationResponse(w, cdc, txBldr.Gas())
			return
		}
	}
	
	stdMsg, err := txBldr.BuildSignMsg(msgs)
	if err != nil {
		rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	
	stdTx := auth.NewStdTx(stdMsg.Msgs, stdMsg.Fee, nil, stdMsg.Memo)
	
	stdTx, err = SignStdTxFromRest(txBldr, cliCtx, cliCtx.GetFromName(), stdTx, true, false, password)
	if err != nil {
		rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	
	BroadcastRest(w, cliCtx, cdc, stdTx, mode)
	
}
