package types

import (
"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
	    MinterList: []Minter{},
// this line is used by starport scaffolding # genesis/types/default
	    Params:	DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
    // Check for duplicated ID in minter
minterIdMap := make(map[uint64]bool)
minterCount := gs.GetMinterCount()
for _, elem := range gs.MinterList {
	if _, ok := minterIdMap[elem.Id]; ok {
		return fmt.Errorf("duplicated id for minter")
	}
	if elem.Id >= minterCount {
		return fmt.Errorf("minter id should be lower or equal than the last id")
	}
	minterIdMap[elem.Id] = true
}
// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
