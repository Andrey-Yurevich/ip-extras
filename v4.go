package ipextras

import (
	"net"
)

// PreviousIPv4 returns the IPv4 address immediately preceding the given IP.
func PreviousIPv4(ip net.IP) net.IP {
	ip = ip.To4()
	prev := make(net.IP, len(ip))
	copy(prev, ip)
	for i := len(prev) - 1; i >= 0; i-- {
		if prev[i] == 0 {
			prev[i] = 255
		} else {
			prev[i]--
			break
		}
	}
	return prev
}

// NextIPv4 returns the IPv4 address immediately following the given IP.
func NextIPv4(ip net.IP) net.IP {
	ip = ip.To4()
	next := make(net.IP, len(ip))
	copy(next, ip)
	for i := len(next) - 1; i >= 0; i-- {
		next[i]++
		if next[i] != 0 {
			break
		}
	}
	return next
}

// FirstIPv4 calculates the first (network) IPv4 address within the provided subnet.
func FirstIPv4(network *net.IPNet) net.IP {
	ip := network.IP.To4()
	if ip == nil {
		return nil
	}

	first := make(net.IP, len(ip))
	copy(first, ip)

	for i := range ip {
		first[i] &= network.Mask[i]
	}
	return first
}

// LastIPv4 calculates the last IPv4 address within the provided subnet.
func LastIPv4(network *net.IPNet) net.IP {
	ip := network.IP.To4()
	if ip == nil {
		return nil
	}

	last := make(net.IP, len(ip))
	copy(last, ip)

	for i := range ip {
		last[i] |= ^network.Mask[i]
	}
	return last
}

// ContainsSubnetV4 reports whether the supernet fully contains the sub subnet.
func ContainsSubnetV4(super, sub *net.IPNet) bool {
	return super.Contains(sub.IP) && super.Contains(LastIPv4(sub))
}

// IsSingleV4HostMask reports whether the mask corresponds to a single-host IPv4 mask (/32).
func IsSingleV4HostMask(mask net.IPMask) bool {
	if mask[0] == 255 && mask[1] == 255 && mask[2] == 255 && mask[3] == 255 {
		return true
	}
	return false
}
