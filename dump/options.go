package dump

import (
	"io"
)

// Options for dumper
type Options struct {
	// Output the output writer
	Output io.Writer
	// NoType don't show data type TODO
	NoType bool
	// NoColor don't with color
	NoColor bool
	// IndentLen width. default is 2
	IndentLen int
	// IndentChar default is one space
	IndentChar byte
	// MaxDepth for nested print
	MaxDepth int
	// ShowFlag for display caller position
	ShowFlag int
	// CallerSkip skip for call runtime.Caller()
	CallerSkip int
	// ColorTheme for print result.
	ColorTheme Theme
	// SkipNilField value dump on map, struct.
	SkipNilField bool
	// SkipPrivate field dump on struct.
	SkipPrivate bool
	// BytesAsString dump handle.
	BytesAsString bool
	// ShowLen display length information for string, slice, array, map
	ShowLen bool
	// MoreLenNL array/slice elements length > MoreLenNL, will wrap new line
	// MoreLenNL int
	// MaxElementsNum for a long-long slice or array. The excess will be displayed `...`
	MaxElementsNum int
}

// OptionFunc type
type OptionFunc func(opts *Options)

// NewDefaultOptions create.
func NewDefaultOptions(out io.Writer, skip int) *Options { _ = "STUB: not implemented"; return nil }

// ---

// MoreLenNL: 8,
// ---

// ---
// show length by default for backward compatibility

// SkipNilField setting.
func SkipNilField() OptionFunc { _ = "STUB: not implemented"; return *new(OptionFunc) }

// SkipPrivate field dump on struct.
func SkipPrivate() OptionFunc { _ = "STUB: not implemented"; return *new(OptionFunc) }

// BytesAsString setting.
func BytesAsString() OptionFunc { _ = "STUB: not implemented"; return *new(OptionFunc) }

// WithCallerSkip on print caller position information.
func WithCallerSkip(skip int) OptionFunc { _ = "STUB: not implemented"; return *new(OptionFunc) }

// WithoutPosition dont print call dump position information.
func WithoutPosition() OptionFunc { _ = "STUB: not implemented"; return *new(OptionFunc) }

// WithoutOutput setting.
func WithoutOutput(out io.Writer) OptionFunc { _ = "STUB: not implemented"; return *new(OptionFunc) }

// WithoutColor setting.
func WithoutColor() OptionFunc { _ = "STUB: not implemented"; return *new(OptionFunc) }

// WithoutType setting.
func WithoutType() OptionFunc { _ = "STUB: not implemented"; return *new(OptionFunc) }

// WithoutLen setting. hide length information for string, slice, array, map.
func WithoutLen() OptionFunc { _ = "STUB: not implemented"; return *new(OptionFunc) }
