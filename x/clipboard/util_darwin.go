//go:build darwin

package clipboard

// GetWriterBin program name
func GetWriterBin() string {
	_ = "STUB: not implemented"

	// GetReaderBin program name
	return ""
}

func GetReaderBin() string { _ = "STUB: not implemented"; return "" }

func available() bool { _ = "STUB: not implemented"; return false }
