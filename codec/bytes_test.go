package codec

import (
	"testing"
)

func TestBytesFromHex(t *testing.T) {
	bytes, err := NewBytes("0x183A4D8E")
	if err != nil {
		t.Error(err)
	}

	expected := []byte{24, 58, 77, 142}
	remaining := bytes.GetRemaining()

	for i, b := range remaining {
		if expected[i] != b {
			t.Error("Wrong value")
		}
	}
}

func TestBytes(t *testing.T) {
	expected := []byte{24, 58, 77, 142}
	bytes, err := NewBytes(expected)
	if err != nil {
		t.Error(err)
	}

	remaining := bytes.GetRemaining()
	for i, b := range remaining {
		if expected[i] != b {
			t.Error("Wrong value")
		}
	}

}

func TestOffset(t *testing.T) {
	expected := []byte{24, 58, 77, 142}
	bytes, err := NewBytes(expected)
	if err != nil {
		t.Error(err)
	}

	nb, _ := bytes.GetNextBytes(1)
	if expected[0] != nb[0] {
		t.Error("Wrong next value")
	}

	if bytes.offset != 1 {
		t.Error("Wrong offset value")
	}

	if bytes.GetRemainingLength() != 3 {
		t.Error("Wrong remaining length")
	}

	b, err := bytes.GetNextByte()
	if expected[1] != b {
		t.Error("Wrong next value")
	}

	if bytes.offset != 2 {
		t.Error("Wrong offset value")
	}

	if bytes.GetRemainingLength() != 2 {
		t.Error("Wrong remaining length")
	}

	remaining := bytes.GetRemaining()
	for i, b := range remaining {
		if expected[i+2] != b {
			t.Error("Wrong value")
		}
	}

	if bytes.offset != len(expected) {
		t.Error("Wrong offset value")
	}

	bytes.Reset()
	if bytes.offset != 0 {
		t.Error("Wrong offset value after reset")
	}

}
