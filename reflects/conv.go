package reflects

import (
	"reflect"
)

// BaseTypeVal convert custom type or intX,uintX,floatX to generic base type.
func BaseTypeVal(v reflect.Value) (value any, err error) {
	_ = "STUB: not implemented"
	return *

	// ToBaseVal convert custom type or intX,uintX,floatX to generic base type.
	//
	//	intX 	    => int64
	//	unitX 	    => uint64
	//	floatX      => float64
	//	string 	    => string
	//
	// returns int64,string,float or error
	new(any), nil
}

func ToBaseVal(v reflect.Value) (value any, err error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// always return int64

// ConvToType convert and create reflect.Value by give reflect.Type
func ConvToType(val any, typ reflect.Type) (rv reflect.Value, err error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), nil
}

// ValueByType create reflect.Value by give reflect.Type
func ValueByType(val any, typ reflect.Type) (rv reflect.Value, err error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), nil
}

// fix: check newRv is valid

// check the same type. like map

// handle kind: string, bool, intX, uintX, floatX

// try the auto convert slice type

// ValueByKind convert and create reflect.Value by give reflect.Kind
func ValueByKind(val any, kind reflect.Kind) (reflect.Value, error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), nil
}

// ConvToKind convert and create reflect.Value by give reflect.Kind
//
// TIPs:
//
//	Only support kind: string, bool, intX, uintX, floatX
func ConvToKind(val any, kind reflect.Kind, fallback ...ConvFunc) (rv reflect.Value, err error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), nil
}

// call fallback func

// ConvSlice make new type slice from old slice, will auto convert element type.
//
// TIPs:
//
//	Only support kind: string, bool, intX, uintX, floatX
func ConvSlice(oldSlRv reflect.Value, newElemTyp reflect.Type) (rv reflect.Value, err error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), nil
}

// do not need convert type

// String convert
func String(rv reflect.Value) string { _ = "STUB: not implemented"; return "" }

// ToString convert
func ToString(rv reflect.Value) (str string, err error) { _ = "STUB: not implemented"; return "", nil }

// ValToString convert handle
func ValToString(rv reflect.Value, defaultAsErr bool) (str string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ToTimeOrDuration convert string to time.Time or time.Duration type
//
// If the target type is not match, return the input string.
func ToTimeOrDuration(str string, typ reflect.Type) (any, error) {
	_ = "STUB: not implemented"
	// datetime, time, duration string should not greater than 64
	return *new(any), nil
}

// time.Time date string
