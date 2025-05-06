package mint

import (
	"time"

	"cosmossdk.io/math"
	"github.com/BlueChip23/bluechip/x/mint/keeper"
	"github.com/BlueChip23/bluechip/x/mint/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// BeginBlocker mints new tokens for the previous block.
func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	// fetch stored minter
	minter := k.GetMinter(ctx)

	// inflation phase end
	// if minter.Inflation.Equal(sdk.ZeroDec()) {
	// 	return
	// }

	// fetch stored params
	params := k.GetParams(ctx)
	currentBlock := uint64(ctx.BlockHeight())

	// fetch current total supply
	totalSupply := k.TokenSupply(ctx, params.MintDenom)

	// check if we need to change phase
	nextPhase := minter.NextPhase(params, currentBlock)

	if nextPhase != minter.Phase {
		// store new inflation rate by phase
		newInflation := minter.InflationcalculationFn(nextPhase)
		//minter.Inflation = newInflation
		minter.Phase = nextPhase
		minter.StartPhaseBlock = currentBlock
		v, _ := newInflation.Mul(math.NewDecFromInt64(100))
		z, _ := v.Quo(math.NewDecFromInt64(totalSupply.Int64()))
		minter.Inflation = z
		minter.AnnualProvisions = newInflation
		annualProvisions, _ := minter.AnnualProvisions.Int64()
		minter.TargetSupply = totalSupply.Add(math.NewInt(annualProvisions))
		k.SetMinter(ctx, minter)

		// inflation phase end
		// if minter.Inflation.Equal(sdk.ZeroDec()) {
		// 	return
		// }
	}

	// mint coins, update supply
	mintedCoin := minter.BlockProvision(params, totalSupply)
	mintedCoins := sdk.NewCoins(mintedCoin)

	err := k.MintCoins(ctx, mintedCoins)
	if err != nil {
		panic(err)
	}

	// send the minted coins to the fee collector account
	err = k.AddCollectedFees(ctx, mintedCoins)
	if err != nil {
		panic(err)
	}

	if mintedCoin.Amount.IsInt64() {
		defer telemetry.ModuleSetGauge(types.ModuleName, float32(mintedCoin.Amount.Int64()), "minted_tokens")
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeMint,
			sdk.NewAttribute(types.AttributeKeyInflation, minter.Inflation.String()),
			sdk.NewAttribute(types.AttributeKeyAnnualProvisions, minter.AnnualProvisions.String()),
			sdk.NewAttribute(sdk.AttributeKeyAmount, mintedCoin.Amount.String()),
		),
	)
}
