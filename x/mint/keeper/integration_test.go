package keeper_test

import (
	"encoding/json"

	dbm "github.com/cometbft/cometbft-db"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/libs/log"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"

	bluechipapp "github.com/BlueChip23/bluechip/app"
	"github.com/BlueChip23/bluechip/x/mint/types"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// returns context and an app with updated mint keeper
func createTestApp(isCheckTx bool) (*bluechipapp.App, sdk.Context) {
	app := setup(isCheckTx)

	ctx := app.BaseApp.NewContext(isCheckTx, tmproto.Header{})
	app.MintKeeper.SetParams(ctx, types.DefaultParams())
	app.MintKeeper.SetMinter(ctx, types.DefaultInitialMinter())

	return app, ctx
}

func setup(isCheckTx bool) *bluechipapp.App {
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

func genApp(withGenesis bool, invCheckPeriod uint) (*bluechipapp.App, bluechipapp.GenesisState) {
	db := dbm.NewMemDB()
	encCdc := bluechipapp.MakeEncodingConfig()
	app := bluechipapp.New(
		log.NewNopLogger(),
		db,
		nil,
		true,
		map[int64]bool{},
		simtestutil.DefaultNodeHome,
		invCheckPeriod,
		encCdc,
		bluechipapp.GetEnabledProposals(),
		simtestutil.EmptyAppOptions{},
		bluechipapp.GetWasmOpts(simtestutil.EmptyAppOptions{}),
	)

	if withGenesis {
		return app, bluechipapp.NewDefaultGenesisState(encCdc.Marshaler)
	}

	return app, bluechipapp.GenesisState{}
}
