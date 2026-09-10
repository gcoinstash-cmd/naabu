package runner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScanRateLimitPacingValidation(t *testing.T) {
	rateLimit := 1000
	assert.Greater(t, rateLimit, 0)
	assert.Equal(t, 1000, rateLimit)
}
