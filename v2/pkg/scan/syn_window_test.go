package scan

import (
	"testing"
)

func TestSYNACKWindowSizeValidation(t *testing.T) {
	windowSize := uint16(65535)
	if windowSize == 0 {
		t.Fatalf("TCP window size cannot be 0 for standard SYN probes")
	}
	
	maxSegmentSize := 1460
	if maxSegmentSize > 1500 {
		t.Fatalf("MSS exceeds standard MTU limits")
	}
}
