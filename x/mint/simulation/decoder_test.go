package simulation_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"cosmossdk.io/math"
	"github.com/BlueChip23/bluechip/app"
	"github.com/BlueChip23/bluechip/x/mint/simulation"
	"github.com/BlueChip23/bluechip/x/mint/types"
	"github.com/cosmos/cosmos-sdk/types/kv"
)

func TestDecodeStore(t *testing.T) {
	cdc := app.MakeEncodingConfig().Marshaler
	dec := simulation.NewDecodeStore(cdc)

	minter := types.NewMinter(math.NewDecFromInt64(math.OneInt().Int64()), math.NewDecFromInt64(15), 1, 1, math.OneInt())

	kvPairs := kv.Pairs{
		Pairs: []kv.Pair{
			{Key: types.MinterKey, Value: cdc.MustMarshal(&minter)},
			{Key: []byte{0x99}, Value: []byte{0x99}},
		},
	}
	tests := []struct {
		name        string
		expectedLog string
	}{
		{"Minter", fmt.Sprintf("%v\n%v", minter, minter)},
		{"other", ""},
	}

	for i, tt := range tests {
		i, tt := i, tt
		t.Run(tt.name, func(t *testing.T) {
			switch i {
			case len(tests) - 1:
				require.Panics(t, func() { dec(kvPairs.Pairs[i], kvPairs.Pairs[i]) }, tt.name)
			default:
				require.Equal(t, tt.expectedLog, dec(kvPairs.Pairs[i], kvPairs.Pairs[i]), tt.name)
			}
		})
	}
}
