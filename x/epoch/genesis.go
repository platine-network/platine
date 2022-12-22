package epoch

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/platine-network/platine/x/epoch/keeper"
	"github.com/platine-network/platine/x/epoch/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the epoch
	for _, elem := range genState.Epochs {
		err := k.AddEpoch(ctx, elem)
		if err != nil {
			panic(err)
		}
	}
	// this line is used by starport scaffolding # genesis/module/init
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	genesis.Epochs = k.AllEpoch(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
