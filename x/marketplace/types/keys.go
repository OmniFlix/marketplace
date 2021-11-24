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
