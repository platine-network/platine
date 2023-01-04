package types

import (
	"errors"
	"fmt"
	"strings"

	epochtypes "github.com/platine-network/platine/x/epoch/types"
	yaml "gopkg.in/yaml.v2"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var (
	KeyMintDenom                    = []byte("MintDenom")
	KeyGenesisEpochProvision        = []byte("GenesisEpochProvision")
	KeyEpochIdentifier              = []byte("EpochIdentifier")
	KeyReductionPeriodEpochs        = []byte("ReductionPeriodEpochs")
	KeyReductionFactor              = []byte("ReductionFactor")
	KeyAllocationRatio              = []byte("AllocationRatio")
	KeyRewardDistributionStartEpoch = []byte("RewardDistributionStartEpoch")
	KeyEcosystemPoolAddress         = []byte("EcosystemPoolAddress")
	KeyDeveloperPoolAddress         = []byte("DeveloperPoolAddress")
	KeyRewardPoolAddress            = []byte("RewardPoolAddress")
)

func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

func NewParams(
	mintDenom string, genesisEpochProvision sdk.Dec, epochIdentifier string, reductionFactor sdk.Dec,
	reductionPeriodEpochs int64, distribution Distribution, rewardDistributionStartEpoch int64,
	ecosystemPoolAddress string, developerPoolAddress string, rewardPoolAddress string,
) Params {
	return Params{
		MintDenom:                    mintDenom,
		GenesisEpochProvision:        genesisEpochProvision,
		EpochIdentifier:              epochIdentifier,
		ReductionPeriodEpochs:        reductionPeriodEpochs,
		ReductionFactor:              reductionFactor,
		Distribution:                 distribution,
		RewardDistributionStartEpoch: rewardDistributionStartEpoch,
		EcosystemPoolAddress:         ecosystemPoolAddress,
		DeveloperPoolAddress:         developerPoolAddress,
		RewardPoolAddress:            rewardPoolAddress,
	}
}

func DefaultParams() Params {
	return Params{
		MintDenom:             "uplc",
		GenesisEpochProvision: sdk.NewDec(1000000000),
		EpochIdentifier:       "minute",
		ReductionPeriodEpochs: 6,
		ReductionFactor:       sdk.NewDecWithPrec(666666666666666666, 18), //2/3
		Distribution: Distribution{
			StakingPool:   sdk.NewDecWithPrec(4, 1), // 0.4
			EcosystemPool: sdk.NewDecWithPrec(2, 1), // 0.2
			DeveloperPool: sdk.NewDecWithPrec(2, 1), // 0.2
			RewardPool:    sdk.NewDecWithPrec(1, 1), // 0.1
			CommunityPool: sdk.NewDecWithPrec(1, 1), //0.1
		},
		RewardDistributionStartEpoch: 1,
		EcosystemPoolAddress:         "",
		DeveloperPoolAddress:         "",
		RewardPoolAddress:            "",
	}
}

func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyMintDenom, &p.MintDenom, validateMintDenom),
		paramtypes.NewParamSetPair(KeyGenesisEpochProvision, &p.GenesisEpochProvision, validateGenesisEpochProvision),
		paramtypes.NewParamSetPair(KeyEpochIdentifier, &p.EpochIdentifier, epochtypes.ValidateEpochIdentifierInterface),
		paramtypes.NewParamSetPair(KeyReductionPeriodEpochs, &p.ReductionPeriodEpochs, validateReductionPeriodEpochs),
		paramtypes.NewParamSetPair(KeyReductionFactor, &p.ReductionFactor, validateReductionFactor),
		paramtypes.NewParamSetPair(KeyAllocationRatio, &p.Distribution, validateDistribution),
		paramtypes.NewParamSetPair(KeyRewardDistributionStartEpoch, &p.RewardDistributionStartEpoch, validateRewardDistributionStartEpoch),
		paramtypes.NewParamSetPair(KeyEcosystemPoolAddress, &p.EcosystemPoolAddress, validateBech32Address),
		paramtypes.NewParamSetPair(KeyDeveloperPoolAddress, &p.DeveloperPoolAddress, validateBech32Address),
		paramtypes.NewParamSetPair(KeyRewardPoolAddress, &p.RewardPoolAddress, validateBech32Address),
	}
}

