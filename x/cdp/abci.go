package cdp

import (
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/lcnem/eurx/x/cdp/keeper"
	"github.com/lcnem/eurx/x/pricefeed/types"
)

// BeginBlocker compounds the debt in outstanding cdps and liquidates cdps that are below the required collateralization ratio
func BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock, k keeper.Keeper) {
	params := k.GetParams(ctx)

	for _, cp := range params.CollateralParams {
		ok := k.UpdatePricefeedStatus(ctx, cp.SpotMarketId)
		if !ok {
			continue
		}

		ok = k.UpdatePricefeedStatus(ctx, cp.LiquidationMarketId)
		if !ok {
			continue
		}

		err := k.AccumulateInterest(ctx, cp.Type)
		if err != nil {
			panic(err)
		}

		err = k.SynchronizeInterestForRiskyCdps(ctx, cp.CheckCollateralizationIndexCount, sdk.MaxSortableDec, cp.Type)
		if err != nil {
			panic(err)
		}

		err = k.LiquidateCdps(ctx, cp.LiquidationMarketId, cp.Type, cp.LiquidationRatio, cp.CheckCollateralizationIndexCount)
		if err != nil && !errors.Is(err, types.ErrNoValidPrice) {
			panic(err)
		}
	}

	err := k.RunSurplusAndDebtAuctions(ctx)
	if err != nil {
		panic(err)
	}
}
