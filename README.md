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
- List NFT

    usage:
    ```shell
    marketplaced tx marketplace list-nft [flags]
    ```
   Example:
    ```shell
    marketplaced tx marketplace list-nft --nft-id=<nft-id> --denom-id=<denom-id> --price=<price> --from=<key-name> --chain-id=<chain-id> --fees=<fee>
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
