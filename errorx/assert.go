package errorx

import (
	"github.com/gookit/goutil/comdef"
)

// IsTrue assert result is true, otherwise will return error
func IsTrue(result bool, fmtAndArgs ...any) error { _ = "STUB: not implemented"; return nil }

// IsFalse assert result is false, otherwise will return error
func IsFalse(result bool, fmtAndArgs ...any) error { _ = "STUB: not implemented"; return nil }

// IsIn value should be in the list, otherwise will return error
func IsIn[T comdef.ScalarType](value T, list []T, fmtAndArgs ...any) error {
	_ = "STUB: not implemented"
	return nil
}

// NotIn value should not be in the list, otherwise will return error
func NotIn[T comdef.ScalarType](value T, list []T, fmtAndArgs ...any) error {
	_ = "STUB: not implemented"
	return nil
}

func formatErrMsg(defMsg string, fmtAndArgs []any) string { _ = "STUB: not implemented"; return "" }
