package pageinflation

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/smartdev0328/bluechip/x/pageinflation/keeper"
	"github.com/smartdev0328/bluechip/x/pageinflation/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set if defined
	if genState.MintDenom != nil {
		k.SetMintDenom(ctx, *genState.MintDenom)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
	newMintDenom := types.MintDenom{
		Value: "ubluechip",
	}
	k.SetMintDenom(ctx, newMintDenom)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// Get all mintDenom
	mintDenom, found := k.GetMintDenom(ctx)
	if found {
		genesis.MintDenom = &mintDenom
	}
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
