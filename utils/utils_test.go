package utils

import (
	"log"
	"testing"
)

func TestNetIdToString(t *testing.T) {
	netStr := NetIdToString([]byte{0x00, 0x00, 0x13})
	log.Print(netStr)
}
