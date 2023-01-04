package keeper

import (
	"fmt"

	epochtypes "github.com/platine-network/platine/x/epoch/types"
	"github.com/platine-network/platine/x/treasury/types"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) BeforeEpochStart(ctx sdk.Context, identifier string, epochNumber int64) {

}

func (k Keeper) AfterEpochEnd(ctx sdk.Context, identifier string, epochNumber int64) {
	params := k.GetParams(ctx)

	if identifier == params.EpochIdentifier {
		if epochNumber < params.RewardDistributionStartEpoch {
			return
		} else if epochNumber == params.RewardDistributionStartEpoch {
			k.SetLastReductionEpochNum(ctx, epochNumber)
		}

		minter := k.GetMinter(ctx)
		params := k.GetParams(ctx)

		// Check if we have hit an epoch where we update the inflation parameter.
		// We measure time between reductions in number of epochs.
		// This avoids issues with measuring in block numbers, as epochs have fixed intervals, with very
		// low variance at the relevant sizes. As a result, it is safe to store the epoch number
		// of the last reduction to be later retrieved for comparison.
		if epochNumber >= k.GetParams(ctx).ReductionPeriodEpochs+k.GetLastReductionEpochNum(ctx) {
			// Reduce the reward per reduction period
			minter.Provision = minter.NextEpochProvision(params)
			k.SetMinter(ctx, minter)
			k.SetLastReductionEpochNum(ctx, epochNumber)
		}

		// mint coins, update supply
		mintedCoin := minter.EpochProvision(params)
		mintedCoins := sdk.NewCoins(mintedCoin)

		err := k.MintCoins(ctx, mintedCoins)
		if err != nil {
			panic(err)
		}

		if mintedCoin.Amount.IsInt64() {
			defer telemetry.ModuleSetGauge(types.ModuleName, float32(mintedCoin.Amount.Int64()), "minted_tokens")
		}

		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.ModuleName,
				sdk.NewAttribute(types.AttributeEpochNumber, fmt.Sprintf("%d", epochNumber)),
				sdk.NewAttribute(types.AttributeEpochProvision, minter.Provision.String()),
				sdk.NewAttribute(sdk.AttributeKeyAmount, mintedCoin.Amount.String()),
			),
		)
	}
}

type Hooks struct {
	k Keeper
}

var _ epochtypes.EpochHooks = Hooks{}

func (k Keeper) Hooks() Hooks {
	return Hooks{k}
}

func (h Hooks) AfterEpochEnd(ctx sdk.Context, identifier string, epochNumber int64) error {
	h.k.AfterEpochEnd(ctx, identifier, epochNumber)
	return nil
}

func (h Hooks) BeforeEpochStart(ctx sdk.Context, identifier string, epochNumber int64) error {
	h.k.BeforeEpochStart(ctx, identifier, epochNumber)

	return nil
}
