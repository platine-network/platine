package types

import (
	"fmt"

	"github.com/platine-network/platine/utils"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type EpochHooks interface {
	AfterEpochEnd(ctx sdk.Context, identifier string, epochNumber int64) error
	BeforeEpochStart(ctx sdk.Context, identifier string, epochNumber int64) error
}
type MultiEpochHooks []EpochHooks

var _ EpochHooks = MultiEpochHooks{}

func NewMultiEpochHooks(hooks ...EpochHooks) MultiEpochHooks {
	return hooks
}

func (h MultiEpochHooks) AfterEpochEnd(ctx sdk.Context, identifier string, epochNumber int64) error {
	for i := range h {
		panicCatchingEpochHook(ctx, h[i].AfterEpochEnd, identifier, epochNumber)
	}

	return nil
}

func (h MultiEpochHooks) BeforeEpochStart(ctx sdk.Context, identifier string, epochNumber int64) error {
	for i := range h {
		panicCatchingEpochHook(ctx, h[i].BeforeEpochStart, identifier, epochNumber)
	}

	return nil
}

func panicCatchingEpochHook(
	ctx sdk.Context,
	hookFn func(ctx sdk.Context, identifier string, epochNumber int64) error,
	identifier string,
	epochNumber int64,
) {
	wrappedHookFn := func(ctx sdk.Context) error {
		return hookFn(ctx, identifier, epochNumber)
	}

	err := utils.ApplyFuncIfNoError(ctx, wrappedHookFn)
	if err != nil {
		ctx.Logger().Error(fmt.Sprintf("error in epoch hook %v", err))
	}
}
