package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/smartdev0328/bluechip/x/pageinflation/types"
)

// SetMintDenom set mintDenom in the store
func (k Keeper) SetMintDenom(ctx sdk.Context, mintDenom types.MintDenom) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MintDenomKey))
	b := k.cdc.MustMarshal(&mintDenom)
	store.Set([]byte{0}, b)
}

// GetMintDenom returns mintDenom
func (k Keeper) GetMintDenom(ctx sdk.Context) (val types.MintDenom, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MintDenomKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveMintDenom removes mintDenom from the store
func (k Keeper) RemoveMintDenom(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MintDenomKey))
	store.Delete([]byte{0})
}
