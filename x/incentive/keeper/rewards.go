package keeper

import (
	"math"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	cdptypes "github.com/lcnem/eurx/x/cdp/types"
	"github.com/lcnem/eurx/x/incentive/types"
)

// AccumulateEURXMintingRewards updates the rewards accumulated for the input reward period
func (k Keeper) AccumulateEURXMintingRewards(ctx sdk.Context, rewardPeriod types.RewardPeriod) error {
	previousAccrualTime, found := k.GetPreviousEURXMintingAccrualTime(ctx, rewardPeriod.CollateralType)
	if !found {
		k.SetPreviousEURXMintingAccrualTime(ctx, rewardPeriod.CollateralType, ctx.BlockTime())
		return nil
	}
	timeElapsed := CalculateTimeElapsed(rewardPeriod.Start, rewardPeriod.End, ctx.BlockTime(), previousAccrualTime)
	if timeElapsed.IsZero() {
		return nil
	}
	if rewardPeriod.RewardsPerSecond.Amount.IsZero() {
		k.SetPreviousEURXMintingAccrualTime(ctx, rewardPeriod.CollateralType, ctx.BlockTime())
		return nil
	}
	totalPrincipal := k.cdpKeeper.GetTotalPrincipal(ctx, rewardPeriod.CollateralType, types.PrincipalDenom).ToDec()
	if totalPrincipal.IsZero() {
		k.SetPreviousEURXMintingAccrualTime(ctx, rewardPeriod.CollateralType, ctx.BlockTime())
		return nil
	}
	newRewards := timeElapsed.Mul(rewardPeriod.RewardsPerSecond.Amount)
	cdpFactor, found := k.cdpKeeper.GetInterestFactor(ctx, rewardPeriod.CollateralType)
	if !found {
		k.SetPreviousEURXMintingAccrualTime(ctx, rewardPeriod.CollateralType, ctx.BlockTime())
		return nil
	}
	rewardFactor := newRewards.ToDec().Mul(cdpFactor).Quo(totalPrincipal)

	previousRewardFactor, found := k.GetEURXMintingRewardFactor(ctx, rewardPeriod.CollateralType)
	if !found {
		previousRewardFactor = sdk.ZeroDec()
	}
	newRewardFactor := previousRewardFactor.Add(rewardFactor)
	k.SetEURXMintingRewardFactor(ctx, rewardPeriod.CollateralType, newRewardFactor)
	k.SetPreviousEURXMintingAccrualTime(ctx, rewardPeriod.CollateralType, ctx.BlockTime())
	return nil
}

// InitializeEURXMintingClaim creates or updates a claim such that no new rewards are accrued, but any existing rewards are not lost.
// this function should be called after a cdp is created. If a user previously had a cdp, then closed it, they shouldn't
// accrue rewards during the period the cdp was closed. By setting the reward factor to the current global reward factor,
// any unclaimed rewards are preserved, but no new rewards are added.
func (k Keeper) InitializeEURXMintingClaim(ctx sdk.Context, cdp cdptypes.CDP) {
	_, found := k.GetEURXMintingRewardPeriod(ctx, cdp.Type)
	if !found {
		// this collateral type is not incentivized, do nothing
		return
	}
	rewardFactor, found := k.GetEURXMintingRewardFactor(ctx, cdp.Type)
	if !found {
		rewardFactor = sdk.ZeroDec()
	}
	claim, found := k.GetEURXMintingClaim(ctx, cdp.Owner.AccAddress())
	if !found { // this is the owner's first eurx minting reward claim
		claim = types.NewEURXMintingClaim(cdp.Owner.AccAddress(), sdk.NewCoin(types.EURXMintingRewardDenom, sdk.ZeroInt()), types.RewardIndexes{types.NewRewardIndex(cdp.Type, rewardFactor)})
		k.SetEURXMintingClaim(ctx, claim)
		return
	}
	// the owner has an existing eurx minting reward claim
	index, hasRewardIndex := claim.HasRewardIndex(cdp.Type)
	if !hasRewardIndex { // this is the owner's first eurx minting reward for this collateral type
		claim.RewardIndexes = append(claim.RewardIndexes, types.NewRewardIndex(cdp.Type, rewardFactor))
	} else { // the owner has a previous eurx minting reward for this collateral type
		claim.RewardIndexes[index] = types.NewRewardIndex(cdp.Type, rewardFactor)
	}
	k.SetEURXMintingClaim(ctx, claim)
}

// SynchronizeEURXMintingReward updates the claim object by adding any accumulated rewards and updating the reward index value.
// this should be called before a cdp is modified, immediately after the 'SynchronizeInterest' method is called in the cdp module
func (k Keeper) SynchronizeEURXMintingReward(ctx sdk.Context, cdp cdptypes.CDP) {
	_, found := k.GetEURXMintingRewardPeriod(ctx, cdp.Type)
	if !found {
		// this collateral type is not incentivized, do nothing
		return
	}

	globalRewardFactor, found := k.GetEURXMintingRewardFactor(ctx, cdp.Type)
	if !found {
		globalRewardFactor = sdk.ZeroDec()
	}
	claim, found := k.GetEURXMintingClaim(ctx, cdp.Owner.AccAddress())
	if !found {
		claim = types.NewEURXMintingClaim(cdp.Owner.AccAddress(), sdk.NewCoin(types.EURXMintingRewardDenom, sdk.ZeroInt()), types.RewardIndexes{types.NewRewardIndex(cdp.Type, globalRewardFactor)})
		k.SetEURXMintingClaim(ctx, claim)
		return
	}

	// the owner has an existing eurx minting reward claim
	index, hasRewardIndex := claim.HasRewardIndex(cdp.Type)
	if !hasRewardIndex { // this is the owner's first eurx minting reward for this collateral type
		claim.RewardIndexes = append(claim.RewardIndexes, types.NewRewardIndex(cdp.Type, globalRewardFactor))
		k.SetEURXMintingClaim(ctx, claim)
		return
	}
	userRewardFactor := claim.RewardIndexes[index].RewardFactor
	rewardsAccumulatedFactor := globalRewardFactor.Sub(userRewardFactor)
	if rewardsAccumulatedFactor.IsZero() {
		return
	}
	claim.RewardIndexes[index].RewardFactor = globalRewardFactor
	newRewardsAmount := rewardsAccumulatedFactor.Mul(cdp.GetTotalPrincipal().Amount.ToDec()).RoundInt()
	if newRewardsAmount.IsZero() {
		k.SetEURXMintingClaim(ctx, claim)
		return
	}
	newRewardsCoin := sdk.NewCoin(types.EURXMintingRewardDenom, newRewardsAmount)
	claim.Reward = claim.Reward.Add(newRewardsCoin)
	k.SetEURXMintingClaim(ctx, claim)
	return
}

