package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"
	"strings"

	"github.com/OmniFlix/marketplace/x/marketplace/types"
	"github.com/cosmos/cosmos-sdk/client"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	marketplaceTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	marketplaceTxCmd.AddCommand(
		GetCmdListNft(),
		GetCmdEditListing(),
		GetCmdDeListNft(),
		GetCmdBuyNft(),
	)

	return marketplaceTxCmd
}

// GetCmdListNft implements the list-nft command
func GetCmdListNft() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "list-nft",
		Long: "lists an nft on marketplace",
		Example: fmt.Sprintf(
			"$ %s tx marketplace list-nft "+
				"--nft-id=<nft-id> "+
				"--denom-id=<nft-id> "+
				"--price=\"1000000uflix\" "+
				"--from=<key-name> "+
				"--chain-id=<chain-id> "+
				"--fees=<fee>",
			version.AppName,
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			owner := clientCtx.GetFromAddress()
			denomId, err := cmd.Flags().GetString(FlagDenomId)
			if err != nil {
				return err
			}
			nftId, err := cmd.Flags().GetString(FlagNftId)
			if err != nil {
				return err
			}
			priceStr, err := cmd.Flags().GetString(FlagPrice)
			if err != nil {
				return err
			}
			price, err := sdk.ParseCoinNormalized(priceStr)
			if err != nil {
				return fmt.Errorf("failed to parse price: %s", price)
			}

			msg := types.NewMsgListNFT(denomId, nftId, price, owner)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().AddFlagSet(FsListNft)
	_ = cmd.MarkFlagRequired(FlagDenomId)
	_ = cmd.MarkFlagRequired(FlagNftId)
	_ = cmd.MarkFlagRequired(FlagPrice)
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdEditListing implements the edit-listing command
func GetCmdEditListing() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "edit-listing",
		Long: "Edit an existing marketplace listing ",
		Example: fmt.Sprintf(
			"$ %s tx marketplace edit-listing [listing-id] "+
				"--price=\"1000000uflix\" "+
				"--from=<key-name> "+
				"--chain-id=<chain-id> "+
				"--fees=<fee>",
			version.AppName,
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			owner := clientCtx.GetFromAddress()

			listingId := strings.TrimSpace(args[0])

			priceStr, err := cmd.Flags().GetString(FlagPrice)
			if err != nil {
				return err
			}
			price, err := sdk.ParseCoinNormalized(priceStr)
			if err != nil {
				return fmt.Errorf("failed to parse price: %s", price)
			}

			msg := types.NewMsgEditListing(listingId, price, owner)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().AddFlagSet(FsEditListing)
	_ = cmd.MarkFlagRequired(FlagPrice)
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdDeListNft implements the de-list-nft command
func GetCmdDeListNft() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "de-list-nft",
		Long: "de-list an existing listing from marketplace",
		Example: fmt.Sprintf(
			"$ %s tx marketplace de-list-nft [listing-id] "+
				"--from=<key-name> "+
				"--chain-id=<chain-id> "+
				"--fees=<fee>",
			version.AppName,
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			owner := clientCtx.GetFromAddress()

			listingId := strings.TrimSpace(args[0])

			msg := types.NewMsgDeListNFT(listingId, owner)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdBuyNft implements the buy-nft command
func GetCmdBuyNft() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "buy-nft",
		Short: "Buy an nft from marketplace",
		Example: fmt.Sprintf(
			"$ %s tx marketplace buy-nft [listing-id]"+
				"--from=<key-name> "+
				"--chain-id=<chain-id> "+
				"--fees=<fee>",
			version.AppName,
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			buyer := clientCtx.GetFromAddress()
			listingId := strings.TrimSpace(args[0])

			priceStr, err := cmd.Flags().GetString(FlagPrice)
			if err != nil {
				return err
			}
			price, err := sdk.ParseCoinNormalized(priceStr)
			if err != nil {
				return fmt.Errorf("failed to parse price: %s", price)
			}

			msg := types.NewMsgBuyNFT(listingId, price, buyer)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().AddFlagSet(FsBuyNFT)
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
