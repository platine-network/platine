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

func createNMinter(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Minter {
	items := make([]types.Minter, n)
	for i := range items {
		items[i].Id = keeper.AppendMinter(ctx, items[i])
	}
	return items
}

func TestMinterGet(t *testing.T) {
	keeper, ctx := keepertest.TreasuryKeeper(t)
	items := createNMinter(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetMinter(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestMinterRemove(t *testing.T) {
	keeper, ctx := keepertest.TreasuryKeeper(t)
	items := createNMinter(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveMinter(ctx, item.Id)
		_, found := keeper.GetMinter(ctx, item.Id)
		require.False(t, found)
	}
}

func TestMinterGetAll(t *testing.T) {
	keeper, ctx := keepertest.TreasuryKeeper(t)
	items := createNMinter(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllMinter(ctx)),
	)
}

func TestMinterCount(t *testing.T) {
	keeper, ctx := keepertest.TreasuryKeeper(t)
	items := createNMinter(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetMinterCount(ctx))
}
