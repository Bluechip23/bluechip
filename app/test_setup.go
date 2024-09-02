package app

import (
	"encoding/json"

	dbm "github.com/cometbft/cometbft-db"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/libs/log"

	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
)

func Setup(isCheckTx bool) *App {
	app, genesisState := genApp(!isCheckTx, 5)
	if !isCheckTx {
		// init chain must be called to stop deliverState from being nil
		stateBytes, err := json.MarshalIndent(genesisState, "", " ")
		if err != nil {
			panic(err)
		}

		// Initialize the chain
		app.InitChain(
			abci.RequestInitChain{
				Validators:      []abci.ValidatorUpdate{},
				ConsensusParams: simtestutil.DefaultConsensusParams,
				AppStateBytes:   stateBytes,
			},
		)
	}

	return app
}

func genApp(withGenesis bool, invCheckPeriod uint) (*App, GenesisState) {
	db := dbm.NewMemDB()
	encCdc := MakeEncodingConfig()
	app := New(
		log.NewNopLogger(),
		db,
		nil,
		true,
		map[int64]bool{},
		DefaultNodeHome,
		invCheckPeriod,
		encCdc,
		GetEnabledProposals(),
		simtestutil.EmptyAppOptions{},
		GetWasmOpts(simtestutil.EmptyAppOptions{}),
	)

	if withGenesis {
		return app, NewDefaultGenesisState(encCdc.Marshaler)
	}

	return app, GenesisState{}
}
