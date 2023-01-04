package keeper

import (
	"context"
	"errors"

	"github.com/platine-network/platine/x/epoch/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) EpochAll(c context.Context, req *types.QueryAllEpochRequest) (*types.QueryAllEpochResponse, error) {
	var epochs []types.Epoch
	ctx := sdk.UnwrapSDKContext(c)
	epochs = k.AllEpoch(ctx)

	return &types.QueryAllEpochResponse{Epochs: epochs}, nil
}

func (k Keeper) CurrentEpoch(c context.Context, req *types.QueryCurrentEpochRequest) (*types.QueryCurrentEpochResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	epoch := k.GetEpoch(ctx, req.Identifier)
	if epoch.Identifier != req.Identifier {
		return nil, errors.New("epoch with identifier not found")
	}

	return &types.QueryCurrentEpochResponse{CurrentEpoch: epoch.CurrentEpoch}, nil
}
