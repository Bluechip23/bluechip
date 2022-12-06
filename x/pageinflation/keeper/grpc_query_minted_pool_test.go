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

func TestMintedPoolQuery(t *testing.T) {
	keeper, ctx := keepertest.PageinflationKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := createTestMintedPool(keeper, ctx)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetMintedPoolRequest
		response *types.QueryGetMintedPoolResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetMintedPoolRequest{},
			response: &types.QueryGetMintedPoolResponse{MintedPool: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.MintedPool(wctx, tc.request)
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
