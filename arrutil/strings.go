package arrutil

import (
	"github.com/gookit/goutil/comdef"
)

// StringsToAnys convert []string to []any
func StringsToAnys(ss []string) []any { _ = "STUB: not implemented"; return nil }

// StringsToSlice convert []string to []any. alias of StringsToAnys()
func StringsToSlice(ss []string) []any { _ = "STUB: not implemented"; return nil }

// StringsAsInts convert and ignore error
func StringsAsInts(ss []string) []int { _ = "STUB: not implemented"; return nil }

// StringsToInts string slice to int slice
func StringsToInts(ss []string) (ints []int, err error) {
	_ = "STUB: not implemented"
	return nil,

		// StringsTryInts string slice to int slice
		nil
}

func StringsTryInts(ss []string) (ints []int, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// StringsUnique unique string slice
func StringsUnique(ss []string) []string { _ = "STUB: not implemented"; return nil }

// StringsContains check string slice contains string
func StringsContains(ss []string, s string) bool { _ = "STUB: not implemented"; return false }

// StringsRemove value form a string slice
func StringsRemove(ss []string, s string) []string { _ = "STUB: not implemented"; return nil }

// StringsFilter given strings, default will filter emtpy string.
//
// Usage:
//
//	// output: [a, b]
//	ss := arrutil.StringsFilter([]string{"a", "", "b", ""})
func StringsFilter(ss []string, filter ...comdef.StringMatchFunc) []string {
	_ = "STUB: not implemented"
	return nil
}

// StringsMap handle each string item, map to new strings
func StringsMap(ss []string, mapFn func(s string) string) []string {
	_ = "STUB: not implemented"
	return nil
}

// TrimStrings trim string slice item.
//
// Usage:
//
//	// output: [a, b, c]
//	ss := arrutil.TrimStrings([]string{",a", "b.", ",.c,"}, ",.")
func TrimStrings(ss []string, cutSet ...string) []string { _ = "STUB: not implemented"; return nil }
