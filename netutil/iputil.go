package netutil

import (
	"net"
)

// AllLocalIPv4 Get all non-loop IPv4 addresses. 获取所有非回环 IPv4 地址
func AllLocalIPv4() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func filterIpv4(addrs []net.Addr) []string { _ = "STUB: not implemented"; return nil }

// GetLocalIPs get local IPs(ipv4+ipv6), will panic on error.
func GetLocalIPs() (ips []string) { _ = "STUB: not implemented"; return nil }

// InternalIP get local first internal IP v4/v6 addr.
func InternalIP() string { _ = "STUB: not implemented"; return "" }

// InternalIPv4 get internal IPv4 for host.
func InternalIPv4() string {
	_ = "STUB: not implemented"

	// IPv4 get local first internal IPv4 addr.
	return ""
}

func IPv4() string { _ = "STUB: not implemented"; return "" }

// InternalIPv6 get first internal IPv6 addr
func InternalIPv6() string {
	_ = "STUB: not implemented"

	// IPv6 get local first internal IPv6 addr
	return ""
}

func IPv6() string { _ = "STUB: not implemented"; return "" }

// MustIPv4 get first internal IP v4 addr, will panic on error or not found.
func MustIPv4() (ip string) { _ = "STUB: not implemented"; return "" }

// get local first internal IP v4/v6 addr.
//
// ver: 4, 6 or not limit
func getLocalIP(ver uint8) string { _ = "STUB: not implemented"; return "" }

func getFirstIP(ver uint8, addrs []net.Addr) string { _ = "STUB: not implemented"; return "" }

// need ip v4

// HostIP returns the IP addresses of the local hostname.
func HostIP() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

// IncrIP 将IP地址递增1 eg: 192.168.1.1 -> 192.168.1.2
func IncrIP(ip net.IP) { _ = "STUB: not implemented"; return }
