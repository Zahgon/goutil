// Package strutil provide some string,char,byte util functions
package strutil

import (
	"github.com/gookit/goutil/comdef"
)

// OrCond return s1 on cond is True, OR return s2.
// Like: cond ? s1 : s2
func OrCond(cond bool, s1, s2 string) string { _ = "STUB: not implemented"; return "" }

// BlankOr return default value on val is blank, otherwise return val
func BlankOr(val, defVal string) string { _ = "STUB: not implemented"; return "" }

// ZeroOr return default value on val is zero, otherwise return val. same of OrElse()
func ZeroOr[T ~string](val, defVal T) T { _ = "STUB: not implemented"; return *new(T) }

// ErrorOr return default value on err is not nil, otherwise return s
// func ErrorOr(s string, err error, defVal string) string {
// 	if err != nil {
// 		return defVal
// 	}
// 	return s
// }

// OrElse return default value on s is empty, otherwise return s
func OrElse(s, orVal string) string { _ = "STUB: not implemented"; return "" }

// OrElseNilSafe return default value on s is nil, otherwise return s
func OrElseNilSafe(s *string, orVal string) string { _ = "STUB: not implemented"; return "" }

// OrHandle return fn(s) on s is not empty.
func OrHandle(s string, fn comdef.StringHandleFunc) string { _ = "STUB: not implemented"; return "" }

// Valid return first not empty element.
func Valid(ss ...string) string { _ = "STUB: not implemented"; return "" }

// SubstrCount returns the number of times the substr substring occurs in the s string.
// Actually, it comes from strings.Count().
//
//   - s The string to search in
//   - substr The substring to search for
//   - params[0] The offset where to start counting.
//   - params[1] The maximum length after the specified offset to search for the substring.
func SubstrCount(s, substr string, params ...uint64) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
