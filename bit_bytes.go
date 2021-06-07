package bb

import (
	"encoding/binary"
)

const (
	// BytesByteLength 字节的二进制字节长度
	BytesByteLength = 8

	// BitInt16Length int16的字节长度
	BitInt16Length = 2
	// BytesInt16Length int16的二进制字节长度
	BytesInt16Length = 16

	// BitIntLength int的字节长度
	BitIntLength = 4
	// BytesIntLength int的二进制字节长度
	BytesIntLength = 32

	// BitInt64Length int64的字节长度
	BitInt64Length = 8
	// BytesInt64Length int64的二进制字节长度
	BytesInt64Length = 64
)

// DefaultEndPont 默认大端
var DefaultEndPont binary.ByteOrder = binary.BigEndian

// InitIntBytes 初始化二进制字节切片
func InitIntBytes(length int) []byte {
	switch length {
	case BytesByteLength:
		return Byte2Bytes(byte(0xff))
	case BytesInt16Length:
		return Uint162Bytes(uint16(0xffff))
	case BytesIntLength:
		return Uint2Bytes(uint32(0xffffffff))
	case BytesInt64Length:
		return Uint642Bytes(uint64(0xffffffffffffffff))
	}
	panic("error length")
}

// Revert 字节反转
func Revert(bytes []byte) []byte {
	total := len(bytes)
	half := total / 2
	for i := 0; i < half; i++ {
		tmp := bytes[i]
		bytes[i] = bytes[total-i-1]
		bytes[total-i-1] = tmp
	}
	return bytes
}

// Bytes2Byte 二进制字节切片转换成字节
func Bytes2Byte(bytes []byte) (b byte) {
	length := len(bytes)
	if length < BytesByteLength {
		bytes = append(make([]byte, BytesByteLength-length), bytes...)
	} else if length > BytesByteLength {
		bytes = bytes[:BytesByteLength]
	}
	for k, v := range bytes {
		b += v << (BytesByteLength - k - 1)
	}
	return
}

// Byte2Bytes 字节转换成二进制字节切片
func Byte2Bytes(b byte) (bytes []byte) {
	bytes = make([]byte, BytesByteLength)
	for k := range bytes {
		bytes[BytesByteLength-k-1] = b >> k & 1
	}
	return
}

// Bytes2Int16 二进制字节切片转换成int16
func Bytes2Int16(bytes []byte) (i int16) {
	length := len(bytes)
	if length < BytesInt16Length {
		bytes = append(make([]byte, BytesInt16Length-length), bytes...)
	} else if length > BytesInt16Length {
		bytes = bytes[:BytesInt16Length]
	}
	for k, v := range bytes {
		i += int16(v) << (BytesInt16Length - k - 1)
	}
	return
}

// Int162Bytes int16 转换成二进制字节切片
func Int162Bytes(i int16) (bytes []byte) {
	return Uint162Bytes(uint16(i))
}

// Uint162Bytes uint16 转换成二进制字节切片
func Uint162Bytes(i uint16) (bytes []byte) {
	var b = make([]byte, BitInt16Length)
	bytes = make([]byte, BytesInt16Length)
	DefaultEndPont.PutUint16(b, i)
	for k, v := range b {
		copy(bytes[k*BytesByteLength:(k+1)*BytesByteLength], Byte2Bytes(v))
	}
	return
}

// Bytes2Int 二进制字节切片转换成 int 整型
func Bytes2Int(bytes []byte) (i int) {
	length := len(bytes)
	if length < BytesIntLength {
		bytes = append(make([]byte, BytesIntLength-length), bytes...)
	} else if length > BytesIntLength {
		bytes = bytes[:BytesIntLength]
	}
	for k, v := range bytes {
		i += int(v) << (BytesIntLength - k - 1)
	}
	return
}

// Int2Bytes int转换成二进制字节切片
func Int2Bytes(i int) (bytes []byte) {
	return Uint2Bytes(uint32(i))
}

// Uint2Bytes uint转换成二进制字节切片
func Uint2Bytes(i uint32) (bytes []byte) {
	var b = make([]byte, BitIntLength)
	bytes = make([]byte, BytesIntLength)
	DefaultEndPont.PutUint32(b, i)
	for k, v := range b {
		copy(bytes[k*BytesByteLength:(k+1)*BytesByteLength], Byte2Bytes(v))
	}
	return
}

// Bytes2Int64 二进制字节切片转换int64
func Bytes2Int64(bytes []byte) (i int64) {
	length := len(bytes)
	if length < BytesInt64Length {
		bytes = append(make([]byte, BytesInt64Length-length), bytes...)
	} else if length > BytesInt64Length {
		bytes = bytes[:BytesInt64Length]
	}
	for k, v := range bytes {
		i += int64(v) << (BytesInt64Length - k - 1)
	}
	return
}

// Int642Bytes int64 转换成 二进制字节切片
func Int642Bytes(i int64) (bytes []byte) {
	return Uint642Bytes(uint64(i))
}

// Uint642Bytes uint16 转换成二进制字节切片
func Uint642Bytes(i uint64) (bytes []byte) {
	var b = make([]byte, BitInt64Length)
	bytes = make([]byte, BytesInt64Length)
	DefaultEndPont.PutUint64(b, i)
	for k, v := range b {
		copy(bytes[k*BytesByteLength:(k+1)*BytesByteLength], Byte2Bytes(v))
	}
	return
}
