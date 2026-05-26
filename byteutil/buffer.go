package byteutil

import (
	"bytes"
)

// Buffer wrap and extends the bytes.Buffer, add some useful methods
// and implements the io.Writer, io.Closer and stdio.Flusher interfaces
type Buffer struct {
	bytes.Buffer
	// custom error for testing
	CloseErr error
	FlushErr error
	SyncErr  error
}

// NewBuffer instance
func NewBuffer() *Buffer {
	_ = "STUB: not implemented"

	// PrintByte to buffer, ignore error. alias of WriteByte()
	return nil
}

func (b *Buffer) PrintByte(c byte) {
	_ = "STUB: not implemented"

	// WriteStr1 quiet write one string to buffer
	return
}

func (b *Buffer) WriteStr1(s string) { _ = "STUB: not implemented"; return }

// WriteStr1Nl quiet write one string and end with newline
func (b *Buffer) WriteStr1Nl(s string) { _ = "STUB: not implemented"; return }

// writeStringNl quiet write one string and end with newline
func (b *Buffer) writeStringNl(s string, nl bool) { _ = "STUB: not implemented"; return }

// WriteStr quiet write strings to buffer
func (b *Buffer) WriteStr(ss ...string) { _ = "STUB: not implemented"; return }

// WriteStrings to buffer, ignore error.
func (b *Buffer) WriteStrings(ss []string) { _ = "STUB: not implemented"; return }

// WriteStringNl write message to buffer and end with newline
func (b *Buffer) WriteStringNl(ss ...string) { _ = "STUB: not implemented"; return }

// writeStringsNl to buffer, ignore error.
func (b *Buffer) writeStringsNl(ss []string, nl bool) { _ = "STUB: not implemented"; return }

// WriteAny type value to buffer
func (b *Buffer) WriteAny(vs ...any) { _ = "STUB: not implemented"; return }

// Writeln write values to buffer and end with newline
func (b *Buffer) Writeln(vs ...any) { _ = "STUB: not implemented"; return }

// WriteAnyNl type value to buffer and end with newline
func (b *Buffer) WriteAnyNl(vs ...any) { _ = "STUB: not implemented"; return }

// WriteAnyLn type value to buffer and end with newline
func (b *Buffer) writeAnysWithNl(vs []any, nl bool) { _ = "STUB: not implemented"; return }

// Writef write message to buffer, ignore error. alias of Printf()
func (b *Buffer) Writef(tpl string, vs ...any) { _ = "STUB: not implemented"; return }

// Printf quick write message to buffer, ignore error.
func (b *Buffer) Printf(tpl string, vs ...any) { _ = "STUB: not implemented"; return }

// Println quick write message with newline to buffer, will ignore error.
func (b *Buffer) Println(vs ...any) { _ = "STUB: not implemented"; return }

// ResetGet buffer string. alias of ResetAndGet()
func (b *Buffer) ResetGet() string { _ = "STUB: not implemented"; return "" }

// ResetAndGet buffer string.
func (b *Buffer) ResetAndGet() string { _ = "STUB: not implemented"; return "" }

// Close buffer
func (b *Buffer) Close() error {
	_ = "STUB: not implemented"

	// Flush buffer
	return nil
}

func (b *Buffer) Flush() error {
	_ = "STUB: not implemented"

	// Sync anf flush buffer
	return nil
}

func (b *Buffer) Sync() error { _ = "STUB: not implemented"; return nil }