// ZeroEURXMintingClaim zeroes out the claim object's rewards and returns the updated claim object
func (k Keeper) ZeroEURXMintingClaim(ctx sdk.Context, claim types.EURXMintingClaim) types.EURXMintingClaim {
	claim.Reward = sdk.NewCoin(claim.Reward.Denom, sdk.ZeroInt())
	k.SetEURXMintingClaim(ctx, claim)
	return claim
}

// SynchronizeEURXMintingClaim updates the claim object by adding any rewards that have accumulated.
// Returns the updated claim object
func (k Keeper) SynchronizeEURXMintingClaim(ctx sdk.Context, claim types.EURXMintingClaim) (types.EURXMintingClaim, error) {
	for _, ri := range claim.RewardIndexes {
		cdp, found := k.cdpKeeper.GetCdpByOwnerAndCollateralType(ctx, claim.Owner.AccAddress(), ri.CollateralType)
		if !found {
			// if the cdp for this collateral type has been closed, no updates are needed
			continue
		}
		claim = k.synchronizeRewardAndReturnClaim(ctx, cdp)
	}
	return claim, nil
}

// this function assumes a claim already exists, so don't call it if that's not the case
func (k Keeper) synchronizeRewardAndReturnClaim(ctx sdk.Context, cdp cdptypes.CDP) types.EURXMintingClaim {
	k.SynchronizeEURXMintingReward(ctx, cdp)
	claim, _ := k.GetEURXMintingClaim(ctx, cdp.Owner.AccAddress())
	return claim
}

// CalculateTimeElapsed calculates the number of reward-eligible seconds that have passed since the previous
// time rewards were accrued, taking into account the end time of the reward period
func CalculateTimeElapsed(start, end, blockTime time.Time, previousAccrualTime time.Time) sdk.Int {
	if (end.Before(blockTime) &&
		(end.Before(previousAccrualTime) || end.Equal(previousAccrualTime))) ||
		(start.After(blockTime)) ||
		(start.Equal(blockTime)) {
		return sdk.ZeroInt()
	}
	if start.After(previousAccrualTime) && start.Before(blockTime) {
		previousAccrualTime = start
	}

	if end.Before(blockTime) {
		return sdk.MaxInt(sdk.ZeroInt(), sdk.NewInt(int64(math.RoundToEven(
			end.Sub(previousAccrualTime).Seconds(),
		))))
	}
	return sdk.MaxInt(sdk.ZeroInt(), sdk.NewInt(int64(math.RoundToEven(
		blockTime.Sub(previousAccrualTime).Seconds(),
	))))
}

// SimulateEURXMintingSynchronization calculates a user's outstanding EURX minting rewards by simulating reward synchronization
func (k Keeper) SimulateEURXMintingSynchronization(ctx sdk.Context, claim types.EURXMintingClaim) types.EURXMintingClaim {
	for _, ri := range claim.RewardIndexes {
		_, found := k.GetEURXMintingRewardPeriod(ctx, ri.CollateralType)
		if !found {
			continue
		}

		globalRewardFactor, found := k.GetEURXMintingRewardFactor(ctx, ri.CollateralType)
		if !found {
			globalRewardFactor = sdk.ZeroDec()
		}

		// the owner has an existing eurx minting reward claim
		index, hasRewardIndex := claim.HasRewardIndex(ri.CollateralType)
		if !hasRewardIndex { // this is the owner's first eurx minting reward for this collateral type
			claim.RewardIndexes = append(claim.RewardIndexes, types.NewRewardIndex(ri.CollateralType, globalRewardFactor))
		}
		userRewardFactor := claim.RewardIndexes[index].RewardFactor
		rewardsAccumulatedFactor := globalRewardFactor.Sub(userRewardFactor)
		if rewardsAccumulatedFactor.IsZero() {
			continue
		}

		claim.RewardIndexes[index].RewardFactor = globalRewardFactor

		cdp, found := k.cdpKeeper.GetCdpByOwnerAndCollateralType(ctx, claim.GetOwner(), ri.CollateralType)
		if !found {
			continue
		}
		newRewardsAmount := rewardsAccumulatedFactor.Mul(cdp.GetTotalPrincipal().Amount.ToDec()).RoundInt()
		if newRewardsAmount.IsZero() {
			continue
		}
		newRewardsCoin := sdk.NewCoin(types.EURXMintingRewardDenom, newRewardsAmount)
		claim.Reward = claim.Reward.Add(newRewardsCoin)
	}

	return claim
}
