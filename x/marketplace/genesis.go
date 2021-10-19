package market

import (
	"github.com/OmniFlix/marketplace/x/marketplace/keeper"
	"github.com/OmniFlix/marketplace/x/marketplace/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	if err := types.ValidateGenesis(genState); err != nil {
		panic(err.Error())
	}
	for _, l := range genState.Listings {
		k.SetListing(ctx, l)
	}
	k.SetListingCount(ctx, uint64(len(genState.Listings)))
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return types.NewGenesisState(k.GetAllListings(ctx))
}

func DefaultGenesisState() *types.GenesisState {
	return types.NewGenesisState([]types.Listing{})
}
