package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	testkeeper "github.com/platine-network/platine/testutil/keeper"
	"github.com/platine-network/platine/x/token/types"
	"github.com/stretchr/testify/require"
)

func TestParamsQuery(t *testing.T) {
	keeper, ctx := testkeeper.TokenKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	params := types.DefaultParams()
	keeper.SetParams(ctx, params)

	response, err := keeper.Params(wctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
