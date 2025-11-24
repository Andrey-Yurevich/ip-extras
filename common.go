package ipextras

import (
	"bytes"
	"math/big"
	"net"
)

// ipToBigInt converts an IP (IPv4 or IPv6) into a big.Int numeric representation.
func ipToBigInt(ip net.IP) *big.Int {
	if ip4 := ip.To4(); ip4 != nil {
		// IPv4
		return new(big.Int).SetBytes(ip4)
	}
	// IPv6
	ip16 := ip.To16()
	if ip16 == nil {
		return nil
	}
	return new(big.Int).SetBytes(ip16)
}

// IsIPInRange reports whether ipToCheck falls within the byte-range [startIP, endIP].
func IsIPInRange(startIP, endIP, ipToCheck []byte) bool {

	if endIP == nil && startIP != nil {
		return bytes.Equal(startIP, ipToCheck)
	}

	if startIP == nil && endIP != nil {
		return bytes.Equal(endIP, ipToCheck)
	}

	return bytes.Compare(ipToCheck, startIP) >= 0 && bytes.Compare(ipToCheck, endIP) <= 0
}

// MaxIP returns the lexicographically greater of two IP addresses.
func MaxIP(a, b net.IP) net.IP {

	if a == nil && b != nil {
		return b
	}

	if b == nil && a != nil {
		return a
	}

	if bytes.Compare(a, b) > 0 {
		return a
	}
	return b
}

// MinIP returns the lexicographically smaller of two IP addresses.
func MinIP(a, b net.IP) net.IP {

	if a == nil && b != nil {
		return b
	}

	if b == nil && a != nil {
		return a
	}

	if bytes.Compare(a, b) < 0 {
		return a
	}
	return b
}
