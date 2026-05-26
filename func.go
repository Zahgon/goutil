package goutil

// Go is a basic promise implementation: it wraps calls a function in a goroutine
// and returns a channel which will later return the function's return value.
func Go(f func() error) error { _ = "STUB: not implemented"; return nil }

// ErrFunc type
type ErrFunc func() error

// CallOn call func on condition is true
func CallOn(cond bool, fn ErrFunc) error { _ = "STUB: not implemented"; return nil }

// IfElseFn call okFunc() on condition is true, else call elseFn()
func IfElseFn(cond bool, okFn, elseFn ErrFunc) error { _ = "STUB: not implemented"; return nil }

// CallOrElse call okFunc() on condition is true, else call elseFn()
func CallOrElse(cond bool, okFn, elseFn ErrFunc) error { _ = "STUB: not implemented"; return nil }

// SafeRun sync run a func. If the func panics, the panic value is returned as an error.
func SafeRun(fn func()) (err error) { _ = "STUB: not implemented"; return nil }

// SafeRunWithError sync run a func with error.
// If the func panics, the panic value is returned as an error.
func SafeRunWithError(fn func() error) (err error) { _ = "STUB: not implemented"; return nil }

// SafeGo async run a func.
// If the func panics, the panic value will be handle by errHandler.
func SafeGo(fn func(), errHandler func(error)) { _ = "STUB: not implemented"; return }

// SafeGoWithError async run a func with error.
// If the func panics, the panic value will be handle by errHandler.
func SafeGoWithError(fn func() error, errHandler func(error)) { _ = "STUB: not implemented"; return }
