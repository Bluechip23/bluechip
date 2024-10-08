package cli_test

import (
	"fmt"
	"testing"

	tmcli "github.com/cometbft/cometbft/libs/cli"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/status"

	"github.com/BlueChip23/bluechip/testutil/network"
	"github.com/BlueChip23/bluechip/testutil/nullify"
	"github.com/BlueChip23/bluechip/x/pageinflation/client/cli"
	"github.com/BlueChip23/bluechip/x/pageinflation/types"
)

func networkWithMintDenomObjects(t *testing.T) (*network.Network, types.MintDenom) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &state))

	mintDenom := &types.MintDenom{}
	nullify.Fill(&mintDenom)
	state.MintDenom = mintDenom
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg), *state.MintDenom
}

func TestShowMintDenom(t *testing.T) {
	net, obj := networkWithMintDenomObjects(t)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc string
		args []string
		err  error
		obj  types.MintDenom
	}{
		{
			desc: "get",
			args: common,
			obj:  obj,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			var args []string
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowMintDenom(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetMintDenomResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.MintDenom)
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&resp.MintDenom),
				)
			}
		})
	}
}
