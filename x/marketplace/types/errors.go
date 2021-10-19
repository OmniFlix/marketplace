package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Listing module errors
var (
	ErrListingNotExists     = sdkerrors.Register(ModuleName, 2, "Listing does not exist")
	ErrInvalidOwner         = sdkerrors.Register(ModuleName, 3, "invalid Listing owner")
	ErrInvalidPrice         = sdkerrors.Register(ModuleName, 4, "invalid amount")
	ErrInvalidListing       = sdkerrors.Register(ModuleName, 5, "invalid Listing")
	ErrListingAlreadyExists = sdkerrors.Register(ModuleName, 6, "Listing already exists")
	ErrNotEnoughAmount      = sdkerrors.Register(ModuleName, 7, "amount is not enough to buy")
	ErrInvalidPriceDenom    = sdkerrors.Register(ModuleName, 8, "invalid price denom")
	ErrInvalidListingId     = sdkerrors.Register(ModuleName, 9, "invalid Listing id")
	ErrInvalidNftId         = sdkerrors.Register(ModuleName, 10, "invalid nft id")
	ErrNftNotExists         = sdkerrors.Register(ModuleName, 11, "nft not exists with given details")
	ErrUnauthorized         = sdkerrors.Register(ModuleName, 12, "unauthorized")
	ErrNftNonTransferable   = sdkerrors.Register(ModuleName, 13, "non-transferable nft")
	ErrListingDoesNotExists = sdkerrors.Register(ModuleName, 14, "listing doesn't exists")
)
