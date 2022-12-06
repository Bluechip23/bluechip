package keeper

import (
	"github.com/Smartdev0328/bluechip/x/pageinflation/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetStartBlock set startBlock in the store
func (k Keeper) SetStartBlock(ctx sdk.Context, startBlock types.StartBlock) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StartBlockKey))
	b := k.cdc.MustMarshal(&startBlock)
	store.Set([]byte{0}, b)
}

// GetStartBlock returns startBlock
func (k Keeper) GetStartBlock(ctx sdk.Context) (val types.StartBlock, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StartBlockKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveStartBlock removes startBlock from the store
func (k Keeper) RemoveStartBlock(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StartBlockKey))
	store.Delete([]byte{0})
}
