package maths

import (
	"encoding/binary"
	"math/big"
	"net"
)

// GetFirstIPAddress calculates the first IP address in a given network range
func GetFirstIPAddress(network *net.IPNet) net.IP {
	// Convert the network IP to a full 16-byte representation
	ip := network.IP.To16()

	// Check if the IP address is IPv4 or IPv6
	if ip.To4() != nil {
		// IPv4
		ipInt := binary.BigEndian.Uint32(ip[12:])
		firstIPInt := ipInt + 1
		firstIP := make(net.IP, 4)
		binary.BigEndian.PutUint32(firstIP, firstIPInt)
		return firstIP
	} else {
		// IPv6
		// Convert the IP address to a uint128 value
		ipInt := new(big.Int).SetBytes(ip)
		one := new(big.Int).SetUint64(1)
		// Calculate the first host address by adding 1 to the network address
		firstIPInt := new(big.Int).Add(ipInt, one)
		firstIP := firstIPInt.Bytes()
		return firstIP
	}
}

// GetSecondIPAddress calculates the second IP address in a given network range
func GetSecondIPAddress(network *net.IPNet) net.IP {
	// Convert the network IP to a full 16-byte representation
	ip := network.IP.To16()

	// Check if the IP address is IPv4 or IPv6
	if ip.To4() != nil {
		// IPv4
		ipInt := binary.BigEndian.Uint32(ip[12:])
		firstIPInt := ipInt + 2
		firstIP := make(net.IP, 4)
		binary.BigEndian.PutUint32(firstIP, firstIPInt)
		return firstIP
	} else {
		// IPv6
		// Convert the IP address to a uint128 value
		ipInt := new(big.Int).SetBytes(ip)
		two := new(big.Int).SetUint64(2)
		// Calculate the first host address by adding 2 to the network address
		firstIPInt := new(big.Int).Add(ipInt, two)
		firstIP := firstIPInt.Bytes()
		return firstIP
	}
}
