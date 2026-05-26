//go:build !windows

package fsutil

// Realpath returns the shortest path name equivalent to path by purely lexical processing.
// Will expand ~ to home dir, and join with workdir if path is relative path.
func Realpath(pathStr string) string { _ = "STUB: not implemented"; return "" }
