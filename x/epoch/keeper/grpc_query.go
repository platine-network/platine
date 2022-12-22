package keeper

import (
	"github.com/platine-network/platine/x/epoch/types"
)

var _ types.QueryServer = Keeper{}
