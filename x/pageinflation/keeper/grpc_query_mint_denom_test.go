package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/smartdev0328/bluechip/testutil/keeper"
	"github.com/smartdev0328/bluechip/testutil/nullify"
	"github.com/smartdev0328/bluechip/x/pageinflation/types"
)

func TestMintDenomQuery(t *testing.T) {
	keeper, ctx := keepertest.PageinflationKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := createTestMintDenom(keeper, ctx)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetMintDenomRequest
		response *types.QueryGetMintDenomResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetMintDenomRequest{},
			response: &types.QueryGetMintDenomResponse{MintDenom: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.MintDenom(wctx, tc.request)
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
