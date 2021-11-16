package keeper

import (
	"context"
	"fmt"

	"github.com/OmniFlix/marketplace/x/marketplace/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	gogotypes "github.com/gogo/protobuf/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) Listing(goCtx context.Context, req *types.QueryListingRequest) (*types.QueryListingResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	listing, found := k.GetListing(ctx, req.Id)
	if !found {
		return nil, status.Errorf(codes.NotFound, "listing %d not found", req.Id)
	}

	return &types.QueryListingResponse{Listing: &listing}, nil
}

func (k Keeper) Listings(goCtx context.Context, req *types.QueryListingsRequest) (*types.QueryListingsResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	var listings []types.Listing
	var pageRes *query.PageResponse
	store := ctx.KVStore(k.storeKey)

	var owner sdk.AccAddress
	var err error
	if len(req.Owner) > 0 {
		owner, err = sdk.AccAddressFromBech32(req.Owner)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("invalid owner address (%s)", err))
		}
		listingStore := prefix.NewStore(store, append(types.PrefixListingOwner, owner.Bytes()...))
		pageRes, err = query.Paginate(listingStore, req.Pagination, func(key []byte, value []byte) error {
			var listingId gogotypes.StringValue
			k.cdc.MustUnmarshal(value, &listingId)
			listing, found := k.GetListing(ctx, listingId.Value)
			if found {
				listings = append(listings, listing)
			}
			return nil
		})

	} else {
		listingStore := prefix.NewStore(store, types.PrefixListingId)
		pageRes, err = query.Paginate(listingStore, req.Pagination, func(key []byte, value []byte) error {
			var listing types.Listing
			k.cdc.MustUnmarshal(value, &listing)
			listings = append(listings, listing)
			return nil
		})
	}
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "paginate: %v", err)
	}

	return &types.QueryListingsResponse{Listings: listings, Pagination: pageRes}, nil
}

func (k Keeper) ListingsByOwner(goCtx context.Context, req *types.QueryListingsByOwnerRequest) (*types.QueryListingsByOwnerResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	var owner sdk.AccAddress
	var err error
	if len(req.Owner) > 0 {
		owner, err = sdk.AccAddressFromBech32(req.Owner)
		if err != nil || owner == nil {
			return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("invalid owner address (%s)", err))
		}
	}

	var listings []types.Listing
	var pageRes *query.PageResponse
	store := ctx.KVStore(k.storeKey)

	listingStore := prefix.NewStore(store, append(types.PrefixListingOwner, owner.Bytes()...))
	pageRes, err = query.Paginate(listingStore, req.Pagination, func(key []byte, value []byte) error {
		var listingId gogotypes.StringValue
		k.cdc.MustUnmarshal(value, &listingId)
		listing, found := k.GetListing(ctx, listingId.Value)
		if found {
			listings = append(listings, listing)
		}
		return nil
	})
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "paginate: %v", err)
	}

	return &types.QueryListingsByOwnerResponse{Listings: listings, Pagination: pageRes}, nil
}
