package types
import (
	"errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	errNilEpochProvision = errors.New("epoch provision was nil in genesis")
	errNegativeEpochProvision = errors.New("epoch provision must be positive")
)

func NewMinter(epochProvision sdk.Dec) Minter {
	return Minter {
		Provision: epochProvision,
	}
}

func InitialMinter() Minter {
	return NewMinter(sdk.NewDec(0))
}

func DefaultInitialMinter() Minter {
	return InitialMinter()
}

func (m Minter) Validate() error {
	if m.Provision.IsNil() {
		return errNilEpochProvision
	}

	if m.Provision.IsNegative() {
		return errNegativeEpochProvision
	}

	return nil
}

func (m Minter) NextEpochProvision(params Params) sdk.Dec {
	return m.Provision.Mul(params.ReductionFactor)
}

func (m Minter) EpochProvision(params Params) sdk.Coin {
	provisionAmount := m.Provision
	return sdk.NewCoin(params.MintDenom, provisionAmount.TruncateInt())
}