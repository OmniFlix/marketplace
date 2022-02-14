package marketplace

import (
	"fmt"

	"github.com/OmniFlix/marketplace/x/marketplace/keeper"
	"github.com/OmniFlix/marketplace/x/marketplace/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	if err := genState.ValidateGenesis(); err != nil {
		panic(err.Error())
	}
	for _, l := range genState.Listings {
		k.SetListing(ctx, l)
		k.SetWithOwner(ctx, l.GetOwner(), l.GetId())
		k.SetWithNFTID(ctx, l.GetNftId(), l.GetId())
		k.SetWithPriceDenom(ctx, l.Price.GetDenom(), l.GetId())
	}
	k.SetListingCount(ctx, genState.ListingCount)
	k.SetParams(ctx, genState.Params)

	// check if the module account exists
	moduleAcc := k.GetMarketplaceAccount(ctx)
	if moduleAcc == nil {
		panic(fmt.Sprintf("%s module account has not been set", types.ModuleName))
	}
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return types.NewGenesisState(k.GetAllListings(ctx), k.GetListingCount(ctx), k.GetParams(ctx))
}

func DefaultGenesisState() *types.GenesisState {
	return types.NewGenesisState([]types.Listing{}, 0, types.Params{
		SaleCommission: sdk.NewDecWithPrec(1, 2), // "0.01" or 1%
		Distribution: types.Distribution{
			Staking: sdk.NewDecWithPrec(50, 2), // "0.50" or 50%
			CommunityPool: sdk.NewDecWithPrec(50, 2), // "0.50" or 50%
		},
	})
}
