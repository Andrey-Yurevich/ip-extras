package ipextras

import (
	"net"
)

// PreviousIPv6 returns the IPv6 address immediately preceding the given IP.
func PreviousIPv6(ip net.IP) net.IP {
	ip = ip.To16()
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

// NextIPv6 returns the IPv6 address immediately following the given IP.
func NextIPv6(ip net.IP) net.IP {
	ip = ip.To16()
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

// FirstIPv6 calculates the first (network) IPv6 address within the provided subnet.
func FirstIPv6(network *net.IPNet) net.IP {
	ip := network.IP.To16()
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

// LastIPv6 calculates the last IPv6 address within the provided subnet.
func LastIPv6(network *net.IPNet) net.IP {
	ip := network.IP.To16()
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

// ContainsSubnetV6 reports whether the supernet fully contains the sub subnet.
func ContainsSubnetV6(super, sub *net.IPNet) bool {
	return super.Contains(sub.IP) && super.Contains(LastIPv6(sub))
}

// IsSingleV6HostMask reports whether the mask corresponds to a single-host IPv6 mask (/128).
func IsSingleV6HostMask(mask net.IPMask) bool {
	if len(mask) != net.IPv6len {
		return false
	}
	for _, b := range mask {
		if b != 0xFF {
			return false
		}
	}
	return true
}

// IsIPv6LinkLocal reports whether the given IPv6 address is a link-local address (fe80::/10).
func IsIPv6LinkLocal(ip net.IP) bool {
	if ip == nil || ip.To4() != nil {
		return false
	}
	ip = ip.To16()
	// checking first 10 bits: 1111111010 == fe80::/10
	return ip[0] == 0xfe && (ip[1]&0xc0) == 0x80
}
