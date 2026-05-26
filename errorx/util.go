package errorx

// E new a raw go error. alias of errors.New()
func E(msg string) error { _ = "STUB: not implemented"; return nil }

// Err new a raw go error. alias of errors.New()
func Err(msg string) error { _ = "STUB: not implemented"; return nil }

// Raw new a raw go error. alias of errors.New()
func Raw(msg string) error { _ = "STUB: not implemented"; return nil }

// Ef new a raw go error. alias of fmt.Errorf
func Ef(tpl string, vars ...any) error { _ = "STUB: not implemented"; return nil }

// Errf new a raw go error. alias of fmt.Errorf
func Errf(tpl string, vars ...any) error { _ = "STUB: not implemented"; return nil }

// Rf new a raw go error. alias of fmt.Errorf
func Rf(tpl string, vs ...any) error { _ = "STUB: not implemented"; return nil }

// Rawf new a raw go error. alias of fmt.Errorf
func Rawf(tpl string, vs ...any) error { _ = "STUB: not implemented"; return nil }

/*************************************************************
 * helper func for error
 *************************************************************/

// Cause returns the first cause error by call err.Cause().
// Otherwise, will returns current error.
func Cause(err error) error { _ = "STUB: not implemented"; return nil }

// Unwrap returns previous error by call err.Unwrap().
// Otherwise, will returns nil.
func Unwrap(err error) error { _ = "STUB: not implemented"; return nil }

// Previous alias of Unwrap()
func Previous(err error) error {
	_ = "STUB: not implemented"

	// IsErrorX check
	return nil
}

func IsErrorX(err error) (ok bool) { _ = "STUB: not implemented"; return false }

// ToErrorX convert check. like errors.As()
func ToErrorX(err error) (ex *ErrorX, ok bool) { _ = "STUB: not implemented"; return nil, false }

// MustEX convert error to *ErrorX, panic if err check failed.
func MustEX(err error) *ErrorX { _ = "STUB: not implemented"; return nil }

// Has contains target error, or err is eq target.
// alias of errors.Is()
func Has(err, target error) bool { _ = "STUB: not implemented"; return false }

// Is alias of errors.Is()
func Is(err, target error) bool { _ = "STUB: not implemented"; return false }

// To try convert err to target, returns is result.
//
// NOTICE: target must be ptr and not nil. alias of errors.As()
//
// Usage:
//
//	var ex *errorx.ErrorX
//	err := doSomething()
//	if errorx.To(err, &ex) {
//		fmt.Println(ex.GoString())
//	}
func To(err error, target any) bool { _ = "STUB: not implemented"; return false }

// As same of the To(), alias of errors.As()
//
// NOTICE: target must be ptr and not nil
func As(err error, target any) bool { _ = "STUB: not implemented"; return false }
