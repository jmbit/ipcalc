package maths

import (
	"encoding/binary"
	"fmt"
	"net"
)

func GetBroadcastAddress(network *net.IPNet) (net.IP, error) {
	ip := network.IP.To16()

	if ip.To4() != nil {
		ipInt := binary.BigEndian.Uint32(ip[12:])
		mask := binary.BigEndian.Uint32(network.Mask)
		addresses := ^mask + 1
		broadcastIPInt := ipInt | (addresses - 1)
		broadcastIP := make(net.IP, 4)
		binary.BigEndian.PutUint32(broadcastIP, broadcastIPInt)
		return broadcastIP, nil
	} else {
		return nil, fmt.Errorf("IPv6 has no network broadcast address.")
	}

}
