package maths

import (
	"encoding/binary"
	"math/big"
	"net"
)

// getLastIPAddress calculates the last IP address in a given network range
func GetLastIPAddress(network *net.IPNet) net.IP {
	// Convert the network IP to a full 16-byte representation
	ip := network.IP.To16()

	// Check if the IP address is IPv4 or IPv6
	if ip.To4() != nil {
		// IPv4
		ipInt := binary.BigEndian.Uint32(ip[12:])
		mask := binary.BigEndian.Uint32(network.Mask)
		addresses := ^mask + 1
		lastIPInt := ipInt | (addresses - 2)
		lastIP := make(net.IP, 4)
		binary.BigEndian.PutUint32(lastIP, lastIPInt)
		return lastIP
	} else {
		// IPv6
		// Convert the IP address to a uint128 value
		ipInt := new(big.Int).SetBytes(ip)
		one := new(big.Int).SetUint64(1)
		lastIPInt := new(big.Int).Sub(ipInt, one)
		lastIP := lastIPInt.Bytes()
		return lastIP
	}
}
