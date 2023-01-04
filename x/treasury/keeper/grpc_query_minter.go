package keeper

import (
	"context"

	"github.com/platine-network/platine/x/treasury/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) EpochProvision(c context.Context, _ *types.QueryEpochProvisionRequest) (*types.QueryEpochProvisionResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	minter := k.GetMinter(ctx)

	return &types.QueryEpochProvisionResponse{EpochProvision: minter.Provision}, nil
}
