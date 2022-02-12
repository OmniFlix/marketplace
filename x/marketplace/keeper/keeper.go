package keeper

import (
	"fmt"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/OmniFlix/marketplace/x/marketplace/types"
)

type Keeper struct {
	storeKey sdk.StoreKey
	cdc      codec.BinaryCodec

	accountKeeper      types.AccountKeeper
	bankKeeper         types.BankKeeper
	nftKeeper          types.NftKeeper
	distributionKeeper types.DistributionKeeper
	paramSpace         paramstypes.Subspace
}

func NewKeeper(
	cdc codec.BinaryCodec,
	key sdk.StoreKey,

	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
	nftKeeper types.NftKeeper,
	distrKeeper types.DistributionKeeper,
	paramSpace paramstypes.Subspace,
) Keeper {
	// ensure marketplace module account is set
	if addr := accountKeeper.GetModuleAddress(types.ModuleName); addr == nil {
		panic(fmt.Sprintf("%s module account has not been set", types.ModuleName))
	}

	// set KeyTable if it has not already been set
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		storeKey:           key,
		cdc:                cdc,
		accountKeeper:      accountKeeper,
		bankKeeper:         bankKeeper,
		nftKeeper:          nftKeeper,
		distributionKeeper: distrKeeper,
		paramSpace:         paramSpace,
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

	err := k.nftKeeper.TransferOwnership(ctx,
		listing.GetDenomId(), listing.GetNftId(), listing.GetOwner(),
		k.accountKeeper.GetModuleAddress(types.ModuleName))

	if err != nil {
		return err
	}
	// set listing
	k.SetListing(ctx, listing)
	if len(listing.GetOwner()) != 0 {
		// set listing id with owner prefix
		k.SetWithOwner(ctx, listing.GetOwner(), listing.GetId())
	}
	// Update listing count
	count := k.GetListingCount(ctx)
	k.SetListingCount(ctx, count+1)
	k.SetWithNFTID(ctx, listing.NftId, listing.Id)

	if len(listing.Price.Denom) > 0 {
		k.SetWithPriceDenom(ctx, listing.Price.Denom, listing.Id)
	}
	return nil
}

func (k Keeper) DeleteListing(ctx sdk.Context, listing types.Listing) {
	k.RemoveListing(ctx, listing.GetId())
	k.UnsetWithOwner(ctx, listing.GetOwner(), listing.GetId())
	k.UnsetWithNFTID(ctx, listing.GetNftId())
	k.UnsetWithPriceDenom(ctx, listing.Price.Denom, listing.GetId())
}

func (k Keeper) Buy(ctx sdk.Context, listing types.Listing, buyer sdk.AccAddress) error {
	owner, err := sdk.AccAddressFromBech32(listing.Owner)
	if err != nil {
		return err
	}

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, buyer, types.ModuleName, sdk.NewCoins(listing.Price))
	if err != nil {
		return err
	}
	err = k.nftKeeper.TransferOwnership(ctx, listing.GetDenomId(), listing.GetNftId(),
		k.accountKeeper.GetModuleAddress(types.ModuleName), buyer)
	if err != nil {
		_ = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, buyer, sdk.NewCoins(listing.Price))
		return err
	}
	saleCommission := k.GetSaleCommission(ctx)
	marketplaceCoin := k.GetProportions(ctx, listing.Price, saleCommission)
	listingSaleCoin := listing.Price.Sub(marketplaceCoin)
	err = k.distributionKeeper.FundCommunityPool(
		ctx,
		sdk.NewCoins(marketplaceCoin),
		k.accountKeeper.GetModuleAddress(types.ModuleName),
	)
	if err != nil {
		return err
	}

	if len(listing.SplitShares) > 0 {
		for _, share := range listing.SplitShares {
			sharePortionCoins := sdk.NewCoins(k.GetProportions(ctx, listingSaleCoin, share.Weight))
			if share.Address == "" {
				err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, owner, sharePortionCoins)
				if err != nil {
					return err
				}
			} else {
				saleSplitAddr, err := sdk.AccAddressFromBech32(share.Address)
				if err != nil {
					return err
				}
				err = k.bankKeeper.SendCoinsFromModuleToAccount(
					ctx, types.ModuleName, saleSplitAddr, sharePortionCoins)
				if err != nil {
					return err
				}
			}
		}
	} else {
		err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, owner, sdk.NewCoins(listingSaleCoin))
		if err != nil {
			return err
		}
	}

	k.DeleteListing(ctx, listing)
	return nil
}

func (k Keeper) GetProportions(ctx sdk.Context, totalCoin sdk.Coin, ratio sdk.Dec) sdk.Coin {
	return sdk.NewCoin(totalCoin.Denom, totalCoin.Amount.ToDec().Mul(ratio).TruncateInt())
}
