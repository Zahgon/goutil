package errorx

import (
	"fmt"
)

// ErrorCoder interface
type ErrorCoder interface {
	error
	Code() int
}

// ErrorR useful for web service replay/response.
// code == 0 is successful. otherwise, is failed.
type ErrorR interface {
	ErrorCoder
	fmt.Stringer
	IsSuc() bool
	IsFail() bool
}

// error reply struct
type errorR struct {
	code int
	msg  string
}

// NewR code with error response
func NewR(code int, msg string) ErrorR { _ = "STUB: not implemented"; return *new(ErrorR) }

// Fail code with error response
func Fail(code int, msg string) ErrorR { _ = "STUB: not implemented"; return *new(ErrorR) }

// Failf code with error response
func Failf(code int, tpl string, v ...any) ErrorR { _ = "STUB: not implemented"; return *new(ErrorR) }

// Suc success response reply
func Suc(msg string) ErrorR { _ = "STUB: not implemented"; return *new(ErrorR) }

// IsSuc code value check
func (e *errorR) IsSuc() bool {
	_ = "STUB: not implemented"

	// IsFail code value check
	return false
}

func (e *errorR) IsFail() bool {
	_ = "STUB: not implemented"

	// Code value
	return false
}

func (e *errorR) Code() int {
	_ = "STUB: not implemented"

	// Error string
	return 0
}

func (e *errorR) Error() string {
	_ = "STUB: not implemented"

	// String get
	return ""
}

func (e *errorR) String() string { _ = "STUB: not implemented"; return "" }

// GoString get.
func (e *errorR) GoString() string {
	_ = "STUB: not implemented"

	// ErrorM multi error map
	return ""
}

type ErrorM map[string]error

// ErrMap alias of ErrorM
type ErrMap = ErrorM

// Error string
func (e ErrorM) Error() string { _ = "STUB: not implemented"; return "" }

// ErrorOrNil error
func (e ErrorM) ErrorOrNil() error { _ = "STUB: not implemented"; return nil }

// IsEmpty error
func (e ErrorM) IsEmpty() bool {
	_ = "STUB: not implemented"

	// One error
	return false
}

func (e ErrorM) One() error { _ = "STUB: not implemented"; return nil }

// Errors multi error list
type Errors []error

// ErrList alias for Errors
type ErrList = Errors

// Error string
func (es Errors) Error() string { _ = "STUB: not implemented"; return "" }

// ErrorOrNil error
func (es Errors) ErrorOrNil() error { _ = "STUB: not implemented"; return nil }

// IsEmpty error
func (es Errors) IsEmpty() bool { _ = "STUB: not implemented"; return false }

// First error
func (es Errors) First() error { _ = "STUB: not implemented"; return nil }
