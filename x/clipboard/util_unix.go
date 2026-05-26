//go:build !windows && !darwin

package clipboard

// GetWriterBin program name
func GetWriterBin() string {
	_ = "STUB: not implemented"

	// GetReaderBin program name
	return ""
}

func GetReaderBin() string { _ = "STUB: not implemented"; return "" }

func available() bool {
	_ = "STUB: not implemented"
	// X clipboard is unavailable when not under X.
	return false
}
