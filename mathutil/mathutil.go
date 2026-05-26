// Package mathutil provide math(int, number) util functions. eg: convert, math calc, random
package mathutil

import (
	"github.com/gookit/goutil/comdef"
)

// Mul computes the `a*b` value, rounding the result.
func Mul[T1, T2 comdef.Number](a T1, b T2) float64 { _ = "STUB: not implemented"; return 0 }

// MulF2i computes the float64 type a * b value, rounding the result to an integer.
func MulF2i(a, b float64) int { _ = "STUB: not implemented"; return 0 }

// Div computes the `a/b` value, result uses a round handle.
func Div[T1, T2 comdef.Number](a T1, b T2) float64 { _ = "STUB: not implemented"; return 0 }

// DivInt computes the int type a / b value, rounding the result to an integer.
func DivInt[T comdef.Integer](a, b T) int { _ = "STUB: not implemented"; return 0 }

// DivF2i computes the float64 type a / b value, rounding the result to an integer.
func DivF2i(a, b float64) int { _ = "STUB: not implemented"; return 0 }

// Percent returns a value percentage of the total. eg: 1/100 = 1.0%
func Percent(val, total int) float64 { _ = "STUB: not implemented"; return 0 }

// Range a number range expression, and handle each value. eg: "1-100,123,124"
func Range(expr string, handle func(val int)) error { _ = "STUB: not implemented"; return nil }

// is range, eg: "1-100", "-20-2"

// range number

// Expand a number range expression to int[]. eg: "1-100,123,124"
func Expand(expr string) ([]int, error) { _ = "STUB: not implemented"; return nil, nil }

// is range, eg: "1-100", "-20-2"

// 处理范围格式
// eg: "1-30" -> [1, 30], "-20-2" -> [-20, 2]
func parseIntRange(value string, sepIdx int) (min int, max int, err error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

// swap min and max

// 将 "1-30" 转换为 int 列表
func expandIntRange(value string, sepIdx int) ([]int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
