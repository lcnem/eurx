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
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

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

// GetEurxMintingClaim returns the claim in the store corresponding the the input address collateral type and id and a boolean for if the claim was found
func (k Keeper) GetEurxMintingClaim(ctx sdk.Context, addr sdk.AccAddress) (types.EurxMintingClaim, bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.EurxMintingClaimKeyPrefix)
	bz := store.Get(addr)
	if bz == nil {
		return types.EurxMintingClaim{}, false
	}
	var c types.EurxMintingClaim
	k.cdc.MustUnmarshalBinaryBare(bz, &c)
	return c, true
}

// SetEurxMintingClaim sets the claim in the store corresponding to the input address, collateral type, and id
func (k Keeper) SetEurxMintingClaim(ctx sdk.Context, c types.EurxMintingClaim) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.EurxMintingClaimKeyPrefix)
	bz := k.cdc.MustMarshalBinaryBare(&c)
	store.Set(c.Owner, bz)

}

// DeleteEurxMintingClaim deletes the claim in the store corresponding to the input address, collateral type, and id
func (k Keeper) DeleteEurxMintingClaim(ctx sdk.Context, owner sdk.AccAddress) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.EurxMintingClaimKeyPrefix)
	store.Delete(owner)
}

// IterateEurxMintingClaims iterates over all claim  objects in the store and preforms a callback function
func (k Keeper) IterateEurxMintingClaims(ctx sdk.Context, cb func(c types.EurxMintingClaim) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.EurxMintingClaimKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var c types.EurxMintingClaim
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &c)
		if cb(c) {
			break
		}
	}
}

// GetAllEurxMintingClaims returns all Claim objects in the store
func (k Keeper) GetAllEurxMintingClaims(ctx sdk.Context) types.EurxMintingClaims {
	cs := types.EurxMintingClaims{}
	k.IterateEurxMintingClaims(ctx, func(c types.EurxMintingClaim) (stop bool) {
		cs = append(cs, c)
		return false
	})
	return cs
}

// GetPreviousEurxMintingAccrualTime returns the last time a collateral type accrued Eurx minting rewards
func (k Keeper) GetPreviousEurxMintingAccrualTime(ctx sdk.Context, ctype string) (blockTime time.Time, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PreviousEurxMintingRewardAccrualTimeKeyPrefix)
	bz := store.Get([]byte(ctype))
	if bz == nil {
		return time.Time{}, false
	}
	blockTime.UnmarshalBinary(bz)

	return blockTime, true
}

// SetPreviousEurxMintingAccrualTime sets the last time a collateral type accrued Eurx minting rewards
func (k Keeper) SetPreviousEurxMintingAccrualTime(ctx sdk.Context, ctype string, blockTime time.Time) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PreviousEurxMintingRewardAccrualTimeKeyPrefix)
	bz, _ := blockTime.MarshalBinary()
	store.Set([]byte(ctype), bz)
}

// IterateEurxMintingAccrualTimes iterates over all previous Eurx minting accrual times and preforms a callback function
func (k Keeper) IterateEurxMintingAccrualTimes(ctx sdk.Context, cb func(string, time.Time) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PreviousEurxMintingRewardAccrualTimeKeyPrefix)
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

// GetEurxMintingRewardFactor returns the current reward factor for an individual collateral type
func (k Keeper) GetEurxMintingRewardFactor(ctx sdk.Context, ctype string) (factor sdk.Dec, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.EurxMintingRewardFactorKeyPrefix)
	bz := store.Get([]byte(ctype))
	if bz == nil {
		return sdk.ZeroDec(), false
	}
	factor.Unmarshal(bz)

	return factor, true
}

// SetEurxMintingRewardFactor sets the current reward factor for an individual collateral type
func (k Keeper) SetEurxMintingRewardFactor(ctx sdk.Context, ctype string, factor sdk.Dec) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.EurxMintingRewardFactorKeyPrefix)
	bz, _ := factor.Marshal()
	store.Set([]byte(ctype), bz)
}
