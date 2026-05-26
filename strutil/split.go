package strutil

import (
	"regexp"
)

// BeforeFirst get substring before first sep.
func BeforeFirst(s, sep string) string { _ = "STUB: not implemented"; return "" }

// AfterFirst get substring after first sep.
func AfterFirst(s, sep string) string { _ = "STUB: not implemented"; return "" }

// BeforeLast get substring before last sep.
func BeforeLast(s, sep string) string { _ = "STUB: not implemented"; return "" }

// AfterLast get substring after last sep.
func AfterLast(s, sep string) string { _ = "STUB: not implemented"; return "" }

/*************************************************************
 * String split operation
 *************************************************************/

// Cut alias of the strings.Cut
func Cut(s, sep string) (before string, after string, found bool) {
	_ = "STUB: not implemented"
	return "", "", false
}

// QuietCut always returns two substring.
func QuietCut(s, sep string) (before string, after string) {
	_ = "STUB: not implemented"
	return "", ""
}

// MustCut always returns two substring.
func MustCut(s, sep string) (before string, after string) { _ = "STUB: not implemented"; return "", "" }

// TrimCut always returns two substring and trim space for items.
func TrimCut(s, sep string) (string, string) { _ = "STUB: not implemented"; return "", "" }

// SplitKV split string to key and value.
func SplitKV(s, sep string) (string, string) {
	_ = "STUB: not implemented"
	return "",

		// SplitValid string to slice. will trim each item and filter empty string node.
		""
}

func SplitValid(s, sep string) (ss []string) { _ = "STUB: not implemented"; return nil }

// Split string to slice. will trim each item and filter empty string node.
func Split(s, sep string) (ss []string) { _ = "STUB: not implemented"; return nil }

// SplitNValid string to slice. will filter empty string node.
func SplitNValid(s, sep string, n int) (ss []string) { _ = "STUB: not implemented"; return nil }

// SplitN string to slice. will filter empty string node.
func SplitN(s, sep string, n int) (ss []string) { _ = "STUB: not implemented"; return nil }

// SplitTrimmed split string to slice.
// will trim space for each node, but not filter empty
func SplitTrimmed(s, sep string) (ss []string) { _ = "STUB: not implemented"; return nil }

// SplitNTrimmed split string to slice.
// will trim space for each node, but not filter empty
func SplitNTrimmed(s, sep string, n int) (ss []string) { _ = "STUB: not implemented"; return nil }

// 根据空白字符（空格，TAB，换行等）分隔字符串
var whitespaceRegexp = regexp.MustCompile("\\s+")

// SplitByWhitespace Separate strings by whitespace characters (space, TAB, newline, etc.)
func SplitByWhitespace(s string) []string { _ = "STUB: not implemented"; return nil }

// Substr for a string. NOTE: strLn := len(runes)
// if length <= 0, return pos to end.
func Substr(s string, pos, length int) string { _ = "STUB: not implemented"; return "" }

// pos is too large

// SplitInlineComment for an inline text string. default is strict mode.
func SplitInlineComment(val string, strict ...bool) (string, string) {
	_ = "STUB: not implemented"
	// strict check: must with a space
	return "", ""
}

// FirstLine from command output
func FirstLine(output string) string { _ = "STUB: not implemented"; return "" }
