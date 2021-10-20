package keeper

import (
	"context"
	"github.com/OmniFlix/marketplace/x/marketplace/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type msgServer struct {
	Keeper
}

var _ types.MsgServer = msgServer{}

// NewMsgServerImpl returns an implementation of the token MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}
func (m msgServer) ListNFT(goCtx context.Context, msg *types.MsgListNFT) (*types.MsgListNFTResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nil, err
	}

	nft, err := m.nftKeeper.GetONFT(ctx, msg.DenomId, msg.NftId)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrNftNotExists,
			"invalid nft and or denomId, nftId %s, denomId %s", msg.NftId, msg.DenomId)
	}
	if owner.String() != nft.GetOwner().String() {
		return nil, sdkerrors.Wrapf(types.ErrUnauthorized, "unauthorized address %s", owner)
	}
	if !nft.IsTransferable() {
		return nil, sdkerrors.Wrapf(
			types.ErrNftNonTransferable, "non-transferable nfts not allowed to list in marketplace")
	}

	listing := types.NewListing(msg.Id, msg.NftId, msg.DenomId, msg.Price, sdk.AccAddress(msg.Owner))
	_ = m.Keeper.AddListing(ctx, listing)

	ctx.EventManager().EmitTypedEvent(
		&types.EventListNFT{
			Id:      listing.Id,
			NftId:   listing.NftId,
			DenomId: listing.DenomId,
			Owner:   listing.Owner,
		},
	)

	return &types.MsgListNFTResponse{}, nil
}

func (m msgServer) EditListing(goCtx context.Context,
	msg *types.MsgEditListing) (*types.MsgEditListingResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nil, err
	}

	listing, found := m.Keeper.GetListing(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrListingDoesNotExists, "listing id %s not exists", listing.Id)
	}
	if owner.String() != listing.Owner {
		return nil, sdkerrors.Wrapf(types.ErrUnauthorized, "unauthorized address %s", owner)
	}
	if err := types.ValidatePrice(msg.Price); err != nil {
		return nil, err
	}
	listing.Price = msg.Price
	m.Keeper.SetListing(ctx, listing)

	ctx.EventManager().EmitTypedEvent(
		&types.EventEditListing{
			Id:      listing.Id,
			NftId:   listing.NftId,
			DenomId: listing.DenomId,
			Owner:   listing.Owner,
		},
	)

	return &types.MsgEditListingResponse{}, nil
}

func (m msgServer) DeListNFT(goCtx context.Context,
	msg *types.MsgDeListNFT) (*types.MsgDeListNFTResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nil, err
	}
	listing, found := m.Keeper.GetListing(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrListingDoesNotExists, "listing id %s not exists", listing.Id)
	}
	if owner.String() != listing.Owner {
		return nil, sdkerrors.Wrapf(types.ErrUnauthorized, "unauthorized address %s", owner)
	}
	m.Keeper.DeleteListing(ctx, listing)

	ctx.EventManager().EmitTypedEvent(
		&types.EventDeListNFT{
			Id:      listing.Id,
			NftId:   listing.NftId,
			DenomId: listing.DenomId,
			Owner:   listing.Owner,
		},
	)

	return &types.MsgDeListNFTResponse{}, nil
}

func (m msgServer) BuyNFT(goCtx context.Context, msg *types.MsgBuyNFT) (*types.MsgBuyNFTResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	buyer, err := sdk.AccAddressFromBech32(msg.Buyer)
	if err != nil {
		return nil, err
	}

	listing, found := m.Keeper.GetListing(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrListingDoesNotExists, "listing id %s not exists", listing.Id)
	}
	if err := types.ValidatePrice(msg.Price); err != nil {
		return nil, err
	}
	if !msg.Price.Equal(listing.Price) {
		if msg.Price.Denom != listing.Price.Denom {
			return nil, sdkerrors.Wrapf(types.ErrInvalidPriceDenom, "invalid price denom %s", msg.Price.Denom)
		}
		if msg.Price.Amount.LT(listing.Price.Amount) {
			return nil, sdkerrors.Wrapf(types.ErrNotEnoughAmount,
				"%s is not enough, to buy %s required", msg.Price.String(), listing.Price.String())
		}
		return nil, sdkerrors.Wrapf(types.ErrInvalidPrice,
			"price %s not matched with listing price", msg.Price.String())
	}
	err = m.Keeper.Buy(ctx, listing, buyer)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitTypedEvent(
		&types.EventBuyNFT{
			Id:      listing.Id,
			NftId:   listing.NftId,
			DenomId: listing.DenomId,
			Owner:   listing.Owner,
			Buyer:   msg.Buyer,
		},
	)

	return &types.MsgBuyNFTResponse{}, nil
}