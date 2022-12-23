package types

import (
	"errors"
	"time"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

func NewGenesisState(epochs []Epoch) *GenesisState {
	return &GenesisState{Epochs: epochs}
}

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
		epochs := []Epoch{
			NewGenesisEpoch("minute", time.Minute),
			NewGenesisEpoch("day", time.Hour * 24),
			NewGenesisEpoch("hour", time.Hour),
			NewGenesisEpoch("week", time.Hour * 24 * 7),
		}

		return NewGenesisState(epochs)
		// this line is used by starport scaffolding # genesis/types/default
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in epoch
	epochIdentifiers := make(map[string]bool)
	for _, epoch := range gs.Epochs {
		if err := epoch.Validate(); err != nil {
			return err
		}
		if epochIdentifiers[epoch.Identifier] {
			return errors.New("epoch identifier should be unique")
		}
		epochIdentifiers[epoch.Identifier] = true
	}

	return nil
	// this line is used by starport scaffolding # genesis/types/validate
}

func (epoch Epoch) Validate() error {
	if epoch.Identifier == "" {
		return errors.New("epoch identifier should not be empty")
	}

	if epoch.Duration == 0 {
		return errors.New("epoch duration should not be 0")
	}

	if epoch.CurrentEpoch < 0 {
		return errors.New("currrent epoch should not be negative")
	}

	if epoch.CurrentEpochStartHeight < 0 {
		return errors.New("currrent epoch start height should not be negative")
	}

	return nil
}

func NewGenesisEpoch(identifier string, duration time.Duration) Epoch {
	return Epoch{
		Identifier: identifier,
		StartTime: time.Time{},
		Duration: duration,
		CurrentEpoch: 0,
		CurrentEpochStartHeight: 0,
		CurrentEpochStartTime: time.Time{},
		EpochCountingStarted: false,
	}
}
