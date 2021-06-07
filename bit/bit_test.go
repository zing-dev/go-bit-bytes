package bit_test

import (
	"encoding/binary"
	"fmt"
	"github.com/zing-dev/go-bit-bytes/bit"
	"log"
	"testing"
)

func TestToUint32(t *testing.T) {
	log.Println(fmt.Sprintf("%0x", bit.ToUint32([]byte{0, 0, 0, 0xff})))
	log.Println(fmt.Sprintf("%0x", bit.ToUint32([]byte{0, 0, 0xff, 0xff})))
	log.Println(fmt.Sprintf("%0x", bit.ToUint32([]byte{0, 0xff, 0xff, 0xff})))
	log.Println(fmt.Sprintf("%0x", bit.ToUint32([]byte{0xff, 0xff, 0xff, 0xff})))
	log.Println(fmt.Sprintf("%0x", bit.ToInt32([]byte{0xff, 0xff, 0xff, 0xff})))
	log.Println(fmt.Sprintf("%0x", bit.ToInt([]byte{0xff, 0xff, 0xff, 0xff})))
}

func TestFromUint32(t *testing.T) {
	log.Println(fmt.Sprintf("%v", bit.FromUint32(0xffff)))
	bit.DefaultEndPont = binary.LittleEndian
	log.Println(fmt.Sprintf("%v", bit.FromUint32(0xffff)))
}
