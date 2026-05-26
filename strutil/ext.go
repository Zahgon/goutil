package strutil

import (
	"strings"
)

// SimilarComparator definition
//
// links:
//
//	https://github.com/mkideal/cli/blob/master/fuzzy.go
type SimilarComparator struct {
	src, dst string
}

// NewComparator create
func NewComparator(src, dst string) *SimilarComparator { _ = "STUB: not implemented"; return nil }

// Similarity calc for two string.
//
// Usage:
//
//	rate, ok := Similarity("hello", "he")
func Similarity(s, t string, rate float32) (float32, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

// Similar by minDifferRate
//
// Usage:
//
//	c := NewComparator("hello", "he")
//	rate, ok :c.Similar(0.3)
func (c *SimilarComparator) Similar(minDifferRate float32) (float32, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func (c *SimilarComparator) editDistance(s, t []byte) float32 { _ = "STUB: not implemented"; return 0 }

// Builder struct
type Builder struct {
	strings.Builder
}

// Write bytes and no error report
func (b *Builder) Write(p []byte) { _ = "STUB: not implemented"; return }

// WriteRune and no error report
func (b *Builder) WriteRune(r rune) { _ = "STUB: not implemented"; return }

// WriteByteNE write byte and no error report
func (b *Builder) WriteByteNE(c byte) {
	_ = "STUB: not implemented"

	// WriteString to builder
	return
}

func (b *Builder) WriteString(s string) { _ = "STUB: not implemented"; return }

// Writef write string by fmt.Sprintf formatted
func (b *Builder) Writef(tpl string, vs ...any) { _ = "STUB: not implemented"; return }

// Writeln write string with newline.
func (b *Builder) Writeln(s string) { _ = "STUB: not implemented"; return }

// WriteAny write any type value.
func (b *Builder) WriteAny(v any) { _ = "STUB: not implemented"; return }

// WriteAnys write any type values.
func (b *Builder) WriteAnys(vs ...any) { _ = "STUB: not implemented"; return }

// WriteMulti write multi byte at once.
func (b *Builder) WriteMulti(bs ...byte) { _ = "STUB: not implemented"; return }

// WriteStrings write multi string at once.
func (b *Builder) WriteStrings(ss ...string) { _ = "STUB: not implemented"; return }

// ResetGet return current string and reset builder
func (b *Builder) ResetGet() string { _ = "STUB: not implemented"; return "" }
