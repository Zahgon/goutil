package envutil

import (
	"io"
)

// IsWin system. linux windows darwin
func IsWin() bool { _ = "STUB: not implemented"; return false }

// IsWindows system. alias of IsWin
func IsWindows() bool { _ = "STUB: not implemented"; return false }

// IsMac system
func IsMac() bool { _ = "STUB: not implemented"; return false }

// IsLinux system
func IsLinux() bool { _ = "STUB: not implemented"; return false }

// IsMSys msys(MINGW64) env. alias of the sysutil.IsMSys()
func IsMSys() bool { _ = "STUB: not implemented"; return false }

// IsTerminal isatty check
//
// Usage:
//
//	envutil.IsTerminal(os.Stdout.Fd())
func IsTerminal(fd uintptr) bool {
	_ = "STUB: not implemented"
	// return isatty.IsTerminal(fd) // "github.com/mattn/go-isatty"
	return false
}

// StdIsTerminal os.Stdout is terminal
func StdIsTerminal() bool { _ = "STUB: not implemented"; return false }

// IsConsole check out is console env. alias of the sysutil.IsConsole()
func IsConsole(out io.Writer) bool { _ = "STUB: not implemented"; return false }

// HasShellEnv has shell env check.
//
// Usage:
//
//	HasShellEnv("sh")
//	HasShellEnv("bash")
func HasShellEnv(shell string) bool { _ = "STUB: not implemented"; return false }

// Support color:
//
//	"TERM=xterm"
//	"TERM=xterm-vt220"
//	"TERM=xterm-256color"
//	"TERM=screen-256color"
//	"TERM=tmux-256color"
//	"TERM=rxvt-unicode-256color"
//
// Don't support color:
//
//	"TERM=cygwin"
var specialColorTerms = map[string]bool{
	"alacritty": true,
}

// IsSupportColor check current console is support color.
//
// Supported:
//
//	linux, mac, or Windows's ConEmu, Cmder, putty, git-bash.exe
//
// Not support:
//
//	windows cmd.exe, powerShell.exe
func IsSupportColor() bool { _ = "STUB: not implemented"; return false }

// it's special color term

// like on ConEmu software, e.g "ConEmuANSI=ON"

// like on ConEmu software, e.g "ANSICON=189x2000 (189x43)"

// up: if support 256-color, can also support basic color.

// IsSupport256Color render
func IsSupport256Color() bool {
	_ = "STUB: not implemented"
	// "TERM=xterm-256color"
	// "TERM=screen-256color"
	// "TERM=tmux-256color"
	// "TERM=rxvt-unicode-256color"
	return false
}

// up: if support true-color, can also support 256-color.

// IsSupportTrueColor render. IsSupportRGBColor
func IsSupportTrueColor() bool {
	_ = "STUB: not implemented"
	// "COLORTERM=truecolor"
	return false
}

// IsGithubActions env
func IsGithubActions() bool { _ = "STUB: not implemented"; return false }
