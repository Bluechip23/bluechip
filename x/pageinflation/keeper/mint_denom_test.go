package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/Smartdev0328/bluechip/testutil/keeper"
	"github.com/Smartdev0328/bluechip/testutil/nullify"
	"github.com/Smartdev0328/bluechip/x/pageinflation/keeper"
	"github.com/Smartdev0328/bluechip/x/pageinflation/types"
)

func createTestMintDenom(keeper *keeper.Keeper, ctx sdk.Context) types.MintDenom {
	item := types.MintDenom{}
	keeper.SetMintDenom(ctx, item)
	return item
}

func TestMintDenomGet(t *testing.T) {
	keeper, ctx := keepertest.PageinflationKeeper(t)
	item := createTestMintDenom(keeper, ctx)
	rst, found := keeper.GetMintDenom(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestMintDenomRemove(t *testing.T) {
	keeper, ctx := keepertest.PageinflationKeeper(t)
	createTestMintDenom(keeper, ctx)
	keeper.RemoveMintDenom(ctx)
	_, found := keeper.GetMintDenom(ctx)
	require.False(t, found)
}
