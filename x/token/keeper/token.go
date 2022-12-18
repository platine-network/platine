package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/platine-network/platine/x/token/types"
)

// GetTokenCount get the total number of token
func (k Keeper) GetTokenCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.TokenCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetTokenCount set the total number of token
func (k Keeper) SetTokenCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.TokenCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendToken appends a token in the store with a new id and update the count
func (k Keeper) AppendToken(
	ctx sdk.Context,
	token types.Token,
) uint64 {
	// Create the token
	count := k.GetTokenCount(ctx)

	// Set the ID of the appended value
	token.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TokenKey))
	appendedValue := k.cdc.MustMarshal(&token)
	store.Set(GetTokenIDBytes(token.Id), appendedValue)

	// Update token count
	k.SetTokenCount(ctx, count+1)

	return count
}

// SetToken set a specific token in the store
func (k Keeper) SetToken(ctx sdk.Context, token types.Token) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TokenKey))
	b := k.cdc.MustMarshal(&token)
	store.Set(GetTokenIDBytes(token.Id), b)
}

// GetToken returns a token from its id
func (k Keeper) GetToken(ctx sdk.Context, id uint64) (val types.Token, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TokenKey))
	b := store.Get(GetTokenIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveToken removes a token from the store
func (k Keeper) RemoveToken(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TokenKey))
	store.Delete(GetTokenIDBytes(id))
}

// GetAllToken returns all token
func (k Keeper) GetAllToken(ctx sdk.Context) (list []types.Token) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TokenKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Token
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetTokenIDBytes returns the byte representation of the ID
func GetTokenIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetTokenIDFromBytes returns ID in uint64 format from a byte array
func GetTokenIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
