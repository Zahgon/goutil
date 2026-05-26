package sysutil

import (
	"io"
)

// os names
const (
	Windows = "windows"
	Linux   = "linux"
	Darwin  = "darwin"
	FreeBSD = "freebsd"
)

// IsMSys msys(MINGW64) env，不一定支持颜色
func IsMSys() bool {
	_ = "STUB: not implemented"
	// "MSYSTEM=MINGW64"
	return false
}

// IsWSL system env
func IsWSL() bool { _ = "STUB: not implemented"; return false }

// IsConsole check out is in stderr/stdout/stdin
//
// Usage:
//
//	sysutil.IsConsole(os.Stdout)
func IsConsole(out io.Writer) bool { _ = "STUB: not implemented"; return false }

// fix: cannot use 'o == os.Stdout' to compare

// IsTerminal isatty check
//
// Usage:
//
//	sysutil.IsTerminal(os.Stdout.Fd())
func IsTerminal(fd uintptr) bool {
	_ = "STUB: not implemented"
	// return isatty.IsTerminal(fd) // "github.com/mattn/go-isatty"
	return false
}

// StdIsTerminal os.Stdout is terminal
func StdIsTerminal() bool { _ = "STUB: not implemented"; return false }

// Hostname is alias of os.Hostname, but ignore error
func Hostname() string { _ = "STUB: not implemented"; return "" }

// CurrentShell get current used shell env file.
//
// eg "/bin/zsh" "/bin/bash".
// if onlyName=true, will return "zsh", "bash"
func CurrentShell(onlyName bool, fallbackShell ...string) string {
	_ = "STUB: not implemented"
	return ""
}

// HasShellEnv has shell env check.
//
// Usage:
//
//	HasShellEnv("sh")
//	HasShellEnv("bash")
func HasShellEnv(shell string) bool {
	_ = "STUB: not implemented"
	// can also use: "echo $0"
	return false
}

// IsShellSpecialVar reports whether the character identifies a special
// shell variable such as $*.
func IsShellSpecialVar(c uint8) bool { _ = "STUB: not implemented"; return false }

// FindExecutable in the system
//
// Usage:
//
//	sysutil.FindExecutable("bash")
func FindExecutable(binName string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// Executable find in the system, alias of FindExecutable()
//
// Usage:
//
//	sysutil.Executable("bash")
func Executable(binName string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// HasExecutable in the system
//
// Usage:
//
//	HasExecutable("bash")
func HasExecutable(binName string) bool { _ = "STUB: not implemented"; return false }

// Getenv get ENV value by key name, can with default value
func Getenv(name string, def ...string) string { _ = "STUB: not implemented"; return "" }

// Environ like os.Environ, but will returns key-value map[string]string data.
func Environ() map[string]string { _ = "STUB: not implemented"; return nil }

// EnvMapWith like os.Environ, but will return key-value map[string]string data.
func EnvMapWith(newEnv map[string]string) map[string]string { _ = "STUB: not implemented"; return nil }

// EnvPaths get and split $PATH to []string
func EnvPaths() []string { _ = "STUB: not implemented"; return nil }

// ToEnvPATH convert []string to $PATH string.
func ToEnvPATH(paths []string) string { _ = "STUB: not implemented"; return "" }

// SearchPathOption settings for SearchPath
type SearchPathOption struct {
	// 限制查找的扩展名(Windows) such as ".exe", ".bat", ".cmd"
	LimitExt []string
}

// SearchPath search executable files in the system $PATH
//
// Usage:
//
//	sysutil.SearchPath("go")
func SearchPath(keywords string, limit int, opts ...SearchPathOption) []string {
	_ = "STUB: not implemented"
	return nil
}

// if windows, will limit with .exe, .bat, .cmd

// Unix shell semantics: path element "" means "."

// mark dir is checked

// if windows, will limit with .exe, .bat, .cmd

// limit result size
