package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/platine-network/platine/x/treasury/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) DistributionAll(c context.Context, req *types.QueryAllDistributionRequest) (*types.QueryAllDistributionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var distributions []types.Distribution
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	distributionStore := prefix.NewStore(store, types.KeyPrefix(types.DistributionKey))

	pageRes, err := query.Paginate(distributionStore, req.Pagination, func(key []byte, value []byte) error {
		var distribution types.Distribution
		if err := k.cdc.Unmarshal(value, &distribution); err != nil {
			return err
		}

		distributions = append(distributions, distribution)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllDistributionResponse{Distribution: distributions, Pagination: pageRes}, nil
}

func (k Keeper) Distribution(c context.Context, req *types.QueryGetDistributionRequest) (*types.QueryGetDistributionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	distribution, found := k.GetDistribution(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetDistributionResponse{Distribution: distribution}, nil
}
