package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/platine-network/platine/x/token/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) TokenList(c context.Context, req *types.QueryTokenListRequest) (*types.QueryTokenListResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TokenOwnerKey))
	accountStore := prefix.NewStore(store, []byte(req.Owner))
	iterator := accountStore.Iterator(nil, nil)

	defer iterator.Close()

	var tokens []string
	for ; iterator.Valid(); iterator.Next(){
		tokenId := string(iterator.Key()[:])
		tokens = append(tokens, tokenId)
	}


	return &types.QueryTokenListResponse{Token: tokens}, nil
}

