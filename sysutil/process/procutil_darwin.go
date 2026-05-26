//go:build darwin

package process

// Exists check process running by given pid
func Exists(pid int) bool {
	_ = "STUB: not implemented"
	// OS X does not have a proc filesystem.
	// Use kill -0 pid to judge if the process exists.
	return false
}
