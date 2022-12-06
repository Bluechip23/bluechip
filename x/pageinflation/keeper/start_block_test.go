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

func createTestStartBlock(keeper *keeper.Keeper, ctx sdk.Context) types.StartBlock {
	item := types.StartBlock{}
	keeper.SetStartBlock(ctx, item)
	return item
}

func TestStartBlockGet(t *testing.T) {
	keeper, ctx := keepertest.PageinflationKeeper(t)
	item := createTestStartBlock(keeper, ctx)
	rst, found := keeper.GetStartBlock(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestStartBlockRemove(t *testing.T) {
	keeper, ctx := keepertest.PageinflationKeeper(t)
	createTestStartBlock(keeper, ctx)
	keeper.RemoveStartBlock(ctx)
	_, found := keeper.GetStartBlock(ctx)
	require.False(t, found)
}
