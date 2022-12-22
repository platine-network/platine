package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ModuleName defines the module name
	ModuleName = "token"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_token"
)

const (
	TokenKey      = "TokenKey"
	TokenOwnerKey = "TokenOwnerKey"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

func TokenKeyPrefix(tokenId string) []byte {
	var key []byte
	tokenIdBytes := []byte(tokenId)
	key = append(key, tokenIdBytes...)
	key = append(key, []byte("/")...)

	return key
}
