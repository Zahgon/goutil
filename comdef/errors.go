package comdef

import (
	"errors"
)

// ErrConvType error
var ErrConvType = errors.New("convert value type error")

// Errors multi error list
type Errors []error

// Error string
func (es Errors) Error() string { _ = "STUB: not implemented"; return "" }

// ErrOrNil error
func (es Errors) ErrOrNil() error { _ = "STUB: not implemented"; return nil }

// First error
func (es Errors) First() error { _ = "STUB: not implemented"; return nil }
