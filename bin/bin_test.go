package bin_test

import (
	"encoding/binary"
	"fmt"
	"github.com/zing-dev/go-bit-bytes/bin"
	"log"
	"math"
	"testing"
)

func TestInitInt(t *testing.T) {
	log.Println(bin.Init(bin.BytesByteLength))
	log.Println(bin.Init(bin.BytesInt16Length))
	log.Println(bin.Init(bin.BytesIntLength))
	log.Println(bin.Init(bin.BytesInt64Length))
}

func TestRevert(t *testing.T) {
	log.Println(bin.Revert([]byte{1, 2, 3, 4, 5, 6, 7, 8}))
	log.Println(bin.Revert([]byte{1, 2, 3, 4}))
	log.Println(bin.Revert([]byte{1}))
	log.Println(bin.Revert([]byte{}))
}

func TestFromByte(t *testing.T) {
	for i := 0; i <= math.MaxUint8; i++ {
		log.Println(fmt.Sprintf("%x", i), i, bin.FromByte(byte(i)), bin.Revert(bin.FromByte(byte(i))))
	}
}

func TestToByte(t *testing.T) {
	for i := 0; i <= math.MaxUint8; i++ {
		log.Println(fmt.Sprintf("%x", i), i, bin.FromByte(byte(i)), bin.Revert(bin.FromByte(byte(i))))
	}
}

func TestFromInt16(t *testing.T) {
	bin.DefaultEndPont = binary.LittleEndian
	log.Println(bin.FromInt16(1))
	log.Println(bin.FromInt16(0xf))
	log.Println(bin.FromInt16(0xff))
	log.Println(bin.FromInt16(0xfff))
	log.Println(bin.FromInt16(0xfff << 2))
	log.Println(bin.FromInt16(0xfff << 3))

	bin.DefaultEndPont = binary.BigEndian
	log.Println(bin.FromInt16(1))
	log.Println(bin.FromInt16(0xf))
	log.Println(bin.FromInt16(0xff))
	log.Println(bin.FromInt16(0xfff))
	log.Println(bin.FromInt16(0xfff << 2))
	log.Println(bin.FromInt16(0xfff << 3))
}

func TestToInt16(t *testing.T) {
	log.Println(1<<1, bin.FromUint16(1<<1), bin.ToInt16(bin.FromUint(1<<1)))
	log.Println(1<<1, bin.FromUint16(1<<4), bin.ToInt16(bin.FromUint(1<<4)))
	log.Println(1<<1, bin.FromUint16(1<<8), bin.ToInt16(bin.FromUint(1<<8)))
	log.Println(1<<1, bin.FromUint16(1<<14), bin.ToInt16(bin.FromUint(1<<14)))
}

func TestFromInt(t *testing.T) {
	log.Println(bin.FromInt(0xf))
	log.Println(bin.FromInt(0xf0))
	log.Println(bin.FromInt(0xff))
	log.Println(bin.FromInt(0xff0))
	log.Println(bin.FromInt(0xfff))
	log.Println(bin.FromInt(0xfff0))
	log.Println(bin.FromInt(0xffff))
	log.Println(bin.FromInt(0xfffff))
	log.Println(bin.FromInt(0xffffff))
	log.Println(bin.FromInt(0xfffffff))
	log.Println(bin.FromInt(0xffffffaa))
	log.Println(bin.FromInt(0xffffffff))
}

func TestBytes2Int(t *testing.T) {
	log.Println(bin.FromInt(0xf), bin.ToInt(bin.FromInt(0xf)))
	log.Println(bin.FromInt(0xf0), bin.ToInt(bin.FromInt(0xf0)))
	log.Println(bin.FromInt(0xff0), bin.ToInt(bin.FromInt(0xff0)))
	log.Println(bin.FromInt(0xff0), bin.ToInt(bin.FromInt(0xff0)))
	log.Println(bin.FromInt(0xfff), bin.ToInt(bin.FromInt(0xfff)))
	log.Println(bin.FromInt(0xfff0), bin.ToInt(bin.FromInt(0xfff0)))
	log.Println(bin.FromInt(0xffff), bin.ToInt(bin.FromInt(0xffff)))
	log.Println(bin.FromInt(0xfffff), bin.ToInt(bin.FromInt(0xfffff)))
	log.Println(bin.FromInt(0xffffff), bin.ToInt(bin.FromInt(0xffffff)))
	log.Println(bin.FromInt(0xfffffff), bin.ToInt(bin.FromInt(0xfffffff)))
	log.Println(bin.FromInt(0xffffffaa), bin.ToInt(bin.FromInt(0xffffffaa)))
	log.Println(bin.FromInt(0xffffffff), bin.ToInt(bin.FromInt(0xffffffff)))
	log.Println(bin.ToInt([]byte{0, 1, 1, 1, 1}))
}

func TestFromInt64(t *testing.T) {
	log.Println(bin.FromInt64(0xff))
	log.Println(bin.FromInt64(0xffffffff))
}

func TestFromUint64(t *testing.T) {
	log.Println(bin.FromUint64(0xff))
	log.Println(bin.FromUint64(0xffffffff))
	log.Println(bin.FromUint64(0xffffffffffffffff))
}

func TestToInt64(t *testing.T) {
	log.Println(fmt.Sprintf("%x", math.MaxInt8), bin.FromInt64(math.MaxInt8), bin.ToInt64(bin.FromUint64(math.MaxInt8)))
	log.Println(fmt.Sprintf("%x", math.MaxInt16), bin.FromInt64(math.MaxInt16), bin.ToInt64(bin.FromUint64(math.MaxInt16)))
	log.Println(fmt.Sprintf("%x", math.MaxInt32), bin.FromInt64(math.MaxInt32), bin.ToInt64(bin.FromUint64(math.MaxInt32)))
	log.Println(fmt.Sprintf("%x", math.MaxInt64), bin.FromInt64(math.MaxInt64), bin.ToInt64(bin.FromUint64(math.MaxInt64)))
}
