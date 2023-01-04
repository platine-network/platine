package keeper_test

import (
	"testing"

	keepertest "github.com/platine-network/platine/testutil/keeper"
	"github.com/platine-network/platine/testutil/nullify"
	"github.com/platine-network/platine/x/epoch/keeper"
	"github.com/platine-network/platine/x/epoch/types"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func createNEpoch(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Epoch {
	items := make([]types.Epoch, n)
	for i := range items {
		items[i].Id = keeper.AppendEpoch(ctx, items[i])
	}
	return items
}

func TestEpochGet(t *testing.T) {
	keeper, ctx := keepertest.EpochKeeper(t)
	items := createNEpoch(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetEpoch(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestEpochRemove(t *testing.T) {
	keeper, ctx := keepertest.EpochKeeper(t)
	items := createNEpoch(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveEpoch(ctx, item.Id)
		_, found := keeper.GetEpoch(ctx, item.Id)
		require.False(t, found)
	}
}

func TestEpochGetAll(t *testing.T) {
	keeper, ctx := keepertest.EpochKeeper(t)
	items := createNEpoch(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllEpoch(ctx)),
	)
}

func TestEpochCount(t *testing.T) {
	keeper, ctx := keepertest.EpochKeeper(t)
	items := createNEpoch(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetEpochCount(ctx))
}
