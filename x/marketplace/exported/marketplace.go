package exported

import sdk "github.com/cosmos/cosmos-sdk/types"

type ListingI interface {
	GetId() uint64
	GetNftDenomId() string
	GetNftId() string
	GetPrice() sdk.Coin
	GetOwner() sdk.AccAddress
}
