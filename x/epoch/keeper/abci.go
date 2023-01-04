package keeper

import (
	"fmt"
	"time"

	"github.com/platine-network/platine/x/epoch/types"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) BeginBlocker(ctx sdk.Context) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)
	k.IterateEpochs(ctx, func(index int64, epoch types.Epoch) (stop bool) {
		logger := k.Logger(ctx)

		// If blocktime < initial epoch start time, return
		if ctx.BlockTime().Before(epoch.StartTime) {
			return
		}

		// if epoch counting hasn't started, signal we need to start.
		needInitialEpochStart := !epoch.EpochCountingStarted
		epochEndTime := epoch.CurrentEpochStartTime.Add(epoch.Duration)
		needEpochStart := (ctx.BlockTime().After(epochEndTime)) || needInitialEpochStart

		if !needEpochStart {
			return false
		}

		epoch.CurrentEpochStartHeight = ctx.BlockHeight()

		if needInitialEpochStart {
			epoch.EpochCountingStarted = true
			epoch.CurrentEpoch = 1
			epoch.CurrentEpochStartTime = epoch.StartTime

			logger.Info(fmt.Sprintf("Starting new epoch with identifier %s epoch number %d", epoch.Identifier, epoch.CurrentEpoch))
		} else {
			ctx.EventManager().EmitEvent(
				sdk.NewEvent(
					types.EventTypeEpochEnd,
					sdk.NewAttribute(types.AttributeEpochNumber, fmt.Sprintf("%d", epoch.CurrentEpoch)),
				),
			)
			k.AfterEpochEnd(ctx, epoch.Identifier, epoch.CurrentEpoch)
			epoch.CurrentEpoch += 1
			epoch.CurrentEpochStartTime = epoch.CurrentEpochStartTime.Add(epoch.Duration)
			logger.Info(fmt.Sprintf("Starting epoch with identifier %s epoch number %d", epoch.Identifier, epoch.CurrentEpoch))
		}
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeEpochStart,
				sdk.NewAttribute(types.AttributeEpochNumber, fmt.Sprintf("%d", epoch.CurrentEpoch)),
				sdk.NewAttribute(types.AttributeEpochStartTime, fmt.Sprintf("%d", epoch.CurrentEpochStartTime.Unix())),
			),
		)

		k.SetEpoch(ctx, epoch)
		k.BeforeEpochStart(ctx, epoch.Identifier, epoch.CurrentEpoch)

		return false
	})
}
