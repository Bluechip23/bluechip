package simulation

import (
	"math/rand"

	"github.com/BlueChip23/bluechip/x/pageinflation/keeper"
	"github.com/BlueChip23/bluechip/x/pageinflation/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgCreatorPoolMint(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgCreatorPoolMint{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the CreatorPoolMint simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "CreatorPoolMint simulation not implemented"), nil, nil
	}
}
