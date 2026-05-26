// Package assert Provides commonly asserts functions for help write Go testing.
//
// inspired the package: github.com/stretchr/testify/assert
package assert

// TestingT is an interface wrapper around *testing.T
type TestingT interface {
	Helper()
	Name() string
	Error(args ...any)
}

//
// -------------------- render error --------------------
//

var (
	// ShowFullPath on show error trace
	ShowFullPath = true
	// EnableColor on show error trace
	EnableColor = true
	// FailFast fail fast, stop test when first error(will call testing.T.FailNow())
	FailFast = false
)

// DisableColor render
func DisableColor() {
	_ = "STUB: not implemented"

	// HideFullPath render
	return
}

func HideFullPath() { _ = "STUB: not implemented"; return }

// SetFailFast set fail fast
func SetFailFast(enable bool) {
	_ = "STUB: not implemented"

	// fail reports a failure through
	return
}

func fail(t TestingT, failMsg string, fmtAndArgs []any) bool {
	_ = "STUB: not implemented"
	return false
}

// user custom message

// fail fast handle

//
// -------------------- required --------------------
//

// Must assert that the given condition is true. alias of Require()
//
// If it's not, it calls t.FailNow() to terminate the test.
//
// Usage:
//
//	assert.Must(t, assert.True(false))
func Must(t TestingT, condition bool, fmtAndArgs ...any) { _ = "STUB: not implemented"; return }

// Require asserts that the given condition is true.
//
// If it's not, it calls t.FailNow() to terminate the test.
//
// Usage:
//
//	assert.Require(t, assert.True(false))
func Require(t TestingT, condition bool, fmtAndArgs ...any) { _ = "STUB: not implemented"; return }

//
// -------------------- fail --------------------
//

// Fail reports a failure through
func Fail(t TestingT, failMsg string, fmtAndArgs ...any) bool {
	_ = "STUB: not implemented"
	return false
}

type failNower interface {
	FailNow()
}

// FailNow fails test
func FailNow(t TestingT, failMsg string, fmtAndArgs ...any) bool {
	_ = "STUB: not implemented"
	return false
}
