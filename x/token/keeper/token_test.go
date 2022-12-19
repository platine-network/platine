package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/platine-network/platine/testutil/keeper"
	"github.com/platine-network/platine/testutil/nullify"
	"github.com/platine-network/platine/x/token/keeper"
	"github.com/platine-network/platine/x/token/types"
	"github.com/stretchr/testify/require"
)

func createNToken(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Token {
	items := make([]types.Token, n)
	for i := range items {
		items[i].Id = keeper.AppendToken(ctx, items[i])
	}
	return items
}

func TestTokenGet(t *testing.T) {
	keeper, ctx := keepertest.TokenKeeper(t)
	items := createNToken(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetToken(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestTokenRemove(t *testing.T) {
	keeper, ctx := keepertest.TokenKeeper(t)
	items := createNToken(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveToken(ctx, item.Id)
		_, found := keeper.GetToken(ctx, item.Id)
		require.False(t, found)
	}
}

func TestTokenGetAll(t *testing.T) {
	keeper, ctx := keepertest.TokenKeeper(t)
	items := createNToken(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllToken(ctx)),
	)
}

func TestTokenCount(t *testing.T) {
	keeper, ctx := keepertest.TokenKeeper(t)
	items := createNToken(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetTokenCount(ctx))
}
