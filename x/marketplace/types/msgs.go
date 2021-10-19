package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	MsgRoute = "marketplace"

	TypeMsgListNFT     = "list_nft"
	TypeMsgEditListing = "edit_listing"
	TypeMsgDeListNFT   = "de_list_nft"
	TypeMsgBuyNFT      = "buy_nft"

	// DoNotModify used to indicate that some field should not be updated
	DoNotModify = "[do-not-modify]"
	IdPrefix    = "list"
)

var (
	_ sdk.Msg = &MsgListNFT{}
	_ sdk.Msg = &MsgEditListing{}
	_ sdk.Msg = &MsgDeListNFT{}
	_ sdk.Msg = &MsgBuyNFT{}
)

func NewMsgListNFT(denomId, nftId string, price sdk.Coin, owner sdk.AccAddress) *MsgListNFT {
	return &MsgListNFT{
		Id:      GenUniqueID(IdPrefix),
		NftId:   nftId,
		DenomId: denomId,
		Price:   price,
		Owner:   owner.String(),
	}
}

func (msg MsgListNFT) Route() string { return MsgRoute }

func (msg MsgListNFT) Type() string { return TypeMsgListNFT }

func (msg MsgListNFT) ValidateBasic() error {
	return ValidateListing(
		NewListing(
			msg.Id,
			msg.NftId,
			msg.DenomId,
			msg.Price,
			sdk.AccAddress(msg.Owner),
		),
	)
}

// GetSignBytes Implements Msg.
func (msg MsgListNFT) GetSignBytes() []byte {
	b, err := ModuleCdc.MarshalJSON(&msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners Implements Msg.
func (msg MsgListNFT) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

func NewMsgEditListing(id string, price sdk.Coin, owner sdk.AccAddress) *MsgEditListing {
	return &MsgEditListing{
		Id:    id,
		Price: price,
		Owner: owner.String(),
	}
}

func (msg MsgEditListing) Route() string { return MsgRoute }

func (msg MsgEditListing) Type() string { return TypeMsgEditListing }

func (msg MsgEditListing) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	return ValidatePrice(msg.Price)
}

// GetSignBytes Implements Msg.
func (msg MsgEditListing) GetSignBytes() []byte {
	b, err := ModuleCdc.MarshalJSON(&msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners Implements Msg.
func (msg MsgEditListing) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

// NewMsgDeListNFT
func NewMsgDeListNFT(id string, owner sdk.AccAddress) *MsgDeListNFT {
	return &MsgDeListNFT{
		Id:    id,
		Owner: owner.String(),
	}
}

// Route Implements Msg.
func (msg MsgDeListNFT) Route() string { return MsgRoute }

// Type Implements Msg.
func (msg MsgDeListNFT) Type() string { return TypeMsgDeListNFT }

// ValidateBasic Implements Msg.
func (msg MsgDeListNFT) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgDeListNFT) GetSignBytes() []byte {
	b, err := ModuleCdc.MarshalJSON(&msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners Implements Msg.
func (msg MsgDeListNFT) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

// NewMsgBuyNFT
func NewMsgBuyNFT(id string, price sdk.Coin, buyer sdk.AccAddress) *MsgBuyNFT {
	return &MsgBuyNFT{
		Id:    id,
		Price: price,
		Buyer: buyer.String(),
	}
}

// Route Implements Msg.
func (msg MsgBuyNFT) Route() string { return MsgRoute }

// Type Implements Msg.
func (msg MsgBuyNFT) Type() string { return TypeMsgBuyNFT }

// ValidateBasic Implements Msg.
func (msg MsgBuyNFT) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Buyer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return ValidatePrice(msg.Price)
}

// GetSignBytes Implements Msg.
func (msg MsgBuyNFT) GetSignBytes() []byte {
	b, err := ModuleCdc.MarshalJSON(&msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners Implements Msg.
func (msg MsgBuyNFT) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Buyer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}
