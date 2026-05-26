package fsutil

import (
	"io"
)

// DetectMime detect a file mime type. alias of MimeType()
func DetectMime(path string) string { _ = "STUB: not implemented"; return "" }

// MimeType get file mime type name. eg "image/png"
func MimeType(path string) (mime string) { _ = "STUB: not implemented"; return "" }

// ReaderMimeType get the io.Reader mimeType
//
// Usage:
//
//	file, err := os.Open(filepath)
//	if err != nil {
//		return
//	}
//	mime := ReaderMimeType(file)
func ReaderMimeType(r io.Reader) (mime string) { _ = "STUB: not implemented"; return "" }
