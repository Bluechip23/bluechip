package keeper_test

import (
	"testing"

	testkeeper "github.com/smartdev0328/bluechip/testutil/keeper"
	"github.com/smartdev0328/bluechip/x/pageinflation/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.PageinflationKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
