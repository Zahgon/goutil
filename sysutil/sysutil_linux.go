package sysutil

// OsName system name. like runtime.GOOS. allow: linux, windows, darwin
const OsName = Linux

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

	// There are multiple possible providers to open a browser on linux
	// One of them is xdg-open, another is x-www-browser, then there's www-browser, etc.
	// Look for one that exists and run it
	return false
}

var openBins = []string{"xdg-open", "x-www-browser", "www-browser"}

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
func OpenURL(URL string) error { _ = "STUB: not implemented"; return nil }
