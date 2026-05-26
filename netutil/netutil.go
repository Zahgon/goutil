// Package netutil provide some network util functions.
package netutil

// FreePort returns a free port.
func FreePort() (port int, err error) { _ = "STUB: not implemented"; return 0, nil }

// AllMacAddrs get all mac addresses
func AllMacAddrs() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

// 跳过回环接口（如 lo）和未启用的接口

// FirstMacAddr 获取第一个非lo网卡的MAC地址
func FirstMacAddr() (string, error) { _ = "STUB: not implemented"; return "", nil }

// 跳过回环接口（如 lo）和未启用的接口

// 获取MAC地址
