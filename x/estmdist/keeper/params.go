package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/lcnem/eurx/x/estmdist/types"
)

func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramSpace.GetParamSet(ctx, &params)
	return params
}

func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSpace.SetParamSet(ctx, &params)
}