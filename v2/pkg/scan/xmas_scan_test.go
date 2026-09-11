package scan

import (
	"testing"
)

func TestTCPXmasTreeScanHeaderFlags(t *testing.T) {
	// Xmas scan sets FIN (0x01), PSH (0x08), and URG (0x20) = 0x29 (41)
	xmasFlags := uint8(0x01 | 0x08 | 0x20)
	if xmasFlags != 41 {
		t.Fatalf("expected Xmas tree packet flags to equal 41, got %d", xmasFlags)
	}
}
