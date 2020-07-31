package keeper_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/lcnem/eurx/app"
	"github.com/lcnem/eurx/x/pricefeed/types"
)

// TestKeeper_SetGetMarket tests adding markets to the pricefeed, getting markets from the store
func TestKeeper_SetGetMarket(t *testing.T) {
	tApp := app.NewTestApp()
	ctx := tApp.NewContext(true, abci.Header{})
	keeper := tApp.GetPriceFeedKeeper()

	mp := types.Params{
		Markets: types.Markets{
			types.Market{MarketID: "tsteur", BaseAsset: "tst", QuoteAsset: "eur", Oracles: []sdk.AccAddress{}, Active: true},
		},
	}
	keeper.SetParams(ctx, mp)
	markets := keeper.GetMarkets(ctx)
	require.Equal(t, len(markets), 1)
	require.Equal(t, markets[0].MarketID, "tsteur")

	_, found := keeper.GetMarket(ctx, "tsteur")
	require.Equal(t, found, true)

	mp = types.Params{
		Markets: types.Markets{
			types.Market{MarketID: "tsteur", BaseAsset: "tst", QuoteAsset: "eur", Oracles: []sdk.AccAddress{}, Active: true},
			types.Market{MarketID: "tst2eur", BaseAsset: "tst2", QuoteAsset: "eur", Oracles: []sdk.AccAddress{}, Active: true},
		},
	}
	keeper.SetParams(ctx, mp)
	markets = keeper.GetMarkets(ctx)
	require.Equal(t, len(markets), 2)
	require.Equal(t, markets[0].MarketID, "tsteur")
	require.Equal(t, markets[1].MarketID, "tst2eur")

	_, found = keeper.GetMarket(ctx, "nan")
	require.Equal(t, found, false)
}

// TestKeeper_GetSetPrice Test Posting the price by an oracle
func TestKeeper_GetSetPrice(t *testing.T) {
	_, addrs := app.GeneratePrivKeyAddressPairs(2)
	tApp := app.NewTestApp()
	ctx := tApp.NewContext(true, abci.Header{})
	keeper := tApp.GetPriceFeedKeeper()

	mp := types.Params{
		Markets: types.Markets{
			types.Market{MarketID: "tsteur", BaseAsset: "tst", QuoteAsset: "eur", Oracles: []sdk.AccAddress{}, Active: true},
		},
	}
	keeper.SetParams(ctx, mp)
	// Set price by oracle 1
	_, err := keeper.SetPrice(
		ctx, addrs[0], "tsteur",
		sdk.MustNewDecFromStr("0.33"),
		time.Now().Add(1*time.Hour))
	require.NoError(t, err)
	// Get raw prices
	rawPrices, err := keeper.GetRawPrices(ctx, "tsteur")
	require.NoError(t, err)
	require.Equal(t, len(rawPrices), 1)
	require.Equal(t, rawPrices[0].Price.Equal(sdk.MustNewDecFromStr("0.33")), true)
	// Set price by oracle 2
	_, err = keeper.SetPrice(
		ctx, addrs[1], "tsteur",
		sdk.MustNewDecFromStr("0.35"),
		time.Now().Add(time.Hour*1))
	require.NoError(t, err)

	rawPrices, err = keeper.GetRawPrices(ctx, "tsteur")
	require.NoError(t, err)
	require.Equal(t, len(rawPrices), 2)
	require.Equal(t, rawPrices[1].Price.Equal(sdk.MustNewDecFromStr("0.35")), true)

	// Update Price by Oracle 1
	_, err = keeper.SetPrice(
		ctx, addrs[0], "tsteur",
		sdk.MustNewDecFromStr("0.37"),
		time.Now().Add(time.Hour*1))
	require.NoError(t, err)
	rawPrices, err = keeper.GetRawPrices(ctx, "tsteur")
	require.NoError(t, err)
	require.Equal(t, rawPrices[0].Price.Equal(sdk.MustNewDecFromStr("0.37")), true)
}

// TestKeeper_GetSetCurrentPrice Test Setting the median price of an Asset
func TestKeeper_GetSetCurrentPrice(t *testing.T) {
	_, addrs := app.GeneratePrivKeyAddressPairs(4)
	tApp := app.NewTestApp()
	ctx := tApp.NewContext(true, abci.Header{})
	keeper := tApp.GetPriceFeedKeeper()

	mp := types.Params{
		Markets: types.Markets{
			types.Market{MarketID: "tsteur", BaseAsset: "tst", QuoteAsset: "eur", Oracles: []sdk.AccAddress{}, Active: true},
		},
	}
	keeper.SetParams(ctx, mp)
	keeper.SetPrice(
		ctx, addrs[0], "tsteur",
		sdk.MustNewDecFromStr("0.33"),
		time.Now().Add(time.Hour*1))
	keeper.SetPrice(
		ctx, addrs[1], "tsteur",
		sdk.MustNewDecFromStr("0.35"),
		time.Now().Add(time.Hour*1))
	keeper.SetPrice(
		ctx, addrs[2], "tsteur",
		sdk.MustNewDecFromStr("0.34"),
		time.Now().Add(time.Hour*1))
	// Set current price
	err := keeper.SetCurrentPrices(ctx, "tsteur")
	require.NoError(t, err)
	// Get Current price
	price, err := keeper.GetCurrentPrice(ctx, "tsteur")
	require.Nil(t, err)
	require.Equal(t, price.Price.Equal(sdk.MustNewDecFromStr("0.34")), true)

	// Even number of oracles
	keeper.SetPrice(
		ctx, addrs[3], "tsteur",
		sdk.MustNewDecFromStr("0.36"),
		time.Now().Add(time.Hour*1))
	err = keeper.SetCurrentPrices(ctx, "tsteur")
	require.NoError(t, err)
	price, err = keeper.GetCurrentPrice(ctx, "tsteur")
	require.Nil(t, err)
	require.Equal(t, price.Price.Equal(sdk.MustNewDecFromStr("0.345")), true)
}
