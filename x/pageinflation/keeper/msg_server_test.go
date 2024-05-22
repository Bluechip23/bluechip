package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/BlueChip23/bluechip/testutil/keeper"
	"github.com/BlueChip23/bluechip/x/pageinflation/keeper"
	"github.com/BlueChip23/bluechip/x/pageinflation/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.PageinflationKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
