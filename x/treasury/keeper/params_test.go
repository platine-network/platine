package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/platine-network/platine/testutil/keeper"
	"github.com/platine-network/platine/x/treasury/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.TreasuryKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
