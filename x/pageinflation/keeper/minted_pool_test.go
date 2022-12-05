package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/smartdev0328/bluechip/testutil/keeper"
	"github.com/smartdev0328/bluechip/testutil/nullify"
	"github.com/smartdev0328/bluechip/x/pageinflation/keeper"
	"github.com/smartdev0328/bluechip/x/pageinflation/types"
)

func createTestMintedPool(keeper *keeper.Keeper, ctx sdk.Context) types.MintedPool {
	item := types.MintedPool{}
	keeper.SetMintedPool(ctx, item)
	return item
}

func TestMintedPoolGet(t *testing.T) {
	keeper, ctx := keepertest.PageinflationKeeper(t)
	item := createTestMintedPool(keeper, ctx)
	rst, found := keeper.GetMintedPool(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestMintedPoolRemove(t *testing.T) {
	keeper, ctx := keepertest.PageinflationKeeper(t)
	createTestMintedPool(keeper, ctx)
	keeper.RemoveMintedPool(ctx)
	_, found := keeper.GetMintedPool(ctx)
	require.False(t, found)
}
