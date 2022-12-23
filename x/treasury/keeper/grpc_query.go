package keeper

import (
	"github.com/platine-network/platine/x/treasury/types"
)

var _ types.QueryServer = Keeper{}
