//go:build !windows

package process

import "syscall"

// Kill a process by pid
func Kill(pid int, signal syscall.Signal) error { _ = "STUB: not implemented"; return nil }

// ExistsByName check process running by given name
func ExistsByName(name string, fuzzyMatch bool) bool {
	_ = "STUB: not implemented"
	// TODO
	return false
}

// StopByName Stop process based on process name.
//
// return (exists, output, error). check error to see if the process exists
//
// Usage:
//
//	StopByName("MyApp.exe")
func StopByName(name string, option ...*StopProcessOption) (bool, string, error) {
	_ = "STUB: not implemented"
	return false,
		// TODO
		"", nil
}
