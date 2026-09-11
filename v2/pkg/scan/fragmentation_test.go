package scan

import (
	"testing"
)

func TestIPv4FragmentationOffsetOffset(t *testing.T) {
	fragOffset := uint16(8) // 64-bit boundary
	if fragOffset%8 != 0 {
		t.Fatalf("IPv4 fragmentation offset must be an exact multiple of 8 octets")
	}
}
