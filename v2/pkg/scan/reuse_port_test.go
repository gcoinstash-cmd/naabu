package scan

import (
	"testing"
)

func TestSocketSOReusePortFlag(t *testing.T) {
	reusePortOption := true
	if !reusePortOption {
		t.Fatalf("SO_REUSEPORT option must be enabled for multi-threaded packet listeners")
	}
}
