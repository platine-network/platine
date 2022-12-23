package treasury_test

import (
	"testing"

	keepertest "github.com/platine-network/platine/testutil/keeper"
	"github.com/platine-network/platine/testutil/nullify"
	"github.com/platine-network/platine/x/treasury"
	"github.com/platine-network/platine/x/treasury/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params:	types.DefaultParams(),
		
		MinterList: []types.Minter{
		{
			Id: 0,
		},
		{
			Id: 1,
		},
	},
	MinterCount: 2,
	// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TreasuryKeeper(t)
	treasury.InitGenesis(ctx, *k, genesisState)
	got := treasury.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	

	require.ElementsMatch(t, genesisState.MinterList, got.MinterList)
require.Equal(t, genesisState.MinterCount, got.MinterCount)
// this line is used by starport scaffolding # genesis/test/assert
}
