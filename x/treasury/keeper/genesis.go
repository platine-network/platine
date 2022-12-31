package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/platine-network/platine/x/treasury/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func (k Keeper) InitGenesis(ctx sdk.Context, data *types.GenesisState) {
	if data == nil {
		panic("genesis state is nil")
	}

	if data.Minter.Provision.IsZero() || data.Minter.Provision.IsNil() {
		data.Minter.Provision = data.Params.GenesisEpochProvision
	}
	k.SetMinter(ctx, data.Minter)
	k.SetParams(ctx, data.Params)

	// The call to GetModuleAccount creates a module account if it does not exist.
	k.accountKeeper.GetModuleAccount(ctx, types.ModuleName)
	k.SetLastReductionEpochNum(ctx, data.ReductionStartEpoch)
}

// ExportGenesis returns the module's exported genesis
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	minter := k.GetMinter(ctx)
	params := k.GetParams(ctx)

	lastHalvenEpoch := k.GetLastReductionEpochNum(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return types.NewGenesisState(minter, params, lastHalvenEpoch)
}
