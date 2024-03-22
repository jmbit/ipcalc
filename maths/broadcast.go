package maths

import (
	"encoding/binary"
	"fmt"
	"net"
)

// GetBroadcastAddress returns the broadcast address for a given network
func GetBroadcastAddress(network *net.IPNet) (net.IP, error) {
	ip := network.IP.To16()

	// Check if the IP address is IPv4
	if ip.To4() != nil {
		//IPv4
		// Convert the network IP to a full 16-byte representation
		ipInt := binary.BigEndian.Uint32(ip[12:])
		mask := binary.BigEndian.Uint32(network.Mask)
		// Calculate the broadcast address
		addresses := ^mask + 1
		broadcastIPInt := ipInt | (addresses - 1)
		broadcastIP := make(net.IP, 4)
		binary.BigEndian.PutUint32(broadcastIP, broadcastIPInt)
		return broadcastIP, nil
	} else {
		//IPv6
		return nil, fmt.Errorf("IPv6 has no network broadcast address.")
	}

}
