package genesis

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/tendermint/go-amino"
	cfg "github.com/tendermint/tendermint/config"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/libs/common"
	"github.com/tendermint/tendermint/p2p"
	"github.com/tendermint/tendermint/privval"
	"github.com/tendermint/tendermint/types"
	
	"ibc-marketplace/app"
	
	"github.com/cosmos/cosmos-sdk/server"
)

func ExportGenesisFile(
	genFile, chainID string, validators []types.GenesisValidator, appState json.RawMessage,
) error {
	
	genDoc := types.GenesisDoc{
		ChainID:    chainID,
		Validators: validators,
		AppState:   appState,
	}
	
	if err := genDoc.ValidateAndComplete(); err != nil {
		return err
	}
	
	return genDoc.SaveAs(genFile)
}

// InitializeNodeValidatorFiles creates private validator and p2p configuration files.
func InitializeNodeValidatorFiles(config *cfg.Config,
) (nodeID string, valPubKey crypto.PubKey, err error) {
	
	nodeKey, err := p2p.LoadOrGenNodeKey(config.NodeKeyFile())
	if err != nil {
		return nodeID, valPubKey, err
	}
	
	nodeID = string(nodeKey.ID())
	server.UpgradeOldPrivValFile(config)
	
	pvKeyFile := config.PrivValidatorKeyFile()
	if err := common.EnsureDir(filepath.Dir(pvKeyFile), 0777); err != nil {
		return nodeID, valPubKey, nil
	}
	
	pvStateFile := config.PrivValidatorStateFile()
	if err := common.EnsureDir(filepath.Dir(pvStateFile), 0777); err != nil {
		return nodeID, valPubKey, nil
	}
	
	valPubKey = privval.LoadOrGenFilePV(pvKeyFile, pvStateFile).GetPubKey()
	
	return nodeID, valPubKey, nil
}

func initializeEmptyGenesis(
	cdc *codec.Codec, genFile, chainID string, overwrite bool,
) (appState json.RawMessage, err error) {
	
	if !overwrite && common.FileExists(genFile) {
		return nil, fmt.Errorf("genesis.json file already exists: %v", genFile)
	}
	
	return codec.MarshalJSONIndent(cdc, app.NewDefaultGenesisState())
}

// LoadGenesisDoc reads and unmarshals GenesisDoc from the given file.
func LoadGenesisDoc(cdc *amino.Codec, genFile string) (genDoc types.GenesisDoc, err error) {
	genContents, err := ioutil.ReadFile(genFile)
	if err != nil {
		return genDoc, err
	}
	
	if err := cdc.UnmarshalJSON(genContents, &genDoc); err != nil {
		return genDoc, err
	}
	
	return genDoc, err
}
