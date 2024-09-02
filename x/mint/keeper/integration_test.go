package keeper_test

import (
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"

	bluechipapp "github.com/BlueChip23/bluechip/app"
	"github.com/BlueChip23/bluechip/x/mint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// returns context and an app with updated mint keeper
func createTestApp(isCheckTx bool) (*bluechipapp.App, sdk.Context) {
	app := bluechipapp.Setup(isCheckTx)

	ctx := app.BaseApp.NewContext(isCheckTx, tmproto.Header{})
	app.MintKeeper.SetParams(ctx, types.DefaultParams())
	app.MintKeeper.SetMinter(ctx, types.DefaultInitialMinter())

	return app, ctx
}
