package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewMinter returns a new Minter object with the given inflation and annual
// provisions values.
func NewMinter(inflation, annualProvisions sdk.Dec, phase, startPhaseBlock uint64, targetSupply sdk.Int) Minter {
	return Minter{
		Inflation:        inflation,
		AnnualProvisions: annualProvisions,
		Phase:            phase,
		StartPhaseBlock:  startPhaseBlock,
		TargetSupply:     targetSupply,
	}
}

// InitialMinter returns an initial Minter object with a given inflation value.
func InitialMinter(inflation sdk.Dec) Minter {
	return NewMinter(
		inflation,
		sdk.NewDec(0),
		0,
		0,
		sdk.NewInt(0),
	)
}

// DefaultInitialMinter returns a default initial Minter object for a new chain
// which uses an inflation rate of 13%.
func DefaultInitialMinter() Minter {
	return InitialMinter(
		sdk.NewDecWithPrec(13, 2),
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
func (m Minter) InflationcalculationFn(phase uint64) sdk.Dec {
	InflationAmt := sdk.NewDec(Inflation_Amount_Per_Year)

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
func (m Minter) NextAnnualProvisions(_ Params, totalSupply sdk.Int) sdk.Dec {
	return m.Inflation.MulInt(totalSupply)
}

// BlockProvision returns the provisions for a block based on the annual
// provisions rate.
func (m Minter) BlockProvision(params Params, totalSupply sdk.Int) sdk.Coin {
	provisionAmt := m.AnnualProvisions.QuoInt(sdk.NewInt(int64(params.BlocksPerYear)))

	// Because of rounding, we might mint too many tokens in this phase, let's limit it
	futureSupply := totalSupply.Add(provisionAmt.TruncateInt())
	if futureSupply.GT(m.TargetSupply) {
		return sdk.NewCoin(params.MintDenom, m.TargetSupply.Sub(totalSupply))
	}

	return sdk.NewCoin(params.MintDenom, provisionAmt.TruncateInt())
}
