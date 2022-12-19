package keeper_test

import (
	"testing"

	testkeeper "github.com/platine-network/platine/testutil/keeper"
	"github.com/platine-network/platine/x/token/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.TokenKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
