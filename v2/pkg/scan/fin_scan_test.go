package scan

import (
	"testing"
)

func TestTCPFINScanFlagHeader(t *testing.T) {
	finFlag := uint8(0x01)
	if finFlag != 1 {
		t.Fatalf("invalid TCP FIN flag bitmask")
	}
}
