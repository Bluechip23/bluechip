package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/smartdev0328/bluechip/x/pageinflation/types"
)

func (k msgServer) CreatorPoolMint(goCtx context.Context, msg *types.MsgCreatorPoolMint) (*types.MsgCreatorPoolMintResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	currentBlock := uint64(ctx.BlockHeight())

	storeMintDenom, isFound := k.GetMintDenom(ctx)
	mintDenom := ""
	if isFound {
		mintDenom = storeMintDenom.Value
	} else {
		mintDenom = "ubluechip"
	}
	///
	storeStartBlock, isFound := k.GetStartBlock(ctx)
	startBlock := uint64(0)
	if isFound {
		startBlock, _ = strconv.ParseUint(storeStartBlock.Value, 0, 64)
	} else {
		startBlock = currentBlock
	}
	mintedAmount := sdk.NewDec(4500000000000000).Quo(k.bankKeeper.GetSupply(ctx, mintDenom).Amount.ToDec().Quo(sdk.NewDec(100))).Mul(sdk.NewDec(1000000))
	mintedCoin := sdk.NewCoin(mintDenom, mintedAmount.TruncateInt())
	coins := sdk.NewCoins(mintedCoin)

	if coins.Empty() {
		panic("err")
	}

	err := k.bankKeeper.MintCoins(ctx, types.ModuleName, coins)

	if err != nil {
		panic(err)
	}

	recipient, _ := sdk.AccAddressFromBech32(msg.Creator)
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, recipient, coins)

	if err != nil {
		panic(err)
	}
	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreatorPoolMintResponse{}, nil
}
