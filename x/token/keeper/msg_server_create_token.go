package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/platine-network/platine/x/token/types"
)

func (k msgServer) CreateToken(goCtx context.Context, msg *types.MsgCreateToken) (*types.MsgCreateTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateTokenResponse{}, nil
}
