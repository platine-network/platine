package keeper

import (
	"encoding/binary"
	"fmt"
	"regexp"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/platine-network/platine/x/token/types"
)

const (
	maxSymbolLength = 20
	maxNameLength   = 40
	maxTokenSupply  = 10000000000000000
	nativeTokenName = "uplc"
)

func (k Keeper) CheckCommonError(tokenId string, symbol string, name string, supply uint64) error {
	if len(symbol) > maxSymbolLength {
		errMsg := fmt.Sprintf("token symbol can not exceed %d caracters", maxSymbolLength)
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, errMsg)
	}

	if len(name) > maxNameLength {
		errMsg := fmt.Sprintf("token name can not exceed %d caracters", maxNameLength)
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, errMsg)
	}

	if tokenId == nativeTokenName {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "token name can not be same than native token")
	}

	re := regexp.MustCompile("^[a-zA-Z0-9_-]*$")
	if !re.MatchString(tokenId) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "only alpahabet, underscore, numbers are allowed for token id")
	}

	if supply > maxTokenSupply {
		errMsg := fmt.Sprintf("token total supply can not exceed %d", maxTokenSupply)
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, errMsg)
	}

	if supply <= 0 {
		errMsg := fmt.Sprintf("token total supply can not be less or equal to %d", 0)
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, errMsg)
	}

	return nil
}

// SetToken set a specific token in the store
func (k Keeper) SetToken(ctx sdk.Context, token types.Token) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TokenKey))
	b := k.cdc.MustMarshal(&token)
	store.Set(types.TokenKeyPrefix(token.Id), b)
	k.SetTokenToAccount(ctx, token)
}

func (k Keeper) SetTokenToAccount(ctx sdk.Context, token types.Token) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TokenOwnerKey))
	accountStore := prefix.NewStore(store, []byte(token.Owner))
	accountStore.Set([]byte(token.Id), GetBytesFromUint64(1))
}

// GetToken returns a token from its id
func (k Keeper) GetToken(ctx sdk.Context, id string) (val types.Token, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TokenKey))
	b := store.Get(types.TokenKeyPrefix(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveToken removes a token from the store
func (k Keeper) RemoveToken(ctx sdk.Context, id string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TokenKey))
	store.Delete(types.TokenKeyPrefix(id))
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

// GetBytesFromUint64 returns the byte representation of the ID
func GetBytesFromUint64(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetUInt64FromBytes returns ID in uint64 format from a byte array
func GetUInt64FromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
