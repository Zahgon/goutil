package strutil

import (
	"strings"
)

/*************************************************************
 * change string case
 *************************************************************/

// methods aliases
var (
	UpWords = UpperWord
	LoFirst = LowerFirst
	UpFirst = UpperFirst

	Snake = SnakeCase
)

// Title alias of the strings.ToTitle()
func Title(s string) string { _ = "STUB: not implemented"; return "" }

// Lower alias of the strings.ToLower()
func Lower(s string) string { _ = "STUB: not implemented"; return "" }

// Lowercase alias of the strings.ToLower()
func Lowercase(s string) string { _ = "STUB: not implemented"; return "" }

// Upper alias of the strings.ToUpper()
func Upper(s string) string { _ = "STUB: not implemented"; return "" }

// Uppercase alias of the strings.ToUpper()
func Uppercase(s string) string { _ = "STUB: not implemented"; return "" }

// UpperWord Change the first character of each word to uppercase
func UpperWord(s string) string { _ = "STUB: not implemented"; return "" }

// LowerFirst lower first char
func LowerFirst(s string) string { _ = "STUB: not implemented"; return "" }

// UpperFirst upper first char
func UpperFirst(s string) string { _ = "STUB: not implemented"; return "" }

// SnakeCase convert. eg "RangePrice" -> "range_price"
func SnakeCase(s string, sep ...string) string { _ = "STUB: not implemented"; return "" }

// Camel alias of the CamelCase
func Camel(s string, sep ...string) string { _ = "STUB: not implemented"; return "" }

// CamelCase convert string to camel case.
//
// Support:
//
//	"range_price" -> "rangePrice"
//	"range price" -> "rangePrice"
//	"range-price" -> "rangePrice"
func CamelCase(s string, sep ...string) string { _ = "STUB: not implemented"; return "" }

// Not contains sep char

// Get regexp instance

//
// Indent format multi line text
// from package: github.com/kr/text
//

// Indent inserts prefix at the beginning of each non-empty line of s. The
// end-of-line marker is NL.
func Indent(s, prefix string) string { _ = "STUB: not implemented"; return "" }

// IndentBytes inserts prefix at the beginning of each non-empty line of b.
// The end-of-line marker is NL.
func IndentBytes(b, prefix []byte) []byte { _ = "STUB: not implemented"; return nil }

// Replaces replace multi strings
//
//	pairs: {old1: new1, old2: new2, ...}
//
// Can also use:
//
//	strings.NewReplacer("old1", "new1", "old2", "new2").Replace(str)
func Replaces(str string, pairs map[string]string) string { _ = "STUB: not implemented"; return "" }

// ReplaceVars replaces simple variables in a string. format: {varName}
//
// Usage:
//
//	strutil.ReplaceVars("{name}, age is {age}", map[string]string{
//		"name": "Joe",
//		"age": "18"
//	})
func ReplaceVars(s string, vars map[string]string) string { _ = "STUB: not implemented"; return "" }

// format var name to {name}

// NewReplacer instance
func NewReplacer(pairs map[string]string) *strings.Replacer { _ = "STUB: not implemented"; return nil }

// WrapTag for given string.
func WrapTag(s, tag string) string { _ = "STUB: not implemented"; return "" }
