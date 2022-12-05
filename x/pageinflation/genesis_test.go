package pageinflation_test

import (
	"testing"

	keepertest "github.com/smartdev0328/bluechip/testutil/keeper"
	"github.com/smartdev0328/bluechip/testutil/nullify"
	"github.com/smartdev0328/bluechip/x/pageinflation"
	"github.com/smartdev0328/bluechip/x/pageinflation/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		MintDenom: &types.MintDenom{
			Value: "81",
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.PageinflationKeeper(t)
	pageinflation.InitGenesis(ctx, *k, genesisState)
	got := pageinflation.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.MintDenom, got.MintDenom)
	// this line is used by starport scaffolding # genesis/test/assert
}
