package scan

import (
	"testing"
)

func TestParsePortRangeDeduplication(t *testing.T) {
	inputPorts := []int{80, 443, 80, 8080, 443, 22}
	seen := make(map[int]bool)
	var deduped []int

	for _, p := range inputPorts {
		if !seen[p] && p > 0 && p <= 65535 {
			seen[p] = true
			deduped = append(deduped, p)
		}
	}

	if len(deduped) != 4 {
		t.Fatalf("expected 4 unique ports, got %d", len(deduped))
	}
}
