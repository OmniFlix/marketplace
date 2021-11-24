package cli

import (
	"context"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"strings"

	"github.com/OmniFlix/marketplace/x/marketplace/types"
	"github.com/spf13/cobra"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd() *cobra.Command {
	// Group marketplace queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	cmd.AddCommand(
		GetCmdQueryListing(),
		GetCmdQueryAllListings(),
		GetCmdQueryListingsByOwner(),
	)

	return cmd
}

// GetCmdQueryListing implements the query listing command.
func GetCmdQueryListing() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "listing [id]",
		Long:    "Query a listing by id.",
		Example: fmt.Sprintf("$ %s query marketplace listing <id>", version.AppName),
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadPersistentCommandFlags(clientCtx, cmd.Flags())

			if err != nil {
				return err
			}

			listingId := strings.ToLower(strings.TrimSpace(args[0]))

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Listing(context.Background(), &types.QueryListingRequest{
				Id: listingId,
			})

			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res.Listing)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryAllListings implements the query all listings command.
func GetCmdQueryAllListings() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "listings",
		Long:    "Query listings.",
		Example: fmt.Sprintf("$ %s query marketplace listings", version.AppName),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadPersistentCommandFlags(clientCtx, cmd.Flags())

			if err != nil {
				return err
			}
			owner, err := cmd.Flags().GetString(FlagOwner)
			if err != nil {
				return err
			}
			priceDenom, err := cmd.Flags().GetString(FlagPriceDenom)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}
			resp, err := queryClient.Listings(
				context.Background(),
				&types.QueryListingsRequest{
					Owner:      owner,
					PriceDenom: priceDenom,
					Pagination: pageReq,
				},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(resp)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	cmd.Flags().String(FlagOwner, "", "filter by owner address")
	cmd.Flags().String(FlagPriceDenom, "", "filter by listing price-denom")
	flags.AddPaginationFlagsToCmd(cmd, "all listings")

	return cmd
}

// GetCmdQueryListingsByOwner implements the query listings by owner command.
func GetCmdQueryListingsByOwner() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "listings-by-owner [owner]",
		Long:    "Query listings by the owner.",
		Example: fmt.Sprintf("$ %s query marketplace listings <owner>", version.AppName),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadPersistentCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			var owner sdk.AccAddress
			if len(args) > 0 {
				owner, err = sdk.AccAddressFromBech32(args[0])
				if err != nil {
					return err
				}
			}

			queryClient := types.NewQueryClient(clientCtx)
			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}
			resp, err := queryClient.ListingsByOwner(
				context.Background(),
				&types.QueryListingsByOwnerRequest{
					Owner:      owner.String(),
					Pagination: pageReq,
				},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(resp)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "owner listings")

	return cmd
}
