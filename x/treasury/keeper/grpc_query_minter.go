package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/platine-network/platine/x/treasury/types"
)

func (k Keeper) EpochProvision(c context.Context, _ *types.QueryEpochProvisionRequest) (*types.QueryEpochProvisionResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	minter := k.GetMinter(ctx)

	return &types.QueryEpochProvisionResponse{EpochProvision: minter.Provision}, nil
}
