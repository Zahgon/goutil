package panics

// IsTrue assert result is true, otherwise will panic
func IsTrue(result bool, fmtAndArgs ...any) { _ = "STUB: not implemented"; return }

// IsFalse assert result is false, otherwise will panic
func IsFalse(result bool, fmtAndArgs ...any) { _ = "STUB: not implemented"; return }

// IsNil assert result is nil, otherwise will panic
func IsNil(result any, fmtAndArgs ...any) { _ = "STUB: not implemented"; return }

// NotNil assert result is non-nil, otherwise will panic
func NotNil(result any, fmtAndArgs ...any) { _ = "STUB: not implemented"; return }

// IsEmpty assert result is empty, otherwise will panic
func IsEmpty(result any, fmtAndArgs ...any) { _ = "STUB: not implemented"; return }

// NotEmpty assert result is empty, otherwise will panic
func NotEmpty(result any, fmtAndArgs ...any) { _ = "STUB: not implemented"; return }

func panicWithMsg(errMsg string, fmtAndArgs []any) { _ = "STUB: not implemented"; return }

// IsEmpty value check
func isEmpty(v any) bool { _ = "STUB: not implemented"; return false }
