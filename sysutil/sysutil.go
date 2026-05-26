// Package sysutil provide some system util functions. eg: sysenv, exec, user, process
package sysutil

// Workdir get
func Workdir() string { _ = "STUB: not implemented"; return "" }

// BinDir get
func BinDir() string { _ = "STUB: not implemented"; return "" }

// BinName get
func BinName() string { _ = "STUB: not implemented"; return "" }

// BinFile get
func BinFile() string {
	_ = "STUB: not implemented"

	// Open file or url address
	return ""
}

func Open(fileOrURL string) error { _ = "STUB: not implemented"; return nil }

// OpenBrowser file or url address
func OpenBrowser(fileOrURL string) error { _ = "STUB: not implemented"; return nil }

// OpenFile open files browser window for the file path.
func OpenFile(path string) error { _ = "STUB: not implemented"; return nil }
