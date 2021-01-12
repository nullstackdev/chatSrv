package utils

import (
	"encoding/binary"
	"sync"
	"unsafe"
)

const (
	uint16Len = 2
	record = 0xAABB
)

var (
	nativeEndian binary.ByteOrder
	once         = &sync.Once{}
)

func load() {

	buf := [uint16Len]byte{}
	*(*uint16)(unsafe.Pointer(&buf[0])) = uint16(record)

	switch buf {
	case [uint16Len]byte{record & 0x00FF, record >> 8}:
		nativeEndian = binary.LittleEndian
	case [uint16Len]byte{record >> 8, record & 0x00FF}:
		nativeEndian = binary.BigEndian
	default:
		panic("could not determine native endianness.")
	}
}

func GetEndian() binary.ByteOrder {
	
	once.Do(func() {
		load()
	})
	return nativeEndian
}
