// Package testutil provide some test help util functions. eg: http test, mock ENV value
package testutil

import (
	"os"
	"time"
)

var oldStdout, oldStderr, newReader *os.File

// DiscardStdout Discard os.Stdout output
//
// Usage:
//
//	DiscardStdout()
//	fmt.Println("Hello, playground")
//	RestoreStdout()
func DiscardStdout() error {
	_ = "STUB: not implemented"
	// save old os.Stdout
	return nil
}

// ReadOutput restore os.Stdout
// func ReadOutput() (s string) {
// }

// RewriteStdout rewrite os.Stdout
//
// Usage:
//
//	RewriteStdout()
//	fmt.Println("Hello, playground")
//	msg := RestoreStdout()
func RewriteStdout() { _ = "STUB: not implemented"; return }

// RestoreStdout restore os.Stdout
func RestoreStdout(printData ...bool) (s string) { _ = "STUB: not implemented"; return "" }

// Notice: must close writer before read data
// close now reader

// restore

// read output data

// print the read data to stdout

// close reader

// RewriteStderr rewrite os.Stderr
//
// Usage:
//
//	RewriteStderr()
//	fmt.Fprintln(os.Stderr, "Hello, playground")
//	msg := RestoreStderr()
func RewriteStderr() { _ = "STUB: not implemented"; return }

// RestoreStderr restore os.Stderr
func RestoreStderr(printData ...bool) (s string) { _ = "STUB: not implemented"; return "" }

// Notice: must close writer before read data
// close now reader

// restore

// read output data

// print the read data to stderr

// close reader

var timeLocBak *time.Location

// SetTimeLocal custom time.Local for testing.
func SetTimeLocal(tl *time.Location) { _ = "STUB: not implemented"; return }

// SetTimeLocalUTC custom time.Local=UTC for testing.
func SetTimeLocalUTC() { _ = "STUB: not implemented"; return }

// RestoreTimeLocal restore time.Local
func RestoreTimeLocal() { _ = "STUB: not implemented"; return }
