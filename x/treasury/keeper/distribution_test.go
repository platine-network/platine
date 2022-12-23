package keeper_test

import (
	"testing"

    "github.com/platine-network/platine/x/treasury/keeper"
    "github.com/platine-network/platine/x/treasury/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/platine-network/platine/testutil/keeper"
	"github.com/platine-network/platine/testutil/nullify"
	"github.com/stretchr/testify/require"
)

func createNDistribution(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Distribution {
	items := make([]types.Distribution, n)
	for i := range items {
		items[i].Id = keeper.AppendDistribution(ctx, items[i])
	}
	return items
}

func TestDistributionGet(t *testing.T) {
	keeper, ctx := keepertest.TreasuryKeeper(t)
	items := createNDistribution(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetDistribution(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestDistributionRemove(t *testing.T) {
	keeper, ctx := keepertest.TreasuryKeeper(t)
	items := createNDistribution(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveDistribution(ctx, item.Id)
		_, found := keeper.GetDistribution(ctx, item.Id)
		require.False(t, found)
	}
}

func TestDistributionGetAll(t *testing.T) {
	keeper, ctx := keepertest.TreasuryKeeper(t)
	items := createNDistribution(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllDistribution(ctx)),
	)
}

func TestDistributionCount(t *testing.T) {
	keeper, ctx := keepertest.TreasuryKeeper(t)
	items := createNDistribution(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetDistributionCount(ctx))
}
