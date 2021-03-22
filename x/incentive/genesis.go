package incentive

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lcnem/eurx/x/incentive/keeper"
	"github.com/lcnem/eurx/x/incentive/types"
)

// InitGenesis initializes the store state from a genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, accountKeeper types.AccountKeeper, cdpKeeper types.CdpKeeper, gs types.GenesisState) {

	// check if the module account exists
	moduleAcc := accountKeeper.GetModuleAccount(ctx, types.IncentiveMacc)
	if moduleAcc == nil {
		panic(fmt.Sprintf("%s module account has not been set", types.IncentiveMacc))
	}

	if err := gs.Validate(); err != nil {
		panic(fmt.Sprintf("failed to validate %s genesis state: %s", types.ModuleName, err))
	}

	for _, rp := range gs.Params.EurxMintingRewardPeriods {
		_, found := cdpKeeper.GetCollateral(ctx, rp.CollateralType)
		if !found {
			panic(fmt.Sprintf("eurx minting collateral type %s not found in cdp collateral types", rp.CollateralType))
		}
		k.SetEurxMintingRewardFactor(ctx, rp.CollateralType, sdk.ZeroDec())
	}

	k.SetParams(ctx, gs.Params)

	for _, gat := range gs.EurxAccumulationTimes {
		k.SetPreviousEurxMintingAccrualTime(ctx, gat.CollateralType, gat.PreviousAccumulationTime)
	}

	for i, claim := range gs.EurxMintingClaims {
		for j, ri := range claim.RewardIndexes {
			if ri.RewardFactor != sdk.ZeroDec() {
				gs.EurxMintingClaims[i].RewardIndexes[j].RewardFactor = sdk.ZeroDec()
			}
		}
		k.SetEurxMintingClaim(ctx, claim)
	}
}

// ExportGenesis export genesis state for incentive module
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types.GenesisState {
	params := k.GetParams(ctx)

	eurxClaims := k.GetAllEurxMintingClaims(ctx)

	synchronizedEurxClaims := types.EurxMintingClaims{}

	for _, eurxClaim := range eurxClaims {
		claim, err := k.SynchronizeEurxMintingClaim(ctx, eurxClaim)
		if err != nil {
			panic(err)
		}
		for i := range claim.RewardIndexes {
			claim.RewardIndexes[i].RewardFactor = sdk.ZeroDec()
		}
		synchronizedEurxClaims = append(synchronizedEurxClaims, claim)
	}

	var eurxMintingGats types.GenesisAccumulationTimes
	for _, rp := range params.EurxMintingRewardPeriods {
		pat, found := k.GetPreviousEurxMintingAccrualTime(ctx, rp.CollateralType)
		if !found {
			panic(fmt.Sprintf("expected previous eurx minting reward accrual time to be set in state for %s", rp.CollateralType))
		}
		gat := types.NewGenesisAccumulationTime(rp.CollateralType, pat)
		eurxMintingGats = append(eurxMintingGats, gat)
	}

	return types.NewGenesisState(params, eurxMintingGats, synchronizedEurxClaims)
}
