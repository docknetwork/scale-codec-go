package base

import (
	"encoding/hex"
	"fmt"
)

//Bytes is a wrapper for []bytes that's keeping offset for reading purposes
type Bytes struct {
	data   []byte
	offset int
}

//GetNextBytes returns `length` number of bytes
func (sb *Bytes) GetNextBytes(length int) (bytes []byte, err error) {
	calcLength := MinInt(length, sb.GetRemainingLength())
	bytes = sb.data[sb.offset:calcLength]
	if len(bytes) == 0 {
		err = fmt.Errorf("out of range")
	}
	sb.offset += calcLength
	return
}

//GetNextByte returns next byte
func (sb *Bytes) GetNextByte() (b byte, err error) {
	length := sb.GetRemainingLength()
	if sb.offset+length > len(sb.data) {
		err = fmt.Errorf("out of range")
		return
	}
	b = sb.data[sb.offset]
	sb.offset++
	return
}

//GetRemaining returns remaining bytes
func (sb *Bytes) GetRemaining() (bytes []byte) {
	length := sb.GetRemainingLength()
	if length == 0 {
		return
	}
	bytes = sb.data[sb.offset : sb.offset+length]
	sb.offset += length
	return
}

//GetAll resets offset and returns all bytes
func (sb *Bytes) GetAll() (bytes []byte) {
	sb.Reset()
	bytes = sb.GetRemaining()
	return
}

//GetRemainingLength returns number of remaining bytes
func (sb *Bytes) GetRemainingLength() (length int) {
	length = len(sb.data) - sb.offset
	return
}

//Reset resets offset
func (sb *Bytes) Reset() {
	sb.offset = 0
}

//Check checks if extra bytes exist
func (sb *Bytes) Check() (err error) {
	dataLength := len(sb.data)
	if sb.offset != dataLength {
		err = fmt.Errorf("current offset: %v and current length: %v", sb.offset, dataLength)
	}
	return
}

func (sb *Bytes) ToHex() (res string) {
	res = "0x" + hex.EncodeToString(sb.GetAll())
	return
}

func (sb *Bytes) ToCompact() (res Bytes, err error) {
	bytes := BytesToCompactBytes(sb.GetAll())
	res, err = NewBytes(bytes)
	return
}

//NewBytes creates Bytes struct
func NewBytes(data interface{}) (sb Bytes, err error) {
	switch data.(type) {
	case Bytes:
		sb = data.(Bytes)
	case []byte:
		sb = Bytes{
			data: data.([]byte),
		}
	case string:
		sb, err = func(data string) (sb Bytes, err error) {
			if data[:2] == "0x" { //Respect the 60's
				data = data[2:]
			}
			bytes, err := hex.DecodeString(data)
			if err == nil {
				sb = Bytes{
					data: bytes,
				}
			}
			return
		}(data.(string))
	case nil:
		sb = Bytes{}
	default:
		err = fmt.Errorf("unknown type %T", data)
	}
	return
}
