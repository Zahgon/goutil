//go:build !windows && !darwin
// +build !windows,!darwin

package process

// Exists check process running by given pid
func Exists(pid int) bool { _ = "STUB: not implemented"; return false }
