package stdio

import (
	"io"
)

// WriteWrapper warp io.Writer support more operate methods.
type WriteWrapper struct {
	Out io.Writer
}

// WrapW instance
func WrapW(w io.Writer) *WriteWrapper { _ = "STUB: not implemented"; return nil }

// NewWriteWrapper instance
func NewWriteWrapper(w io.Writer) *WriteWrapper { _ = "STUB: not implemented"; return nil }

// Write bytes data
func (w *WriteWrapper) Write(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0,

		// Writef data to output
		nil
}

func (w *WriteWrapper) Writef(tpl string, vs ...any) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// WriteByte data
func (w *WriteWrapper) WriteByte(c byte) error { _ = "STUB: not implemented"; return nil }

// WriteString data
func (w *WriteWrapper) WriteString(s string) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// String get write data string
func (w *WriteWrapper) String() string { _ = "STUB: not implemented"; return "" }
