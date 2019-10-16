package codec

import (
	"fmt"
	"testing"
)

func ExampleDecode() {
	bytes, _ := NewBytes("0xc15d")
	value, _ := bytes.ToCompactUInt32()
	fmt.Println(value)
	// Output: 6000
}

func TestDecodeCompactUInt32(t *testing.T) {
	testTable := map[string]U32{
		"0x02093d00": 1000000,
		"0x18":       6,
		"0xc15d":     6000,
	}

	for h, expected := range testTable {
		bytes, err := NewBytes(h)
		if err != nil {
			t.Error(err)
		}
		i, err := bytes.ToCompactUInt32()
		if i != expected {
			t.Error(fmt.Errorf("%v != %v", i, expected))
		}
	}
}

func TestDecodeBool(t *testing.T) {
	bytes, err := NewBytes("0x01")
	if err != nil {
		t.Error(err)
	}
	b, err := bytes.ToBool()
	if err != nil {
		t.Error(err)
	}
	if !b {
		t.Error(fmt.Errorf("%v != %v", b, true))
	}

	bytes, err = NewBytes("0x00")
	if err != nil {
		t.Error(err)
	}
	b, err = bytes.ToBool()
	if err != nil {
		t.Error(err)
	}
	if b {
		t.Error(fmt.Errorf("%v != %v", b, false))
	}

}
