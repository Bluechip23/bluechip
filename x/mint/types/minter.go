package types

import (
	"fmt"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewMinter returns a new Minter object with the given inflation and annual
// provisions values.
func NewMinter(inflation, annualProvisions math.Dec, phase, startPhaseBlock uint64, targetSupply math.Int) Minter {
	return Minter{
		Inflation:        inflation,
		AnnualProvisions: annualProvisions,
		Phase:            phase,
		StartPhaseBlock:  startPhaseBlock,
		TargetSupply:     targetSupply,
	}
}

// InitialMinter returns an initial Minter object with a given inflation value.
func InitialMinter(inflation math.Dec) Minter {
	return NewMinter(
		inflation,
		math.NewDecFromInt64(0),
		0,
		0,
		math.NewInt(0),
	)
}

// DefaultInitialMinter returns a default initial Minter object for a new chain
// which uses an inflation rate of 13%.
func DefaultInitialMinter() Minter {
	return InitialMinter(
		// TODO: MIGHT NEED PRECISION LegacyNewDecFromBigIntWithPrec
		math.NewDecFromInt64(13),
	)
}

// validate minter
func ValidateMinter(minter Minter) error {
	if minter.Inflation.IsNegative() {
		return fmt.Errorf("mint parameter Inflation should be positive, is %s",
			minter.Inflation.String())
	}
	return nil
}

// PhaseInflationRate returns the inflation rate by phase.
func (m Minter) InflationcalculationFn(phase uint64) math.Dec {
	InflationAmt := math.NewDecFromInt64(Inflation_Amount_Per_Year)

	return InflationAmt
}

// NextPhase returns the new phase.
func (m Minter) NextPhase(params Params, currentBlock uint64) uint64 {
	nonePhase := m.Phase == 0
	if nonePhase {
		return 1
	}

	blockNewPhase := m.StartPhaseBlock + params.BlocksPerYear
	if blockNewPhase > currentBlock {
		return m.Phase
	}

	return m.Phase + 1
}

// NextAnnualProvisions returns the annual provisions based on current total
// supply and inflation rate.
func (m Minter) NextAnnualProvisions(_ Params, totalSupply math.Int) (math.Dec, error) {
	return m.Inflation.Mul(math.NewDecFromInt64(totalSupply.Int64()))
}

// BlockProvision returns the provisions for a block based on the annual
// provisions rate.
func (m Minter) BlockProvision(params Params, totalSupply math.Int) sdk.Coin {
	provisionAmt := math.NewDecFromInt64(1000000)
	// Because of rounding, we might mint too many tokens in this phase, let's limit it
	x, _ := provisionAmt.Int64()
	futureSupply := totalSupply.Add(math.NewInt(x))
	if futureSupply.GT(m.TargetSupply) {
		return sdk.NewCoin(params.MintDenom, m.TargetSupply.Sub(totalSupply))
	}

	return sdk.NewCoin(params.MintDenom, (math.NewInt(x)))
}
