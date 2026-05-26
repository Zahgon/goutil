// Package stdio provide some standard IO util functions.
package stdio

import (
	"bufio"
	"io"
)

// DiscardReader anything from the reader
func DiscardReader(src io.Reader) { _ = "STUB: not implemented"; return }

// ReadString read contents from io.Reader, return empty string on error
func ReadString(r io.Reader) string { _ = "STUB: not implemented"; return "" }

// MustReadReader read contents from io.Reader, will panic on error
func MustReadReader(r io.Reader) []byte { _ = "STUB: not implemented"; return nil }

// NewIOReader instance by input: string, bytes, io.Reader
func NewIOReader(in any) io.Reader { _ = "STUB: not implemented"; return *new(io.Reader) }

// NewScanner instance by input data or reader
func NewScanner(in any) *bufio.Scanner { _ = "STUB: not implemented"; return nil }

// SafeClose close io.Closer, ignore error
func SafeClose(c io.Closer) {
	_ = "STUB: not implemented"

	// WriteByte to stdout, will ignore error
	return
}

func WriteByte(b byte) { _ = "STUB: not implemented"; return }

// WriteBytes to stdout, will ignore error
func WriteBytes(bs []byte) { _ = "STUB: not implemented"; return }

// WritelnBytes to stdout, will ignore error
func WritelnBytes(bs []byte) { _ = "STUB: not implemented"; return }

// WriteString to stdout. will ignore error
func WriteString(s string) { _ = "STUB: not implemented"; return }

// Writeln string to stdout. will ignore error
func Writeln(s string) { _ = "STUB: not implemented"; return }
