package exported

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"time"
)

type ListingI interface {
	GetId() string
	GetDenomId() string
	GetNftId() string
	GetPrice() sdk.Coin
	GetOwner() sdk.AccAddress
	GetSplitShares() interface{}
}

type AuctionListingI interface {
	GetId() uint64
	GetDenomId() string
	GetNftId() string
	GetStartPrice() sdk.Coin
	GetStartTime() time.Time
	GetIncrementPercentage() sdk.Dec
	GetOwner() sdk.AccAddress
	GetSplitShares() interface{}
}
