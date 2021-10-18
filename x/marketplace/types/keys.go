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
	PrefixListingById     = []byte{0x01}
	PrefixListingsByOwner = []byte{0x02}
)

func KeyListingsById(id uint64) []byte {
	return append(PrefixListingById, sdk.Uint64ToBigEndian(id)...)
}

func KeyListingsByOwner(owner sdk.AccAddress, id uint64) []byte {
	return append(append(PrefixListingsByOwner, owner.Bytes()...), sdk.Uint64ToBigEndian(id)...)
}
