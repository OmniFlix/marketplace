package keeper

import (
	"github.com/OmniFlix/marketplace/x/marketplace/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetListing(ctx sdk.Context, id string) {

}

func (k Keeper) GetListings(ctx sdk.Context) []types.Listing {
	return []types.Listing{}
}
