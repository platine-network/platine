package keeper

import (
	"github.com/platine-network/platine/x/token/types"
)

var _ types.QueryServer = Keeper{}
