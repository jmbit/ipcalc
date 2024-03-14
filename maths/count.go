package maths

import (
	"encoding/binary"
	"math/big"
	"net"
)

func AddressCount(network *net.IPNet) int {
	if network.IP.To4() != nil {
		first := binary.BigEndian.Uint32(GetFirstIPAddress(network))
		last := binary.BigEndian.Uint32(GetLastIPAddress(network))
		return int(last - first)
	} else {
		first := new(big.Int).SetBytes(GetFirstIPAddress(network))
		last := new(big.Int).SetBytes(GetLastIPAddress(network))
		// WHY DOES THIS NOT WORK?
		return int(last.Sub(last, first).Uint64())
	}

}
