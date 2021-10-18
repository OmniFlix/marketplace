package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	MsgRoute = "marketplace"

	TypeMsgListNFT   = "list_nft"
	TypeMsgDeListNFT = "de_list_nft"
	TypeMsgBuyNFT    = "buy_nft"

	// DoNotModify used to indicate that some field should not be updated
	DoNotModify = "[do-not-modify]"
)

var (
	_ sdk.Msg = &MsgListNFT{}
	_ sdk.Msg = &MsgDeListNFT{}
	_ sdk.Msg = &MsgBuyNFT{}
)

func NewMsgListNFT() *MsgListNFT {
	return &MsgListNFT{}
}

func (msg MsgListNFT) Route() string { return MsgRoute }

func (msg MsgListNFT) Type() string { return TypeMsgListNFT }

func (msg MsgListNFT) ValidateBasic() error {
	return nil
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
	from, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

// NewMsgDeListNFT
func NewMsgDeListNFT() *MsgDeListNFT {
	return &MsgDeListNFT{}
}

// Route Implements Msg.
func (msg MsgDeListNFT) Route() string { return MsgRoute }

// Type Implements Msg.
func (msg MsgDeListNFT) Type() string { return TypeMsgDeListNFT }

// ValidateBasic Implements Msg.
func (msg MsgDeListNFT) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
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
	from, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

// NewMsgBuyNFT
func NewMsgBuyNFT() *MsgBuyNFT {
	return &MsgBuyNFT{}
}

// Route Implements Msg.
func (msg MsgBuyNFT) Route() string { return MsgRoute }

// Type Implements Msg.
func (msg MsgBuyNFT) Type() string { return TypeMsgBuyNFT }

// ValidateBasic Implements Msg.
func (msg MsgBuyNFT) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
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
	from, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}
