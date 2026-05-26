package arrutil

import (
	"io"

	"github.com/gookit/goutil/comdef"
)

// ArrFormatter struct
type ArrFormatter struct {
	comdef.BaseFormatter
	// Prefix string for each element
	Prefix string
	// Indent string for format each element
	Indent string
	// ClosePrefix on before end char: ]
	ClosePrefix string
}

// NewFormatter instance
func NewFormatter(arr any) *ArrFormatter { _ = "STUB: not implemented"; return nil }

// FormatIndent array data to string.
func FormatIndent(arr any, indent string) string { _ = "STUB: not implemented"; return "" }

// WithFn for config self
func (f *ArrFormatter) WithFn(fn func(f *ArrFormatter)) *ArrFormatter {
	_ = "STUB: not implemented"

	// WithIndent string
	return nil
}

func (f *ArrFormatter) WithIndent(indent string) *ArrFormatter {
	_ = "STUB: not implemented"
	return nil
}

// FormatTo to custom buffer
func (f *ArrFormatter) FormatTo(w io.Writer) { _ = "STUB: not implemented"; return }

// Format to string
func (f *ArrFormatter) String() string {
	_ = "STUB: not implemented"

	// Format to string
	return ""
}

func (f *ArrFormatter) Format() string { _ = "STUB: not implemented"; return "" }

// Format to string
//
//goland:noinspection GoUnhandledErrorResult
func (f *ArrFormatter) doFormat() { _ = "STUB: not implemented"; return }

// if f.AfterReset {
// 	defer f.Reset()
// }

// sb.Grow(arrLn * 4)

// no indent, with space
