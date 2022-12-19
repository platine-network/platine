package token

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/platine-network/platine/x/token/keeper"
	"github.com/platine-network/platine/x/token/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the token
	for _, elem := range genState.TokenList {
		k.SetToken(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	genesis.TokenList = k.GetAllToken(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
