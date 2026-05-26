package strutil

import (
	"errors"
	"regexp"
	"time"

	"github.com/gookit/goutil/internal/comfunc"
)

var (
	// ErrDateLayout error
	ErrDateLayout = errors.New("invalid date layout string")
	// ErrInvalidParam error
	ErrInvalidParam = errors.New("invalid input for parse time")

	// some regex for convert string.
	toSnakeReg  = regexp.MustCompile("[A-Z][a-z]")
	toCamelRegs = map[string]*regexp.Regexp{
		" ": regexp.MustCompile(" +[a-zA-Z]"),
		"-": regexp.MustCompile("-+[a-zA-Z]"),
		"_": regexp.MustCompile("_+[a-zA-Z]"),
	}
)

// Internal func refers:
// strconv.QuoteRune()
// strconv.QuoteToASCII()
// strconv.AppendQuote()
// strconv.AppendQuoteRune()

// Quote alias of strings.Quote
func Quote(s string) string { _ = "STUB: not implemented"; return "" }

// Unquote remove start and end quotes by single-quote or double-quote
//
// tip: strconv.Unquote cannot unquote single-quote
func Unquote(s string) string { _ = "STUB: not implemented"; return "" }

// exclude quotes

// strconv.Unquote cannot unquote single-quote
// if ns, err := strconv.Unquote(s); err == nil {
// 	return ns
// }

// Join alias of strings.Join
func Join(sep string, ss ...string) string { _ = "STUB: not implemented"; return "" }

// JoinList alias of strings.Join
func JoinList(sep string, ss []string) string { _ = "STUB: not implemented"; return "" }

// JoinComma quick join strings by comma
func JoinComma(ss []string) string { _ = "STUB: not implemented"; return "" }

// JoinAny type to string
func JoinAny(sep string, parts ...any) string { _ = "STUB: not implemented"; return "" }

// Implode alias of strings.Join
func Implode(sep string, ss ...string) string { _ = "STUB: not implemented"; return "" }

/*************************************************************
 * region value to string
 *************************************************************/

// String convert value to string, return error on failed
func String(val any) (string, error) {
	_ = "STUB: not implemented"
	return "",

		// ToString convert value to string, return error on failed
		nil
}

func ToString(val any) (string, error) {
	_ = "STUB: not implemented"
	return "",

		// StringOrErr convert value to string, return error on failed
		nil
}

func StringOrErr(val any) (string, error) {
	_ = "STUB: not implemented"
	return "",

		// QuietString convert value to string, will ignore error. same as SafeString()
		nil
}

func QuietString(val any) string { _ = "STUB: not implemented"; return "" }

// SafeString convert value to string. Will ignore error
func SafeString(in any) string { _ = "STUB: not implemented"; return "" }

// StringOrPanic convert value to string, will panic on error
func StringOrPanic(val any) string { _ = "STUB: not implemented"; return "" }

// MustString convert value to string. will panic on error
func MustString(val any) string { _ = "STUB: not implemented"; return "" }

// StringOrDefault convert any value to string, return default value on failed
func StringOrDefault(val any, defVal string) string { _ = "STUB: not implemented"; return "" }

// StringOr convert any value to string, return default value on failed
func StringOr(val any, defVal string) string { _ = "STUB: not implemented"; return "" }

