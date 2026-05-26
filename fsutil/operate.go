package fsutil

import (
	"io/fs"
	"os"
)

// Mkdir alias of os.MkdirAll()
func Mkdir(dirPath string, perm fs.FileMode) error { _ = "STUB: not implemented"; return nil }

// MkdirQuick with default permission 0755.
func MkdirQuick(dirPath string) error { _ = "STUB: not implemented"; return nil }

// EnsureDir creates a directory if it doesn't exist
func EnsureDir(path string) error { _ = "STUB: not implemented"; return nil }

// MkDirs batch makes multi dirs at once
func MkDirs(perm fs.FileMode, dirPaths ...string) error { _ = "STUB: not implemented"; return nil }

// MkSubDirs batch makes multi sub-dirs at once
func MkSubDirs(perm fs.FileMode, parentDir string, subDirs ...string) error {
	_ = "STUB: not implemented"
	return nil
}

// MkParentDir quickly create parent dir for a given path.
func MkParentDir(fpath string) error { _ = "STUB: not implemented"; return nil }

// ************************************************************
//	options for open file
// ************************************************************

// OpenOption for open file
type OpenOption struct {
	// file open flag. see FsCWTFlags
	Flag int
	// file perm. see DefaultFilePerm
	Perm os.FileMode
}

// OpenOptionFunc for open/write file
type OpenOptionFunc func(*OpenOption)

// NewOpenOption create a new OpenOption instance
//
// Defaults:
//   - open flags: FsCWTFlags (override write)
//   - file Perm: DefaultFilePerm
func NewOpenOption(optFns ...OpenOptionFunc) *OpenOption { _ = "STUB: not implemented"; return nil }

// OpenOptOrNew create a new OpenOption instance if opt is nil
func OpenOptOrNew(opt *OpenOption) *OpenOption { _ = "STUB: not implemented"; return nil }

// WithFlag set file open flag
func WithFlag(flag int) OpenOptionFunc { _ = "STUB: not implemented"; return *new(OpenOptionFunc) }

// WithPerm set file perm
func WithPerm(perm os.FileMode) OpenOptionFunc {
	_ = "STUB: not implemented"
	return *new(OpenOptionFunc)
}

// ************************************************************
//	open/create files
// ************************************************************

// some commonly flag consts for open file.
const (
	FsCWAFlags = os.O_CREATE | os.O_WRONLY | os.O_APPEND // create, append write-only
	FsCWTFlags = os.O_CREATE | os.O_WRONLY | os.O_TRUNC  // create, override write-only
	FsCWFlags  = os.O_CREATE | os.O_WRONLY               // create, write-only
	FsRWFlags  = os.O_RDWR                               // read-write, dont create.
	FsRFlags   = os.O_RDONLY                             // read-only
)

// OpenFile like os.OpenFile, but will auto create dir.
//
// Usage:
//
//	file, err := OpenFile("path/to/file.txt", FsCWFlags, 0666)
func OpenFile(filePath string, flag int, perm os.FileMode) (*os.File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MustOpenFile like os.OpenFile, but will auto create dir.
//
// Usage:
//
//	file := MustOpenFile("path/to/file.txt", FsCWFlags, 0666)
func MustOpenFile(filePath string, flag int, perm os.FileMode) *os.File {
	_ = "STUB: not implemented"
	return nil
}

// QuickOpenFile like os.OpenFile, open for append write. if not exists, will create it.
//
// Alias of OpenAppendFile()
func QuickOpenFile(filepath string, fileFlag ...int) (*os.File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// OpenAppendFile like os.OpenFile, open for append write. if not exists, will create it.
func OpenAppendFile(filepath string, filePerm ...os.FileMode) (*os.File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// OpenTruncFile like os.OpenFile, open for override write. if not exists, will create it.
func OpenTruncFile(filepath string, filePerm ...os.FileMode) (*os.File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// OpenReadFile like os.OpenFile, open file for read contents
func OpenReadFile(filepath string) (*os.File, error) { _ = "STUB: not implemented"; return nil, nil }

// CreateFile create file if not exists
//
// Usage:
//
//	CreateFile("path/to/file.txt", 0664, 0666)
func CreateFile(fpath string, filePerm, dirPerm os.FileMode, fileFlag ...int) (*os.File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MustCreateFile create file, will panic on error
func MustCreateFile(filePath string, filePerm, dirPerm os.FileMode) *os.File {
	_ = "STUB: not implemented"
	return nil
}

// ************************************************************
//	remove files
// ************************************************************

// alias methods
var (
	// MustRm removes the named file or (empty) directory.
	MustRm = MustRemove
	// QuietRm removes the named file or (empty) directory.
	QuietRm = QuietRemove
)

// Remove removes the named file or (empty) directory.
func Remove(fPath string) error { _ = "STUB: not implemented"; return nil }

// MustRemove removes the named file or (empty) directory.
// NOTICE: will panic on error
func MustRemove(fPath string) { _ = "STUB: not implemented"; return }

// QuietRemove removes the named file or (empty) directory.
//
// NOTICE: will ignore error
func QuietRemove(fPath string) { _ = "STUB: not implemented"; return }

// SafeRemoveAll removes path and any children it contains. will ignore error
func SafeRemoveAll(path string) { _ = "STUB: not implemented"; return }

// RmIfExist removes the named file or (empty) directory on existing.
func RmIfExist(fPath string) error { _ = "STUB: not implemented"; return nil }

// DeleteIfExist removes the named file or (empty) directory on existing.
func DeleteIfExist(fPath string) error { _ = "STUB: not implemented"; return nil }

// RmFileIfExist removes the named file on existing.
func RmFileIfExist(fPath string) error { _ = "STUB: not implemented"; return nil }

// DeleteIfFileExist removes the named file on existing.
func DeleteIfFileExist(fPath string) error { _ = "STUB: not implemented"; return nil }

// RemoveSub removes all sub files and dirs of dirPath, but not remove dirPath.
func RemoveSub(dirPath string, fns ...FilterFunc) error { _ = "STUB: not implemented"; return nil }

// skip rm not empty subdir

// ************************************************************
//	other operates
// ************************************************************

// Unzip a zip archive
// from https://blog.csdn.net/wangshubo1989/article/details/71743374
func Unzip(archive, targetDir string) (err error) { _ = "STUB: not implemented"; return nil }

// close all
