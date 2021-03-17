package keeper

import (
	"fmt"
	"time"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/lcnem/eurx/x/incentive/types"
)

type (
	Keeper struct {
		cdc           codec.Marshaler
		storeKey      sdk.StoreKey
		memKey        sdk.StoreKey
		paramSpace    paramtypes.Subspace
		accountKeeper types.AccountKeeper
		bankKeeper    types.BankKeeper
		cdpKeeper     types.CdpKeeper
	}
)

func NewKeeper(cdc codec.Marshaler, storeKey, memKey sdk.StoreKey,
	paramSpace paramtypes.Subspace, accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
	cdpKeeper types.CdpKeeper) Keeper {
	return Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		memKey:        memKey,
		paramSpace:    paramSpace,
		accountKeeper: accountKeeper,
		bankKeeper:    bankKeeper,
		cdpKeeper:     cdpKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// GetEURXMintingClaim returns the claim in the store corresponding the the input address collateral type and id and a boolean for if the claim was found
func (k Keeper) GetEURXMintingClaim(ctx sdk.Context, addr sdk.AccAddress) (types.EURXMintingClaim, bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.EURXMintingClaimKeyPrefix)
	bz := store.Get(addr)
	if bz == nil {
		return types.EURXMintingClaim{}, false
	}
	var c types.EURXMintingClaim
	k.cdc.MustUnmarshalBinaryBare(bz, &c)
	return c, true
}

// SetEURXMintingClaim sets the claim in the store corresponding to the input address, collateral type, and id
func (k Keeper) SetEURXMintingClaim(ctx sdk.Context, c types.EURXMintingClaim) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.EURXMintingClaimKeyPrefix)
	bz := k.cdc.MustMarshalBinaryBare(&c)
	store.Set(c.Owner, bz)

}

// DeleteEURXMintingClaim deletes the claim in the store corresponding to the input address, collateral type, and id
func (k Keeper) DeleteEURXMintingClaim(ctx sdk.Context, owner sdk.AccAddress) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.EURXMintingClaimKeyPrefix)
	store.Delete(owner)
}

// IterateEURXMintingClaims iterates over all claim  objects in the store and preforms a callback function
func (k Keeper) IterateEURXMintingClaims(ctx sdk.Context, cb func(c types.EURXMintingClaim) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.EURXMintingClaimKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var c types.EURXMintingClaim
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &c)
		if cb(c) {
			break
		}
	}
}

// GetAllEURXMintingClaims returns all Claim objects in the store
func (k Keeper) GetAllEURXMintingClaims(ctx sdk.Context) types.EURXMintingClaims {
	cs := types.EURXMintingClaims{}
	k.IterateEURXMintingClaims(ctx, func(c types.EURXMintingClaim) (stop bool) {
		cs = append(cs, c)
		return false
	})
	return cs
}

// GetPreviousEURXMintingAccrualTime returns the last time a collateral type accrued EURX minting rewards
func (k Keeper) GetPreviousEURXMintingAccrualTime(ctx sdk.Context, ctype string) (blockTime time.Time, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PreviousEURXMintingRewardAccrualTimeKeyPrefix)
	bz := store.Get([]byte(ctype))
	if bz == nil {
		return time.Time{}, false
	}
	blockTime.UnmarshalBinary(bz)

	return blockTime, true
}

// SetPreviousEURXMintingAccrualTime sets the last time a collateral type accrued EURX minting rewards
func (k Keeper) SetPreviousEURXMintingAccrualTime(ctx sdk.Context, ctype string, blockTime time.Time) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PreviousEURXMintingRewardAccrualTimeKeyPrefix)
	bz, _ := blockTime.MarshalBinary()
	store.Set([]byte(ctype), bz)
}

// IterateEURXMintingAccrualTimes iterates over all previous EURX minting accrual times and preforms a callback function
func (k Keeper) IterateEURXMintingAccrualTimes(ctx sdk.Context, cb func(string, time.Time) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PreviousEURXMintingRewardAccrualTimeKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var accrualTime time.Time
		var collateralType string
		collateralType = string(iterator.Value())
		accrualTime.UnmarshalBinary(iterator.Value())
		if cb(collateralType, accrualTime) {
			break
		}
	}
}

// GetEURXMintingRewardFactor returns the current reward factor for an individual collateral type
func (k Keeper) GetEURXMintingRewardFactor(ctx sdk.Context, ctype string) (factor sdk.Dec, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.EURXMintingRewardFactorKeyPrefix)
	bz := store.Get([]byte(ctype))
	if bz == nil {
		return sdk.ZeroDec(), false
	}
	factor.Unmarshal(bz)

	return factor, true
}

// SetEURXMintingRewardFactor sets the current reward factor for an individual collateral type
func (k Keeper) SetEURXMintingRewardFactor(ctx sdk.Context, ctype string, factor sdk.Dec) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.EURXMintingRewardFactorKeyPrefix)
	bz, _ := factor.Marshal()
	store.Set([]byte(ctype), bz)
}
