package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Parameter keys
var (
	ParamStoreKeySaleCommission       = []byte("salecommission")
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable returns the parameter key table.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// DefaultParams returns default marketplace parameters
func DefaultParams() Params {
	return Params{
		SaleCommission:        sdk.NewDecWithPrec(1, 2), // 2%
	}
}

// ParamSetPairs returns the parameter set pairs.
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(ParamStoreKeySaleCommission, &p.SaleCommission, validateSaleCommission),
	}
}

// ValidateBasic performs basic validation on marketplace parameters.
func (p Params) ValidateBasic() error {
	if p.SaleCommission.IsNegative() || p.SaleCommission.GT(sdk.OneDec()) {
		return fmt.Errorf(
			"sale commission should be non-negative and less than one: %s", p.SaleCommission,
		)
	}
	return nil
}

func validateSaleCommission(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() {
		return fmt.Errorf("sale commission must be not nil")
	}
	if v.IsNegative() {
		return fmt.Errorf("sale commission must be positive: %s", v)
	}
	if v.GT(sdk.OneDec()) {
		return fmt.Errorf("sale commission too large: %s", v)
	}

	return nil
}