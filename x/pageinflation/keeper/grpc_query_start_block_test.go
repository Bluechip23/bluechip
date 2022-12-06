package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/Smartdev0328/bluechip/testutil/keeper"
	"github.com/Smartdev0328/bluechip/testutil/nullify"
	"github.com/Smartdev0328/bluechip/x/pageinflation/types"
)

func TestStartBlockQuery(t *testing.T) {
	keeper, ctx := keepertest.PageinflationKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := createTestStartBlock(keeper, ctx)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetStartBlockRequest
		response *types.QueryGetStartBlockResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetStartBlockRequest{},
			response: &types.QueryGetStartBlockResponse{StartBlock: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.StartBlock(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}
