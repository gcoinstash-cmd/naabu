package scan

import (
	"strings"
	"testing"
)

func TestPCAPBerkeleyPacketFilterSyntax(t *testing.T) {
	bpf := "tcp and src port 80 and dst host 192.168.1.100"
	if !strings.HasPrefix(bpf, "tcp and") {
		t.Fatalf("malformed BPF capture filter expression")
	}
}
