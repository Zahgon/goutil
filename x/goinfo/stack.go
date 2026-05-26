package goinfo

import (
	"runtime"
)

// some commonly consts
var (
	DefStackLen = 10000
	MaxStackLen = 100000
)

// GetCallStacks stacks is a wrapper for runtime.
// If all is true, Stack that attempts to recover the data for all goroutines.
//
// from glog package
func GetCallStacks(all bool) []byte {
	_ = "STUB: not implemented"
	// We don't know how big the traces are, so grow a few times if they don't fit.
	// Start large, though.
	return nil
}

// 4<<10 // 4 KB should be enough

// GetCallerInfo get caller func name and with base filename and line.
//
// returns:
//
//	github.com/gookit/goutil/x/goinfo_test.someFunc2(),stack_test.go:26
func GetCallerInfo(skip int) string {
	_ = "STUB: not implemented"
	// ignore current func
	return ""
}

// SimpleCallersInfo returns an array of strings containing
// the func name, file and line number of each stack frame leading.
func SimpleCallersInfo(skip, num int) []string {
	_ = "STUB: not implemented"
	// ignore current func
	return nil
}

// GetCallersInfo returns an array of strings containing
// the func name, file and line number of each stack frame leading.
//
// NOTICE: max should > skip
func GetCallersInfo(skip, max int) []string { _ = "STUB: not implemented"; return nil }

// The breaks below failed to terminate the loop, and we ran off the
// end of the call stack.

// This is a huge edge case, but it will panic if this is the case

// eg: github.com/gookit/goutil/x/goinfo_test.someFunc2(),stack_test.go:26

// Drop the package
// segments := strings.Split(name, ".")
// name = segments[len(segments)-1]

// CallerInfo struct
type CallerInfo struct {
	PC   uintptr
	Fc   *runtime.Func
	File string
	Line int
}

// String convert
func (ci *CallerInfo) String() string { _ = "STUB: not implemented"; return "" }

// CallerFilterFunc type
type CallerFilterFunc func(file string, fc *runtime.Func) bool

// CallersInfos returns an array of the CallerInfo, can with filters
//
// Usage:
//
//	cs := sysutil.CallersInfos(3, 2)
//	for _, ci := range cs {
//		fc := runtime.FuncForPC(pc)
//		// maybe need check fc = nil
//		fnName = fc.Name()
//	}
func CallersInfos(skip, num int, filters ...CallerFilterFunc) []*CallerInfo {
	_ = "STUB: not implemented"
	return nil
}

// The breaks below failed to terminate the loop, and we ran off the
// end of the call stack.

// filter - return false for skip

// collecting
