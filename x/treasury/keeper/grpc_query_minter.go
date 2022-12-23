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

func (k Keeper) MinterAll(c context.Context, req *types.QueryAllMinterRequest) (*types.QueryAllMinterResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var minters []types.Minter
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	minterStore := prefix.NewStore(store, types.KeyPrefix(types.MinterKey))

	pageRes, err := query.Paginate(minterStore, req.Pagination, func(key []byte, value []byte) error {
		var minter types.Minter
		if err := k.cdc.Unmarshal(value, &minter); err != nil {
			return err
		}

		minters = append(minters, minter)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllMinterResponse{Minter: minters, Pagination: pageRes}, nil
}

func (k Keeper) Minter(c context.Context, req *types.QueryGetMinterRequest) (*types.QueryGetMinterResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	minter, found := k.GetMinter(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetMinterResponse{Minter: minter}, nil
}
