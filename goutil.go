// Package goutil 💪 Useful utils for Go: byte, int, string, array/slice, map, struct, reflect, error, time, format, CLI, ENV, filesystem,
// system, testing, debug and more.
package goutil

import (
	"github.com/gookit/goutil/structs"
)

// Value alias of structs.Value
type Value = structs.Value

// Panicf format panic message use fmt.Sprintf
func Panicf(format string, v ...any) { _ = "STUB: not implemented"; return }

// PanicIf if cond = true, panics with an error message
func PanicIf(cond bool, fmtAndArgs ...any) { _ = "STUB: not implemented"; return }

// PanicErr if error is not empty, will panic.
// Alias of basefn.PanicErr()
func PanicErr(err error) { _ = "STUB: not implemented"; return }

// PanicIfErr if error is not empty, will panic.
// Alias of basefn.PanicErr()
func PanicIfErr(err error) {
	_ = "STUB: not implemented"

	// MustOK if error is not empty, will panic.
	// Alias of basefn.MustOK()
	return
}

func MustOK(err error) {
	_ = "STUB: not implemented"

	// MustIgnore for return like (v, error). Ignore return v and will panic on error.
	//
	// Useful for io, file operation func: (n int, err error)
	//
	// Usage:
	//
	//	// old
	//	_, err := fn()
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	// new
	//	goutil.MustIgnore(fn())
	return
}

func MustIgnore(_ any, err error) {
	_ = "STUB: not implemented"

	// Must return like (v, error). will panic on error, otherwise return v.
	//
	// Usage:
	//
	//	// old
	//	v, err := fn()
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	// new
	//	v := goutil.Must(fn())
	return
}

func Must[T any](v T, err error) T { _ = "STUB: not implemented"; return *new(T) }

// ErrOnFail return input error on cond is false, otherwise return nil
func ErrOnFail(cond bool, err error) error { _ = "STUB: not implemented"; return nil }

// OrError return input error on cond is false, otherwise return nil
func OrError(cond bool, err error) error { _ = "STUB: not implemented"; return nil }

// OrValue get. like: if cond { okVal } else { elVal }
func OrValue[T any](cond bool, okVal, elVal T) T { _ = "STUB: not implemented"; return *new(T) }

// OrReturn call okFunc() on condition is true, else call elseFn()
func OrReturn[T any](cond bool, okFn, elseFn func() T) T { _ = "STUB: not implemented"; return *new(T) }

//
// ------------------------- check functions -------------------------
//

// IsNil value check
func IsNil(v any) bool { _ = "STUB: not implemented"; return false }

// IsZero value check, alias of the IsEmpty()
var IsZero = IsEmpty

// IsEmpty value check
func IsEmpty(v any) bool { _ = "STUB: not implemented"; return false }

// IsZeroReal Alias of the IsEmptyReal()
var IsZeroReal = IsEmptyReal

// IsEmptyReal checks for empty given value and also real empty value if the passed value is a pointer
func IsEmptyReal(v any) bool { _ = "STUB: not implemented"; return false }

// IsFunc value
func IsFunc(val any) bool { _ = "STUB: not implemented"; return false }

// IsEqual determines if two objects are considered equal.
//
// TIP: cannot compare a function type
func IsEqual(src, dst any) bool { _ = "STUB: not implemented"; return false }

// cannot compare a function type

// Contains try loop over the data check if the data includes the element.
// alias of the IsContains
//
// TIP: only support types: string, map, array, slice
//
//	map         - check key exists
//	string 	    - check sub-string exists
//	array,slice - check sub-element exists
func Contains(data, elem any) bool { _ = "STUB: not implemented"; return false }

// IsContains try loop over the data check if the data includes the element.
//
// TIP: only support types: string, map, array, slice
//
//	map         - check key exists
//	string 	    - check sub-string exists
//	array,slice - check sub-element exists
func IsContains(data, elem any) bool { _ = "STUB: not implemented"; return false }

//
// ------------------------- goinfo functions -------------------------
//

// FuncName get func name
func FuncName(f any) string { _ = "STUB: not implemented"; return "" }

// PkgName get the current package name. alias of goinfo.PkgName()
//
// Usage:
//
//	funcName := goutil.FuncName(fn)
//	pgkName := goutil.PkgName(funcName)
func PkgName(funcName string) string { _ = "STUB: not implemented"; return "" }
