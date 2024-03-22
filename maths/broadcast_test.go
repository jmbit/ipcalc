package maths_test

import (
	"net"
	"testing"

	"github.com/jmbit/ipcalc/maths"
)

func TestGetBroadcastAddress(t *testing.T) {
	// Create a new network
	network := &net.IPNet{
		IP:   net.ParseIP("192.168.178.0"),
		Mask: net.CIDRMask(24, 32),
	}

	// Calculate the broadcast address
	broadcast, err := maths.GetBroadcastAddress(network)
	if err != nil {
		t.Error(err)

	}
	if broadcast.String() != "192.168.178.255" {
		t.Errorf("Expected 192.168.178.255, but got %s", broadcast.String())
	}
}
