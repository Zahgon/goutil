//go:build !windows

package termenv

// detect special term color support on macOS, linux, unix
func detectSpecialTermColor(termVal string) (ColorLevel, bool) {
	_ = "STUB: not implemented"

	// detect WSL as it has True Color support
	// on Windows WSL:
	// - runtime.GOOS == "Linux"
	// - support true-color
	return *new(ColorLevel), false
}

// on TERM=screen:
// - support 256, not support true-color. test on macOS

func syscallStdinFd() int { _ = "STUB: not implemented"; return 0 }

func syscallStdoutFd() int { _ = "STUB: not implemented"; return 0 }
