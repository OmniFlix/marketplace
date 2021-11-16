package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strings"
)

var (
	allowedDenoms = []string{}
)
// ValidateListing checks listing is valid or not
func ValidateListing(listing Listing) error {
	if len(listing.Owner) > 0 {
		if _, err := sdk.AccAddressFromBech32(listing.Owner); err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
		}
	}
	if err := ValidateId(listing.Id); err != nil {
		return err
	}
	if err := ValidatePrice(listing.Price); err != nil {
		return err
	}
	return nil
}

// ValidatePrice
func ValidatePrice(price sdk.Coin) error {
	if price.IsZero() {
		return sdkerrors.Wrapf(ErrInvalidPrice, "invalid price %d, only accepts positive amount", price)
	}
        /*
	if !StringInSlice(price.Denom, allowedDenoms) {
		return sdkerrors.Wrapf(ErrInvalidPriceDenom, "invalid denom %s", price.Denom)
	}
        */
	return nil
}

func ValidateId(id string) error {
	id = strings.TrimSpace(id)
	if len(id) < MinListingIdLength || len(id) > MaxListingIdLength {

		return sdkerrors.Wrapf(
			ErrInvalidListingId,
			"invalid id %s, only accepts value [%d, %d]", id, MinListingIdLength, MaxListingIdLength,
			)
	}
	if !IsBeginWithAlpha(id) || !IsAlphaNumeric(id) {
		return sdkerrors.Wrapf(ErrInvalidListingId, "invalid id %s, only accepts alphanumeric characters,and begin with an english letter", id)
	}
	return nil
}
