package market

import (
	"github.com/OmniFlix/marketplace/x/marketplace/keeper"
	"github.com/OmniFlix/marketplace/x/marketplace/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {

}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return types.NewGenesisState(k.GetListings(ctx))
}

func DefaultGenesisState() *types.GenesisState {
	return types.NewGenesisState([]types.Listing{})
}
