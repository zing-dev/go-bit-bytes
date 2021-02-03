package bb_test

import (
	"fmt"
	"github.com/zing-dev/go-bit-bytes"
	"log"
	"math"
	"testing"
)

func TestRevert(t *testing.T) {
	log.Println(bb.Revert([]byte{1, 2, 3, 4, 5, 6, 7, 8}))
	log.Println(bb.Revert([]byte{1, 2, 3, 4}))
	log.Println(bb.Revert([]byte{1}))
	log.Println(bb.Revert([]byte{}))
}

func TestByte2Bytes(t *testing.T) {
	for i := 0; i <= math.MaxUint8; i++ {
		log.Println(fmt.Sprintf("%x", i), i, bb.Byte2Bytes(byte(i)), bb.Revert(bb.Byte2Bytes(byte(i))))
	}
}

func TestBytes2Byte(t *testing.T) {
	for i := 0; i <= math.MaxUint8; i++ {
		log.Println(fmt.Sprintf("%x", i), i, bb.Byte2Bytes(byte(i)), bb.Bytes2Byte(bb.Byte2Bytes(byte(i))))
	}
}

func TestInt162Bytes(t *testing.T) {
	log.Println(bb.Int162Bytes(1))
	log.Println(bb.Int162Bytes(0xf))
	log.Println(bb.Int162Bytes(0xff))
	log.Println(bb.Int162Bytes(0xfff))
	log.Println(bb.Int162Bytes(0xfff << 2))
	log.Println(bb.Int162Bytes(0xfff << 3))
}

func TestBytes2Int16(t *testing.T) {
	log.Println(1<<1, bb.Int162Bytes(1<<1), bb.Bytes2Int16(bb.Int162Bytes(1<<1)))
	log.Println(1<<4, bb.Int162Bytes(1<<4), bb.Bytes2Int16(bb.Int162Bytes(1<<4)))
	log.Println(1<<8, bb.Int162Bytes(1<<8), bb.Bytes2Int16(bb.Int162Bytes(1<<8)))
	log.Println(1<<14, bb.Int162Bytes(1<<14), bb.Bytes2Int16(bb.Int162Bytes(1<<14)))
}

func TestInt2Bytes(t *testing.T) {
	log.Println(bb.Int2Bytes(0xf))
	log.Println(bb.Int2Bytes(0xf0))
	log.Println(bb.Int2Bytes(0xff))
	log.Println(bb.Int2Bytes(0xff0))
	log.Println(bb.Int2Bytes(0xfff))
	log.Println(bb.Int2Bytes(0xfff0))
	log.Println(bb.Int2Bytes(0xffff))
	log.Println(bb.Int2Bytes(0xfffff))
	log.Println(bb.Int2Bytes(0xffffff))
	log.Println(bb.Int2Bytes(0xfffffff))
	log.Println(bb.Int2Bytes(0xffffffaa))
	log.Println(bb.Int2Bytes(0xffffffff))
}

func TestBytes2Int(t *testing.T) {
	log.Println(bb.Int2Bytes(0xf), bb.Bytes2Int(bb.Int2Bytes(0xf)))
	log.Println(bb.Int2Bytes(0xf0), bb.Bytes2Int(bb.Int2Bytes(0xf0)))
	log.Println(bb.Int2Bytes(0xff0), bb.Bytes2Int(bb.Int2Bytes(0xff0)))
	log.Println(bb.Int2Bytes(0xff0), bb.Bytes2Int(bb.Int2Bytes(0xff0)))
	log.Println(bb.Int2Bytes(0xfff), bb.Bytes2Int(bb.Int2Bytes(0xfff)))
	log.Println(bb.Int2Bytes(0xfff0), bb.Bytes2Int(bb.Int2Bytes(0xfff0)))
	log.Println(bb.Int2Bytes(0xffff), bb.Bytes2Int(bb.Int2Bytes(0xffff)))
	log.Println(bb.Int2Bytes(0xfffff), bb.Bytes2Int(bb.Int2Bytes(0xfffff)))
	log.Println(bb.Int2Bytes(0xffffff), bb.Bytes2Int(bb.Int2Bytes(0xffffff)))
	log.Println(bb.Int2Bytes(0xfffffff), bb.Bytes2Int(bb.Int2Bytes(0xfffffff)))
	log.Println(bb.Int2Bytes(0xffffffaa), bb.Bytes2Int(bb.Int2Bytes(0xffffffaa)))
	log.Println(bb.Int2Bytes(0xffffffff), bb.Bytes2Int(bb.Int2Bytes(0xffffffff)))
	log.Println(bb.Bytes2Int([]byte{0, 1, 1, 1, 1}))
}

func TestInt642Bytes(t *testing.T) {
	log.Println(bb.Int642Bytes(0xff))
	log.Println(bb.Int642Bytes(0xffffffff))
}

func TestUint642Bytes(t *testing.T) {
	log.Println(bb.Uint642Bytes(0xff))
	log.Println(bb.Uint642Bytes(0xffffffff))
	log.Println(bb.Uint642Bytes(0xffffffffffffffff))
}

func TestBytes2Int64(t *testing.T) {
	log.Println(fmt.Sprintf("%x", math.MaxInt8), bb.Int642Bytes(math.MaxInt8), bb.Bytes2Int64(bb.Int642Bytes(math.MaxInt8)))
	log.Println(fmt.Sprintf("%x", math.MaxInt16), bb.Int642Bytes(math.MaxInt16), bb.Bytes2Int64(bb.Int642Bytes(math.MaxInt16)))
	log.Println(fmt.Sprintf("%x", math.MaxInt32), bb.Int642Bytes(math.MaxInt32), bb.Bytes2Int64(bb.Int642Bytes(math.MaxInt32)))
	log.Println(fmt.Sprintf("%x", math.MaxInt64), bb.Int642Bytes(math.MaxInt64), bb.Bytes2Int64(bb.Int642Bytes(math.MaxInt64)))
}
