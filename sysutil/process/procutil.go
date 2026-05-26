// Package process Provide some process handle util functions
package process

import (
	"os"
	"syscall"
)

// PID get current process ID
func PID() int {
	_ = "STUB: not implemented"

	// Start starts a new process with the program, arguments and attributes
	// specified by name, argv and attr.
	//
	// alias of os.StartProcess()
	return 0
}

func Start(name string, argv []string, attr *os.ProcAttr) (*os.Process, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ProcInfo looks for a running process by its pid.
//
// alias of os.FindProcess()
func ProcInfo(pid int) (*os.Process, error) {
	_ = "STUB: not implemented"
	return nil,

		// PIDByName get PID by process name match(by pgrep)
		nil
}

func PIDByName(keywords string) int {
	_ = "STUB: not implemented"
	// pgrep keywords
	return 0
}

// KillByName kill process by name match
func KillByName(keywords string, sig syscall.Signal) error { _ = "STUB: not implemented"; return nil }

// StopProcessOption stop process option
type StopProcessOption struct {
	// Check if the process exists before stopping it
	CheckExist bool
	// Whether to force exit the process
	ForceKill bool
	// Whether to wait for the process to exit
	WaitExit bool
	// How long to wait for the process to exit
	ExitTimeout int
	// Signal to send to the process(non-win)
	Signal syscall.Signal
}
