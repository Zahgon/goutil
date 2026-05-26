package fakeobj

import (
	"github.com/gookit/goutil/byteutil"
)

// IOWriter only implements the io.Writer
type IOWriter struct {
	Buf []byte
	// ErrOnWrite return error on write, useful for testing
	ErrOnWrite bool
}

// NewIOWriter instance
func NewIOWriter() *IOWriter { _ = "STUB: not implemented"; return nil }

// Write implements
func (w *IOWriter) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

// Reset the buffer
func (w *IOWriter) Reset() {
	_ = "STUB: not implemented"

	// Reader implements the io.Reader, io.Closer
	return
}

type Reader struct {
	byteutil.Buffer
	// ErrOnRead return error on read, useful for testing
	ErrOnRead bool
}

// NewReader instance
func NewReader() *Reader {
	_ = "STUB: not implemented"

	// NewStrReader instance
	return nil
}

func NewStrReader(s string) *Reader { _ = "STUB: not implemented"; return nil }

// SetErrOnRead mark
func (r *Reader) SetErrOnRead() {
	_ = "STUB: not implemented"

	// Read implements the io.Reader
	return
}

func (r *Reader) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

// Writer implements the io.Writer, stdio.Flusher, io.Closer.
type Writer struct {
	byteutil.Buffer
	// ErrOnWrite return error on write, useful for testing
	ErrOnWrite bool
	// ErrOnFlush return error on flush, useful for testing
	ErrOnFlush bool
	// ErrOnSync return error on flush, useful for testing
	ErrOnSync bool
	// ErrOnClose return error on close, useful for testing
	ErrOnClose bool
}

// NewBuffer instance. alias of NewWriter()
func NewBuffer() *Writer {
	_ = "STUB: not implemented"

	// NewWriter instance
	return nil
}

func NewWriter() *Writer {
	_ = "STUB: not implemented"

	// SetErrOnWrite method
	return nil
}

func (w *Writer) SetErrOnWrite() *Writer { _ = "STUB: not implemented"; return nil }

// SetErrOnFlush method
func (w *Writer) SetErrOnFlush() *Writer { _ = "STUB: not implemented"; return nil }

// SetErrOnSync method
func (w *Writer) SetErrOnSync() *Writer { _ = "STUB: not implemented"; return nil }

// SetErrOnClose method
func (w *Writer) SetErrOnClose() *Writer { _ = "STUB: not implemented"; return nil }

// ResetGet buffer string.
func (w *Writer) ResetGet() string { _ = "STUB: not implemented"; return "" }

// Write implements
func (w *Writer) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

// Close implements
func (w *Writer) Close() error { _ = "STUB: not implemented"; return nil }

// Flush implements stdio.Flusher
func (w *Writer) Flush() error { _ = "STUB: not implemented"; return nil }

// Sync implements stdio.Syncer
func (w *Writer) Sync() error { _ = "STUB: not implemented"; return nil }
