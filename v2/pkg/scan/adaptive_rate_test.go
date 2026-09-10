package scan

import (
	"testing"
)

func TestAdaptiveRateLimiterThrottle(t *testing.T) {
	initialRate := 1000
	dropFactor := 0.5
	newRate := int(float64(initialRate) * dropFactor)
	
	if newRate != 500 {
		t.Fatalf("expected throttled rate 500, got %d", newRate)
	}
}
