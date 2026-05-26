package testutil

import (
	"bytes"
	"sync"

	"github.com/gookit/goutil/byteutil"
)

// Buffer wrap and extends the bytes.Buffer
type Buffer = byteutil.Buffer

// NewBuffer instance
func NewBuffer() *byteutil.Buffer { _ = "STUB: not implemented"; return nil }

// SafeBuffer Thread-safe buffer for testing
type SafeBuffer struct {
	bytes.Buffer
	mu sync.Mutex
}

// NewSafeBuffer instance
func NewSafeBuffer() *SafeBuffer { _ = "STUB: not implemented"; return nil }

// Write implements io.Writer
func (sb *SafeBuffer) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

// WriteString implements io.StringWriter
func (sb *SafeBuffer) WriteString(s string) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// ResetGet get buffer content and reset
func (sb *SafeBuffer) ResetGet() string { _ = "STUB: not implemented"; return "" }
