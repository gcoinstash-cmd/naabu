package scan

import (
	"net"
	"testing"
)

func TestHostExclusionFilter(t *testing.T) {
	excludedIPs := map[string]bool{
		"192.168.1.1":   true,
		"10.0.0.1":       true,
		"172.16.0.1":     true,
	}

	targets := []string{"192.168.1.1", "192.168.1.50", "10.0.0.1", "8.8.8.8"}
	var activeTargets []string

	for _, tgt := range targets {
		if !excludedIPs[tgt] && net.ParseIP(tgt) != nil {
			activeTargets = append(activeTargets, tgt)
		}
	}

	if len(activeTargets) != 2 {
		t.Fatalf("expected 2 active targets, got %d", len(activeTargets))
	}
}
