package codec

import (
	"fmt"
	"math/big"
	"testing"
)

func TestStructs(t *testing.T) {

	bytes, err := NewBytes("0x0c00")
	if err != nil {
		t.Error(err)
	}
	prefs, err := bytes.ToValidatorPrefsLegacy()
	if err != nil {
		t.Error(err)
	}
	if prefs.UnstakeThreshold != 3 {
		t.Errorf("%v != %v", prefs.UnstakeThreshold, 3)
	}

	bigPayment := big.Int(prefs.ValidatorPayment)
	if big.NewInt(0).Cmp(&bigPayment) != 0 {
		t.Errorf("%v != %v", prefs.UnstakeThreshold, 0)
	}

	fmt.Println(prefs.ValidatorPayment)

}
