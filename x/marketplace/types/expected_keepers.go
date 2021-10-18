package types

import (
	nft "github.com/OmniFlix/onft/exported"
	nftypes "github.com/OmniFlix/onft/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

type AccountKeeper interface {
	// Methods imported from account should be defined here
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) authtypes.AccountI
}

type BankKeeper interface {
	// Methods imported from bank should be defined here
	GetAllBalances(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
	LockedCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	SendCoins(ctx sdk.Context, from sdk.AccAddress, to sdk.AccAddress, amount sdk.Coins) error
}

type NftKeeper interface {
	// methods imported from nft should be defined here
	GetONFT(ctx sdk.Context, denomID, onftID string) (nft nft.ONFT, err error)
	GetDenom(ctx sdk.Context, denomId string) (nftypes.Denom, error)
	TransferOwnership(ctx sdk.Context, denomId, nftId string, srcOwner, dstOwner sdk.AccAddress) error
}
