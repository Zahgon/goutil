package fakeobj

import (
	"io/fs"
	"time"
)

// DirEntry implements the fs.DirEntry
type DirEntry struct {
	Dir bool
	Nam string // basename
	Mod fs.FileMode
	Fi  fs.FileInfo
	Err error // for test info error
}

// NewDirEntry create a fs.DirEntry
func NewDirEntry(fPath string, isDir ...bool) *DirEntry { _ = "STUB: not implemented"; return nil }

// Name get
func (d *DirEntry) Name() string {
	_ = "STUB: not implemented"

	// IsDir get
	return ""
}

func (d *DirEntry) IsDir() bool {
	_ = "STUB: not implemented"

	// Type get
	return false
}

func (d *DirEntry) Type() fs.FileMode {
	_ = "STUB: not implemented"

	// Info get
	return *new(fs.FileMode)
}

func (d *DirEntry) Info() (fs.FileInfo, error) {
	_ = "STUB: not implemented"
	return *new(fs.FileInfo), nil
}

// FileInfo implements the fs.FileInfo, fs.File
type FileInfo struct {
	Dir bool
	Nam string // basename
	Mod fs.FileMode
	Mt  time.Time

	// Path full path
	Path string

	Contents string
	CloseErr error
	offset   int
}

// NewFile instance
func NewFile(fPath string) *FileInfo { _ = "STUB: not implemented"; return nil }

// NewFileInfo instance
func NewFileInfo(fPath string, isDir ...bool) *FileInfo { _ = "STUB: not implemented"; return nil }

// WithBody set file body contents
func (f *FileInfo) WithBody(s string) *FileInfo { _ = "STUB: not implemented"; return nil }

// WithMtime set file modify time
func (f *FileInfo) WithMtime(mt time.Time) *FileInfo { _ = "STUB: not implemented"; return nil }

// Reset prepares a FileInfo for reuse.
func (f *FileInfo) Reset() *FileInfo { _ = "STUB: not implemented"; return nil }

// fs.File methods.

// Stat returns the FileInfo structure describing file.
func (f *FileInfo) Stat() (fs.FileInfo, error) {
	_ = "STUB: not implemented"

	// Read reads up to len(p) bytes into p.
	return *new(fs.FileInfo), nil
}

func (f *FileInfo) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// Close closes the file
func (f *FileInfo) Close() error {
	_ = "STUB: not implemented"

	// fs.FileInfo methods.
	return nil
}

// Name returns the base name of the file.
func (f *FileInfo) Name() string {
	_ = "STUB: not implemented"

	// Size returns the length in bytes for regular files; system-dependent for others.
	return ""
}

func (f *FileInfo) Size() int64 { _ = "STUB: not implemented"; return 0 }

// Mode returns file mode bits.
func (f *FileInfo) Mode() fs.FileMode {
	_ = "STUB: not implemented"

	// ModTime returns the modification time.
	return *new(fs.FileMode)
}

func (f *FileInfo) ModTime() time.Time {
	_ = "STUB: not implemented"

	// IsDir returns true if the file is a directory.
	return *new(time.Time)
}

func (f *FileInfo) IsDir() bool {
	_ = "STUB: not implemented"

	// Sys returns underlying data source (can return nil).
	return false
}

func (f *FileInfo) Sys() any { _ = "STUB: not implemented"; return *new(any) }
