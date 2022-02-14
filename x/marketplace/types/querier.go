package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	QueryParams          = "params"
	QueryListing         = "listing"
	QueryAllListings     = "listings"
	QueryListingsByOwner = "listings-by-owner"
)

// QueryListingParams is the query parameters for '/marketplace/listings/{id}'
type QueryListingParams struct {
	Id string
}

// NewQueryListingParams
func NewQueryListingParams(id string) QueryListingParams {
	return QueryListingParams{
		Id: id,
	}
}

// QueryAllListingsParams is the query parameters for 'marketplace/listings'
type QueryAllListingsParams struct {
}

// NewQueryAllListingsParams
func NewQueryAllListingsParams() QueryAllListingsParams {
	return QueryAllListingsParams{}
}

// QueryListingsByOwnerParams is the query parameters for 'marketplace/listings/{owner}'
type QueryListingsByOwnerParams struct {
	Owner sdk.AccAddress
}

// NewQueryListingsByOwnerParams
func NewQueryListingsByOwnerParams(owner sdk.AccAddress) QueryListingsByOwnerParams {
	return QueryListingsByOwnerParams{
		Owner: owner,
	}
}
