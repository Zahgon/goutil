// Package jsonutil provide some util functions for quick operate JSON data
package jsonutil

import (
	"regexp"
)

// WriteFile write data to JSON file
func WriteFile(filePath string, data any) error { _ = "STUB: not implemented"; return nil }

// WritePretty write pretty data to JSON file
func WritePretty(filePath string, data any) error { _ = "STUB: not implemented"; return nil }

// ReadFile Read JSON file data
func ReadFile(filePath string, v any) error { _ = "STUB: not implemented"; return nil }

// Pretty JSON string and return
func Pretty(v any) (string, error) { _ = "STUB: not implemented"; return "", nil }

// MustPretty data to JSON string, will panic on error
func MustPretty(v any) string { _ = "STUB: not implemented"; return "" }

// Mapping src data(map,struct) to dst struct use json tags.
//
// On src, dst both is struct, equivalent to merging two structures (src should be a subset of dsc)
func Mapping(src, dst any) error { _ = "STUB: not implemented"; return nil }

// IsJSON check if the string is valid JSON. (Note: uses json.Valid)
func IsJSON(s string) bool { _ = "STUB: not implemented"; return false }

// IsJSONFast simple and fast check input is valid JSON array or object.
func IsJSONFast(s string) bool { _ = "STUB: not implemented"; return false }

// object

// array

// IsArray check if the string is valid JSON array.
func IsArray(s string) bool { _ = "STUB: not implemented"; return false }

// IsObject check if the string is valid JSON object.
func IsObject(s string) bool { _ = "STUB: not implemented"; return false }

// object

// `(?s:` enable match multi line
var jsonMLComments = regexp.MustCompile(`(?s:/\*.*?\*/\s*)`)

// StripComments strip comments for a JSON string
func StripComments(src string) string {
	_ = "STUB: not implemented"
	// multi line comments
	return ""
}

// single line comments

// strip inline comments

// don't skip comments

// } else {
// fmt.Printf("%s: %s\n", s.Position, txt)
