package treasury

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/platine-network/platine/x/treasury/keeper"
	"github.com/platine-network/platine/x/treasury/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
    // Set all the minter
for _, elem := range genState.MinterList {
	k.SetMinter(ctx, elem)
}

// Set minter count
k.SetMinterCount(ctx, genState.MinterCount)
// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

    genesis.MinterList = k.GetAllMinter(ctx)
genesis.MinterCount = k.GetMinterCount(ctx)
// this line is used by starport scaffolding # genesis/module/export

    return genesis
}
