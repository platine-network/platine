package keeper

import (
	"fmt"
	"github.com/tendermint/tendermint/libs/log"
	"github.com/platine-network/platine/x/treasury/types"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

type (
	Keeper struct {
		cdc      	codec.BinaryCodec
		storeKey 	storetypes.StoreKey
		paramSpace	paramtypes.Subspace
    accountKeeper types.AccountKeeper
    bankKeeper types.BankKeeper
    distrKeeper types.DistrKeeper
    epochKeeper types.EpochKeeper
    hooks types.MintHooks
    feeCollectorName string
	}
)

func NewKeeper(
    cdc codec.BinaryCodec, key storetypes.StoreKey, ps paramtypes.Subspace,
    ak types.AccountKeeper, bk types.BankKeeper, dk types.DistrKeeper, ek types.EpochKeeper, 
		feeCollectorName string,
) Keeper {
	// ensure mint module account is set
	if addr := ak.GetModuleAddress(types.ModuleName); addr == nil{
		panic("treasury module account is not set")
	}
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		cdc:      	cdc,
		storeKey: 	key,
		paramSpace:	ps,
    accountKeeper: ak,
    bankKeeper: bk,
    distrKeeper: dk,
    epochKeeper: ek,
    feeCollectorName: feeCollectorName,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k *Keeper) SetHooks(mh types.MintHooks) *Keeper{
	if k.hooks != nil {
		panic("Can not set mint hooks twice")
	}

	k.hooks = mh

	return k
}

func (k Keeper) GetLastReductionEpochNum(ctx sdk.Context) int64 {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.LastReductionEpochKey)
	if b == nil {
		return 0
	}

	return int64(sdk.BigEndianToUint64(b))
}

func (k Keeper) SetLastReductionEpochNum(ctx sdk.Context, epochNum int64) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.LastReductionEpochKey, sdk.Uint64ToBigEndian(uint64(epochNum)))
}

func (k Keeper) GetMinter(ctx sdk.Context) (minter types.Minter) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.MinterKey)
	if b == nil {
		panic("stored minter should be set")
	}
	k.cdc.MustUnmarshal(b, &minter)
	return
}

func (k Keeper) SetMinter(ctx sdk.Context, minter types.Minter) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshal(&minter)
	store.Set(types.MinterKey, b)
}

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramSpace.GetParamSet(ctx, &params)
	return params
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSpace.SetParamSet(ctx, &params)
}


func (k Keeper) MintCoins(ctx sdk.Context, newCoins sdk.Coins) error {
	if newCoins.Empty() {
		return nil
	}

	return k.bankKeeper.MintCoins(ctx, types.ModuleName, newCoins)
}


func (k Keeper) GetDistributions(ctx sdk.Context, mintedCoin sdk.Coin, ratio sdk.Dec) sdk.Coin {
	return sdk.NewCoin(mintedCoin.Denom, sdk.NewDecFromInt(mintedCoin.Amount).Mul(ratio).TruncateInt())
}

func (k Keeper) DistributeMintedCoin(ctx sdk.Context, mintedCoin sdk.Coin) error {
	params := k.GetParams(ctx)
	distributions := params.Distribution

	stakingCoins := sdk.NewCoins(k.GetDistributions(ctx, mintedCoin, distributions.StakingPool))
	err := k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, k.feeCollectorName, stakingCoins)
	if err != nil {
		return err
	}

	ecosystemCoins := sdk.NewCoins(k.GetDistributions(ctx, mintedCoin, distributions.EcosystemPool))
	if params.EcosystemPoolAddress == "" {
		err = k.distrKeeper.FundCommunityPool(ctx, ecosystemCoins, k.accountKeeper.GetModuleAddress(types.ModuleName))
		if err != nil {
			return err
		}
	} else {
		ecosystemPoolAddress, err := sdk.AccAddressFromBech32(params.EcosystemPoolAddress)
		if err != nil {
			return err
		}
		err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, ecosystemPoolAddress, ecosystemCoins)
		if err != nil {
			return err
		}
	}

	developerCoins := sdk.NewCoins(k.GetDistributions(ctx, mintedCoin, distributions.DeveloperPool))
	if params.DeveloperPoolAddress == "" {
		err = k.distrKeeper.FundCommunityPool(ctx, developerCoins, k.accountKeeper.GetModuleAddress(types.ModuleName))
		if err != nil {
			return err
		}
	} else {
		developerPoolAddress, err := sdk.AccAddressFromBech32(params.DeveloperPoolAddress)
		if err != nil {
			return err
		}
		err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, developerPoolAddress, developerCoins)
		if err != nil {
			return err
		}
	}

	rewardCoins := sdk.NewCoins(k.GetDistributions(ctx, mintedCoin, distributions.RewardPool))
	if params.RewardPoolAddress == "" {
		err = k.distrKeeper.FundCommunityPool(ctx, rewardCoins, k.accountKeeper.GetModuleAddress(types.ModuleName))
		if err != nil {
			return err
		}
	} else {
		rewardPoolAddress, err := sdk.AccAddressFromBech32(params.RewardPoolAddress)
		if err != nil {
			return err
		}
		err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, rewardPoolAddress, rewardCoins)
		if err != nil {
			return err
		}
	}

	/*
	communityCoins := sdk.NewCoins(mintedCoin).Sub(stakingCoins).Sub(ecosystemCoins).Sub(developerCoins).Sub(rewardCoins)
	err = k.distrKeeper.FundCommunityPool(ctx, communityCoins, k.accountKeeper.GetModuleAddress(types.ModuleName))
	if err != nil {
		return err
	}
	*/
	k.hooks.AfterDistributeMintedCoin(ctx, mintedCoin)

	return err
}