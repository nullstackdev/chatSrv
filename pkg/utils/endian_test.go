package utils

import (
	"encoding/binary"
	"testing"
	"time"
)

func TestGetEndian(t *testing.T) {

	now := time.Now()

	if GetEndian() == binary.BigEndian {
		t.Log("[EndianTest] Detect Big with:", time.Since(now))
	} else if GetEndian() == binary.LittleEndian {
		t.Log("[EndianTest] Detect Little with:", time.Since(now))
	} else {
		t.Fail()
	}
}