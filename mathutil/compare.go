package mathutil

import (
	"github.com/gookit/goutil/comdef"
)

// Min compare two value and return max value
func Min[T comdef.Number](x, y T) T { _ = "STUB: not implemented"; return *new(T) }

// Max compare two value and return max value
func Max[T comdef.Number](x, y T) T { _ = "STUB: not implemented"; return *new(T) }

// SwapMin compare and always return [min, max] value
func SwapMin[T comdef.Number](x, y T) (T, T) { _ = "STUB: not implemented"; return *new(T), *new(T) }

// SwapMax compare and always return [max, min] value
func SwapMax[T comdef.Number](x, y T) (T, T) { _ = "STUB: not implemented"; return *new(T), *new(T) }

// MaxInt compare and return max value
func MaxInt(x, y int) int { _ = "STUB: not implemented"; return 0 }

// SwapMaxInt compare and return max, min value
func SwapMaxInt(x, y int) (int, int) { _ = "STUB: not implemented"; return 0, 0 }

// MaxI64 compare and return max value
func MaxI64(x, y int64) int64 { _ = "STUB: not implemented"; return 0 }

// SwapMaxI64 compare and return max, min value
func SwapMaxI64(x, y int64) (int64, int64) { _ = "STUB: not implemented"; return 0, 0 }

// MaxFloat compare and return max value
func MaxFloat(x, y float64) float64 { _ = "STUB: not implemented"; return 0 }

// OrElse return default value on val is zero, else return val
func OrElse[T comdef.Number](val, defVal T) T { _ = "STUB: not implemented"; return *new(T) }

// ZeroOr return default value on val is zero, else return val
func ZeroOr[T comdef.Number](val, defVal T) T { _ = "STUB: not implemented"; return *new(T) }

// LessOr return val on val < max, else return default value.
//
// Example:
//
//	LessOr(11, 10, 1) // 1
//	LessOr(2, 10, 1) // 2
//	LessOr(10, 10, 1) // 1
func LessOr[T comdef.Number](val, max, devVal T) T { _ = "STUB: not implemented"; return *new(T) }

// LteOr return val on val <= max, else return default value.
//
// Example:
//
//	LteOr(11, 10, 1) // 11
//	LteOr(2, 10, 1) // 2
//	LteOr(10, 10, 1) // 10
func LteOr[T comdef.Number](val, max, devVal T) T { _ = "STUB: not implemented"; return *new(T) }

// GreaterOr return val on val > max, else return default value.
//
// Example:
//
//	GreaterOr(23, 0, 2) // 23
//	GreaterOr(0, 0, 2) // 2
func GreaterOr[T comdef.Number](val, min, defVal T) T { _ = "STUB: not implemented"; return *new(T) }

// GteOr return val on val >= max, else return default value.
//
// Example:
//
//	GteOr(23, 0, 2) // 23
//	GteOr(0, 0, 2) // 0
func GteOr[T comdef.Number](val, min, defVal T) T { _ = "STUB: not implemented"; return *new(T) }
