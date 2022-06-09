package keeper

import (
	"fmt"
	"github.com/OmniFlix/marketplace/x/marketplace/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	gogotypes "github.com/gogo/protobuf/types"
)

// GetNextAuctionNumber get the next auction number
func (k Keeper) GetNextAuctionNumber(ctx sdk.Context) uint64 {
	var nextAuctionNumber uint64
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(types.PrefixNextAuctionNumber)
	if bz == nil {
		panic(fmt.Errorf("auction module not initialized -- Should have been done in InitGenesis"))
	} else {
		val := gogotypes.UInt64Value{}

		err := k.cdc.Unmarshal(bz, &val)
		if err != nil {
			panic(err)
		}

		nextAuctionNumber = val.GetValue()
	}
	return nextAuctionNumber
}

// SetNextAuctionNumber set the next auction number
func (k Keeper) SetNextAuctionNumber(ctx sdk.Context, number uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&gogotypes.UInt64Value{Value: number})
	store.Set(types.PrefixNextAuctionNumber, bz)
}

// SetAuctionListing set a specific auction listing in the store
func (k Keeper) SetAuctionListing(ctx sdk.Context, auctionListing types.AuctionListing) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PrefixAuctionId)
	bz := k.cdc.MustMarshal(&auctionListing)
	store.Set(types.KeyAuctionIdPrefix(auctionListing.Id), bz)
}

// GetAuctionListing returns a auction listing by its id
func (k Keeper) GetAuctionListing(ctx sdk.Context, id uint64) (val types.AuctionListing, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PrefixAuctionId)
	b := store.Get(types.KeyAuctionIdPrefix(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetListing returns a listing from its nft id
func (k Keeper) GetAuctionListingIdByNftId(ctx sdk.Context, nftId string) (val uint64, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PrefixAuctionNFTID)
	bz := store.Get(types.KeyAuctionNFTIDPrefix(nftId))
	if bz == nil {
		return val, false
	}
	var auctionId gogotypes.UInt64Value
	k.cdc.MustUnmarshal(bz, &auctionId)
	return auctionId.Value, true
}

// RemoveAuctionListing removes a auction listing from the store
func (k Keeper) RemoveAuctionListing(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PrefixAuctionId)
	store.Delete(types.KeyAuctionIdPrefix(id))
}

// GetAllAuctionListings returns all auction listings
func (k Keeper) GetAllAuctionListings(ctx sdk.Context) (list []types.AuctionListing) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PrefixAuctionId)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.AuctionListing
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetAuctionListingsByOwner returns all auction listings of specific owner
func (k Keeper) GetAuctionListingsByOwner(ctx sdk.Context, owner sdk.AccAddress) (auctionListings []types.AuctionListing) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, append(types.PrefixAuctionOwner, owner.Bytes()...))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var id gogotypes.UInt64Value
		k.cdc.MustUnmarshal(iterator.Value(), &id)
		listing, found := k.GetAuctionListing(ctx, id.Value)
		if !found {
			continue
		}
		auctionListings = append(auctionListings, listing)
	}

	return
}

func (k Keeper) HasAuctionListing(ctx sdk.Context, id uint64) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.KeyAuctionIdPrefix(id))
}

func (k Keeper) SetAuctionListingWithOwner(ctx sdk.Context, owner sdk.AccAddress, id uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&gogotypes.UInt64Value{Value: id})

	store.Set(types.KeyAuctionOwnerPrefix(owner, id), bz)
}
func (k Keeper) UnsetAuctionListingWithOwner(ctx sdk.Context, owner sdk.AccAddress, id uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.KeyAuctionOwnerPrefix(owner, id))
}

func (k Keeper) SetAuctionListingWithNFTID(ctx sdk.Context, nftId string, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PrefixAuctionNFTID)
	bz := k.cdc.MustMarshal(&gogotypes.UInt64Value{Value: id})
	store.Set(types.KeyAuctionNFTIDPrefix(nftId), bz)
}

func (k Keeper) UnsetAuctionListingWithNFTID(ctx sdk.Context, nftId string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PrefixAuctionNFTID)
	store.Delete(types.KeyAuctionNFTIDPrefix(nftId))
}

func (k Keeper) SetAuctionListingWithPriceDenom(ctx sdk.Context, priceDenom string, id uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&gogotypes.UInt64Value{Value: id})

	store.Set(types.KeyAuctionPriceDenomPrefix(priceDenom, id), bz)
}

func (k Keeper) UnsetAuctionListingWithPriceDenom(ctx sdk.Context, priceDenom string, id uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.KeyAuctionPriceDenomPrefix(priceDenom, id))
}
