// Package clipboard provides a simple clipboard read and write operations.
package clipboard

import (
	"bytes"
	"io"
)

// Clipboard struct
type Clipboard struct {
	// TODO add event on write, read
	// buffer for write
	buf *bytes.Buffer

	// print exec command line on run
	verbose bool
	// available - a bin file exists on the OS.
	writeable, readable bool

	readerBin string
	readArgs  []string
	writerBin string
	writeArgs []string
}

// New instance
func New() *Clipboard {
	_ = "STUB: not implemented"
	// special handle on with args
	return nil
}

// SetReader for handle clip
// func (c *Clipboard) SetReader(line string)  {
// }

// SetWriter for handle clip
// func (c *Clipboard) SetWriter(line string)  {
// }

// WithVerbose setting
func (c *Clipboard) WithVerbose(yn bool) *Clipboard { _ = "STUB: not implemented"; return nil }

// Clean the clipboard
func (c *Clipboard) Clean() error {
	_ = "STUB: not implemented"

	// Reset and clean the clipboard
	return nil
}

func (c *Clipboard) Reset() error { _ = "STUB: not implemented"; return nil }

// echo empty string for clean clipboard.
// run: echo '' | pbcopy

//
// ---------------------------------------- write ----------------------------------------
//

// Write bytes data to buffer. should call Flush() to write to clipboard
func (c *Clipboard) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// WriteString data to buffer. should call Flush() to write to clipboard
func (c *Clipboard) WriteString(s string) (int, error) {
	_ = "STUB: not implemented"
	//	if c.addSlashes {
	//		s = strutil.AddSlashes(s)
	//	}
	return 0, nil
}

// Flush buffer contents to clipboard
func (c *Clipboard) Flush() error { _ = "STUB: not implemented"; return nil }

// WriteFromFile contents to clipboard
func (c *Clipboard) WriteFromFile(filepath string) error { _ = "STUB: not implemented"; return nil }

// WriteFrom reader data to clipboard
func (c *Clipboard) WriteFrom(r io.Reader) error { _ = "STUB: not implemented"; return nil }

//
// ---------------------------------------- read ----------------------------------------
//

// Read bytes contents from clipboard
func (c *Clipboard) Read() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// ReadToBuffer read clipboard contents to new buffer.
func (c *Clipboard) ReadToBuffer() (*bytes.Buffer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SafeString read contents as string from clipboard, will return empty string on error
func (c *Clipboard) SafeString() string { _ = "STUB: not implemented"; return "" }

// ReadString contents as string from clipboard
func (c *Clipboard) ReadString() (string, error) { _ = "STUB: not implemented"; return "", nil }

// fix: at Windows will always return end of the "\r\n"

// ReadToFile dump clipboard data to file
func (c *Clipboard) ReadToFile(filepath string) error { _ = "STUB: not implemented"; return nil }

// ReadTo read clipboard contents to writer
func (c *Clipboard) ReadTo(w io.Writer) error { _ = "STUB: not implemented"; return nil }

//
// ---------------------------------------- help ----------------------------------------
//

// Available check
func (c *Clipboard) Available() bool { _ = "STUB: not implemented"; return false }

// Writeable check
func (c *Clipboard) Writeable() bool {
	_ = "STUB: not implemented"

	// Readable check
	return false
}

func (c *Clipboard) Readable() bool { _ = "STUB: not implemented"; return false }

func (c *Clipboard) buffer() *bytes.Buffer { _ = "STUB: not implemented"; return nil }
