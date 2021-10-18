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
	ErrInvalidAmountDenom   = sdkerrors.Register(ModuleName, 8, "invalid amount denom")
)
