// Package fsutil Filesystem util functions, quick create, read and write file. eg: file and dir check, operate
package fsutil

import (
	"os"
)

// PathSep alias of os.PathSeparator
const PathSep = os.PathSeparator

// JoinPaths elements, alias of filepath.Join()
func JoinPaths(elem ...string) string { _ = "STUB: not implemented"; return "" }

// JoinPaths3 elements, like the filepath.Join()
func JoinPaths3(basePath, secPath string, elems ...string) string {
	_ = "STUB: not implemented"
	return ""
}

// JoinSubPaths elements, like the filepath.Join()
func JoinSubPaths(basePath string, elems ...string) string { _ = "STUB: not implemented"; return "" }

// SlashPath alias of filepath.ToSlash
func SlashPath(path string) string { _ = "STUB: not implemented"; return "" }

// UnixPath like of filepath.ToSlash, but always replace
func UnixPath(path string) string { _ = "STUB: not implemented"; return "" }

// ToAbsPath convert path to absolute path.
// Will expand home dir, if empty will return current work dir
//
// TIP: will don't check path is really exists
func ToAbsPath(p string) string {
	_ = "STUB: not implemented"
	// return current work dir
	return ""
}

// expand home dir

// Must2 ok for (any, error) result. if it has error, will panic
func Must2(_ any, err error) { _ = "STUB: not implemented"; return }
