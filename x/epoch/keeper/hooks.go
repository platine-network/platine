package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) AfterEpochEnd(ctx sdk.context, identifier string, epochNumber int64) {
	// Error is not handled as AfterEpochEnd Hooks use utils.ApplyFuncIfNoError()
	_ := k.hooks.AfterEpochEnd(ctx, identifier, epochNumber)
}

func (k Keeper) BeforeEpochStart(ctx sdk.context, identifier string, epochNumber int64) {
	// Error is not handled as BeforeEpochStart Hooks use utils.ApplyFuncIfNoError()
	_ := k.hooks.BeforeEpochStart(ctx, identifier, epochNumber)
}