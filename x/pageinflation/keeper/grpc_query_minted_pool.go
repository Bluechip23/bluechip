package keeper

import (
	"context"

	"github.com/Smartdev0328/bluechip/x/pageinflation/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) MintedPool(c context.Context, req *types.QueryGetMintedPoolRequest) (*types.QueryGetMintedPoolResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetMintedPool(ctx)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetMintedPoolResponse{MintedPool: val}, nil
}
