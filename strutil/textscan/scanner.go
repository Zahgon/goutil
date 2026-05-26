// Package textscan Implemented a parser that quickly scans and analyzes text content.
// It can be used to parse INI, Properties and other formats
package textscan

import (
	"bufio"
)

// ErrScan error on scan or parse contents
type ErrScan struct {
	Msg  string // error message
	Line int    // error line number, start 1
	Text string // text contents on error
}

// Error string
func (e ErrScan) Error() string { _ = "STUB: not implemented"; return "" }

// Matcher interface
type Matcher interface {
	// Match text line by kind, if success returns a new Token
	Match(line string, prev Token) (tok Token, err error)
}

// TextScanner struct.
type TextScanner struct {
	in *bufio.Scanner
	// ks map[Kind]string

	// token matchers
	matchers []Matcher
	prevTok  Token

	line int
	next string // not used
	tok  Token
	err  error
}

// NewScanner instance
func NewScanner(in any) *TextScanner { _ = "STUB: not implemented"; return nil }

// SetInput for scan and parse
func (s *TextScanner) SetInput(in any) {
	_ = "STUB: not implemented"
	// init
	//
	//	if s.ks == nil {
	//		s.ks = make(map[Kind]string, len(kinds))
	//	}
	//
	//	for kind, name := range kinds {
	//		s.ks[kind] = name
	//	}
	return
}

// SetSplit set split func on scan
func (s *TextScanner) SetSplit(fn bufio.SplitFunc) { _ = "STUB: not implemented"; return }

// AddKind register new kind
func (s *TextScanner) AddKind(k Kind, name string) { _ = "STUB: not implemented"; return }

// AddMatchers register token matchers
func (s *TextScanner) AddMatchers(ms ...Matcher) { _ = "STUB: not implemented"; return }

// Each every token by given func
func (s *TextScanner) Each(fn func(t Token)) error { _ = "STUB: not implemented"; return nil }

// Scan source input and parsing.
// Can use Token() get current parsed token value
//
// Usage:
//
//	ts := textscan.NewScanner(`source ...`)
//	for ts.Scan() {
//		tok := ts.Token()
//		// do something...
//	}
//	fmt.Println(ts.Err())
func (s *TextScanner) Scan() bool { _ = "STUB: not implemented"; return false }

// at end.

func (s *TextScanner) matchToken(text string) (ok bool) { _ = "STUB: not implemented"; return false }

// emtpy line, match next valid token

// end EOF

// ScanNext advance and fetch next line text
func (s *TextScanner) ScanNext() (ok bool, text string) {
	_ = "STUB: not implemented"
	return false, ""
}

// SetNext text for scan and parse
func (s *TextScanner) SetNext(text string) {
	_ = "STUB: not implemented"

	// Token get of current scan.
	return
}

func (s *TextScanner) Token() Token {
	_ = "STUB: not implemented"

	// PrevToken get of previous scan.
	return *new(Token)
}

func (s *TextScanner) PrevToken() Token {
	_ = "STUB: not implemented"

	// Line on current
	return *new(Token)
}

func (s *TextScanner) Line() int {
	_ = "STUB: not implemented"

	// Err get
	return 0
}

func (s *TextScanner) Err() error { _ = "STUB: not implemented"; return nil }
