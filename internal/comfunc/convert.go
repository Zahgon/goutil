package comfunc

import (
	"fmt"

	"github.com/gookit/goutil/comdef"
)

// Bool try to convert type to bool
func Bool(v any) bool { _ = "STUB: not implemented"; return false }

// ToBool try to convert type to bool
func ToBool(v any) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// StrToBool parse string to bool. like strconv.ParseBool()
func StrToBool(s string) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// FormatWithArgs format message with args
//
//   - only one element, format to string
//   - first is format: fmt.Sprintf(firstElem, fmtAndArgs[1:]...)
//   - all is args: return fmt.Sprint(fmtAndArgs...)
func FormatWithArgs(fmtAndArgs []any) string { _ = "STUB: not implemented"; return "" }

// is template string.

// ConvOption convert options
type ConvOption struct {
	// if ture: value is nil, will return convert error;
	// if false(default): value is nil, will convert to zero value
	NilAsFail bool
	// HandlePtr auto convert ptr type(int,float,string) value. eg: *int to int
	// 	- if true: will use real type try convert. default is false
	//	- NOTE: current T type's ptr is default support.
	HandlePtr bool
	// set custom fallback convert func for not supported type.
	UserConvFn comdef.ToStringFunc
}

// ConvOptionFn convert option func
type ConvOptionFn func(opt *ConvOption)

// StrBySprintFn convert any value to string by fmt.Sprint
var StrBySprintFn = func(v any) (string, error) {
	return fmt.Sprint(v), nil
}

// WithUserConvFn set ConvOption.UserConvFn option
func WithUserConvFn(fn comdef.ToStringFunc) ConvOptionFn {
	_ = "STUB: not implemented"
	return *new(ConvOptionFn)
}

// NewConvOption create a new ConvOption
func NewConvOption(optFns ...ConvOptionFn) *ConvOption { _ = "STUB: not implemented"; return nil }

// WithOption set convert option
func (opt *ConvOption) WithOption(optFns ...ConvOptionFn) { _ = "STUB: not implemented"; return }

// ToStringWith try to convert value to string. can with some option func, more see ConvOption.
func ToStringWith(in any, optFns ...ConvOptionFn) (str string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

// same as `rune`
