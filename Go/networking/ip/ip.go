package ip

import (
	"net"
)

func parsIp(ip string) net.IP {
	return net.ParseIP(ip)
}

// RFC 1918 Address Allocation for Private Internets
// RFC 4193 Unice Local Ipv6 Unicast Address
func isItPrivate(ip net.IP) bool {
	return ip.IsPrivate()
}
