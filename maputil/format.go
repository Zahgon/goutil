package maputil

import (
	"io"

	"github.com/gookit/goutil/comdef"
)

// MapFormatter struct
type MapFormatter struct {
	comdef.BaseFormatter
	// Prefix string for each element
	Prefix string
	// Indent string for each element
	Indent string
	// ClosePrefix string for last "}"
	ClosePrefix string
	// AfterReset after reset on call Format().
	// AfterReset bool
}

// NewFormatter instance
func NewFormatter(mp any) *MapFormatter { _ = "STUB: not implemented"; return nil }

// WithFn for config self
func (f *MapFormatter) WithFn(fn func(f *MapFormatter)) *MapFormatter {
	_ = "STUB: not implemented"

	// WithIndent string
	return nil
}

func (f *MapFormatter) WithIndent(indent string) *MapFormatter {
	_ = "STUB: not implemented"
	return nil
}

// FormatTo to custom buffer
func (f *MapFormatter) FormatTo(w io.Writer) { _ = "STUB: not implemented"; return }

// Format to string
func (f *MapFormatter) String() string {
	_ = "STUB: not implemented"

	// Format to string
	return ""
}

func (f *MapFormatter) Format() string { _ = "STUB: not implemented"; return "" }

// Format map data to string.
//
//goland:noinspection GoUnhandledErrorResult
func (f *MapFormatter) doFormat() { _ = "STUB: not implemented"; return }

// buf.Grow(ln * 16)

// no indent, with space

// with newline
