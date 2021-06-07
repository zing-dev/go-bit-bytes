package bit

import "encoding/binary"

var DefaultEndPont binary.ByteOrder = binary.BigEndian

// ToInt16 字节切片转换成 int16
func ToInt16(b []byte) (i int16) {
	return int16(ToUint16(b))
}

// ToUint16 字节切片转换成 uint16
func ToUint16(b []byte) uint16 {
	return DefaultEndPont.Uint16(b)
}

// FromUint16  转换成字节数组
func FromUint16(i uint16) []byte {
	data := make([]byte, 2)
	DefaultEndPont.PutUint16(data, i)
	return data
}

// ToInt 字节切片转换成 int
func ToInt(b []byte) (i int) {
	i = int(ToUint32(b))
	return
}

func ToUint(b []byte) (i uint) {
	i = uint(ToUint32(b))
	return
}

func ToInt32(b []byte) (i int32) {
	i = int32(ToUint32(b))
	return
}

// ToUint32 字节切片转换成uint32
func ToUint32(b []byte) (i uint32) {
	i = DefaultEndPont.Uint32(b)
	return
}

// FromInt uint32转换成字节数组
func FromInt(i int) []byte {
	return FromUint32(uint32(i))
}

// FromUint32 uint32转换成字节数组
func FromUint32(i uint32) []byte {
	data := make([]byte, 4)
	DefaultEndPont.PutUint32(data, i)
	return data
}

// ToUint64 字节切片转换成uint32
func ToUint64(b []byte) uint64 {
	return DefaultEndPont.Uint64(b)
}

// FromUint64  uint64 转换成字节数组
func FromUint64(i uint64) []byte {
	data := make([]byte, 8)
	DefaultEndPont.PutUint64(data, i)
	return data
}
