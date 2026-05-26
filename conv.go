package goutil

import (
	"reflect"
)

// Bool convert value to bool
func Bool(v any) bool { _ = "STUB: not implemented"; return false }

// ToBool try to convert type to bool
func ToBool(v any) (bool, error) {
	_ = "STUB: not implemented"
	return false,

		// String func. always converts value to string, will ignore error
		nil
}

func String(v any) string { _ = "STUB: not implemented"; return "" }

// ToString convert value to string, will return error on fail.
func ToString(v any) (string, error) { _ = "STUB: not implemented"; return "", nil }

// Int convert value to int
func Int(v any) int { _ = "STUB: not implemented"; return 0 }

// ToInt try to convert value to int
func ToInt(v any) (int, error) {
	_ = "STUB: not implemented"
	return 0,

		// Int64 convert value to int64
		nil
}

func Int64(v any) int64 { _ = "STUB: not implemented"; return 0 }

// ToInt64 try to convert value to int64
func ToInt64(v any) (int64, error) {
	_ = "STUB: not implemented"
	return 0,

		// Uint convert value to uint
		nil
}

func Uint(v any) uint { _ = "STUB: not implemented"; return 0 }

// ToUint try to convert value to uint
func ToUint(v any) (uint, error) {
	_ = "STUB: not implemented"
	return 0,

		// Uint64 convert value to uint64
		nil
}

func Uint64(v any) uint64 { _ = "STUB: not implemented"; return 0 }

// ToUint64 try to convert value to uint64
func ToUint64(v any) (uint64, error) {
	_ = "STUB: not implemented"
	return 0,

		// BoolString convert bool to string
		nil
}

func BoolString(bl bool) string { _ = "STUB: not implemented"; return "" }

// BaseTypeVal convert custom type or intX,uintX,floatX to generic base type.
//
//	intX 	    => int64
//	unitX 	    => uint64
//	floatX      => float64
//	string 	    => string
//
// returns int64,uint64,string,float or error
func BaseTypeVal(val any) (value any, err error) { _ = "STUB: not implemented"; return *new(any), nil }

// SafeKind convert input any value to given reflect.Kind type.
func SafeKind(val any, kind reflect.Kind) (newVal any) { _ = "STUB: not implemented"; return *new(any) }

// SafeConv convert input any value to given reflect.Kind type.
func SafeConv(val any, kind reflect.Kind) (newVal any) { _ = "STUB: not implemented"; return *new(any) }

// ConvTo convert input any value to given reflect.Kind.
func ConvTo(val any, kind reflect.Kind) (newVal any, err error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// ConvOrDefault convert input any value to given reflect.Kind.
// if fail will return default value.
func ConvOrDefault(val any, kind reflect.Kind, defVal any) any {
	_ = "STUB: not implemented"
	return *new(any)
}

// ToType
// func ToType[T any](val any, kind reflect.Kind, fbFunc func(val any) (T, error)) (newVal T, err error)  {
// 	switch typVal.(type) { // assert ERROR
// 	case string:
// 	}
// }

// ToKind convert input any value to given reflect.Kind type.
//
// TIPs: Only support kind: string, bool, intX, uintX, floatX
//
// Examples:
//
//	val, err := ToKind("123", reflect.Int) // 123
func ToKind(val any, kind reflect.Kind, fbFunc func(val any) (any, error)) (newVal any, err error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}
