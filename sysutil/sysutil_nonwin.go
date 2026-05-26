//go:build !windows

package sysutil

import (
	"syscall"
)

// Kill a process by pid
func Kill(pid int, signal syscall.Signal) error { _ = "STUB: not implemented"; return nil }

// ProcessExists check process exists by pid
func ProcessExists(pid int) bool { _ = "STUB: not implemented"; return false }
