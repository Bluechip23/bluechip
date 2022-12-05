package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/smartdev0328/bluechip/x/pageinflation/types"
)

// SetMintedPool set mintedPool in the store
func (k Keeper) SetMintedPool(ctx sdk.Context, mintedPool types.MintedPool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MintedPoolKey))
	b := k.cdc.MustMarshal(&mintedPool)
	store.Set([]byte{0}, b)
}

// GetMintedPool returns mintedPool
func (k Keeper) GetMintedPool(ctx sdk.Context) (val types.MintedPool, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MintedPoolKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveMintedPool removes mintedPool from the store
func (k Keeper) RemoveMintedPool(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MintedPoolKey))
	store.Delete([]byte{0})
}
