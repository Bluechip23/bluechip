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
	// Set if defined
	if genState.StartBlock != nil {
		k.SetStartBlock(ctx, *genState.StartBlock)
	}
	// Set if defined
	if genState.MintedPool != nil {
		k.SetMintedPool(ctx, *genState.MintedPool)
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
	// Get all startBlock
	startBlock, found := k.GetStartBlock(ctx)
	if found {
		genesis.StartBlock = &startBlock
	}
	// Get all mintedPool
	mintedPool, found := k.GetMintedPool(ctx)
	if found {
		genesis.MintedPool = &mintedPool
	}
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
