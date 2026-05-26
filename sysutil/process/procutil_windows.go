package process

import (
	"syscall"
)

const (
	processQueryLimitedInformation = 0x1000

	stillActive = 259
)

// Kill a process by pid. use taskkill on windows
//
// CMD example:
//
//	taskkill /pid 1234
//	taskkill /pid 1234 /f
func Kill(pid int, signal syscall.Signal) error { _ = "STUB: not implemented"; return nil }

// Exists check a process running by given pid
func Exists(pid int) bool { _ = "STUB: not implemented"; return false }

// ExistsByName Determine whether a process exists based on its name(by tasklist)
//
// Usage:
//
//	ExistsByName("MyApp.exe")
//	// Fuzzy match by input name
//	ExistsByName("MyApp", true)
func ExistsByName(name string, fuzzyMatch ...bool) bool {
	_ = "STUB: not implemented"
	// 按名称模糊匹配
	return false
}

// out, err := sysutil.ShellExec("tasklist /FI \"IMAGENAME eq "+name+"\" /NH", "cmd") // shell执行有问题

// StopByName Stop process based on process name(by taskkill).
//
// return (exists, output, error). check error to see if the process exists
//
// Usage:
//
//	StopByName("MyApp.exe")
func StopByName(name string, option ...*StopProcessOption) (bool, string, error) {
	_ = "STUB: not implemented"
	return false, "", nil
}

// 1. 检查进程是否存在

// cmd: taskkill /IM name.exe /F
