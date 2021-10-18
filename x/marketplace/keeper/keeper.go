package keeper

import (
	"fmt"
	
	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/OmniFlix/marketplace/x/marketplace/types"
)

type Keeper struct {
	storeKey      sdk.StoreKey
	cdc           codec.Codec
	accountKeeper types.AccountKeeper
	bankKeeper    types.BankKeeper
	nftKeeper     types.NftKeeper
}

func NewKeeper(
	cdc codec.Codec,
	key sdk.StoreKey,
	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
	nftKeeper types.NftKeeper,
) Keeper {
	return Keeper{
		storeKey:      key,
		cdc:           cdc,
		accountKeeper: accountKeeper,
		bankKeeper:    bankKeeper,
		nftKeeper:     nftKeeper,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("OmniFlix/%s", types.ModuleName))
}

func (k Keeper) ListNFT(ctx sdk.Context) {

}

func (k Keeper) DeListNFT(ctx sdk.Context) {

}

func (k Keeper) BuyNFT(ctx sdk.Context) {

}
