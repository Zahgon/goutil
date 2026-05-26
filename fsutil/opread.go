package fsutil

import (
	"bufio"
	"io"
	"text/scanner"
)

// NewIOReader instance by input file path or io.Reader
func NewIOReader(in any) (r io.Reader, err error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), nil
}

// as file path

// DiscardReader anything from the reader
func DiscardReader(src io.Reader) { _ = "STUB: not implemented"; return }

// ReadFile read file contents, will panic on error
func ReadFile(filePath string) []byte { _ = "STUB: not implemented"; return nil }

// MustReadFile read file contents, will panic on error
func MustReadFile(filePath string) []byte { _ = "STUB: not implemented"; return nil }

// ReadReader read contents from io.Reader, will panic on error
func ReadReader(r io.Reader) []byte { _ = "STUB: not implemented"; return nil }

// MustReadReader read contents from io.Reader, will panic on error
func MustReadReader(r io.Reader) []byte { _ = "STUB: not implemented"; return nil }

// ReadString read contents from path or io.Reader, will panic on in type error
func ReadString(in any) string { _ = "STUB: not implemented"; return "" }

// ReadStringOrErr read contents from path or io.Reader, will panic on in type error
func ReadStringOrErr(in any) (string, error) { _ = "STUB: not implemented"; return "", nil }

// ReadAll read contents from path or io.Reader, will panic on in type error
func ReadAll(in any) []byte {
	_ = "STUB: not implemented"

	// GetContents read contents from path or io.Reader, will panic on in type error
	return nil
}

func GetContents(in any) []byte {
	_ = "STUB: not implemented"

	// MustRead read contents from path or io.Reader, will panic on in type error
	return nil
}

func MustRead(in any) []byte { _ = "STUB: not implemented"; return nil }

// ReadOrErr read contents from path or io.Reader, will panic on in type error
func ReadOrErr(in any) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// ReadExistFile read file contents if existed, will panic on error
func ReadExistFile(filePath string) []byte { _ = "STUB: not implemented"; return nil }

// TextScanner from filepath or io.Reader, will panic on in type error.
// Will scan parse text to tokens: Ident, Int, Float, Char, String, RawString, Comment, etc.
//
// Usage:
//
//	s := fsutil.TextScanner("/path/to/file")
//	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
//		fmt.Printf("%s: %s\n", s.Position, s.TokenText())
//	}
func TextScanner(in any) *scanner.Scanner { _ = "STUB: not implemented"; return nil }

// LineScanner create from filepath or io.Reader, will panic on in type error.
// Will scan and parse text to lines.
//
//	s := fsutil.LineScanner("/path/to/file")
//	for s.Scan() {
//		fmt.Println(s.Text())
//	}
func LineScanner(in any) *bufio.Scanner { _ = "STUB: not implemented"; return nil }
