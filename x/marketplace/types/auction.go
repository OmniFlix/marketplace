package types

import (
	"github.com/OmniFlix/marketplace/x/marketplace/exported"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
	"time"
)

var (
	_ proto.Message     = &AuctionListing{}
	_ exported.AuctionListingI = &AuctionListing{}
)

func NewAuctionListing(id uint64, nftId, denomId string, start_time *time.Time, start_price sdk.Coin, increment_percentage sdk.Dec,
	owner sdk.AccAddress, splitShares []WeightedAddress) AuctionListing {
	return AuctionListing{
		Id: id,
		NftId:       nftId,
		DenomId:     denomId,
		StartTime: start_time,
		StartPrice:  start_price,
		IncrementPercentage: increment_percentage,
		Owner:       owner.String(),
		SplitShares: splitShares,
	}
}

func (al AuctionListing) GetId() uint64 {
	return al.Id
}

func (al AuctionListing) GetDenomId() string {
	return al.DenomId
}

func (al AuctionListing) GetNftId() string {
	return al.NftId
}
func (al AuctionListing) GetStartTime() time.Time {
	return *al.StartTime
}

func (al AuctionListing) GetStartPrice() sdk.Coin {
	return al.StartPrice
}

func (al AuctionListing) GetIncrementPercentage() sdk.Dec {
	return al.IncrementPercentage
}

func (al AuctionListing) GetOwner() sdk.AccAddress {
	owner, _ := sdk.AccAddressFromBech32(al.Owner)
	return owner
}

func (al AuctionListing) GetSplitShares() interface{} {
	return al.SplitShares
}
