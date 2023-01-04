package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"

	keepertest "github.com/platine-network/platine/testutil/keeper"
	"github.com/platine-network/platine/testutil/nullify"
	"github.com/platine-network/platine/x/token/types"
)

func TestTokenQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.TokenKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNToken(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetTokenRequest
		response *types.QueryGetTokenResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetTokenRequest{Id: msgs[0].Id},
			response: &types.QueryGetTokenResponse{Token: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetTokenRequest{Id: msgs[1].Id},
			response: &types.QueryGetTokenResponse{Token: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetTokenRequest{Id: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Token(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestTokenQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.TokenKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNToken(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllTokenRequest {
		return &types.QueryAllTokenRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.TokenAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Token), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Token),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.TokenAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Token), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Token),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.TokenAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Token),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.TokenAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
