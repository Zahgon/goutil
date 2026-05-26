// Package cliutil provides some util functions for CLI
package cliutil

import (
	"github.com/gookit/goutil/strutil"
)

// SplitMulti split multi string by sep string.
func SplitMulti(ss []string, sep string) []string { _ = "STUB: not implemented"; return nil }

// LineBuild build command line string by given args.
func LineBuild(binFile string, args []string) string { _ = "STUB: not implemented"; return "" }

// BuildLine build command line string by given args.
func BuildLine(binFile string, args []string) string { _ = "STUB: not implemented"; return "" }

// String2OSArgs parse input command line text to os.Args
func String2OSArgs(line string) []string { _ = "STUB: not implemented"; return nil }

// StringToOSArgs parse input command line text to os.Args
func StringToOSArgs(line string) []string { _ = "STUB: not implemented"; return nil }

// ParseLine input command line text. alias of the StringToOSArgs()
func ParseLine(line string) []string { _ = "STUB: not implemented"; return nil }

// QuickExec quick exec a simple command line
func QuickExec(cmdLine string, workDir ...string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ExecLine quick exec an command line string
func ExecLine(cmdLine string, workDir ...string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ExecCmd a CLI bin file and return output.
//
// Usage:
//
//	ExecCmd("ls", []string{"-al"})
func ExecCmd(binName string, args []string, workDir ...string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ExecCommand alias of the ExecCmd()
func ExecCommand(binName string, args []string, workDir ...string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ShellExec exec command by shell
//
// Usage:
// ret, err := cliutil.ShellExec("ls -al")
func ShellExec(cmdLine string, shells ...string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// CurrentShell get current used shell env file. eg "/bin/zsh" "/bin/bash"
func CurrentShell(onlyName bool) (path string) { _ = "STUB: not implemented"; return "" }

// HasShellEnv has shell env check.
//
// Usage:
//
//	HasShellEnv("sh")
//	HasShellEnv("bash")
func HasShellEnv(shell string) bool { _ = "STUB: not implemented"; return false }

// BuildOptionHelpName for render flag help
func BuildOptionHelpName(names []string) string { _ = "STUB: not implemented"; return "" }

// ShellQuote quote a string on contains ', ", SPACE
func ShellQuote(s string) string { _ = "STUB: not implemented"; return "" }

// OutputLines split output to lines
func OutputLines(output string) []string { _ = "STUB: not implemented"; return nil }

// FirstLine from command output
//
// Deprecated: please use strutil.FirstLine
var FirstLine = strutil.FirstLine
