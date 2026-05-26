package fsutil

import (
	"os"
)

// perm for create dir or file
var (
	DefaultDirPerm   os.FileMode = 0775
	DefaultFilePerm  os.FileMode = 0665
	OnlyReadFilePerm os.FileMode = 0444
)

var (
	// DefaultFileFlags for create and write
	DefaultFileFlags = os.O_CREATE | os.O_WRONLY | os.O_APPEND
	// OnlyReadFileFlags open file for read
	OnlyReadFileFlags = os.O_RDONLY
)

// alias methods
var (
	DirExist  = IsDir
	FileExist = IsFile
	PathExist = PathExists
)

// PathExists reports whether the named file or directory exists.
func PathExists(path string) bool { _ = "STUB: not implemented"; return false }

// IsDir reports whether the named directory exists.
func IsDir(path string) bool { _ = "STUB: not implemented"; return false }

// FileExists reports whether the named file or directory exists.
func FileExists(path string) bool { _ = "STUB: not implemented"; return false }

// IsFile reports whether the named file or directory exists.
//
// - NOTE: not support symlink file
func IsFile(path string) bool { _ = "STUB: not implemented"; return false }

// IsSymlink reports whether the named file is a symlink.
func IsSymlink(path string) bool { _ = "STUB: not implemented"; return false }

// IsAbsPath is abs path check
func IsAbsPath(aPath string) bool { _ = "STUB: not implemented"; return false }

// IsAbsPath2 is abs path check. if start withs / OR ~, will direct return true.
func IsAbsPath2(aPath string) bool { _ = "STUB: not implemented"; return false }

// IsEmptyDir reports whether the named directory is empty.
func IsEmptyDir(dirPath string) bool { _ = "STUB: not implemented"; return false }

// ImageMimeTypes refer net/http package
var ImageMimeTypes = map[string]string{
	"bmp": "image/bmp",
	"gif": "image/gif",
	"ief": "image/ief",
	"jpg": "image/jpeg",
	// "jpe":  "image/jpeg",
	"jpeg": "image/jpeg",
	"png":  "image/png",
	"svg":  "image/svg+xml",
	"ico":  "image/x-icon",
	"webp": "image/webp",
}

// IsImageFile check file is image file.
func IsImageFile(path string) bool { _ = "STUB: not implemented"; return false }

// IsZipFile check is zip file.
// from https://blog.csdn.net/wangshubo1989/article/details/71743374
func IsZipFile(filepath string) bool { _ = "STUB: not implemented"; return false }

// PathMatch check for a string. alias of path.Match()
func PathMatch(pattern, s string) bool { _ = "STUB: not implemented"; return false }
