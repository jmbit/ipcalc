package maths

import (
	"fmt"
	"log"
	"math/big"
	"net"
)

// Convert IP to big integer
func ipToBigInt(ip net.IP) *big.Int {
	ipBigInt := big.NewInt(0)
	ipBigInt.SetBytes(ip.To4())
	return ipBigInt
}

// Calculate the smallest subnet. Currently still has some weird edge cases
func CalculateSmallestSubnet(ip1 net.IP, ip2 net.IP) (*net.IPNet, error) {
	ip1BigInt := ipToBigInt(ip1)
	ip2BigInt := ipToBigInt(ip2)

	diff := big.NewInt(0)
	diff.Sub(ip1BigInt, ip2BigInt)
	diff.Abs(diff)

	mask := 32
	for diff.BitLen() > 0 {
		diff.Rsh(diff, 1)
		mask--
	}

	log.Println(ip1, mask)
	_, cidr, err := net.ParseCIDR(fmt.Sprintf("%s/%d", ip1, mask))

	return cidr, err
}
