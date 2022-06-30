package conversion

import (
	"bytes"
	"encoding/binary"
	"strconv"
)

//UTILITIES

//Itob int to byte
func Itob(v int) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(v))
	return b
}

// UI32tob ...
func UI32tob(v uint32) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint32(b, uint32(v))
	return b
}

//BtoI byte to int
func BtoI(v []byte) int {
	if len(v) < 8 {
		return 0
	}
	return int(binary.LittleEndian.Uint64(v))
}

// Btoui32 ....
func Btoui32(v []byte) uint32 {
	return (binary.LittleEndian.Uint32(v))
}

//IdFromString ...
func IdFromString(s *string) []byte {
	v, _ := strconv.Atoi(*s)
	return Itob(v)
}

//B16toi byte to uint16
func B16toi(v []byte) uint16 {
	if len(v) > 0 {
		return binary.LittleEndian.Uint16(v)
	}
	return 0
}

//Btoi64 byte to uint64
func Btoi64(v []byte) uint64 {
	return binary.LittleEndian.Uint64(v)
}

//Uitob16 uint16 to byte
//used for versioning
func Uitob16(v uint16) []byte {
	b := make([]byte, 8)

	binary.LittleEndian.PutUint16(b, v)

	return b
}

//Uitob64 uint to byte64
func Uitob64(v uint64) []byte {
	b := make([]byte, 8)

	binary.LittleEndian.PutUint64(b, v)
	return b
}

//BinBool boolean to byte
func BinBool(val bool) []byte {
	byteVal := []byte{}
	if val {
		return append(byteVal, 1&1)
	}
	return append(byteVal, 0&1)
}

//BoolInB byte to boolean
func BoolInB(v []byte) bool {
	if len(v) > 0 {
		return v[0]&0x1 == 0x1
	}
	return false
}

// Uint8ToByte to convert uint8 value to array of byte
func Uint8ToByte(f uint8) []byte {
	var buf bytes.Buffer
	binary.Write(&buf, binary.BigEndian, f)
	return buf.Bytes()
}

// UintToByte to convert uint value to array of byte
func UintToByte(f uint) []byte {
	var buf bytes.Buffer
	binary.Write(&buf, binary.BigEndian, f)
	return buf.Bytes()
}
