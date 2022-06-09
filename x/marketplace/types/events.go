package types

const (
	EventTypeListNFT                = "list_nft"
	EventTypeEditListing            = "edit_listing"
	EventTypeDeListNFT              = "de_list_nft"
	EventTypeBuyNFT                 = "buy_nft"
	EventTypeRoyaltyShareTransfer   = "royalty_share_transfer"
	EventTypeSplitShareTransfer     = "split_share_transfer"
	EventTypeSaleCommissionTransfer = "sale_commission_transfer"

	EventTypeCreateAuction                = "create_auction"
	EcentTypeCancelAuction                = "cancel_auction"
	EventTypePlaceBid                     = "place_bid"

	AttributeValueCategory = ModuleName
	AttributeKeyListingId  = "listing-id"
	AttributeKeyDenomId    = "denom-id"
	AttributeKeyNftId      = "nft-id"
	AttributeKeyBuyer      = "buyer"
	AttributeKeyOwner      = "owner"
	AttributeKeyRecipient  = "recipient"
	AttributeKeyAmount     = "amount"
	AttributeKeyAuctionId  = "auction-id"
	AttributeKeyStartPrice  = "start-price"
)
