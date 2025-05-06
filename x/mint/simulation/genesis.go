package simulation

// DONTCOVER

import (
	"encoding/json"
	"fmt"
	"math/rand"

	"cosmossdk.io/math"
	"github.com/BlueChip23/bluechip/x/mint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
)

// Simulation parameter constants
const (
	Inflation = "inflation"
)

// GenInflation randomized Inflation
func GenInflation(r *rand.Rand) math.Dec {
	return math.NewDecFromInt64(int64(r.Intn(99)))
}

// GenBlocksPerYear randomized BlocksPerYear
func GenBlocksPerYear(r *rand.Rand) math.Dec {
	return math.NewDecFromInt64(5256000)
}

// RandomizedGenState generates a random GenesisState for mint
func RandomizedGenState(simState *module.SimulationState) {
	// minter
	var inflation math.Dec
	simState.AppParams.GetOrGenerate(
		Inflation, &inflation, simState.Rand,
		func(r *rand.Rand) { inflation = GenInflation(r) },
	)

	// params
	mintDenom := sdk.DefaultBondDenom
	blocksPerYear := uint64(5256000)
	params := types.NewParams(mintDenom, blocksPerYear)

	mintGenesis := types.NewGenesisState(types.InitialMinter(inflation), params)

	bz, err := json.MarshalIndent(&mintGenesis, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Selected randomly generated minting parameters:\n%s\n", bz)
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(mintGenesis)
}
