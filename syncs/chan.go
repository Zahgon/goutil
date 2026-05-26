package syncs

// Go is a basic promise implementation: it wraps calls a function in a goroutine
// and returns a channel which will later return the function's return value.
//
// if panic happen, it will be recovered and return as error
func Go(f func() error) error { _ = "STUB: not implemented"; return nil }

// add recovery handle
