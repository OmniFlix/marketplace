package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	ModuleName               = "marketplace"
	StoreKey          string = ModuleName
	QuerierRoute      string = ModuleName
	RouterKey         string = ModuleName
	DefaultParamspace        = ModuleName
)

var (
	PrefixListingId         = []byte{0x01}
	PrefixListingOwner      = []byte{0x02}
	PrefixListingsCount     = []byte{0x03}
	PrefixListingNFTID      = []byte{0x04}
	PrefixListingPriceDenom = []byte{0x05}
	PrefixAuctionId         = []byte{0x06}
	PrefixAuctionOwner      = []byte{0x07}
	PrefixAuctionNFTID      = []byte{0x08}
	PrefixAuctionPriceDenom = []byte{0x09}
	PrefixNextAuctionNumber = []byte{0x10}
)

func KeyListingIdPrefix(id string) []byte {
	return append(PrefixListingId, []byte(id)...)
}

func KeyListingOwnerPrefix(owner sdk.AccAddress, id string) []byte {
	return append(append(PrefixListingOwner, owner.Bytes()...), []byte(id)...)
}

func KeyListingNFTIDPrefix(nftId string) []byte {
	return append(PrefixListingNFTID, []byte(nftId)...)
}

func KeyListingPriceDenomPrefix(priceDenom, id string) []byte {
	return append(append(PrefixListingPriceDenom, []byte(priceDenom)...), []byte(id)...)
}

func KeyAuctionIdPrefix(id string) []byte {
	return append(PrefixAuctionId, []byte(id)...)
}

func KeyAuctionOwnerPrefix(owner sdk.AccAddress, id string) []byte {
	return append(append(PrefixAuctionOwner, owner.Bytes()...), []byte(id)...)
}

func KeyAuctionNFTIDPrefix(nftId string) []byte {
	return append(PrefixAuctionNFTID, []byte(nftId)...)
}

func KeyAuctionPriceDenomPrefix(priceDenom, id string) []byte {
	return append(append(PrefixAuctionPriceDenom, []byte(priceDenom)...), []byte(id)...)
}
