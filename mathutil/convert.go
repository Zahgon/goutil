package mathutil

import (
	"github.com/gookit/goutil/internal/comfunc"
)

// ConvOption convert options
type ConvOption[T any] struct {
	// if ture: value is nil, will return convert error;
	// if false(default): value is nil, will convert to zero value
	NilAsFail bool
	// HandlePtr auto convert ptr type(int,float,string) value. eg: *int to int
	// 	- if true: will use real type try convert. default is false
	//	- NOTE: current T type's ptr is default support.
	HandlePtr bool
	// StrictMode for convert value. default is false
	//
	// TRUE:
	//  - to int: string, float will return error
	StrictMode bool
	// set custom fallback convert func for not supported type.
	UserConvFn ToTypeFunc[T]
}

// NewConvOption create a new ConvOption
func NewConvOption[T any](optFns ...ConvOptionFn[T]) *ConvOption[T] {
	_ = "STUB: not implemented"
	return nil
}

// WithOption set convert option
func (opt *ConvOption[T]) WithOption(optFns ...ConvOptionFn[T]) { _ = "STUB: not implemented"; return }

// ConvOptionFn convert option func
type ConvOptionFn[T any] func(opt *ConvOption[T])

// WithNilAsFail set ConvOption.NilAsFail option
//
// Example:
//
//	ToIntWithFunc(val, mathutil.WithNilAsFail[int])
func WithNilAsFail[T any](opt *ConvOption[T]) { _ = "STUB: not implemented"; return }

// WithHandlePtr set ConvOption.HandlePtr option
func WithHandlePtr[T any](opt *ConvOption[T]) { _ = "STUB: not implemented"; return }

// WithStrictMode set ConvOption.StrictMode option
func WithStrictMode[T any](opt *ConvOption[T]) { _ = "STUB: not implemented"; return }

// WithUserConvFn set ConvOption.UserConvFn option
func WithUserConvFn[T any](fn ToTypeFunc[T]) ConvOptionFn[T] { _ = "STUB: not implemented"; return nil }

/*************************************************************
 * region Strict to int/uint
 *************************************************************/

// StrictInt check the given value is an integer(intX,uintX), return the int64 value and true if success
func StrictInt(val any) (int64, bool) { _ = "STUB: not implemented"; return 0, false }

// StrictUint strict check value is integer(intX,uintX) and convert to uint64.
func StrictUint(val any) (uint64, bool) { _ = "STUB: not implemented"; return 0, false }

/*************************************************************
 * region convert to float64
 *************************************************************/

// QuietFloat convert value to float64, will ignore error. alias of SafeFloat
func QuietFloat(in any) float64 { _ = "STUB: not implemented"; return 0 }

// SafeFloat convert value to float64, will ignore error
func SafeFloat(in any) float64 { _ = "STUB: not implemented"; return 0 }

// FloatOrPanic convert value to float64, will panic on error
func FloatOrPanic(in any) float64 { _ = "STUB: not implemented"; return 0 }

// MustFloat convert value to float64, will panic on error
func MustFloat(in any) float64 { _ = "STUB: not implemented"; return 0 }

// FloatOrDefault convert value to float64, will return default value on error
func FloatOrDefault(in any, defVal float64) float64 { _ = "STUB: not implemented"; return 0 }

// FloatOr convert value to float64, will return default value on error
func FloatOr(in any, defVal float64) float64 { _ = "STUB: not implemented"; return 0 }

// Float convert value to float64, return error on failed
func Float(in any) (float64, error) {
	_ = "STUB: not implemented"
	return 0,

		// FloatOrErr convert value to float64, return error on failed
		nil
}

func FloatOrErr(in any) (float64, error) {
	_ = "STUB: not implemented"
	return 0,

		// ToFloat convert value to float64, return error on failed
		nil
}

func ToFloat(in any) (float64, error) {
	_ = "STUB: not implemented"
	return 0,

		// ToFloatWith try to convert value to float64. can with some option func, more see ConvOption.
		nil
}

func ToFloatWith(in any, optFns ...ConvOptionFn[float64]) (f64 float64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// default support float64 ptr type

// eg: json.Number

/*************************************************************
 * region intX/floatX to string
 *************************************************************/

// MustString convert intX/floatX value to string, will panic on error
func MustString(val any) string { _ = "STUB: not implemented"; return "" }

// StringOrPanic convert intX/floatX value to string, will panic on error
func StringOrPanic(val any) string { _ = "STUB: not implemented"; return "" }

// StringOrDefault convert intX/floatX value to string, will return default value on error
func StringOrDefault(val any, defVal string) string { _ = "STUB: not implemented"; return "" }

// StringOr convert intX/floatX value to string, will return default value on error
func StringOr(val any, defVal string) string { _ = "STUB: not implemented"; return "" }

// ToString convert intX/floatX value to string, return error on failed
func ToString(val any) (string, error) {
	_ = "STUB: not implemented"
	return "",

		// StringOrErr convert intX/floatX value to string, return error on failed
		nil
}

func StringOrErr(val any) (string, error) {
	_ = "STUB: not implemented"
	return "",

		// QuietString convert intX/floatX value to string, other type convert by fmt.Sprint
		nil
}

func QuietString(val any) string { _ = "STUB: not implemented"; return "" }

// String convert intX/floatX value to string, other type convert by fmt.Sprint
func String(val any) string { _ = "STUB: not implemented"; return "" }

// SafeString convert intX/floatX value to string, other type convert by fmt.Sprint
func SafeString(val any) string { _ = "STUB: not implemented"; return "" }

// TryToString try convert intX/floatX value to string
//
// if defaultAsErr is False, will use fmt.Sprint convert other type
func TryToString(val any, defaultAsErr bool) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ToStringWith try to convert value to string. can with some option func, more see comfunc.ConvOption.
func ToStringWith(in any, optFns ...comfunc.ConvOptionFn) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
