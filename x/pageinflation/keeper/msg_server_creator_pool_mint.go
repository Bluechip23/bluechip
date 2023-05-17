package keeper

import (
	"context"
	"math"

	"github.com/Smartdev0328/bluechip/x/pageinflation/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreatorPoolMint(goCtx context.Context, msg *types.MsgCreatorPoolMint) (*types.MsgCreatorPoolMintResponse, error) {

	///cg
	ctx := sdk.UnwrapSDKContext(goCtx)

	recipient, _ := sdk.AccAddressFromBech32(msg.Creator)
	if k.accountKeeper.GetAccount(ctx, recipient).GetAccountNumber() != 1 {
		panic("Not Backend account")
	}
	storeMintDenom, _ := k.GetMintDenom(ctx)
	mintDenom := storeMintDenom.Value
	///

	currentBlock := ctx.BlockHeight()
	storeStartBlock, _ := k.GetStartBlock(ctx)
	startBlock := storeStartBlock.Value
	if startBlock == 0 {
		newStartBlock := types.StartBlock{
			Value: uint64(currentBlock),
		}
		k.SetStartBlock(ctx, newStartBlock)
		startBlock = uint64(currentBlock)
	}
	///
	storeMintedPool, _ := k.GetMintedPool(ctx)
	mintedPool := storeMintedPool.Value

	///calculate the next mintedAmount
	numerator := float64(0.0064179)*math.Pow(float64(mintedPool), float64(2)) + float64(mintedPool)
	denominator := math.Pow(float64(uint64(currentBlock)-startBlock), float64(1)/float64(1.569)) +
		1200*math.Pow(float64(1.012), float64(mintedPool)/float64(-25))
	mintedAmount := sdk.NewDec(int64(1200000000 - uint64(numerator*1000000/denominator)))
	mintedCoin := sdk.NewCoin(mintDenom, mintedAmount.TruncateInt())
	coins := sdk.NewCoins(mintedCoin)

	if coins.Empty() {
		panic("err")
	}

	err := k.bankKeeper.MintCoins(ctx, types.ModuleName, coins)

	if err != nil {
		panic(err)
	}
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, recipient, coins)

	newMintedPool := types.MintedPool{
		Value: mintedPool + 1,
	}
	k.SetMintedPool(ctx, newMintedPool)

	if err != nil {
		panic(err)
	}
	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreatorPoolMintResponse{}, nil
}
