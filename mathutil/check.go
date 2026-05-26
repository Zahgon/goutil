package mathutil

import "github.com/gookit/goutil/comdef"

// IsNumeric returns true if the given character is a numeric, otherwise false.
func IsNumeric(c byte) bool { _ = "STUB: not implemented"; return false }

// IsInteger strict check the given value is an integer(intX,uintX), otherwise false.
func IsInteger(val any) bool { _ = "STUB: not implemented"; return false }

// IsFloat returns true if the given character is a float(32/64), otherwise false.
func IsFloat(val any) bool { _ = "STUB: not implemented"; return false }

// Compare any intX,floatX value by given op. returns `first op(=,!=,<,<=,>,>=) second`
//
// Usage:
//
//	mathutil.Compare(2, 3, ">") // false
//	mathutil.Compare(2, 1.3, ">") // true
//	mathutil.Compare(2.2, 1.3, ">") // true
//	mathutil.Compare(2.1, 2, ">") // true
func Compare(first, second any, op string) bool { _ = "STUB: not implemented"; return false }

// as int64

// CompInt compare all intX,uintX type value. returns `first op(=,!=,<,<=,>,>=) second`
func CompInt[T comdef.Xint](first, second T, op string) (ok bool) {
	_ = "STUB: not implemented"
	return false
}

// CompInt64 compare int64 value. returns `first op(=,!=,<,<=,>,>=) second`
func CompInt64(first, second int64, op string) bool { _ = "STUB: not implemented"; return false }

// CompFloat compare float64,float32 value. returns `first op(=,!=,<,<=,>,>=) second`
func CompFloat[T comdef.Float](first, second T, op string) (ok bool) {
	_ = "STUB: not implemented"
	return false
}

// CompValue compare intX,uintX,floatX value. returns `first op(=,!=,<,<=,>,>=) second`
func CompValue[T comdef.Number](first, second T, op string) (ok bool) {
	_ = "STUB: not implemented"
	return false
}

// InRange check if val in int/float range [min, max]
func InRange[T comdef.Number](val, min, max T) bool { _ = "STUB: not implemented"; return false }

// OutRange check if val not in int/float range [min, max]
func OutRange[T comdef.Number](val, min, max T) bool { _ = "STUB: not implemented"; return false }

// InUintRange check if val in unit range [min, max]
func InUintRange[T comdef.Uint](val, min, max T) bool { _ = "STUB: not implemented"; return false }

// InDelta Check whether two floating-point numbers are equal within a specified margin of error
//
// Params:
//
//	want - 期望的浮点数值
//	give - 实际给定的浮点数值
//	delta - 允许的误差范围
func InDelta[T comdef.Float](want, give T, delta float64) bool {
	_ = "STUB: not implemented"
	return false
}

// InDeltaAny Check whether two floating-point numbers are equal within a specified margin of error
func InDeltaAny(want, give any, delta float64) bool { _ = "STUB: not implemented"; return false }
