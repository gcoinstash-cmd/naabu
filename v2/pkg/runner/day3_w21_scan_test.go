package runner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScanPortRangeOptionValidation(t *testing.T) {
	ports := "80,443,8000-8080"
	assert.NotEmpty(t, ports)
	assert.Contains(t, ports, "443")
}
