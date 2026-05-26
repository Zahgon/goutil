// Package errorx provide an enhanced error implements for go,
// allow with stacktraces and wrap another error.
package errorx

import (
	"fmt"
	"io"
)

// Causer interface for get first cause error
type Causer interface {
	// Cause returns the first cause error by call err.Cause().
	// Otherwise, will returns current error.
	Cause() error
}

// Unwrapper interface for get previous error
type Unwrapper interface {
	// Unwrap returns previous error by call err.Unwrap().
	// Otherwise, will returns nil.
	Unwrap() error
}

// XErrorFace interface
type XErrorFace interface {
	error
	Causer
	Unwrapper
}

// Exception interface
// type Exception interface {
// 	XErrorFace
// 	Code() string
// 	Message() string
// 	StackString() string
// }

/*************************************************************
 * implements XErrorFace interface
 *************************************************************/

// ErrorX struct
//
// TIPS:
//
//	fmt pkg call order: Format > GoString > Error > String
type ErrorX struct {
	// trace stack
	*stack
	prev error
	msg  string
}

// Cause implements Causer.
func (e *ErrorX) Cause() error { _ = "STUB: not implemented"; return nil }

// Unwrap implements Unwrapper.
func (e *ErrorX) Unwrap() error {
	_ = "STUB: not implemented"

	// Format error, will output stack information.
	return nil
}

func (e *ErrorX) Format(s fmt.State, verb rune) {
	_ = "STUB: not implemented"
	// format current error: only output on have msg
	return
}

// format prev error

// GoString to GO string, contains stack information.
// printing an error with %#v will produce useful information.
func (e *ErrorX) GoString() string {
	_ = "STUB: not implemented"
	// var sb strings.Builder
	return ""
}

// Error msg string, not contains stack information.
func (e *ErrorX) Error() string { _ = "STUB: not implemented"; return "" }

// String error to string, contains stack information.
func (e *ErrorX) String() string { _ = "STUB: not implemented"; return "" }

// WriteTo write the error to a writer, contains stack information.
func (e *ErrorX) WriteTo(w io.Writer) (n int64, err error) {
	_ = "STUB: not implemented"
	// current error: only output on have msg
	return 0, nil
}

// with stack

// with prev error

// Message error message of current
func (e *ErrorX) Message() string {
	_ = "STUB: not implemented"

	// StackString returns error stack string of current.
	return ""
}

func (e *ErrorX) StackString() string { _ = "STUB: not implemented"; return "" }

// writeMsgTo write the error msg to a writer
func (e *ErrorX) writeMsgTo(w io.Writer) {
	_ = "STUB: not implemented"
	// current error
	return
}

// with prev error

// CallerFunc returns the error caller func. if stack is nil, will return nil
func (e *ErrorX) CallerFunc() *Func { _ = "STUB: not implemented"; return nil }

// Location information for the caller func. more please see CallerFunc
//
// Returns eg:
//
//	github.com/gookit/goutil/errorx_test.TestWithPrev(), errorx_test.go:34
func (e *ErrorX) Location() string { _ = "STUB: not implemented"; return "" }

/*************************************************************
 * new error with call stacks
 *************************************************************/

// New error message and with caller stacks
func New(msg string) error { _ = "STUB: not implemented"; return nil }

// Newf error with format message, and with caller stacks.
// alias of Errorf()
func Newf(tpl string, vars ...any) error { _ = "STUB: not implemented"; return nil }

// Errorf error with format message, and with caller stacks
func Errorf(tpl string, vars ...any) error { _ = "STUB: not implemented"; return nil }

// With prev error and error message, and with caller stacks
func With(err error, msg string) error { _ = "STUB: not implemented"; return nil }

// Withf error and with format message, and with caller stacks
func Withf(err error, tpl string, vars ...any) error { _ = "STUB: not implemented"; return nil }

// WithPrev error and message, and with caller stacks. alias of With()
func WithPrev(err error, msg string) error { _ = "STUB: not implemented"; return nil }

// WithPrevf error and with format message, and with caller stacks. alias of Withf()
func WithPrevf(err error, tpl string, vars ...any) error { _ = "STUB: not implemented"; return nil }

/*************************************************************
 * wrap go error with call stacks
 *************************************************************/

// WithStack wrap a go error with a stacked trace. If err is nil, will return nil.
func WithStack(err error) error { _ = "STUB: not implemented"; return nil }

// prev:  err,

// Traced warp a go error and with caller stacks. alias of WithStack()
func Traced(err error) error { _ = "STUB: not implemented"; return nil }

// Stacked warp a go error and with caller stacks. alias of WithStack()
func Stacked(err error) error { _ = "STUB: not implemented"; return nil }

// WithOptions new error with some option func
func WithOptions(msg string, fns ...func(opt *ErrStackOpt)) error {
	_ = "STUB: not implemented"
	return nil
}

/*************************************************************
 * helper func for wrap error without stacks
 *************************************************************/

// Wrap error and with message, but not with stack
func Wrap(err error, msg string) error { _ = "STUB: not implemented"; return nil }

// Wrapf error with format message, but not with stack
func Wrapf(err error, tpl string, vars ...any) error { _ = "STUB: not implemented"; return nil }
