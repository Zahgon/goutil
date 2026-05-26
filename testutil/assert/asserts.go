package assert

import (
	"reflect"

	"github.com/gookit/goutil/comdef"
)

// Nil asserts that the given is a nil value
func Nil(t TestingT, give any, fmtAndArgs ...any) bool { _ = "STUB: not implemented"; return false }

// NotNil asserts that the given is a not nil value
func NotNil(t TestingT, give any, fmtAndArgs ...any) bool { _ = "STUB: not implemented"; return false }

// True asserts that the given is a bool true
func True(t TestingT, give bool, fmtAndArgs ...any) bool { _ = "STUB: not implemented"; return false }

// False asserts that the given is a bool false
func False(t TestingT, give bool, fmtAndArgs ...any) bool { _ = "STUB: not implemented"; return false }

// Empty asserts that the give should be empty
func Empty(t TestingT, give any, fmtAndArgs ...any) bool { _ = "STUB: not implemented"; return false }

// NotEmpty asserts that the give should not be empty
func NotEmpty(t TestingT, give any, fmtAndArgs ...any) bool {
	_ = "STUB: not implemented"
	return false
}

// Zero asserts that the give should be zero value
func Zero(t TestingT, give any, fmtAndArgs ...any) bool { _ = "STUB: not implemented"; return false }

// NotZero asserts that the give should not be zero value
func NotZero(t TestingT, give any, fmtAndArgs ...any) bool { _ = "STUB: not implemented"; return false }

// PanicRunFunc define
type PanicRunFunc func()

// didPanic returns true if the function passed to it panics. Otherwise, it returns false.
func runPanicFunc(f PanicRunFunc) (didPanic bool, message any, stack string) {
	_ = "STUB: not implemented"
	return false, *new(any), ""
}

// call the target function

// Panics asserts that the code inside the specified func panics.
func Panics(t TestingT, fn PanicRunFunc, fmtAndArgs ...any) bool {
	_ = "STUB: not implemented"
	return false
}

// NotPanics asserts that the code inside the specified func NOT panics.
func NotPanics(t TestingT, fn PanicRunFunc, fmtAndArgs ...any) bool {
	_ = "STUB: not implemented"
	return false
}

// PanicsMsg should panic and with a value
func PanicsMsg(t TestingT, fn PanicRunFunc, wantVal any, fmtAndArgs ...any) bool {
	_ = "STUB: not implemented"
	return false
}

// PanicsErrMsg should panic and with error message
func PanicsErrMsg(t TestingT, fn PanicRunFunc, errMsg string, fmtAndArgs ...any) bool {
	_ = "STUB: not implemented"
	return false
}

// Contains asserting that the given data(string,slice,map) should contain element
//
// TIP: only support types: string, map, array, slice
//
//	map         - check key exists
//	string      - check substring exists
//	array,slice - check sub-element exists
func Contains(t TestingT, src, elem any, fmtAndArgs ...any) bool {
	_ = "STUB: not implemented"
	return false
}

// src invalid

// not found

// NotContains asserts that the given data(string,slice,map) should not contain element
//
// TIP: only support types: string, map, array, slice
//
//	map         - check key exists
//	string      - check substring exists
//	array,slice - check sub-element exists
func NotContains(t TestingT, src, elem any, fmtAndArgs ...any) bool {
	_ = "STUB: not implemented"
	return false
}

// src invalid

// found

// ContainsKey asserts that the given map is containing key
func ContainsKey(t TestingT, mp, key any, fmtAndArgs ...any) bool {
	_ = "STUB: not implemented"
	return false
}

// NotContainsKey asserts that the given map is not contains key
func NotContainsKey(t TestingT, mp, key any, fmtAndArgs ...any) bool {
	_ = "STUB: not implemented"
	return false
}

// ContainsKeys asserts that the map is contains all given keys
//
// Usage:
//
//	ContainsKeys(t, map[string]any{...}, []string{"key1", "key2"})
func ContainsKeys(t TestingT, mp any, keys any, fmtAndArgs ...any) bool {
	_ = "STUB: not implemented"
	return false
}

// NotContainsKeys asserts that the map is not contains all given keys
//
// Usage:
//
//	NotContainsKeys(t, map[string]any{...}, []string{"key1", "key2"})
func NotContainsKeys(t TestingT, mp any, keys any, fmtAndArgs ...any) bool {
	_ = "STUB: not implemented"
	return false
}

