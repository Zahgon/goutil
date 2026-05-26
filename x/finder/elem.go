package finder

import (
	"io/fs"
)

// Elem of find file/dir path result
type Elem interface {
	fs.DirEntry
	// Path gets file/dir full path. eg: "/path/to/file.go"
	Path() string
	// Info get file info. like fs.DirEntry.Info(), but will cache result.
	Info() (fs.FileInfo, error)
}

type elem struct {
	fs.DirEntry
	path string
	stat fs.FileInfo
	sErr error
}

// NewElem create a new Elem instance
func NewElem(fPath string, ent fs.DirEntry) Elem { _ = "STUB: not implemented"; return *new(Elem) }

// Path gets full file/dir path. eg: "/path/to/file.go"
func (e *elem) Path() string {
	_ = "STUB: not implemented"

	// Info gets file info, will cache result
	return ""
}

func (e *elem) Info() (fs.FileInfo, error) {
	_ = "STUB: not implemented"
	return *new(fs.FileInfo), nil
}

// String get string representation
func (e *elem) String() string { _ = "STUB: not implemented"; return "" }
