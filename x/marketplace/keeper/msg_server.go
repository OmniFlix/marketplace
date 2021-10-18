package keeper

import (
	"context"
	"github.com/OmniFlix/marketplace/x/marketplace/types"
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
func (m msgServer) ListNFT(goCtx context.Context,
	msg *types.MsgListNFT) (*types.MsgListNFTResponse, error) {
	return &types.MsgListNFTResponse{}, nil
}

func (m msgServer) DeListNFT(goCtx context.Context,
	msg *types.MsgDeListNFT) (*types.MsgDeListNFTResponse, error) {
	return &types.MsgDeListNFTResponse{}, nil
}

func (m msgServer) BuyNFT(goCtx context.Context,
	msg *types.MsgBuyNFT) (*types.MsgBuyNFTResponse, error) {
	return &types.MsgBuyNFTResponse{}, nil
}
