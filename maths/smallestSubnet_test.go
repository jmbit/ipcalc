package maths_test

import (
	"net"
	"testing"

	"github.com/jmbit/ipcalc/maths"
)

func TestSmallestSubnetForTwoIPs(t *testing.T) {
	// Test cases with IPv4 IPs
	testCasesIPv4 := []struct {
		ip1           string
		ip2           string
		expectedCIDR  string
		expectedError bool
	}{
		{
			"192.168.1.1",
			"192.168.1.2",
			"192.168.1.0/31",
			false,
		}, // Test case for two consecutive IPs
		{
			"192.168.1.0",
			"192.168.1.255",
			"192.168.1.0/24",
			false,
		}, // Test case for two IPs in the same /24 subnet
		{
			"192.168.1.1",
			"192.168.2.1",
			"192.168.0.0/22",
			false,
		},
	}

	// Test cases with IPv6 IPs
	testCasesIPv6 := []struct {
		ip1           string
		ip2           string
		expectedCIDR  string
		expectedError bool
	}{
		{
			"2001:db8::1",
			"2001:db8::2",
			"2001:db8::/127",
			false,
		}, // Test case for two consecutive IPs
		{
			"2001:db8::",
			"2001:db8::ffff:ffff",
			"2001:db8::/112",
			false,
		}, // Test case for two IPs in the same /112 subnet
		{
			"2001:db8::1",
			"2001:db8:1::1",
			"2001:db8::/32",
			false,
		},
	}

	for _, tc := range testCasesIPv4 {
		tcpip1 := net.ParseIP(tc.ip1)
		tcpip2 := net.ParseIP(tc.ip2)
		actualCIDR, err := maths.CalculateSmallestSubnet(tcpip1, tcpip2)
		if tc.expectedError {
			if err == nil {
				t.Errorf("Expected error, but got none for IPs %s and %s", tc.ip1, tc.ip2)
			}
			continue
		}
		if err != nil {
			t.Errorf("Unexpected error for IPs %s and %s: %v", tc.ip1, tc.ip2, err)
			continue
		}
		if actualCIDR.String() != tc.expectedCIDR {
			t.Errorf(
				"Incorrect result for IPs %s and %s. Expected: %s, Got: %s",
				tc.ip1,
				tc.ip2,
				tc.expectedCIDR,
				actualCIDR.String(),
			)
		}
	}

	for _, tc := range testCasesIPv6 {
		tcpip1 := net.ParseIP(tc.ip1)
		tcpip2 := net.ParseIP(tc.ip2)
		actualCIDR, err := maths.CalculateSmallestSubnet(tcpip1, tcpip2)
		if tc.expectedError {
			if err == nil {
				t.Errorf("Expected error, but got none for IPs %s and %s", tc.ip1, tc.ip2)
			}
			continue
		}
		if err != nil {
			t.Errorf("Unexpected error for IPs %s and %s: %v", tc.ip1, tc.ip2, err)
			continue
		}
		if actualCIDR.String() != tc.expectedCIDR {
			t.Errorf(
				"Incorrect result for IPs %s and %s. Expected: %s, Got: %s",
				tc.ip1,
				tc.ip2,
				tc.expectedCIDR,
				actualCIDR.String(),
			)
		}
	}
}
