# ip-extras

[![Version](https://img.shields.io/github/v/tag/Andrey-Yurevich/ip-extras?label=version)](https://github.com/Andrey-Yurevich/ip-extras/releases)
[![Go Reference](https://pkg.go.dev/badge/github.com/Andrey-Yurevich/ip-extras.svg)](https://pkg.go.dev/github.com/Andrey-Yurevich/ip-extras)

A small Go library providing utilities for performing arithmetic and structural
operations on IPv4 and IPv6 addresses.

## Documentation

- **PreviousIPv6** — returns the IPv6 address immediately preceding the given IP.  
- **NextIPv6** — returns the IPv6 address immediately following the given IP.  
- **FirstIPv6** — returns the first (network) IPv6 address in the subnet.  
- **LastIPv6** — returns the last IPv6 address in the subnet.  
- **ContainsSubnetV6** — reports whether one IPv6 subnet fully contains another.  
- **IsSingleV6HostMask** — checks whether the mask is a /128 single-host IPv6 mask.  
- **IsIPv6LinkLocal** — reports whether an IPv6 address is link‑local (fe80::/10).  
- **PreviousIPv4** — returns the IPv4 address immediately preceding the given IP.  
- **NextIPv4** — returns the IPv4 address immediately following the given IP.  
- **FirstIPv4** — returns the first (network) IPv4 address in the subnet.  
- **LastIPv4** — returns the last IPv4 address in the subnet.  
- **ContainsSubnetV4** — reports whether one IPv4 subnet fully contains another.  
- **IsSingleV4HostMask** — checks whether the mask is a /32 single-host IPv4 mask.  
- **ipToBigInt** — converts an IPv4 or IPv6 address into a big.Int.  
- **IsIPInRange** — checks whether an IP lies within a byte range.  
- **MaxIP** — returns the greater of two IP addresses.  
- **MinIP** — returns the smaller of two IP addresses.