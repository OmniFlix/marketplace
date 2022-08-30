# marketplace

## CLI Commands

### Queries

Querying commands for the marketplace module

Usage:
```shell
  marketplaced query marketplace [flags]
  marketplaced query marketplace [command]
```

Available Commands:
```shell
  listing           
  listings          
  listings-by-owner 
  auction
  auctions
  auctions-by-owner
  bid
  bids
  params
```
 - Get Listings 
   
    usage:
   ```shell
    marketplaced q marketplace listings [Flags]
   ```
   Flags:
   
    **owner**: get listings of a specfic account address
   
    **pagination flags**: count-toal, limit, offset etc ..


- Get Listing Details

  usage:
  ```shell
   marketplaced q marketplace listing [listingId] [Flags]
  ```
  args:
  
  **listingsId**: listing id

- Get Listings by owner

  usage:
  ```shell
   marketplaced q marketplace listings-by-owner  [owner] [Flags]
  ```
  args:

  **owner**: bech32 account address
  
 - Get Auctions 
   
    usage:
   ```shell
    marketplaced q marketplace auctions [Flags]
   ```
   Flags:
   
    **owner**: get auctions of a specfic account address
   
    **pagination flags**: count-toal, limit, offset etc ..

 - Get Auction
   
    usage:
   ```shell
    marketplaced q marketplace auction <auction-id> [Flags]
   ```
- Get Auction Bid
     usage:
   ```shell
    marketplaced q marketplace bid <auction-id> [Flags]
   ```

### Transactions

marketplace transactions subcommands

Usage:
```shell
marketplaced tx marketplace [flags]
marketplaced tx marketplace [command]
```
Available Commands:
```shell
list-nft     - List an nft on marketplace
buy-nft      - Buy an nft from marketplace
de-list-nft  - DeList an nft from marketplace
edit-listing - Edit active listing on marketplace
```
#### Listings

- List NFT

    usage:
    ```shell
    marketplaced tx marketplace list-nft [flags]
    ```
   Example:
    ```shell
    marketplaced tx marketplace list-nft 
      --nft-id="nft_id"
      --denom-id="denom_id"
      --price="10000000uflix"
      --from=test-key
      --chain-id="chain_id"
      --fees="200uflix"
    ```
   For splitting sale amount between multiple accounts use `split-shares`
   ```shell
   --split-shares="address:percentage,address:percentage"
   ```
   Example:
   ```shell
  --split-shares "omniflix1e49p22vz8w5nyer77gl0nhs2puumu3jdel822w:0.70,omniflix1muyp5qvz7e6qd8wkpxex0h963um962qcd777ez:0.30"
   ```

- Buy NFT

  usage:
    ```shell
    marketplaced tx marketplace buy-nft [flags]
    ```
  Example:
    ```shell
    marketplaced tx marketplace buy-nft [listing-id]--price=<price> --from=<key-name> --chain-id=<chain-id> --fees=<fee>
    ```

- DeList NFT

  usage:
    ```shell
    marketplaced tx marketplace de-list-nft [flags]
    ```
  Example:
    ```shell
    marketplaced tx marketplace de-list-nft [listing-id] --from=<key-name> --chain-id=<chain-id> --fees=<fee>
    ```
    
#### Auctions

- Create Auction

  usage:
    ```shell
    marketplaced tx marketplace create-auction [flags]
    ```
  Example:
    ```shell
    marketplaced tx marketplace create-auction 
      --nft-id=<nft_id>  
      --denom-id=<denom_id>  
      --start-price=1000000uflix 
      --start-time="2022-11-27T17:26:00.000Z" 
      --from=<key-name> 
      --chain-id=<chain-id>
      --fees=<fee>
    ```

- Cancel Auction (only allowed when there are no bids)

  usage:
    ```shell
    marketplaced tx marketplace cancel-auction [auction-id] [flags]
    ```
  Example:
  ```shell
  marketplaced tx marketplace cancel-auction <auction_id> 
    --chain-id=<chain_id> 
    --from=<key-name> 
    --fees=<fee>
  ```
- Place Bid 

  usage:
    ```shell
    marketplaced tx marketplace place-bid [auction-id] [flags]
    ```
  Example:
  ```shell
    marketplaced tx marketplace place-bid <auction-id> 
      --amount=<bid_amount> 
      --chain-id=<chain_id> 
      --from=<key-name> 
      --fees=<fee>
  ```
  
