package keeper

import (
	"fmt"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/OmniFlix/marketplace/x/marketplace/types"
)

type Keeper struct {
	storeKey      sdk.StoreKey
	cdc           codec.Codec
	accountKeeper types.AccountKeeper
	bankKeeper    types.BankKeeper
	nftKeeper     types.NftKeeper
}

func NewKeeper(
	cdc codec.Codec,
	key sdk.StoreKey,
	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
	nftKeeper types.NftKeeper,
) Keeper {
	return Keeper{
		storeKey:      key,
		cdc:           cdc,
		accountKeeper: accountKeeper,
		bankKeeper:    bankKeeper,
		nftKeeper:     nftKeeper,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("OmniFlix/%s", types.ModuleName))
}

// AddListing adds a listing in the store and set owner to listing and updates the count
func (k Keeper) AddListing(ctx sdk.Context, listing types.Listing) error {
	// check listing already exists
	if k.HasListing(ctx, listing.GetId()) {
		return sdkerrors.Wrapf(types.ErrListingAlreadyExists, "listing already exists: %s", listing.GetId())
	}
	// set listing
	k.SetListing(ctx, listing)
	if len(listing.GetOwner()) != 0 {
		// set listing id with owner prefix
		k.setWithOwner(ctx, listing.GetOwner(), listing.GetId())
	}

	// Update listing count
	count := k.GetListingCount(ctx)
	k.SetListingCount(ctx, count+1)

	return nil
}

func (k Keeper) DeleteListing(ctx sdk.Context, listing types.Listing) {
	k.RemoveListing(ctx, listing.GetId())
	k.unsetWithOwner(ctx, listing.GetOwner(), listing.GetId())
}

func (k Keeper) Buy(ctx sdk.Context, listing types.Listing, buyer sdk.AccAddress) error {
	owner, err := sdk.AccAddressFromBech32(listing.Owner)
	if err != nil {
		return err
	}

	err = k.bankKeeper.SendCoins(ctx, buyer, owner, sdk.NewCoins(listing.Price))
	if err != nil {
		return err
	}
	err = k.nftKeeper.TransferOwnership(ctx, listing.DenomId, listing.NftId, owner, buyer)
	if err != nil {
		return err
	}
	k.DeleteListing(ctx, listing)

	return nil
}