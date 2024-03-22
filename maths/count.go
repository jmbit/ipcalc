package maths

import (
	"math/big"
	"net"
)

// AddressCount calculates the number of IP addresses in a given network
func AddressCount(network *net.IPNet) int {
	ones, bits := network.Mask.Size()
	numHosts := new(big.Int).Exp(big.NewInt(2), big.NewInt(int64(bits-ones)), nil)
	if network.IP.To4() != nil {
		numHosts.Sub(numHosts, big.NewInt(2))
	}
	return int(numHosts.Int64())

}
