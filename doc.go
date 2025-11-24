/*
PACKAGE DOCUMENTATION

Package ipextras provides a collection of helper functions for working with
IPv4 and IPv6 addresses, including arithmetic on IPs, subnet boundary
calculations, range checks, and mask inspection utilities.

IP ADDRESS ARITHMETIC

	PreviousIPv6(ip net.IP) net.IP
	    PreviousIPv6 returns the IPv6 address immediately preceding the
	    given IP. If the address is ::0, it wraps around to ffff:...:ffff.

	NextIPv6(ip net.IP) net.IP
	    NextIPv6 returns the IPv6 address immediately following the
	    given IP. If the address is ffff:...:ffff, it wraps around to ::0.

	PreviousIPv4(ip net.IP) net.IP
	    PreviousIPv4 returns the IPv4 address immediately preceding the
	    given IP. If the address is 0.0.0.0, it wraps to 255.255.255.255.

	NextIPv4(ip net.IP) net.IP
	    NextIPv4 returns the IPv4 address immediately following the
	    given IP. If the address is 255.255.255.255, it wraps to 0.0.0.0.

SUBNET BOUNDARY CALCULATIONS

	FirstIPv6(network *net.IPNet) net.IP
	    FirstIPv6 returns the first (network) IPv6 address in the given
	    subnet by applying the network mask.

	LastIPv6(network *net.IPNet) net.IP
	    LastIPv6 returns the last IPv6 address in the given subnet by
	    inverting the mask and applying it to the network address.

	FirstIPv4(network *net.IPNet) net.IP
	    FirstIPv4 returns the first (network) IPv4 address in the subnet.

	LastIPv4(network *net.IPNet) net.IP
	    LastIPv4 returns the last IPv4 address in the subnet.

SUBNET CONTAINMENT

	ContainsSubnetV6(super, sub *net.IPNet) bool
	    ContainsSubnetV6 reports whether the supernet completely contains
	    the sub subnet, inclusive of its last address.

	ContainsSubnetV4(super, sub *net.IPNet) bool
	    ContainsSubnetV4 reports whether the supernet completely contains
	    the sub subnet, inclusive of its last address.

MASK UTILITIES

	IsSingleV6HostMask(mask net.IPMask) bool
	    IsSingleV6HostMask reports whether the mask corresponds to a
	    single-host IPv6 mask (/128).

	IsSingleV4HostMask(mask net.IPMask) bool
	    IsSingleV4HostMask reports whether the mask corresponds to a
	    single-host IPv4 mask (/32).

IP CLASSIFICATION

	IsIPv6LinkLocal(ip net.IP) bool
	    IsIPv6LinkLocal reports whether the IPv6 address falls within the
	    fe80::/10 link-local range.

RANGE AND COMPARISON UTILITIES

	ipToBigInt(ip net.IP) *big.Int
	    ipToBigInt converts an IP address (IPv4 or IPv6) into a big-endian
	    big.Int representation. Returns nil for invalid IPs.

	IsIPInRange(startIP, endIP, ip []byte) bool
	    IsIPInRange reports whether ip is within the inclusive byte range
	    [startIP, endIP]. A nil bound acts as an exact match.

	MaxIP(a, b net.IP) net.IP
	    MaxIP returns the lexicographically greater of two IP addresses.

	MinIP(a, b net.IP) net.IP
	    MinIP returns the lexicographically smaller of two IP addresses.
*/
package ipextras
