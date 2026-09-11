package scan

import (
	"net"
	"testing"
)

func TestARPCacheHardwareAddressResolution(t *testing.T) {
	hwAddr, err := net.ParseMAC("00:1A:2B:3C:4D:5E")
	if err != nil {
		t.Fatalf("failed to parse valid MAC hardware address: %v", err)
	}
	if len(hwAddr) != 6 {
		t.Fatalf("invalid Ethernet MAC address byte length")
	}
}
