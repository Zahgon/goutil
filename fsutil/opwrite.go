package fsutil

import (
	"os"
)

// ************************************************************
//	temp file or dir
// ************************************************************

// OSTempFile create a temp file on os.TempDir()
//
// Usage:
//
//	fsutil.OSTempFile("example.*.txt")
func OSTempFile(pattern string) (*os.File, error) { _ = "STUB: not implemented"; return nil, nil }

// TempFile is like os.CreateTemp, but can custom temp dir.
//
// Usage:
//
//	// create temp file on os.TempDir()
//	fsutil.TempFile("", "example.*.txt")
//	// create temp file on "testdata" dir
//	fsutil.TempFile("testdata", "example.*.txt")
func TempFile(dir, pattern string) (*os.File, error) { _ = "STUB: not implemented"; return nil, nil }

// OSTempDir creates a new temp dir on os.TempDir and return the temp dir path
//
// Usage:
//
//	fsutil.OSTempDir("example.*")
func OSTempDir(pattern string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// TempDir creates a new temp dir and return the temp dir path
//
// Usage:
//
//	fsutil.TempDir("", "example.*")
//	fsutil.TempDir("testdata", "example.*")
func TempDir(dir, pattern string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// ************************************************************
//	write, copy files
// ************************************************************

// MustSave create file and write contents to file, panic on error.
//
//   - data type allow: string, []byte, io.Reader
//
// default option see NewOpenOption()
func MustSave(filePath string, data any, optFns ...OpenOptionFunc) {
	_ = "STUB: not implemented"
	return
}

// SaveFile create file and write contents to file. will auto create dir.
//
//   - data type allow: string, []byte, io.Reader
//
// default option see NewOpenOption()
func SaveFile(filePath string, data any, optFns ...OpenOptionFunc) error {
	_ = "STUB: not implemented"
	return nil
}

// WriteData Quick write any data to file, alias of PutContents
func WriteData(filePath string, data any, fileFlag ...int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// PutContents create file and write contents to file at once. Will auto create dir
//
// data type allows: string, []byte, io.Reader
//
// Tip: file flag default is FsCWTFlags (override write)
//
// Usage:
//
//	fsutil.PutContents(filePath, contents, fsutil.FsCWAFlags) // append write
//	fsutil.Must2(fsutil.PutContents(filePath, contents)) // panic on error
func PutContents(filePath string, data any, fileFlag ...int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// WriteFile create file and write contents to file, can set perm for a file.
//
// data type allows: string, []byte, io.Reader
//
// Tip: file flag default is FsCWTFlags (override write)
//
// Usage:
//
//	fsutil.WriteFile(filePath, contents, fsutil.DefaultFilePerm, fsutil.FsCWAFlags)
func WriteFile(filePath string, data any, perm os.FileMode, fileFlag ...int) error {
	_ = "STUB: not implemented"
	return nil
}

// WriteOSFile write data to give os.File, then close file.
//
// data type allows: string, []byte, io.Reader
func WriteOSFile(f *os.File, data any) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

// eg: buffer

// CopyFile copy a file to another file path.
func CopyFile(srcPath, dstPath string) error { _ = "STUB: not implemented"; return nil }

// create and open file

// MustCopyFile copy file to another path.
func MustCopyFile(srcPath, dstPath string) { _ = "STUB: not implemented"; return }

// UpdateContents read file contents, call handleFn(contents) handle, then write updated contents to file
func UpdateContents(filePath string, handleFn func(bs []byte) []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// read file contents

// CreateSymlink creates a symbolic link
func CreateSymlink(target, linkPath string) error {
	_ = "STUB: not implemented"
	// Check if the link already exists
	return nil
}

// Remove existing link/file
