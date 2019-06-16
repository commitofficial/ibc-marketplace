package rest

import (
	"bytes"
	"fmt"
	
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authtxb "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
)

func SignStdTxFromRest(txBldr authtxb.TxBuilder, cliCtx context.CLIContext, name string, stdTx auth.StdTx, appendSig bool, offline bool, password string) (auth.StdTx, error) {
	
	var signedStdTx auth.StdTx
	
	info, err := txBldr.Keybase().Get(name)
	if err != nil {
		return signedStdTx, err
	}
	
	addr := info.GetPubKey().Address()
	
	// check whether the address is a signer
	if !isTxSigner(sdk.AccAddress(addr), stdTx.GetSigners()) {
		return signedStdTx, fmt.Errorf("%s: %s", client.ErrInvalidSigner, name)
	}
	
	if !offline {
		txBldr, err = populateAccountFromState(txBldr, cliCtx, sdk.AccAddress(addr))
		if err != nil {
			return signedStdTx, err
		}
	}
	
	return txBldr.SignStdTx(name, password, stdTx, appendSig)
}

func isTxSigner(user sdk.AccAddress, signers []sdk.AccAddress) bool {
	for _, s := range signers {
		if bytes.Equal(user.Bytes(), s.Bytes()) {
			return true
		}
	}
	
	return false
}

func populateAccountFromState(
	txBldr authtxb.TxBuilder, cliCtx context.CLIContext, addr sdk.AccAddress,
) (authtxb.TxBuilder, error) {
	
	accNum, err := cliCtx.GetAccountNumber(addr)
	if err != nil {
		return txBldr, err
	}
	
	accSeq, err := cliCtx.GetAccountSequence(addr)
	if err != nil {
		return txBldr, err
	}
	
	return txBldr.WithAccountNumber(accNum).WithSequence(accSeq), nil
}
