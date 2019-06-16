package config

import (
	"bytes"
	"fmt"

	"github.com/tichex-project/hackatom-cosmos-tichex/x/issue/msgs"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
)

func init() {
	//msgs.MsgCdc.RegisterConcrete(Params{}, "issue/Params", nil)
}

var (
	// key for constant fee parameter
	ParamStoreKeyIssueFee         = []byte("IssueFee")
	ParamStoreKeyMintFee          = []byte("MintFee")
	ParamStoreKeyFreezeFee        = []byte("FreezeFee")
	ParamStoreKeyUnFreezeFee      = []byte("UnfreezeFee")
	ParamStoreKeyBurnFee          = []byte("BurnFee")
	ParamStoreKeyBurnFromFee      = []byte("BurnFromFee")
	ParamStoreKeyTransferOwnerFee = []byte("TransferOwnerFee")
	ParamStoreKeyDescribeFee      = []byte("DescribeFee")
)

var _ params.ParamSet = &Params{}

// Param Config issue for issue
type Params struct {
	IssueFee         sdk.Coin `json:"issue_fee"`
	MintFee          sdk.Coin `json:"mint_fee"`
	FreezeFee        sdk.Coin `json:"freeze_fee"`
	UnFreezeFee      sdk.Coin `json:"unfreeze_fee"`
	BurnFee          sdk.Coin `json:"burn_fee"`
	BurnFromFee      sdk.Coin `json:"burn_from_fee"`
	TransferOwnerFee sdk.Coin `json:"transfer_owner_fee"`
	DescribeFee      sdk.Coin `json:"describe_fee"`
}

// ParamKeyTable for auth module
func ParamKeyTable() params.KeyTable {
	return params.NewKeyTable().RegisterParamSet(&Params{})
}

// ParamSetPairs implements the ParamSet interface and returns all the key/value pairs
// pairs of auth module's parameters.
// nolint
func (p *Params) ParamSetPairs() params.ParamSetPairs {
	return params.ParamSetPairs{
		{ParamStoreKeyIssueFee, &p.IssueFee},
		{ParamStoreKeyMintFee, &p.MintFee},
		{ParamStoreKeyFreezeFee, &p.FreezeFee},
		{ParamStoreKeyUnFreezeFee, &p.UnFreezeFee},
		{ParamStoreKeyBurnFee, &p.BurnFee},
		{ParamStoreKeyBurnFromFee, &p.BurnFromFee},
		{ParamStoreKeyTransferOwnerFee, &p.TransferOwnerFee},
		{ParamStoreKeyDescribeFee, &p.DescribeFee},
	}
}

// Checks equality of Params
func (dp Params) Equal(dp2 Params) bool {
	b1 := msgs.MsgCdc.MustMarshalBinaryBare(dp)
	b2 := msgs.MsgCdc.MustMarshalBinaryBare(dp2)
	return bytes.Equal(b1, b2)
}

// DefaultParams returns a default set of parameters.
func DefaultParams(denom string) Params {
	return Params{
		IssueFee:         sdk.NewCoin(denom, sdk.NewIntWithDecimal(20000, 8)),
		MintFee:          sdk.NewCoin(denom, sdk.NewIntWithDecimal(10000, 8)),
		FreezeFee:        sdk.NewCoin(denom, sdk.NewIntWithDecimal(20000, 8)),
		UnFreezeFee:      sdk.NewCoin(denom, sdk.NewIntWithDecimal(20000, 8)),
		BurnFee:          sdk.NewCoin(denom, sdk.NewIntWithDecimal(10000, 8)),
		BurnFromFee:      sdk.NewCoin(denom, sdk.NewIntWithDecimal(10000, 8)),
		TransferOwnerFee: sdk.NewCoin(denom, sdk.NewIntWithDecimal(20000, 8)),
		DescribeFee:      sdk.NewCoin(denom, sdk.NewIntWithDecimal(4000, 8)),
	}
}

func (dp Params) String() string {
	return fmt.Sprintf(`Params:
  IssueFee:			%s
  MintFee:			%s
  FreezeFee:			%s
  UnFreezeFee:			%s
  BurnFee:			%s
  BurnFromFee:			%s
  TransferOwnerFee:			%s
  DescribeFee:			%s`,
		dp.IssueFee.String(),
		dp.MintFee.String(),
		dp.FreezeFee.String(),
		dp.UnFreezeFee.String(),
		dp.BurnFee.String(),
		dp.BurnFromFee.String(),
		dp.TransferOwnerFee.String(),
		dp.DescribeFee.String(),
	)
}