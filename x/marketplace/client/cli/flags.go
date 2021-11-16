package cli

import (
	flag "github.com/spf13/pflag"
)

const (
	FlagDenomId = "denom-id"
	FlagNftId   = "nft-id"
	FlagPrice   = "price"
	FlagOwner   = "owner"
)

var (
	FsListNft     = flag.NewFlagSet("", flag.ContinueOnError)
	FsEditListing = flag.NewFlagSet("", flag.ContinueOnError)
	FsBuyNFT      = flag.NewFlagSet("", flag.ContinueOnError)
)

func init() {
	FsListNft.String(FlagDenomId, "", "nft denom id")
	FsListNft.String(FlagNftId, "", "nft id")
	FsListNft.String(FlagPrice, "", "listing price of nft")

	FsEditListing.String(FlagPrice, "", "listing price of nft")

	FsBuyNFT.String(FlagPrice, "", "buying price of nft")
}
