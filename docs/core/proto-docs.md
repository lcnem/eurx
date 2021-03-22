<!-- This file is auto-generated. Please do not modify it yourself. -->
# Protobuf Documentation
<a name="top"></a>

## Table of Contents

- [auction/auction.proto](#auction/auction.proto)
    - [BaseAuction](#eurx.auction.BaseAuction)
    - [CollateralAuction](#eurx.auction.CollateralAuction)
    - [DebtAuction](#eurx.auction.DebtAuction)
    - [MsgPlaceBid](#eurx.auction.MsgPlaceBid)
    - [Params](#eurx.auction.Params)
    - [SurplusAuction](#eurx.auction.SurplusAuction)
    - [WeightedAddresses](#eurx.auction.WeightedAddresses)
  
- [auction/genesis.proto](#auction/genesis.proto)
    - [GenesisState](#eurx.auction.GenesisState)
  
- [auction/query.proto](#auction/query.proto)
    - [QueryAllAuctionRequest](#eurx.auction.QueryAllAuctionRequest)
    - [QueryAllAuctionResponse](#eurx.auction.QueryAllAuctionResponse)
    - [QueryGetAuctionRequest](#eurx.auction.QueryGetAuctionRequest)
    - [QueryGetAuctionResponse](#eurx.auction.QueryGetAuctionResponse)
    - [QueryParamsRequest](#eurx.auction.QueryParamsRequest)
    - [QueryParamsResponse](#eurx.auction.QueryParamsResponse)
  
    - [Query](#eurx.auction.Query)
  
- [cdp/cdp.proto](#cdp/cdp.proto)
    - [AugmentedCDP](#eurx.cdp.AugmentedCDP)
    - [CDP](#eurx.cdp.CDP)
    - [CollateralParam](#eurx.cdp.CollateralParam)
    - [DebtParam](#eurx.cdp.DebtParam)
    - [Deposit](#eurx.cdp.Deposit)
    - [MsgCreateCDP](#eurx.cdp.MsgCreateCDP)
    - [MsgDeposit](#eurx.cdp.MsgDeposit)
    - [MsgDrawDebt](#eurx.cdp.MsgDrawDebt)
    - [MsgLiquidate](#eurx.cdp.MsgLiquidate)
    - [MsgRepayDebt](#eurx.cdp.MsgRepayDebt)
    - [MsgWithdraw](#eurx.cdp.MsgWithdraw)
    - [Params](#eurx.cdp.Params)
  
- [cdp/genesis.proto](#cdp/genesis.proto)
    - [GenesisAccumulationTime](#eurx.cdp.GenesisAccumulationTime)
    - [GenesisState](#eurx.cdp.GenesisState)
    - [GenesisTotalPrincipal](#eurx.cdp.GenesisTotalPrincipal)
  
- [cdp/query.proto](#cdp/query.proto)
    - [QueryAllCdpRequest](#eurx.cdp.QueryAllCdpRequest)
    - [QueryAllCdpResponse](#eurx.cdp.QueryAllCdpResponse)
    - [QueryGetCdpRequest](#eurx.cdp.QueryGetCdpRequest)
    - [QueryGetCdpResponse](#eurx.cdp.QueryGetCdpResponse)
    - [QueryParamsRequest](#eurx.cdp.QueryParamsRequest)
    - [QueryParamsResponse](#eurx.cdp.QueryParamsResponse)
  
    - [Query](#eurx.cdp.Query)
  
- [estmdist/estmdist.proto](#estmdist/estmdist.proto)
    - [Params](#eurx.estmdist.Params)
    - [Period](#eurx.estmdist.Period)
  
- [estmdist/genesis.proto](#estmdist/genesis.proto)
    - [GenesisState](#eurx.estmdist.GenesisState)
  
- [estmdist/query.proto](#estmdist/query.proto)
    - [QueryGetBalancesRequest](#eurx.estmdist.QueryGetBalancesRequest)
    - [QueryGetBalancesResponse](#eurx.estmdist.QueryGetBalancesResponse)
    - [QueryParamsRequest](#eurx.estmdist.QueryParamsRequest)
    - [QueryParamsResponse](#eurx.estmdist.QueryParamsResponse)
  
    - [Query](#eurx.estmdist.Query)
  
- [incentive/incentive.proto](#incentive/incentive.proto)
    - [BaseClaim](#eurx.incentive.BaseClaim)
    - [BaseMultiClaim](#eurx.incentive.BaseMultiClaim)
    - [EURXMintingClaim](#eurx.incentive.EURXMintingClaim)
    - [MsgClaimEURXMintingReward](#eurx.incentive.MsgClaimEURXMintingReward)
    - [Multiplier](#eurx.incentive.Multiplier)
    - [Params](#eurx.incentive.Params)
    - [RewardIndex](#eurx.incentive.RewardIndex)
    - [RewardPeriod](#eurx.incentive.RewardPeriod)
  
- [incentive/genesis.proto](#incentive/genesis.proto)
    - [GenesisAccumulationTime](#eurx.incentive.GenesisAccumulationTime)
    - [GenesisState](#eurx.incentive.GenesisState)
  
- [incentive/query.proto](#incentive/query.proto)
    - [QueryParamsRequest](#eurx.incentive.QueryParamsRequest)
    - [QueryParamsResponse](#eurx.incentive.QueryParamsResponse)
  
    - [Query](#eurx.incentive.Query)
  
- [pricefeed/pricefeed.proto](#pricefeed/pricefeed.proto)
    - [CurrentPrice](#eurx.pricefeed.CurrentPrice)
    - [Market](#eurx.pricefeed.Market)
    - [MsgPostPrice](#eurx.pricefeed.MsgPostPrice)
    - [Params](#eurx.pricefeed.Params)
    - [PostedPrice](#eurx.pricefeed.PostedPrice)
  
- [pricefeed/genesis.proto](#pricefeed/genesis.proto)
    - [GenesisState](#eurx.pricefeed.GenesisState)
  
- [pricefeed/query.proto](#pricefeed/query.proto)
    - [QueryAllMarketRequest](#eurx.pricefeed.QueryAllMarketRequest)
    - [QueryAllMarketResponse](#eurx.pricefeed.QueryAllMarketResponse)
    - [QueryAllOracleRequest](#eurx.pricefeed.QueryAllOracleRequest)
    - [QueryAllOracleResponse](#eurx.pricefeed.QueryAllOracleResponse)
    - [QueryAllPriceRequest](#eurx.pricefeed.QueryAllPriceRequest)
    - [QueryAllPriceResponse](#eurx.pricefeed.QueryAllPriceResponse)
    - [QueryAllRawPriceRequest](#eurx.pricefeed.QueryAllRawPriceRequest)
    - [QueryAllRawPriceResponse](#eurx.pricefeed.QueryAllRawPriceResponse)
    - [QueryGetPriceRequest](#eurx.pricefeed.QueryGetPriceRequest)
    - [QueryGetPriceResponse](#eurx.pricefeed.QueryGetPriceResponse)
    - [QueryParamsRequest](#eurx.pricefeed.QueryParamsRequest)
    - [QueryParamsResponse](#eurx.pricefeed.QueryParamsResponse)
  
    - [Query](#eurx.pricefeed.Query)
  
- [Scalar Value Types](#scalar-value-types)



<a name="auction/auction.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## auction/auction.proto



<a name="eurx.auction.BaseAuction"></a>

### BaseAuction



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [uint64](#uint64) |  |  |
| `initiator` | [string](#string) |  |  |
| `lot` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `bidder` | [string](#string) |  |  |
| `bid` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `has_received_bids` | [bool](#bool) |  |  |
| `end_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `max_end_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="eurx.auction.CollateralAuction"></a>

### CollateralAuction



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `base_auction` | [BaseAuction](#eurx.auction.BaseAuction) |  |  |
| `corresponding_debt` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `max_bid` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `lot_returns` | [WeightedAddresses](#eurx.auction.WeightedAddresses) |  |  |






<a name="eurx.auction.DebtAuction"></a>

### DebtAuction



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `base_auction` | [BaseAuction](#eurx.auction.BaseAuction) |  |  |
| `corresponding_debt` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="eurx.auction.MsgPlaceBid"></a>

### MsgPlaceBid



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `auction_id` | [uint64](#uint64) |  |  |
| `bidder` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="eurx.auction.Params"></a>

### Params



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `max_auction_duration` | [google.protobuf.Duration](#google.protobuf.Duration) |  |  |
| `bid_duration` | [google.protobuf.Duration](#google.protobuf.Duration) |  |  |
| `increment_surplus` | [string](#string) |  |  |
| `increment_debt` | [string](#string) |  |  |
| `increment_collateral` | [string](#string) |  |  |






<a name="eurx.auction.SurplusAuction"></a>

### SurplusAuction



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `base_auction` | [BaseAuction](#eurx.auction.BaseAuction) |  |  |






<a name="eurx.auction.WeightedAddresses"></a>

### WeightedAddresses



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `addresses` | [string](#string) | repeated |  |
| `weights` | [string](#string) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="auction/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## auction/genesis.proto



<a name="eurx.auction.GenesisState"></a>

### GenesisState
GenesisState defines the auction module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `next_auction_id` | [uint64](#uint64) |  |  |
| `params` | [Params](#eurx.auction.Params) |  |  |
| `auctions` | [google.protobuf.Any](#google.protobuf.Any) | repeated | this line is used by starport scaffolding # genesis/proto/state |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="auction/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## auction/query.proto



<a name="eurx.auction.QueryAllAuctionRequest"></a>

### QueryAllAuctionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="eurx.auction.QueryAllAuctionResponse"></a>

### QueryAllAuctionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `auctions` | [google.protobuf.Any](#google.protobuf.Any) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="eurx.auction.QueryGetAuctionRequest"></a>

### QueryGetAuctionRequest
this line is used by starport scaffolding # 3


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [uint64](#uint64) |  |  |






<a name="eurx.auction.QueryGetAuctionResponse"></a>

### QueryGetAuctionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `auction` | [google.protobuf.Any](#google.protobuf.Any) |  |  |






<a name="eurx.auction.QueryParamsRequest"></a>

### QueryParamsRequest







<a name="eurx.auction.QueryParamsResponse"></a>

### QueryParamsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#eurx.auction.Params) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="eurx.auction.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#eurx.auction.QueryParamsRequest) | [QueryParamsResponse](#eurx.auction.QueryParamsResponse) |  | GET|/eurx/auction/params|
| `Auction` | [QueryGetAuctionRequest](#eurx.auction.QueryGetAuctionRequest) | [QueryGetAuctionResponse](#eurx.auction.QueryGetAuctionResponse) | this line is used by starport scaffolding # 2 | GET|/eurx/auction/auctions/{id}|
| `AuctionAll` | [QueryAllAuctionRequest](#eurx.auction.QueryAllAuctionRequest) | [QueryAllAuctionResponse](#eurx.auction.QueryAllAuctionResponse) |  | GET|/eurx/auction/auctions|

 <!-- end services -->



<a name="cdp/cdp.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cdp/cdp.proto



<a name="eurx.cdp.AugmentedCDP"></a>

### AugmentedCDP



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `cdp` | [CDP](#eurx.cdp.CDP) |  |  |
| `collateral_value` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `collateralization_ratio` | [string](#string) |  |  |






<a name="eurx.cdp.CDP"></a>

### CDP



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [uint64](#uint64) |  |  |
| `owner` | [string](#string) |  |  |
| `type` | [string](#string) |  |  |
| `collateral` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `principal` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `accumulated_fees` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `fees_updated` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `interest_factor` | [string](#string) |  |  |






<a name="eurx.cdp.CollateralParam"></a>

### CollateralParam



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `type` | [string](#string) |  |  |
| `liquidation_ratio` | [string](#string) |  |  |
| `debt_limit` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `stability_fee` | [string](#string) |  |  |
| `auction_size` | [string](#string) |  |  |
| `liquidation_penalty` | [string](#string) |  |  |
| `prefix` | [uint32](#uint32) |  |  |
| `spot_market_id` | [string](#string) |  |  |
| `liquidation_market_id` | [string](#string) |  |  |
| `keeper_reward_percentage` | [string](#string) |  |  |
| `check_collateralization_index_count` | [string](#string) |  |  |
| `conversion_factor` | [string](#string) |  |  |






<a name="eurx.cdp.DebtParam"></a>

### DebtParam



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `reference_asset` | [string](#string) |  |  |
| `conversion_factor` | [string](#string) |  |  |
| `debt_floor` | [string](#string) |  |  |






<a name="eurx.cdp.Deposit"></a>

### Deposit



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `cdp_id` | [uint64](#uint64) |  |  |
| `depositor` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="eurx.cdp.MsgCreateCDP"></a>

### MsgCreateCDP



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `collateral` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `principal` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `collateral_type` | [string](#string) |  |  |






<a name="eurx.cdp.MsgDeposit"></a>

### MsgDeposit



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `depositor` | [string](#string) |  |  |
| `owner` | [string](#string) |  |  |
| `collateral` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `collateral_type` | [string](#string) |  |  |






<a name="eurx.cdp.MsgDrawDebt"></a>

### MsgDrawDebt



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `collateral_type` | [string](#string) |  |  |
| `principal` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="eurx.cdp.MsgLiquidate"></a>

### MsgLiquidate



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `keeper` | [string](#string) |  |  |
| `borrower` | [string](#string) |  |  |
| `collateral_type` | [string](#string) |  |  |






<a name="eurx.cdp.MsgRepayDebt"></a>

### MsgRepayDebt



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `collateral_type` | [string](#string) |  |  |
| `payment` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="eurx.cdp.MsgWithdraw"></a>

### MsgWithdraw



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `depositor` | [string](#string) |  |  |
| `owner` | [string](#string) |  |  |
| `collateral` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `collateral_type` | [string](#string) |  |  |






<a name="eurx.cdp.Params"></a>

### Params



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collateral_params` | [CollateralParam](#eurx.cdp.CollateralParam) | repeated |  |
| `debt_param` | [DebtParam](#eurx.cdp.DebtParam) |  |  |
| `global_debt_limit` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `surplus_auction_threshold` | [string](#string) |  |  |
| `surplus_auction_lot` | [string](#string) |  |  |
| `debt_auction_threshold` | [string](#string) |  |  |
| `debt_auction_lot` | [string](#string) |  |  |
| `circuit_breaker` | [bool](#bool) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cdp/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cdp/genesis.proto



<a name="eurx.cdp.GenesisAccumulationTime"></a>

### GenesisAccumulationTime



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collateral_type` | [string](#string) |  |  |
| `previous_accumulation_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `interest_factor` | [string](#string) |  |  |






<a name="eurx.cdp.GenesisState"></a>

### GenesisState
GenesisState defines the cdp module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#eurx.cdp.Params) |  |  |
| `cdps` | [CDP](#eurx.cdp.CDP) | repeated |  |
| `deposits` | [Deposit](#eurx.cdp.Deposit) | repeated |  |
| `starting_cdp_id` | [uint64](#uint64) |  |  |
| `debt_denom` | [string](#string) |  |  |
| `gov_denom` | [string](#string) |  |  |
| `previous_accumulation_times` | [GenesisAccumulationTime](#eurx.cdp.GenesisAccumulationTime) | repeated |  |
| `total_principals` | [GenesisTotalPrincipal](#eurx.cdp.GenesisTotalPrincipal) | repeated | this line is used by starport scaffolding # genesis/proto/state |






<a name="eurx.cdp.GenesisTotalPrincipal"></a>

### GenesisTotalPrincipal



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collateral_type` | [string](#string) |  |  |
| `total_principal` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cdp/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cdp/query.proto



<a name="eurx.cdp.QueryAllCdpRequest"></a>

### QueryAllCdpRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="eurx.cdp.QueryAllCdpResponse"></a>

### QueryAllCdpResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `Cdp` | [CDP](#eurx.cdp.CDP) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="eurx.cdp.QueryGetCdpRequest"></a>

### QueryGetCdpRequest
this line is used by starport scaffolding # 3


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [uint64](#uint64) |  |  |
| `collateral_type` | [string](#string) |  |  |






<a name="eurx.cdp.QueryGetCdpResponse"></a>

### QueryGetCdpResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `Cdp` | [CDP](#eurx.cdp.CDP) |  |  |






<a name="eurx.cdp.QueryParamsRequest"></a>

### QueryParamsRequest







<a name="eurx.cdp.QueryParamsResponse"></a>

### QueryParamsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#eurx.cdp.Params) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="eurx.cdp.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#eurx.cdp.QueryParamsRequest) | [QueryParamsResponse](#eurx.cdp.QueryParamsResponse) |  | GET|/eurx/cdp/params|
| `Cdp` | [QueryGetCdpRequest](#eurx.cdp.QueryGetCdpRequest) | [QueryGetCdpResponse](#eurx.cdp.QueryGetCdpResponse) | this line is used by starport scaffolding # 2 | GET|/eurx/cdp/cdps/{id}|
| `CdpAll` | [QueryAllCdpRequest](#eurx.cdp.QueryAllCdpRequest) | [QueryAllCdpResponse](#eurx.cdp.QueryAllCdpResponse) |  | GET|/eurx/cdp/cdps|

 <!-- end services -->



<a name="estmdist/estmdist.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## estmdist/estmdist.proto



<a name="eurx.estmdist.Params"></a>

### Params



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `active` | [bool](#bool) |  |  |
| `periods` | [Period](#eurx.estmdist.Period) | repeated |  |






<a name="eurx.estmdist.Period"></a>

### Period



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `start` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `end` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `inflation` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="estmdist/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## estmdist/genesis.proto



<a name="eurx.estmdist.GenesisState"></a>

### GenesisState
GenesisState defines the estmdist module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#eurx.estmdist.Params) |  |  |
| `previous_block_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | this line is used by starport scaffolding # genesis/proto/state |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="estmdist/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## estmdist/query.proto



<a name="eurx.estmdist.QueryGetBalancesRequest"></a>

### QueryGetBalancesRequest







<a name="eurx.estmdist.QueryGetBalancesResponse"></a>

### QueryGetBalancesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `balances` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="eurx.estmdist.QueryParamsRequest"></a>

### QueryParamsRequest







<a name="eurx.estmdist.QueryParamsResponse"></a>

### QueryParamsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#eurx.estmdist.Params) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="eurx.estmdist.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#eurx.estmdist.QueryParamsRequest) | [QueryParamsResponse](#eurx.estmdist.QueryParamsResponse) |  | GET|/eurx/estmdist/params|
| `Balances` | [QueryGetBalancesRequest](#eurx.estmdist.QueryGetBalancesRequest) | [QueryGetBalancesResponse](#eurx.estmdist.QueryGetBalancesResponse) | this line is used by starport scaffolding # 2 | GET|/eurx/estmdist/balances|

 <!-- end services -->



<a name="incentive/incentive.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## incentive/incentive.proto



<a name="eurx.incentive.BaseClaim"></a>

### BaseClaim



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `owner` | [string](#string) |  |  |
| `reward` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="eurx.incentive.BaseMultiClaim"></a>

### BaseMultiClaim



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `owner` | [string](#string) |  |  |
| `reward` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="eurx.incentive.EURXMintingClaim"></a>

### EURXMintingClaim



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `base_claim` | [BaseClaim](#eurx.incentive.BaseClaim) |  |  |
| `reward_indexes` | [RewardIndex](#eurx.incentive.RewardIndex) | repeated |  |






<a name="eurx.incentive.MsgClaimEURXMintingReward"></a>

### MsgClaimEURXMintingReward



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `multiplier_name` | [string](#string) |  |  |






<a name="eurx.incentive.Multiplier"></a>

### Multiplier



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `name` | [string](#string) |  |  |
| `months_lockup` | [int64](#int64) |  |  |
| `factor` | [string](#string) |  |  |






<a name="eurx.incentive.Params"></a>

### Params



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `eurx_minting_reward_periods` | [RewardPeriod](#eurx.incentive.RewardPeriod) | repeated |  |
| `claim_multipliers` | [Multiplier](#eurx.incentive.Multiplier) | repeated |  |
| `claim_end` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="eurx.incentive.RewardIndex"></a>

### RewardIndex



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collateral_type` | [string](#string) |  |  |
| `reward_factor` | [string](#string) |  |  |






<a name="eurx.incentive.RewardPeriod"></a>

### RewardPeriod



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `active` | [bool](#bool) |  |  |
| `collateral_type` | [string](#string) |  |  |
| `start` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `end` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `rewards_per_second` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="incentive/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## incentive/genesis.proto



<a name="eurx.incentive.GenesisAccumulationTime"></a>

### GenesisAccumulationTime



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collateral_type` | [string](#string) |  |  |
| `previous_accumulation_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="eurx.incentive.GenesisState"></a>

### GenesisState
GenesisState defines the incentive module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#eurx.incentive.Params) |  |  |
| `eurx_accumulation_times` | [GenesisAccumulationTime](#eurx.incentive.GenesisAccumulationTime) | repeated |  |
| `eurx_minting_claims` | [EURXMintingClaim](#eurx.incentive.EURXMintingClaim) | repeated | this line is used by starport scaffolding # genesis/proto/state |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="incentive/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## incentive/query.proto



<a name="eurx.incentive.QueryParamsRequest"></a>

### QueryParamsRequest







<a name="eurx.incentive.QueryParamsResponse"></a>

### QueryParamsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#eurx.incentive.Params) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="eurx.incentive.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#eurx.incentive.QueryParamsRequest) | [QueryParamsResponse](#eurx.incentive.QueryParamsResponse) | this line is used by starport scaffolding # 2 | GET|/eurx/incentive/params|

 <!-- end services -->



<a name="pricefeed/pricefeed.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## pricefeed/pricefeed.proto



<a name="eurx.pricefeed.CurrentPrice"></a>

### CurrentPrice



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `market_id` | [string](#string) |  |  |
| `price` | [string](#string) |  |  |






<a name="eurx.pricefeed.Market"></a>

### Market



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `market_id` | [string](#string) |  |  |
| `base_asset` | [string](#string) |  |  |
| `quote_asset` | [string](#string) |  |  |
| `oracles` | [string](#string) | repeated |  |
| `active` | [bool](#bool) |  |  |






<a name="eurx.pricefeed.MsgPostPrice"></a>

### MsgPostPrice



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `from` | [string](#string) |  |  |
| `market_id` | [string](#string) |  |  |
| `price` | [string](#string) |  |  |
| `expiry` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="eurx.pricefeed.Params"></a>

### Params



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `markets` | [Market](#eurx.pricefeed.Market) | repeated |  |






<a name="eurx.pricefeed.PostedPrice"></a>

### PostedPrice



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `market_id` | [string](#string) |  |  |
| `oracle_address` | [string](#string) |  |  |
| `price` | [string](#string) |  |  |
| `expiry` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="pricefeed/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## pricefeed/genesis.proto



<a name="eurx.pricefeed.GenesisState"></a>

### GenesisState
GenesisState defines the pricefeed module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#eurx.pricefeed.Params) |  |  |
| `posted_prices` | [PostedPrice](#eurx.pricefeed.PostedPrice) | repeated | this line is used by starport scaffolding # genesis/proto/state |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="pricefeed/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## pricefeed/query.proto



<a name="eurx.pricefeed.QueryAllMarketRequest"></a>

### QueryAllMarketRequest
this line is used by starport scaffolding # 3


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="eurx.pricefeed.QueryAllMarketResponse"></a>

### QueryAllMarketResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `markets` | [Market](#eurx.pricefeed.Market) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="eurx.pricefeed.QueryAllOracleRequest"></a>

### QueryAllOracleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `market_id` | [string](#string) |  |  |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="eurx.pricefeed.QueryAllOracleResponse"></a>

### QueryAllOracleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `oracles` | [string](#string) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="eurx.pricefeed.QueryAllPriceRequest"></a>

### QueryAllPriceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="eurx.pricefeed.QueryAllPriceResponse"></a>

### QueryAllPriceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `prices` | [CurrentPrice](#eurx.pricefeed.CurrentPrice) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="eurx.pricefeed.QueryAllRawPriceRequest"></a>

### QueryAllRawPriceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `market_id` | [string](#string) |  |  |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="eurx.pricefeed.QueryAllRawPriceResponse"></a>

### QueryAllRawPriceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `prices` | [PostedPrice](#eurx.pricefeed.PostedPrice) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="eurx.pricefeed.QueryGetPriceRequest"></a>

### QueryGetPriceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `market_id` | [string](#string) |  |  |






<a name="eurx.pricefeed.QueryGetPriceResponse"></a>

### QueryGetPriceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `price` | [CurrentPrice](#eurx.pricefeed.CurrentPrice) |  |  |






<a name="eurx.pricefeed.QueryParamsRequest"></a>

### QueryParamsRequest







<a name="eurx.pricefeed.QueryParamsResponse"></a>

### QueryParamsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#eurx.pricefeed.Params) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="eurx.pricefeed.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#eurx.pricefeed.QueryParamsRequest) | [QueryParamsResponse](#eurx.pricefeed.QueryParamsResponse) |  | GET|/eurx/pricefeed/params|
| `MarketAll` | [QueryAllMarketRequest](#eurx.pricefeed.QueryAllMarketRequest) | [QueryAllMarketResponse](#eurx.pricefeed.QueryAllMarketResponse) | this line is used by starport scaffolding # 2 | GET|/eurx/pricefeed/markets|
| `OracleAll` | [QueryAllOracleRequest](#eurx.pricefeed.QueryAllOracleRequest) | [QueryAllOracleResponse](#eurx.pricefeed.QueryAllOracleResponse) |  | GET|/eurx/pricefeed/markets/{market_id}/oracles|
| `Price` | [QueryGetPriceRequest](#eurx.pricefeed.QueryGetPriceRequest) | [QueryGetPriceResponse](#eurx.pricefeed.QueryGetPriceResponse) |  | GET|/eurx/pricefeed/markets/{market_id}/price|
| `PriceAll` | [QueryAllPriceRequest](#eurx.pricefeed.QueryAllPriceRequest) | [QueryAllPriceResponse](#eurx.pricefeed.QueryAllPriceResponse) |  | GET|/eurx/pricefeed/prices|
| `RawPriceAll` | [QueryAllRawPriceRequest](#eurx.pricefeed.QueryAllRawPriceRequest) | [QueryAllRawPriceResponse](#eurx.pricefeed.QueryAllRawPriceResponse) |  | GET|/eurx/pricefeed/markets/{market_id}/raw_prices|

 <!-- end services -->



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |
