package keeper

import (
	"fmt"
	"time"

	"github.com/gogo/protobuf/proto"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/platine-network/platine/x/epoch/types"
)

// GetEpoch returns a epoch from its id
func (k Keeper) GetEpoch(ctx sdk.Context, id string) types.Epoch {
	epoch := types.Epoch{}
	store := ctx.KVStore(k.storeKey)
	b := store.Get(append(types.KeyPrefixEpoch, []byte(id)...))
	if b == nil {
		return epoch
	}
	err := proto.Unmarshal(b, &epoch)
	if err != nil {
		panic(err)
	}
	
	return epoch
}

// SetEpoch set a specific epoch in the store
func (k Keeper) AddEpoch(ctx sdk.Context, epoch types.Epoch) error {
	err := epoch.Validate()
	if err != nil {
		return err
	}

	if (k.GetEpoch(ctx, epoch.Identifier) != types.Epoch{}) {
		return fmt.Errorf("epoch with identifier %s already exists", epoch.Identifier)
	}

	if epoch.StartTime.Equal(time.Time{}) {
		epoch.StartTime = ctx.BlockTime()
		epoch.CurrentEpochStartHeight = ctx.BlockHeight()
	}
	
	k.SetEpoch(ctx, epoch)

	return nil
}

// SetEpoch set a specific epoch in the store
func (k Keeper) SetEpoch(ctx sdk.Context, epoch types.Epoch) {
	store := ctx.KVStore(k.storeKey)
	value, err := proto.Marshal(&epoch)
	if err != nil {
		panic(err)
	}
	store.Set(append(types.KeyPrefixEpoch, []byte(epoch.Identifier)...), value)
}

// DeleteEpoch removes a epoch from the store
func (k Keeper) DeleteEpoch(ctx sdk.Context, id string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(append(types.KeyPrefixEpoch, []byte(id)...))
}

func (k Keeper) IterateEpochs(ctx sdk.Context, fn func(index int64, epoch types.Epoch) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefixEpoch)
	defer iterator.Close()

	i:= int64(0)
	for ; iterator.Valid(); iterator.Next() {
		epoch := types.Epoch{}
		err := proto.Unmarshal(iterator.Value(), &epoch)
		if err != nil {
			panic(err)
		}

		stop := fn(i, epoch)
		if stop {
			break
		}
		i++
	}
}


// AllEpochs returns all epoch
func (k Keeper) AllEpoch(ctx sdk.Context) []types.Epoch {
	epochs := []types.Epoch{}
	k.IterateEpochs(ctx, func(intdex int64, epoch types.Epoch) (stop bool) {
		epochs = append(epochs, epoch)
		return false
	})

	return epochs
}

// TotalBlockAfterEpochStart returns the number of blocks since the epoch started.
// if the epoch started on block N, then calling this during block N (after BeforeEpochStart)
// would return 0.
// Calling it any point in block N+1 (assuming the epoch doesn't increment) would return 1.
func (k Keeper) TotalBlockAfterEpochStart(ctx sdk.Context, identifier string) (int64, error) {
	epoch := k.GetEpoch(ctx, identifier)
	if (epoch == types.Epoch{}) {
		return 0, fmt.Errorf("epoch with identifier %s does not exist", identifier)
	}

	return ctx.BlockHeight() - epoch.CurrentEpochStartHeight, nil
}