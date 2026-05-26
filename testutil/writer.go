package testutil

import (
	"github.com/gookit/goutil/x/fakeobj"
)

// TestWriter struct, useful for testing. alias of fakeobj.Writer
type TestWriter = fakeobj.Writer

// NewTestWriter instance
func NewTestWriter() *TestWriter { _ = "STUB: not implemented"; return nil }

// DirEnt implements the fs.DirEntry
type DirEnt = fakeobj.DirEntry

// NewDirEnt create a fs.DirEntry
func NewDirEnt(fPath string, isDir ...bool) *fakeobj.DirEntry {
	_ = "STUB: not implemented"
	return nil
}
