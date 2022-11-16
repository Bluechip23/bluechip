package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/smartdev0328/bluechip/testutil/keeper"
	"github.com/smartdev0328/bluechip/x/pageinflation/keeper"
	"github.com/smartdev0328/bluechip/x/pageinflation/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.PageinflationKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
