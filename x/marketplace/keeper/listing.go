package keeper

import (
	"encoding/binary"
	"github.com/OmniFlix/marketplace/x/marketplace/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetListingCount get the total number of listings
func (k Keeper) GetListingCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.PrefixListingsCount
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetListingCount set the total number of listings
func (k Keeper) SetListingCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.PrefixListingsCount
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendListing appends a listing in the store with a new id and update the count
func (k Keeper) AppendListing(
	ctx sdk.Context,
	listing types.Listing,
) uint64 {
	// Create the listing
	count := k.GetListingCount(ctx)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PrefixListingId)
	appendedValue := k.cdc.MustMarshal(&listing)
	store.Set(types.KeyListingIdPrefix(listing.Id), appendedValue)

	// Update listing count
	k.SetListingCount(ctx, count+1)

	return count
}

// SetListing set a specific listing in the store
func (k Keeper) SetListing(ctx sdk.Context, listing types.Listing) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PrefixListingId)
	b := k.cdc.MustMarshal(&listing)
	store.Set(types.KeyListingIdPrefix(listing.Id), b)
}

// GetListing returns a listing from its id
func (k Keeper) GetListing(ctx sdk.Context, id string) (val types.Listing, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PrefixListingId)
	b := store.Get(types.KeyListingIdPrefix(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveListing removes a listing from the store
func (k Keeper) RemoveListing(ctx sdk.Context, id string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PrefixListingId)
	store.Delete(types.KeyListingIdPrefix(id))
}

// GetAllListings returns all listings
func (k Keeper) GetAllListings(ctx sdk.Context) (list []types.Listing) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PrefixListingId)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Listing
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