// ContainsElems asserts that the given list should contain sub elements.
func ContainsElems[T comdef.ScalarType](t TestingT, list, sub []T, fmtAndArgs ...any) bool {
	_ = "STUB: not implemented"
	return false
}

// not contains all

// ContainsElemsAny asserts that the given list should contain sub elements.
//   - non-generic wrapper for ContainsElems
func ContainsElemsAny(t TestingT, list, sub any, fmtAndArgs ...any) bool {
	_ = "STUB: not implemented"
	return false
}

// check all elements in sub are contained in list

// StrContains asserts that the given string should contain substring
func StrContains(t TestingT, s, sub string, fmtAndArgs ...any) bool {
	_ = "STUB: not implemented"
	return false
}

// StrNotContains asserts that the given string should not contain substring
func StrNotContains(t TestingT, s, sub string, fmtAndArgs ...any) bool {
	_ = "STUB: not implemented"
	return false
}

// StrContainsAll asserts that the given string should contain all substrings
func StrContainsAll(t TestingT, s string, subs []string, fmtAndArgs ...any) bool {
	_ = "STUB: not implemented"
	return false
}

// StrCount asserts that the given string should contain substring and count
func StrCount(t TestingT, s, sub string, count int, fmtAndArgs ...any) bool {
	_ = "STUB: not implemented"
	return false
}

//
// -------------------- filesystem --------------------
//

// FileExists asserts that the given file exists
func FileExists(t TestingT, filePath string, fmtAndArgs ...any) bool {
	_ = "STUB: not implemented"
	return false
}

// FileNotExists asserts that the given file not exists
func FileNotExists(t TestingT, filePath string, fmtAndArgs ...any) bool {
	_ = "STUB: not implemented"
	return false
}

// DirExists asserts that the given dir exists
func DirExists(t TestingT, dirPath string, fmtAndArgs ...any) bool {
	_ = "STUB: not implemented"
	return false
}

// DirNotExists asserts that the given dir not exists
func DirNotExists(t TestingT, dirPath string, fmtAndArgs ...any) bool {
	_ = "STUB: not implemented"
	return false
}

//
// -------------------- error --------------------
//

// NoError asserts that the given is a nil error. alias of NoError()
func NoError(t TestingT, err error, fmtAndArgs ...any) bool {
	_ = "STUB: not implemented"
	return false
}

// NoErr asserts that the given is a nil error
func NoErr(t TestingT, err error, fmtAndArgs ...any) bool { _ = "STUB: not implemented"; return false }

// Error asserts that the given is a not nil error. alias of Error()
func Error(t TestingT, err error, fmtAndArgs ...any) bool { _ = "STUB: not implemented"; return false }

// Err asserts that the given is a not nil error
func Err(t TestingT, err error, fmtAndArgs ...any) bool { _ = "STUB: not implemented"; return false }

// ErrIs asserts that the given error is equals wantErr
func ErrIs(t TestingT, err, wantErr error, fmtAndArgs ...any) bool {
	_ = "STUB: not implemented"
	return false
}

// ErrMsg asserts that the given is a not nil error and error message equals wantMsg
func ErrMsg(t TestingT, err error, wantMsg string, fmtAndArgs ...any) bool {
	_ = "STUB: not implemented"
	return false
}

// ErrMsgContains asserts that the given is a not nil error and error message contains subMsg
func ErrMsgContains(t TestingT, err error, subMsg string, fmtAndArgs ...any) bool {
	_ = "STUB: not implemented"
	return false
}

// ErrSubMsg asserts that the given is a not nil error and the error message contains subMsg
func ErrSubMsg(t TestingT, err error, subMsg string, fmtAndArgs ...any) bool {
	_ = "STUB: not implemented"
	return false
}

// ErrHasMsg asserts that the given is a not nil error and the error message contains subMsg
func ErrHasMsg(t TestingT, err error, subMsg string, fmtAndArgs ...any) bool {
	_ = "STUB: not implemented"
	return false
}

//
// -------------------- Len --------------------
//

// Len assert given length is equals to wantLn
func Len(t TestingT, give any, wantLn int, fmtAndArgs ...any) bool {
	_ = "STUB: not implemented"
	return false
}

// LenGt assert given length is greater than to minLn
func LenGt(t TestingT, give any, minLn int, fmtAndArgs ...any) bool {
	_ = "STUB: not implemented"
	return false
}

