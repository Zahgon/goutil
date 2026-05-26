package comdef

import (
	"io"
)

// DataFormatter interface
type DataFormatter interface {
	Format() string
	FormatTo(w io.Writer)
}

// BaseFormatter struct
//
// Usage:
//
//	 type YourFormatter struct {
//			comdef.BaseFormatter
//	 }
//	 // implement the DataFormatter interface...
type BaseFormatter struct {
	ow ByteStringWriter
	// Out formatted to the writer
	Out io.Writer
	// Src data(array, map, struct) for format
	Src any
	// MaxDepth limit depth for array, map data TODO
	MaxDepth int
	// Prefix string for each element
	Prefix string
	// Indent string for format each element
	Indent string
	// ClosePrefix string for last "]", "}"
	ClosePrefix string
}

// Reset after format
func (f *BaseFormatter) Reset() { _ = "STUB: not implemented"; return }

// SetOutput writer
func (f *BaseFormatter) SetOutput(out io.Writer) {
	_ = "STUB: not implemented"

	// BsWriter warp the Out, build a ByteStringWriter
	return
}

func (f *BaseFormatter) BsWriter() ByteStringWriter {
	_ = "STUB: not implemented"
	return *new(ByteStringWriter)
}
