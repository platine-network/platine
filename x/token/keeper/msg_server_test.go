package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/platine-network/platine/testutil/keeper"
	"github.com/platine-network/platine/x/token/keeper"
	"github.com/platine-network/platine/x/token/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.TokenKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
