package types

const (
	// ModuleName defines the module name
	ModuleName = "treasury"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_treasury"
)

const (
	AttributeEpochNumber    = "epoch_number"
	AttributeEpochProvision = "epoch_provision"
)

var MinterKey = []byte{0x00}
var LastReductionEpochKey = []byte{0x03}

func KeyPrefix(p string) []byte {
	return []byte(p)
}
