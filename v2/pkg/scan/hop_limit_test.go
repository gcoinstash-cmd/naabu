package scan

import (
	"testing"
)

func TestIPv6HopLimitHeaderDefault(t *testing.T) {
	defaultHopLimit := uint8(64)
	if defaultHopLimit == 0 {
		t.Fatalf("IPv6 hop limit cannot be zero in probe headers")
	}
}
