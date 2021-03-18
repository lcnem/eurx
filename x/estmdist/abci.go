package estmdist

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lcnem/eurx/x/estmdist/keeper"
)

func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	err := k.MintPeriodInflation(ctx)
	if err != nil {
		panic(err)
	}
}
