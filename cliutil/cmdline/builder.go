package cmdline

import (
	"strings"
)

// LineBuilder build command line string.
// codes refer from strings.Builder
type LineBuilder struct {
	strings.Builder
}

// NewBuilder create
func NewBuilder(binFile string, args ...string) *LineBuilder { _ = "STUB: not implemented"; return nil }

// ResetGet value, will reset after get.
func (b *LineBuilder) ResetGet() string { _ = "STUB: not implemented"; return "" }

// AddArg to builder
func (b *LineBuilder) AddArg(arg string) { _ = "STUB: not implemented"; return }

// AddArgs to builder
func (b *LineBuilder) AddArgs(args ...string) {
	_ = "STUB: not implemented"

	// AddArray to builder
	return
}

func (b *LineBuilder) AddArray(args []string) { _ = "STUB: not implemented"; return }

// AddAny args to builder
func (b *LineBuilder) AddAny(args ...any) { _ = "STUB: not implemented"; return }

// WriteString arg string to the builder, will auto quote special string.
// refer strconv.Quote()
func (b *LineBuilder) WriteString(a string) (int, error) {
	_ = "STUB: not implemented"
	// add sep on not-first write.
	return 0, nil
}
