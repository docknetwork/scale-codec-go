package base

import (
	"bytes"
	"fmt"
	"testing"
)

func ExampleDecode() {
	value, _, _ := Decode("Compact<u32>", "0xc15d", false)
	fmt.Println(value)
	// Output: 6000
}

func TestDecodeCompactUInt32(t *testing.T) {
	testTable := map[string]uint32{
		"0x02093d00": 1000000,
		"0x18":       6,
		"0xc15d":     6000,
	}

	for h, expected := range testTable {
		i, _, err := Decode("Compact<u32>", h, false)
		if err != nil {
			t.Error(err)
		} else if i != expected {
			t.Error(fmt.Errorf("%v != %v", i, expected))
		}
	}
}

func TestDecodeBool(t *testing.T) {
	b, _, err := Decode("Bool", "0x01", false)
	if err != nil {
		t.Error(err)
	} else if b != true {
		t.Error(fmt.Errorf("%b != %v", b, true))
	}

	b, _, err = Decode("Bool", "0x00", false)
	if err != nil {
		t.Error(err)
	} else if b != false {
		t.Error(fmt.Errorf("%b != %v", b, false))
	}
}

func TestDecodeVec(t *testing.T) {
	value, _, err := Decode("Bytes", "0xc15d", false)
	if err != nil {
		t.Error(err)
	}
	slice, err := ToSliceOfUint8(value)
	if err != nil {
		t.Error(err)
	}
	expected := []uint8{193, 93}

	if bytes.Compare(expected, slice) != 0 {
		t.Errorf("%v != %v", expected, slice)
	}
}

func TestDecodeOption(t *testing.T) {
	value, isNull, err := Decode("Option<u32>", "0x0118", false)
	if err != nil {
		t.Error(err)
	}
	if isNull {
		t.Errorf("should not be null")
	}
	var expected uint32 = 24
	if value != expected {
		t.Errorf("%v != %v", value, expected)
	}

	_, isNull, err = Decode("Option<u32>", "0x00", false)
	if err != nil {
		t.Error(err)
	}
	if !isNull {
		t.Errorf("should be null")
	}
}