func (p Params) Validate() error {
	if err := validateMintDenom(p.MintDenom); err != nil {
		return err
	}

	if err := validateGenesisEpochProvision(p.GenesisEpochProvision); err != nil {
		return err
	}

	if err := epochtypes.ValidateEpochIdentifierInterface(p.EpochIdentifier); err != nil {
		return err
	}

	if err := validateReductionPeriodEpochs(p.ReductionPeriodEpochs); err != nil {
		return err
	}

	if err := validateReductionFactor(p.ReductionFactor); err != nil {
		return err
	}

	if err := validateDistribution(p.Distribution); err != nil {
		return err
	}

	if err := validateRewardDistributionStartEpoch(p.RewardDistributionStartEpoch); err != nil {
		return err
	}

	if err := validateBech32Address(p.EcosystemPoolAddress); err != nil {
		return err
	}

	if err := validateBech32Address(p.DeveloperPoolAddress); err != nil {
		return err
	}

	if err := validateBech32Address(p.RewardPoolAddress); err != nil {
		return err
	}

	return nil
}

func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

func validateBech32Address(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type %T", i)
	}

	if v != "" {
		_, err := sdk.AccAddressFromBech32(v)
		if err != nil {
			return fmt.Errorf("invalid address %v", i)
		}
	}

	return nil
}

func validateRewardDistributionStartEpoch(i interface{}) error {
	v, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type %T", i)
	}

	if v < 0 {
		return fmt.Errorf("reward distribution start epoch must be positive, got: %d", v)
	}

	return nil
}

func validateGenesisEpochProvision(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type %T", i)
	}

	if v.LT(sdk.ZeroDec()) {
		return fmt.Errorf("genesis epoch provision must be positive")
	}
	return nil
}

func validateReductionFactor(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type %T", i)
	}

	if v.GT(sdk.NewDec(1)) {
		return fmt.Errorf("reduction factor can not be greather than 1")
	}

	if v.IsNegative() {
		return fmt.Errorf("reduction factor must be positive")
	}

	return nil
}

func validateDistribution(i interface{}) error {
	v, ok := i.(Distribution)
	if !ok {
		return fmt.Errorf("invalid parameter type %T", i)
	}

	if v.StakingPool.IsNegative() {
		return fmt.Errorf("staking pool ratio must be positive")
	}

	if v.EcosystemPool.IsNegative() {
		return fmt.Errorf("ecosystem pool ratio must be positive")
	}

	if v.DeveloperPool.IsNegative() {
		return fmt.Errorf("developer pool ratio must be positive")
	}

	if v.RewardPool.IsNegative() {
		return fmt.Errorf("reward pool ratio must be positive")
	}

	if v.CommunityPool.IsNegative() {
		return fmt.Errorf("community pool ratio must be positive")
	}

	total := v.StakingPool.Add(v.EcosystemPool).Add(v.DeveloperPool).Add(v.RewardPool).Add(v.CommunityPool)

	if !total.Equal(sdk.NewDec(1)) {
		return errors.New("total distribution ratio should equal to 1")
	}

	return nil
}

func validateReductionPeriodEpochs(i interface{}) error {
	v, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type %T", i)
	}

	if v <= 0 {
		return fmt.Errorf("reduction period in epoch must be positive, got: %d", v)
	}
	return nil
}

func validateMintDenom(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type %T", i)
	}

	if strings.TrimSpace(v) == "" {
		return errors.New("mint denom can not be blank")
	}

	if err := sdk.ValidateDenom(v); err != nil {
		return err
	}

	return nil
}
