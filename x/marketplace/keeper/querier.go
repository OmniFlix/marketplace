package keeper

import (
	"fmt"
	"strings"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/OmniFlix/marketplace/x/marketplace/types"
)

func NewQuerier(k Keeper, legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err error) {
		switch path[0] {
		case types.QueryListing:
			return queryListing(ctx, req, k, legacyQuerierCdc)
		case types.QueryAllListings:
			return queryAllListings(ctx, req, k, legacyQuerierCdc)
		case types.QueryListingsByOwner:
			return queryListingsByOwner(ctx, req, k, legacyQuerierCdc)
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unknown query path: %s", path[0])
		}
	}
}

func queryListing(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryListingParams

	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	id := strings.ToLower(strings.TrimSpace(params.Id))

	listing, found := k.GetListing(ctx, id)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrListingDoesNotExists, fmt.Sprintf("listing %s does not exist", id))
	}
	return codec.MarshalJSONIndent(legacyQuerierCdc, listing)
}

func queryAllListings(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryAllListingsParams

	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	listings := k.GetAllListings(ctx)

	return codec.MarshalJSONIndent(legacyQuerierCdc, listings)
}

func queryListingsByOwner(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryListingsByOwnerParams

	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	listings := k.GetListingsByOwner(ctx, params.Owner)
	return codec.MarshalJSONIndent(legacyQuerierCdc, listings)
}
