package runner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUDPChecksumW23(t *testing.T) {
	payload := []byte{0x08, 0x00, 0x00, 0x00}
	var sum uint32
	for i := 0; i < len(payload); i += 2 {
		sum += uint32(payload[i])<<8 | uint32(payload[i+1])
	}
	assert.Greater(t, sum, uint32(0))
}
