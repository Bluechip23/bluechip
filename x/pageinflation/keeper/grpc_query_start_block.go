package keeper

import (
	"context"

	"github.com/BlueChip23/bluechip/x/pageinflation/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) StartBlock(c context.Context, req *types.QueryGetStartBlockRequest) (*types.QueryGetStartBlockResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetStartBlock(ctx)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetStartBlockResponse{StartBlock: val}, nil
}
