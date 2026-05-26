package textscan

import (
	"io"
)

// HandleFn for token
type HandleFn func(t Token)

// Parser struct
type Parser struct {
	ts *TextScanner
	// Func for handle tokens
	Func HandleFn
}

// NewParser instance
func NewParser(fn HandleFn) *Parser { _ = "STUB: not implemented"; return nil }

// AddMatchers register token matchers
func (p *Parser) AddMatchers(ms ...Matcher) { _ = "STUB: not implemented"; return }

// Parse input bytes
func (p *Parser) Parse(bs []byte) error { _ = "STUB: not implemented"; return nil }

// ParseText input string
func (p *Parser) ParseText(text string) error { _ = "STUB: not implemented"; return nil }

// ParseFrom input reader
func (p *Parser) ParseFrom(r io.Reader) error { _ = "STUB: not implemented"; return nil }
