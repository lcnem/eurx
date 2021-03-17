package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lcnem/eurx/x/estmdist/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Balances(c context.Context, req *types.QueryGetBalancesRequest) (*types.QueryGetBalancesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var balances sdk.Coins
	ctx := sdk.UnwrapSDKContext(c)

	acc := k.accountKeeper.GetModuleAccount(ctx, types.EstmdistMacc)
	balances = k.bankKeeper.GetAllBalances(ctx, acc.GetAddress())

	return &types.QueryGetBalancesResponse{Balances: balances}, nil
}
