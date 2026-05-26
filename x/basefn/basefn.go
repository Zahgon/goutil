// Package basefn provide some no-dependents util functions
package basefn

// Panicf format panic message use fmt.Sprintf
func Panicf(format string, v ...any) { _ = "STUB: not implemented"; return }

// PanicIf if cond = true, panics with an error message
func PanicIf(cond bool, fmtAndArgs ...any) { _ = "STUB: not implemented"; return }

// PanicErr panics if error is not empty
func PanicErr(err error) { _ = "STUB: not implemented"; return }

// MustOK if error is not empty, will panic
func MustOK(err error) { _ = "STUB: not implemented"; return }

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
func Must[T any](v T, err error) T { _ = "STUB: not implemented"; return *new(T) }

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
//	basefn.MustIgnore(fn())
func MustIgnore(_ any, err error) {
	_ = "STUB: not implemented"

	// ErrOnFail return input error on cond is false, otherwise return nil
	return
}

func ErrOnFail(cond bool, err error) error { _ = "STUB: not implemented"; return nil }

// OrError return input error on cond is false, otherwise return nil
func OrError(cond bool, err error) error { _ = "STUB: not implemented"; return nil }

// FirstOr get first elem or elseVal
func FirstOr[T any](sl []T, elseVal T) T { _ = "STUB: not implemented"; return *new(T) }

// OrValue get. like: if cond { okVal } else { elVal }
func OrValue[T any](cond bool, okVal, elVal T) T { _ = "STUB: not implemented"; return *new(T) }

// OrReturn call okFunc() on condition is true, else call elseFn()
//
// like expr: if cond { okFunc() } else { elseFn() }
func OrReturn[T any](cond bool, okFn, elseFn func() T) T { _ = "STUB: not implemented"; return *new(T) }

// ErrFunc type
type ErrFunc func() error

// CallOn call func on condition is true
func CallOn(cond bool, fn ErrFunc) error { _ = "STUB: not implemented"; return nil }

// CallOrElse call okFunc() on condition is true, else call elseFn()
func CallOrElse(cond bool, okFn, elseFn ErrFunc) error { _ = "STUB: not implemented"; return nil }
