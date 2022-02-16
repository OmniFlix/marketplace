package types

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

func NewGenesisState(listings []Listing, listingCount uint64, params Params) *GenesisState {
	return &GenesisState{
		Listings:     listings,
		ListingCount: listingCount,
		Params:       params,
	}
}

func (m *GenesisState) ValidateGenesis() error {
	for _, l := range m.Listings {
		if l.GetOwner().Empty() {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing nft owner")
		}
		if err := ValidateListing(l); err != nil {
			return err
		}
	}
	return nil
}
