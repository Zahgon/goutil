package sysutil

import (
	"github.com/gookit/goutil/sysutil/cmdr"
)

// NewCmd instance
func NewCmd(bin string, args ...string) *cmdr.Cmd { _ = "STUB: not implemented"; return nil }

// FlushExec command, will flush output to stdout,stderr
func FlushExec(bin string, args ...string) error { _ = "STUB: not implemented"; return nil }

// QuickExec quick exec a simple command line, return combined output.
func QuickExec(cmdLine string, workDir ...string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ExecLine quick exec a command line string, return combined output.
//
//	NOTE: not support | or ; in cmdLine
func ExecLine(cmdLine string, workDir ...string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// create a new Cmd instance

// ExecCmd a command and return combined output.
//
// Usage:
//
//	ExecCmd("ls", []string{"-al"})
func ExecCmd(binName string, args []string, workDir ...string) (string, error) {
	_ = "STUB: not implemented"
	// create a new Cmd instance
	return "", nil
}

// ShellExec exec command by shell cmdLine, return combined output.
//
// shells e.g. "/bin/sh", "bash", "cmd", "cmd.exe", "powershell", "powershell.exe", "pwsh", "pwsh.exe"
//
//	eg: ShellExec("ls -al")
func ShellExec(cmdLine string, shells ...string) (string, error) {
	_ = "STUB: not implemented"
	// shell := "/bin/sh"
	return "", nil
}

// "-c" for bash,sh,zsh shell

// special for Windows shell

// use cmd.exe, mark is "/c"

// "-Command" for powershell
