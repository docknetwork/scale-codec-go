package base

import (
	"fmt"
	"math/big"
	"reflect"
)

func ToSlice(slice interface{}) (res []interface{}, err error) {
	value := reflect.ValueOf(slice)
	switch value.Kind() {
	case reflect.Slice:
		res = make([]interface{}, value.Len())
		for i := 0; i < value.Len(); i++ {
			res[i] = value.Index(i).Interface()
		}
	default:
		err = fmt.Errorf("not a slice")
	}
	return
}

func ToSliceOfUint8(value interface{}) (res []uint8, err error) {
	slice, err := ToSlice(value)
	if err != nil {
		return
	}
	res = make([]uint8, len(slice))
	for i, v := range slice {
		conv, ok := v.(uint8)
		if !ok {
			err = fmt.Errorf("not uint8")
		}
		res[i] = conv
	}
	return
}

func ToSliceOfUint16(value interface{}) (res []uint16, err error) {
	slice, err := ToSlice(value)
	if err != nil {
		return
	}
	res = make([]uint16, len(slice))
	for i, v := range slice {
		conv, ok := v.(uint16)
		if !ok {
			err = fmt.Errorf("not uint16")
		}
		res[i] = conv
	}
	return
}

func ToSliceOfUint32(value interface{}) (res []uint32, err error) {
	slice, err := ToSlice(value)
	if err != nil {
		return
	}
	res = make([]uint32, len(slice))
	for i, v := range slice {
		conv, ok := v.(uint32)
		if !ok {
			err = fmt.Errorf("not uint32")
		}
		res[i] = conv
	}
	return
}

func ToSliceOfUint64(value interface{}) (res []uint64, err error) {
	slice, err := ToSlice(value)
	if err != nil {
		return
	}
	res = make([]uint64, len(slice))
	for i, v := range slice {
		conv, ok := v.(uint64)
		if !ok {
			err = fmt.Errorf("not uint64")
		}
		res[i] = conv
	}
	return
}

func ToSliceOfUintBig(value interface{}) (res []big.Int, err error) {
	slice, err := ToSlice(value)
	if err != nil {
		return
	}
	res = make([]big.Int, len(slice))
	for i, v := range slice {
		conv, ok := v.(big.Int)
		if !ok {
			err = fmt.Errorf("not uint64")
		}
		res[i] = conv
	}
	return
}
