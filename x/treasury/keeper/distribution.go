package keeper

import (
	"encoding/binary"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/platine-network/platine/x/treasury/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
)

// GetDistributionCount get the total number of distribution
func (k Keeper) GetDistributionCount(ctx sdk.Context) uint64 {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.DistributionCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetDistributionCount set the total number of distribution
func (k Keeper) SetDistributionCount(ctx sdk.Context, count uint64)  {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.DistributionCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendDistribution appends a distribution in the store with a new id and update the count
func (k Keeper) AppendDistribution(
    ctx sdk.Context,
    distribution types.Distribution,
) uint64 {
	// Create the distribution
    count := k.GetDistributionCount(ctx)

    // Set the ID of the appended value
    distribution.Id = count

    store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DistributionKey))
    appendedValue := k.cdc.MustMarshal(&distribution)
    store.Set(GetDistributionIDBytes(distribution.Id), appendedValue)

    // Update distribution count
    k.SetDistributionCount(ctx, count+1)

    return count
}

// SetDistribution set a specific distribution in the store
func (k Keeper) SetDistribution(ctx sdk.Context, distribution types.Distribution) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DistributionKey))
	b := k.cdc.MustMarshal(&distribution)
	store.Set(GetDistributionIDBytes(distribution.Id), b)
}

// GetDistribution returns a distribution from its id
func (k Keeper) GetDistribution(ctx sdk.Context, id uint64) (val types.Distribution, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DistributionKey))
	b := store.Get(GetDistributionIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveDistribution removes a distribution from the store
func (k Keeper) RemoveDistribution(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DistributionKey))
	store.Delete(GetDistributionIDBytes(id))
}

// GetAllDistribution returns all distribution
func (k Keeper) GetAllDistribution(ctx sdk.Context) (list []types.Distribution) {
    store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DistributionKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Distribution
		k.cdc.MustUnmarshal(iterator.Value(), &val)
        list = append(list, val)
	}

    return
}

// GetDistributionIDBytes returns the byte representation of the ID
func GetDistributionIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetDistributionIDFromBytes returns ID in uint64 format from a byte array
func GetDistributionIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
