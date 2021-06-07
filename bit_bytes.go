package bb

import (
	"encoding/binary"
)

const (
	bytesByteLength = 8

	bitInt16Length   = 2
	bytesInt16Length = 16

	bitIntLength   = 4
	bytesIntLength = 32

	bitInt64Length   = 8
	bytesInt64Length = 64
)

// DefaultEndPont 默认大端
var DefaultEndPont binary.ByteOrder = binary.BigEndian

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
	if length < bytesByteLength {
		bytes = append(make([]byte, bytesByteLength-length), bytes...)
	} else if length > bytesByteLength {
		bytes = bytes[:bytesByteLength]
	}
	for k, v := range bytes {
		b += v << (bytesByteLength - k - 1)
	}
	return
}

// Byte2Bytes 字节转换成二进制字节切片
func Byte2Bytes(b byte) (bytes []byte) {
	bytes = make([]byte, bytesByteLength)
	for k := range bytes {
		bytes[bytesByteLength-k-1] = b >> k & 1
	}
	return
}

// Bytes2Int16 二进制字节切片转换成int16
func Bytes2Int16(bytes []byte) (i int16) {
	length := len(bytes)
	if length < bytesInt16Length {
		bytes = append(make([]byte, bytesInt16Length-length), bytes...)
	} else if length > bytesInt16Length {
		bytes = bytes[:bytesInt16Length]
	}
	for k, v := range bytes {
		i += int16(v) << (bytesInt16Length - k - 1)
	}
	return
}

// Int162Bytes int16 转换成二进制字节切片
func Int162Bytes(i int16) (bytes []byte) {
	return Uint162Bytes(uint16(i))
}

// Uint162Bytes uint16 转换成二进制字节切片
func Uint162Bytes(i uint16) (bytes []byte) {
	var b = make([]byte, bitInt16Length)
	bytes = make([]byte, bytesInt16Length)
	DefaultEndPont.PutUint16(b, i)
	for k, v := range b {
		copy(bytes[k*bytesByteLength:(k+1)*bytesByteLength], Byte2Bytes(v))
	}
	return
}

// Bytes2Int 二进制字节切片转换成 int 整型
func Bytes2Int(bytes []byte) (i int) {
	length := len(bytes)
	if length < bytesIntLength {
		bytes = append(make([]byte, bytesIntLength-length), bytes...)
	} else if length > bytesIntLength {
		bytes = bytes[:bytesIntLength]
	}
	for k, v := range bytes {
		i += int(v) << (bytesIntLength - k - 1)
	}
	return
}

// Int2Bytes int转换成二进制字节切片
func Int2Bytes(i int) (bytes []byte) {
	return Uint2Bytes(uint32(i))
}

// Uint2Bytes uint转换成二进制字节切片
func Uint2Bytes(i uint32) (bytes []byte) {
	var b = make([]byte, bitIntLength)
	bytes = make([]byte, bytesIntLength)
	DefaultEndPont.PutUint32(b, i)
	for k, v := range b {
		copy(bytes[k*bytesByteLength:(k+1)*bytesByteLength], Byte2Bytes(v))
	}
	return
}

// Bytes2Int64 二进制字节切片转换int64
func Bytes2Int64(bytes []byte) (i int64) {
	length := len(bytes)
	if length < bytesInt64Length {
		bytes = append(make([]byte, bytesInt64Length-length), bytes...)
	} else if length > bytesInt64Length {
		bytes = bytes[:bytesInt64Length]
	}
	for k, v := range bytes {
		i += int64(v) << (bytesInt64Length - k - 1)
	}
	return
}

// Int642Bytes int64 转换成 二进制字节切片
func Int642Bytes(i int64) (bytes []byte) {
	return Uint642Bytes(uint64(i))
}

// Uint642Bytes uint16 转换成二进制字节切片
func Uint642Bytes(i uint64) (bytes []byte) {
	var b = make([]byte, bitInt64Length)
	bytes = make([]byte, bytesInt64Length)
	DefaultEndPont.PutUint64(b, i)
	for k, v := range b {
		copy(bytes[k*bytesByteLength:(k+1)*bytesByteLength], Byte2Bytes(v))
	}
	return
}
