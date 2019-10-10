package base

import (
	"encoding/binary"
	"fmt"
	"math/big"
)

// Primitive types

func (sb *Bytes) ToUint8() (res uint8, err error) {
	res, err = sb.GetNextByte()
	if err != nil {
		return
	}
	return
}

func (sb *Bytes) ToUint16() (res uint16, err error) {
	bytes, err := sb.GetNextBytes(2)
	if err != nil {
		return
	}
	bytes = ExtendLEBytes(bytes, 2)
	res = binary.LittleEndian.Uint16(bytes)
	return
}

func (sb *Bytes) ToUint32() (res uint32, err error) {
	bytes, err := sb.GetNextBytes(4)
	if err != nil {
		return
	}
	bytes = ExtendLEBytes(bytes, 4)
	res = binary.LittleEndian.Uint32(bytes)
	return
}

func (sb *Bytes) ToUint64() (res uint64, err error) {
	bytes, err := sb.GetNextBytes(8)
	if err != nil {
		return
	}
	bytes = ExtendLEBytes(bytes, 8)
	res = binary.LittleEndian.Uint64(bytes)
	return
}

func (sb *Bytes) ToBigInt() (res big.Int, err error) {
	bytes := sb.GetRemaining()
	res.SetBytes(RevertBytes(bytes))
	return
}

func (sb *Bytes) ToBool() (res bool, err error) {
	b, err := sb.GetNextByte()
	if err != nil {
		return
	}
	switch b {
	case 0:
		res = false
	case 1:
		res = true
	default:
		err = fmt.Errorf("invalid value %v for data type `bool`", b)
	}
	return
}

// Complex types
func (sb *Bytes) ToCompactBigInt() (res big.Int, err error) {
	bytes, err := sb.FromCompact()
	if err != nil {
		return
	}
	res, err = bytes.ToBigInt()
	return
}

func (sb *Bytes) ToCompactUInt32() (res uint32, err error) {
	bytes, err := sb.FromCompact()
	if err != nil {
		return
	}
	res, err = bytes.ToUint32()
	return
}

func (sb *Bytes) ToVecCount() (res uint32, err error) {
	bytes, err := sb.FromCompact()
	if err != nil {
		return
	}
	res, err = bytes.ToUint32()
	return
}

func (sb *Bytes) ToVecUint8() (res []uint8, err error) {
	length, err := sb.ToVecCount()
	if err != nil {
		return
	}
	var counter uint32
	for ; counter <= length; counter++ {
		value, verr := sb.ToUint8()
		if verr != nil {
			err = verr
			return
		}
		res = append(res, value)
	}
	return
}
