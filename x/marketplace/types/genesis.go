package types

func NewGenesisState(listings []Listing) *GenesisState {
	return &GenesisState{
		Listings: listings,
	}
}
