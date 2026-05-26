//go:build freebsd || openbsd || netbsd || dragonfly

package sysutil

import (
	"runtime"
)

// OsName system name. like runtime.GOOS.
// For Unix systems, this will be the actual GOOS value (freebsd, openbsd, netbsd, dragonfly)
var OsName = runtime.GOOS

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

	// There are multiple possible providers to open a browser on Unix systems
	// Similar to Linux, try xdg-open and other common browser launchers
	return false
}

var openBins = []string{"xdg-open", "x-www-browser", "www-browser", "firefox", "chrome", "chromium"}

// OpenURL Open file or browser URL
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
//
// Unix (FreeBSD, OpenBSD, etc.):
//
//	Try xdg-open, x-www-browser, or fallback browsers
func OpenURL(URL string) error { _ = "STUB: not implemented"; return nil }
