package types

func NewGenesisState(minter Minter, params Params, reductionStartEpoch int64) *GenesisState {
	return &GenesisState{
		Minter: minter,
		Params: params,
		ReductionStartEpoch: reductionStartEpoch,
	}
}

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		ReductionStartEpoch: 0,
		Minter: DefaultInitialMinter(),
// this line is used by starport scaffolding # genesis/types/default
	    Params:	DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
  if err := gs.Params.Validate(); err != nil {
		return err
	}

	return gs.Minter.Validate()
}
