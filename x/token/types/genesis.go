package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		TokenList: []Token{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in token
	tokenIdMap := make(map[uint64]bool)
	tokenCount := gs.GetTokenCount()
	for _, elem := range gs.TokenList {
		if _, ok := tokenIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for token")
		}
		if elem.Id >= tokenCount {
			return fmt.Errorf("token id should be lower or equal than the last id")
		}
		tokenIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
