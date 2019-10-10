package base

import (
	"fmt"
	"testing"
)

func ExampleEncode() {
	h, _ := Encode("Compact<u32>", uint32(6000))
	fmt.Println(h)
	// Output: 0xc15d
}

func TestEncodeCompactUInt32(t *testing.T) {
	testTable := map[uint32]string{
		1000000: "0x02093d00",
		6:       "0x18",
		6000:    "0xc15d",
	}
	for i, expected := range testTable {
		h, err := Encode("Compact<u32>", i)
		if err != nil {
			t.Error(err)
		} else if expected != h {
			t.Error(fmt.Errorf("%v != %v", expected, h))
		}
	}
}

func TestEncodeBool(t *testing.T) {
	b, err := Encode("Bool", true)
	if err != nil {
		t.Error(err)
	} else if b != "0x01" {
		t.Error(fmt.Errorf("%v != %v", b, true))
	}

	b, err = Encode("Bool", false)
	if err != nil {
		t.Error(err)
	} else if b != "0x00" {
		t.Error(fmt.Errorf("%v != %v", b, false))
	}
}

//func TestEncodeUIntBig_big(t *testing.T) {
//	h := "0x0b0060b7986c88"
//	expected := big.NewInt(int64(150000000000000))
//
//	hh := EncodeUIntBig(*expected)
//	if hh != h {
//		t.Error(fmt.Errorf("%v != %v", h, hh))
//	}
//}
//
//func TestEncodeUIntBig_3bytes(t *testing.T) {
//	h := "0x02093d00"
//	expected := big.NewInt(int64(1000000))
//
//	hh := EncodeUIntBig(*expected)
//	if hh != h {
//		t.Error(fmt.Errorf("%v != %v", h, hh))
//	}
//}
//
//
//func TestEncodeUIntBig_1bytes(t *testing.T) {
//	h := "0x18"
//	expected := big.NewInt(int64(6))
//
//	hh := EncodeUIntBig(*expected)
//	if hh != h {
//		t.Error(fmt.Errorf("%v != %v", h, hh))
//	}
//}
//
//func TestEncodeUIntBig_2bytes(t *testing.T) {
//	h := "0xc15d"
//	expected := big.NewInt(int64(6000))
//
//	hh := EncodeUIntBig(*expected)
//	if hh != h {
//		t.Error(fmt.Errorf("%v != %v", h, hh))
//	}
//}
