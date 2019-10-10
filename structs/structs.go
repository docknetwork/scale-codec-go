package structs

import (
	"math/big"
	"scale/base"
)

type KeyValue struct {
	Key   []uint8
	Value []uint8
}

//NewKeyValue ... (Vec<u8>, Vec<u8>)
func NewKeyValue(data interface{}) (res KeyValue, err error) {
	bytes, err := base.NewBytes(data)
	if err != nil {
		return
	}
	key, err := bytes.ToVecUint8()
	if err != nil {
		return
	}
	value, err := bytes.ToVecUint8()
	if err != nil {
		return
	}

	res = KeyValue{
		Key:   key,
		Value: value,
	}
	err = bytes.Check()
	return
}

type ValidatorPrefs struct {
	ValidatorPayment big.Int
}

//NewValidatorPrefs ... (Compact<Balance>)
func NewValidatorPrefs(data interface{}) (res ValidatorPrefs, err error) {
	bytes, err := base.NewBytes(data)
	if err != nil {
		return
	}

	value, err := bytes.ToCompactBigInt()
	if err != nil {
		return
	}

	res = ValidatorPrefs{
		value,
	}

	err = bytes.Check()
	return
}

type ValidatorPrefsLegacy struct {
	UnstakeThreshold uint32
	ValidatorPayment big.Int
}

//NewValidatorPrefsLegacy ... (Compact<u32>,Compact<Balance>)
func NewValidatorPrefsLegacy(data interface{}) (res ValidatorPrefsLegacy, err error) {
	bytes, err := base.NewBytes(data)
	if err != nil {
		return
	}

	unstakeThreshold, err := bytes.ToCompactUInt32()
	if err != nil {
		return
	}

	validatorPayment, err := bytes.ToCompactBigInt()
	if err != nil {
		return
	}

	res = ValidatorPrefsLegacy{
		UnstakeThreshold: unstakeThreshold,
		ValidatorPayment: validatorPayment,
	}

	err = bytes.Check()
	return
}