//
// -------------------- compare --------------------
//

// Equal asserts that the want should equal to the given.
//
// alias of Eq()
func Equal(t TestingT, want, give any, fmtAndArgs ...any) bool {
	_ = "STUB: not implemented"
	return false
}

// Eq asserts that the want should equal to the given
func Eq(t TestingT, want, give any, fmtAndArgs ...any) bool {
	_ = "STUB: not implemented"
	return false
}

// float 特殊处理 - 默认按 0.009 精度对比

// TODO diff := diff(want, give)

// Neq asserts that the want should not be equal to the given.
//
// alias of NotEq()
func Neq(t TestingT, want, give any, fmtAndArgs ...any) bool {
	_ = "STUB: not implemented"
	return false
}

// NotEqual asserts that the want should not be equal to the given.
//
// alias of NotEq()
func NotEqual(t TestingT, want, give any, fmtAndArgs ...any) bool {
	_ = "STUB: not implemented"
	return false
}

// NotEq asserts that the want should not be equal to the given
func NotEq(t TestingT, want, give any, fmtAndArgs ...any) bool {
	_ = "STUB: not implemented"
	return false
}

// Lt asserts that the give(intX,uintX,floatX) should not be less than max
func Lt(t TestingT, give, max any, fmtAndArgs ...any) bool { _ = "STUB: not implemented"; return false }

// Lte asserts that the give(intX,uintX,floatX) should not be less than or equals to max
func Lte(t TestingT, give, max any, fmtAndArgs ...any) bool {
	_ = "STUB: not implemented"
	return false
}

// Gt asserts that the give(intX,uintX,floatX) should not be greater than min
func Gt(t TestingT, give, min any, fmtAndArgs ...any) bool { _ = "STUB: not implemented"; return false }

// Gte asserts that the give(intX,uintX,floatX) should not be greater than or equals to min
func Gte(t TestingT, give, min any, fmtAndArgs ...any) bool {
	_ = "STUB: not implemented"
	return false
}

// EqInt asserts that the want should equal to the given intX.
//
// NOTE: Will always convert to int64 to compare.
//
// Example:
//
//	assert.Eq(t, uint(1), int(1)) // false
//	assert.EqInt(t, uint(1), int(1)) // true
func EqInt(t TestingT, want, give any, fmtAndArgs ...any) bool {
	_ = "STUB: not implemented"
	return false
}

// EqFloat asserts that the want should equal to the given with delta. alias of InDelta()
func EqFloat(t TestingT, want, give any, delta float64, fmtAndArgs ...any) bool {
	_ = "STUB: not implemented"
	return false
}

// InDelta assert that two floating-point values differ from each other within a certain range.
func InDelta(t TestingT, want, give any, delta float64, fmtAndArgs ...any) bool {
	_ = "STUB: not implemented"
	return false
}

// IsType assert data type equals
//
// Usage:
//
//	assert.IsType(t, 0, val) // assert type is int
func IsType(t TestingT, wantType, give any, fmtAndArgs ...any) bool {
	_ = "STUB: not implemented"
	return false
}

// IsKind assert data reflect.Kind equals.
// If `give` is ptr or interface, will get real kind.
//
// Usage:
//
//	assert.IsKind(t, reflect.Int, val) // assert type is int kind.
func IsKind(t TestingT, wantKind reflect.Kind, give any, fmtAndArgs ...any) bool {
	_ = "STUB: not implemented"
	return false
}

// Same asserts that two pointers reference the same object.
//
//	assert.Same(t, ptr1, ptr2)
//
// Both arguments must be pointer variables. Pointer variable sameness is
// determined based on the equality of both type and value.
func Same(t TestingT, wanted, actual any, fmtAndArgs ...any) bool {
	_ = "STUB: not implemented"
	return false
}

// NotSame asserts that two pointers do not reference the same object.
//
//	assert.NotSame(t, ptr1, ptr2)
//
// Both arguments must be pointer variables. Pointer variable sameness is
// determined based on the equality of both type and value.
func NotSame(t TestingT, want, actual any, fmtAndArgs ...any) bool {
	_ = "STUB: not implemented"
	return false
}

// samePointers compares two generic interface objects and returns whether
// they point to the same object
func samePointers(first, second any) bool { _ = "STUB: not implemented"; return false }

// compare pointer addresses
