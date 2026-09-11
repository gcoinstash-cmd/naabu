package scan

import (
	"net"
	"testing"
)

func TestParseCIDRSubnetBoundary(t *testing.T) {
	_, ipnet, err := net.ParseCIDR("192.168.1.0/24")
	if err != nil {
		t.Fatalf("unexpected CIDR parse error: %v", err)
	}

	testIP := net.ParseIP("192.168.1.100")
	if !ipnet.Contains(testIP) {
		t.Fatalf("expected CIDR to contain %s", testIP)
	}

	outsideIP := net.ParseIP("192.168.2.1")
	if ipnet.Contains(outsideIP) {
		t.Fatalf("expected CIDR to NOT contain %s", outsideIP)
	}
}
