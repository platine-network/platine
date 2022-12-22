package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/platine-network/platine/x/epoch/types"
	"github.com/tendermint/tendermint/libs/log"
)

type (
	Keeper struct {
		storeKey storetypes.StoreKey
		hooks   types.EpochHooks
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
) *Keeper {

	return &Keeper{
		storeKey: storeKey,
	}
}

func (k *Keeper) SetHooks(eh types.EpochHooks) *Keeper{
	if k.hooks != nil {
		panic("Can not set epoch hooks twice")
	}

	k.hooks = eh

	return k
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
