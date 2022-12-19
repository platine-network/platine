package token_test

import (
	"testing"

	keepertest "github.com/platine-network/platine/testutil/keeper"
	"github.com/platine-network/platine/testutil/nullify"
	"github.com/platine-network/platine/x/token"
	"github.com/platine-network/platine/x/token/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		TokenList: []types.Token{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		TokenCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TokenKeeper(t)
	token.InitGenesis(ctx, *k, genesisState)
	got := token.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.TokenList, got.TokenList)
	require.Equal(t, genesisState.TokenCount, got.TokenCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
