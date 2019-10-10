package base

import (
	"fmt"
)

func Decode(typeString string, value interface{}, skipCheck bool) (res interface{}, isNull bool, err error) {
	bytes, err := NewBytes(value)
	if err != nil {
		return
	}
	return DecodeByte(typeString, &bytes, skipCheck)
}

func DecodeByteToByte(typeString string, bytes *Bytes, skipCheck bool) (res Bytes, isNull bool, err error) {
	switch typeString {
	case "option":
		res, isNull, err = bytes.FromOption()
	case "compact":
		res, err = bytes.FromCompact()
	default:
		err = fmt.Errorf("unknown format %s", typeString)
	}
	return
}

func DecodeByte(typeString string, bytes *Bytes, skipCheck bool) (res interface{}, isNull bool, err error) {

	typeString, subTypeString := ExtractSubType(typeString)

	switch typeString {
	case "option":
		fallthrough
	case "compact":
		res, isNull, err = DecodeByteToByte(typeString, bytes, true)
	case "vec":
		// TODO: get compact32 first

		var slice []interface{}
		for bytesLeft := bytes.Check(); bytesLeft != nil; bytesLeft = bytes.Check() {
			v, _, e := DecodeByte(subTypeString, bytes, true)
			if e != nil {
				err = e
				return
			}
			slice = append(slice, v)
		}
		subTypeString = ""
		res = slice
	//case "option":
	//	res, isNull, err = bytes.FromOption()
	//case "compact":
	//	res, err = bytes.FromCompact()
	case "bool":
		res, err = bytes.ToBool()
	case "u8":
		res, err = bytes.ToUint8()
	case "u16":
		res, err = bytes.ToUint16()
	case "u32":
		res, err = bytes.ToUint32()
	case "u64":
		res, err = bytes.ToUint64()
	case "u128":
		res, err = bytes.ToBigInt()
	default:
		err = fmt.Errorf("unknown format %s", typeString)
	}

	if err == nil && !skipCheck {
		err = bytes.Check()
	}

	if err == nil && len(subTypeString) != 0 && !isNull {
		res, isNull, err = Decode(subTypeString, res, skipCheck)
	}

	return
}
