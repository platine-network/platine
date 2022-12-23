package keeper

import (
	"encoding/binary"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/platine-network/platine/x/treasury/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
)

// GetMinterCount get the total number of minter
func (k Keeper) GetMinterCount(ctx sdk.Context) uint64 {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.MinterCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetMinterCount set the total number of minter
func (k Keeper) SetMinterCount(ctx sdk.Context, count uint64)  {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.MinterCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendMinter appends a minter in the store with a new id and update the count
func (k Keeper) AppendMinter(
    ctx sdk.Context,
    minter types.Minter,
) uint64 {
	// Create the minter
    count := k.GetMinterCount(ctx)

    // Set the ID of the appended value
    minter.Id = count

    store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MinterKey))
    appendedValue := k.cdc.MustMarshal(&minter)
    store.Set(GetMinterIDBytes(minter.Id), appendedValue)

    // Update minter count
    k.SetMinterCount(ctx, count+1)

    return count
}

// SetMinter set a specific minter in the store
func (k Keeper) SetMinter(ctx sdk.Context, minter types.Minter) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MinterKey))
	b := k.cdc.MustMarshal(&minter)
	store.Set(GetMinterIDBytes(minter.Id), b)
}

// GetMinter returns a minter from its id
func (k Keeper) GetMinter(ctx sdk.Context, id uint64) (val types.Minter, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MinterKey))
	b := store.Get(GetMinterIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveMinter removes a minter from the store
func (k Keeper) RemoveMinter(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MinterKey))
	store.Delete(GetMinterIDBytes(id))
}

// GetAllMinter returns all minter
func (k Keeper) GetAllMinter(ctx sdk.Context) (list []types.Minter) {
    store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MinterKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Minter
		k.cdc.MustUnmarshal(iterator.Value(), &val)
        list = append(list, val)
	}

    return
}

// GetMinterIDBytes returns the byte representation of the ID
func GetMinterIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetMinterIDFromBytes returns ID in uint64 format from a byte array
func GetMinterIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
