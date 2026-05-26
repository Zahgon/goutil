package fsutil

import (
	"io/fs"
)

// FilePathInDirs get full file path in dirs. return empty string if not found.
//
// Params:
//   - file: can be relative path, file name, full path.
//   - dirs: dir paths
func FilePathInDirs(fPath string, dirs ...string) string { _ = "STUB: not implemented"; return "" }

// not found

// FirstExists check multi paths and return first exists path.
func FirstExists(paths ...string) string { _ = "STUB: not implemented"; return "" }

// FirstExistsDir check multi paths and return first exists dir.
func FirstExistsDir(paths ...string) string { _ = "STUB: not implemented"; return "" }

// FirstExistsFile check multi paths and return first exists file.
func FirstExistsFile(paths ...string) string { _ = "STUB: not implemented"; return "" }

// MatchPaths given paths by custom mather func.
func MatchPaths(paths []string, matcher PathMatchFunc) []string {
	_ = "STUB: not implemented"
	return nil
}

// MatchFirst filter paths by filter func and return first match path.
func MatchFirst(paths []string, matcher PathMatchFunc, defaultPath string) string {
	_ = "STUB: not implemented"
	return ""
}

// FindParentOption options
type FindParentOption struct {
	MaxLevel int // default: 10
	// NeedDir true: find dirs; false(default): find files
	NeedDir bool
	OnlyOne bool // only find one, default: true
	// Collector func
	Collector func(fullPath string)
	// MatchFunc custom matcher func. return false to stop find.
	MatchFunc func(currentDir string) bool
}

// FindParentOptFn find parent option func
type FindParentOptFn func(opt *FindParentOption)

// FindAllInParentDirs looks for all match file(default)/dir in the current directory and parent directories
func FindAllInParentDirs(dirPath, name string, optFns ...FindParentOptFn) []string {
	_ = "STUB: not implemented"
	return nil
}

// FindOneInParentDirs looks for a file(default)/dir in the current directory and parent directories
func FindOneInParentDirs(dirPath, name string, optFns ...FindParentOptFn) string {
	_ = "STUB: not implemented"
	return ""
}

// FindNameInParentDirs looks for file(default)/dir in the current directory and parent directories
func FindNameInParentDirs(dirPath, name string, collectFn func(fullPath string), optFns ...FindParentOptFn) {
	_ = "STUB: not implemented"
	return
}

// FindInParentDirs looks for file/dir in the current directory and parent directories
//   - MatchFunc custom matcher func. return false to stop find.
func FindInParentDirs(dirPath string, matchFunc func(dir string) bool, maxLevel int) {
	_ = "STUB: not implemented"
	return
}

// Check if the file exists in the current directory

// check find level

// Get parent directory

// Reached the root, file not found

// Move to parent directory

// SearchNameUp find file/dir name in dirPath or parent dirs,
// return the name of directory path
//
// Usage:
//
//	repoDir := fsutil.SearchNameUp("/path/to/dir", ".git")
func SearchNameUp(dirPath, name string) string { _ = "STUB: not implemented"; return "" }

// SearchNameUpx find file/dir name in dirPath or parent dirs,
// return the name of directory path and dir is changed.
func SearchNameUpx(dirPath, name string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// WalkDir walks the file tree rooted at root, calling fn for each file or
// directory in the tree, including root.
//
// TIP: will recursively found in sub dirs.
func WalkDir(dir string, fn fs.WalkDirFunc) error { _ = "STUB: not implemented"; return nil }

// Glob finds files by glob path pattern. alias of filepath.Glob()
// and support filter matched files by name.
//
// Usage:
//
//	files := fsutil.Glob("/path/to/dir/*.go")
func Glob(pattern string, fls ...NameMatchFunc) []string { _ = "STUB: not implemented"; return nil }

// GlobWithFunc find files by glob path pattern, then handle matched file
func GlobWithFunc(pattern string, fn func(filePath string) error) (err error) {
	_ = "STUB: not implemented"
	return nil
}

type (
	// FilterFunc type for FindInDir
	//
	// - return False will skip handle the file.
	FilterFunc func(fPath string, ent fs.DirEntry) bool

	// HandleFunc type for FindInDir
	HandleFunc func(fPath string, ent fs.DirEntry) error
)

// OnlyFindDir on find
func OnlyFindDir(_ string, ent fs.DirEntry) bool {
	_ = "STUB: not implemented"

	// OnlyFindFile on find
	return false
}

func OnlyFindFile(_ string, ent fs.DirEntry) bool {
	_ = "STUB: not implemented"

	// ExcludeNames on find
	return false
}

func ExcludeNames(names ...string) FilterFunc { _ = "STUB: not implemented"; return *new(FilterFunc) }

// IncludeSuffix on find
func IncludeSuffix(ss ...string) FilterFunc { _ = "STUB: not implemented"; return *new(FilterFunc) }

// ExcludeDotFile on find
func ExcludeDotFile(_ string, ent fs.DirEntry) bool { _ = "STUB: not implemented"; return false }

// ExcludeSuffix on find
func ExcludeSuffix(ss ...string) FilterFunc { _ = "STUB: not implemented"; return *new(FilterFunc) }

// ApplyFilters handle
func ApplyFilters(fPath string, ent fs.DirEntry, filters []FilterFunc) bool {
	_ = "STUB: not implemented"
	return false
}

// FindInDir code refer the go pkg: path/filepath.glob()
//
// - TIP: default will be not found in sub-dir.
//
// filters: return false will skip the file.
func FindInDir(dir string, handleFn HandleFunc, filters ...FilterFunc) (e error) {
	_ = "STUB: not implemented"
	return nil
}

// ignore I/O error

// remove the last '/' char

// apply filters

// FileInDirs returns the first file path in the given dirs.
func FileInDirs(paths []string, names ...string) string { _ = "STUB: not implemented"; return "" }