// AnyToString convert any value to string.
//
// For defaultAsErr:
//
//   - False  will use fmt.Sprint convert unsupported type
//   - True   will return error on convert fail.
func AnyToString(val any, defaultAsErr bool) (s string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ToStringWith try to convert value to string. can with some option func, more see comfunc.ConvOption.
func ToStringWith(in any, optFns ...comfunc.ConvOptionFn) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

/*************************************************************
 * region string value to bool
 *************************************************************/

// ToBool convert string to bool
func ToBool(s string) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// QuietBool convert to bool, will ignore error
func QuietBool(s string) bool {
	_ = "STUB: not implemented"

	// SafeBool convert to bool and will ignore error
	return false
}

func SafeBool(s string) bool { _ = "STUB: not implemented"; return false }

// MustBool convert to bool and will panic on error
func MustBool(s string) bool { _ = "STUB: not implemented"; return false }

// Bool parse string to bool. like strconv.ParseBool()
func Bool(s string) (bool, error) { _ = "STUB: not implemented"; return false, nil }

/*************************************************************
 * region string value to int
 *************************************************************/

// Int convert string to int, alias of ToInt()
func Int(s string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// ToInt convert string to int, return error on fail
func ToInt(s string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// IntOrDefault convert string to int, return default value on fail
func IntOrDefault(s string, defVal int) int { _ = "STUB: not implemented"; return 0 }

// IntOr convert string to int, return default value on fail
func IntOr(s string, defVal int) int { _ = "STUB: not implemented"; return 0 }

// SafeInt convert string to int, will ignore error
func SafeInt(s string) int { _ = "STUB: not implemented"; return 0 }

// QuietInt convert string to int, will ignore error
func QuietInt(s string) int {
	_ = "STUB: not implemented"

	// MustInt convert string to int, will panic on error
	return 0
}

func MustInt(s string) int { _ = "STUB: not implemented"; return 0 }

// IntOrPanic convert value to int, will panic on error
func IntOrPanic(s string) int { _ = "STUB: not implemented"; return 0 }

/*************************************************************
 * region convert string to int64
 *************************************************************/

// Int64 convert string to int, will ignore error
func Int64(s string) int64 {
	_ = "STUB: not implemented"

	// QuietInt64 convert string to int, will ignore error
	return 0
}

func QuietInt64(s string) int64 {
	_ = "STUB: not implemented"

	// SafeInt64 convert string to int, will ignore error
	return 0
}

func SafeInt64(s string) int64 { _ = "STUB: not implemented"; return 0 }

// ToInt64 convert string to int, return error on fail
func ToInt64(s string) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// Int64OrDefault convert string to int, return default value on fail
func Int64OrDefault(s string, defVal int64) int64 { _ = "STUB: not implemented"; return 0 }

// Int64Or convert string to int, return default value on fail
func Int64Or(s string, defVal int64) int64 { _ = "STUB: not implemented"; return 0 }

// Int64OrErr convert string to int, return error on fail
func Int64OrErr(s string) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// MustInt64 convert value to int, will panic on error
func MustInt64(s string) int64 { _ = "STUB: not implemented"; return 0 }

// Int64OrPanic convert value to int, will panic on error
func Int64OrPanic(s string) int64 { _ = "STUB: not implemented"; return 0 }

/*************************************************************
 * region string value to uint
 *************************************************************/

// Uint convert string to uint, will ignore error
func Uint(s string) uint64 {
	_ = "STUB: not implemented"

	// SafeUint convert string to uint, will ignore error
	return 0
}

func SafeUint(s string) uint64 { _ = "STUB: not implemented"; return 0 }

// ToUint convert string to uint, return error on fail. alias of UintOrErr()
func ToUint(s string) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

// UintOrErr convert string to uint, return error on fail
func UintOrErr(s string) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

// MustUint convert value to uint, will panic on error. alias of UintOrPanic()
func MustUint(s string) uint64 { _ = "STUB: not implemented"; return 0 }

// UintOrPanic convert value to uint, will panic on error
func UintOrPanic(s string) uint64 { _ = "STUB: not implemented"; return 0 }

// UintOrDefault convert string to uint, return default value on fail
func UintOrDefault(s string, defVal uint64) uint64 { _ = "STUB: not implemented"; return 0 }

// UintOr convert string to uint, return default value on fail
func UintOr(s string, defVal uint64) uint64 { _ = "STUB: not implemented"; return 0 }

/*************************************************************
 * region string value to byte
 * refer from https://github.com/valyala/fastjson/blob/master/util.go
 *************************************************************/

// Byte2str convert bytes to string
func Byte2str(b []byte) string { _ = "STUB: not implemented"; return "" }

// Byte2string convert bytes to string
func Byte2string(b []byte) string { _ = "STUB: not implemented"; return "" }

// ToBytes convert string to bytes
func ToBytes(s string) (b []byte) { _ = "STUB: not implemented"; return nil }

/*************************************************************
 * region string to int/string slice, time.Time
 *************************************************************/

// Ints alias of the ToIntSlice(). default sep is comma(,)
func Ints(s string, sep ...string) []int { _ = "STUB: not implemented"; return nil }

// ToInts alias of the ToIntSlice(). default sep is comma(,)
func ToInts(s string, sep ...string) ([]int, error) {
	_ = "STUB: not implemented"
	return nil,

		// ToIntSlice split string to slice and convert item to int.
		//
		// Default sep is comma
		nil
}

func ToIntSlice(s string, sep ...string) (ints []int, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ToArray alias of the ToSlice()
func ToArray(s string, sep ...string) []string { _ = "STUB: not implemented"; return nil }

// Strings alias of the ToSlice()
func Strings(s string, sep ...string) []string { _ = "STUB: not implemented"; return nil }

// ToStrings alias of the ToSlice()
func ToStrings(s string, sep ...string) []string { _ = "STUB: not implemented"; return nil }

// ToSlice split string to array.
func ToSlice(s string, sep ...string) []string { _ = "STUB: not implemented"; return nil }

// ToDuration parses a duration string. such as "300ms", "-1.5h" or "2h45m".
// Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
func ToDuration(s string) (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}
