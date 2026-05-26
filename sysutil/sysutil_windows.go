//go:build windows
// +build windows

package sysutil

import (
	"syscall"
)

// OsName system name. like runtime.GOOS. only allow: linux, windows, darwin
const OsName = Windows

// IsWin system. linux windows darwin
func IsWin() bool {
	_ = "STUB: not implemented"

	// IsWindows system. linux windows darwin
	return false
}

func IsWindows() bool {
	_ = "STUB: not implemented"

	// IsMac system
	return false
}

func IsMac() bool {
	_ = "STUB: not implemented"

	// IsDarwin system
	return false
}

func IsDarwin() bool {
	_ = "STUB: not implemented"

	// IsLinux system
	return false
}

func IsLinux() bool {
	_ = "STUB: not implemented"

	// Kill a process by pid
	return false
}

func Kill(pid int, signal syscall.Signal) error { _ = "STUB: not implemented"; return nil }

// ProcessExists check process exists by pid
func ProcessExists(pid int) bool { _ = "STUB: not implemented"; return false }

// OpenURL Open file or browser URL
//
// - refers https://github.com/pkg/browser
//
// Mac：
//
//	open 'https://github.com/inhere'
//
// Linux:
//
//	xdg-open URL
//	x-www-browser 'https://github.com/inhere'
//
// Windows:
//
//	cmd /c start https://github.com/inhere
func OpenURL(url string) error {
	_ = "STUB: not implemented"
	// return exec.Command("cmd", "/C", "start", URL).Run()
	return nil
}
