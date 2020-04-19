package ethtool

import (
	"math/big"
)

type ValueUnit int64

const (
	Kwei  ValueUnit = 1e3
	Mwei  ValueUnit = 1e6
	Gwei  ValueUnit = 1e9
	Micon ValueUnit = 1e12
	Milli ValueUnit = 1e15
	Ether ValueUnit = 1e18
)

func ToWei(value *big.Float, unit ValueUnit) *big.Int {
	d := big.NewFloat(float64(unit))
	o, _ := new(big.Float).Mul(value, d).Int(new(big.Int))
	return o
}

func FromWei(value *big.Int, unit ValueUnit) *big.Float {
	v := new(big.Float).SetInt(value)
	d := big.NewFloat(float64(unit))
	return new(big.Float).Quo(v, d)
}
