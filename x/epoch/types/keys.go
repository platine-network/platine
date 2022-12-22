package types

const (
	// ModuleName defines the module name
	ModuleName = "epoch"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_epoch"
)

const (
	EventTypeEpochEnd = "epoch_end"
	EventTypeEpochStart = "epoch_start"

	AttributeEpochNumber = "epoch_number"
	AttributeEpochStartTime = "start_time"
)

var KeyPrefixEpoch = []byte{0x01}

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	EpochKey      = "Epoch/value/"
	EpochCountKey = "Epoch/count/"
)
