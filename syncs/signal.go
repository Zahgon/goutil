package syncs

import (
	"context"
	"os"
)

// WaitCloseSignals for some huang program.
//
// Usage:
//
//	// do something. eg: start a http server
//
//	syncs.WaitCloseSignals(func(sig os.Signal) {
//		// do something on shutdown. eg: close db, flush logs
//	})
func WaitCloseSignals(onClose func(sig os.Signal), sigCh ...chan os.Signal) {
	_ = "STUB: not implemented"
	return
}

// block until a signal is received.

// SignalHandler returns an actor, i.e. an execute and interrupt func, that
// terminates with SignalError when the process receives one of the provided
// signals, or the parent context is canceled.
//
// from https://github.com/oklog/run/blob/master/actors.go
func SignalHandler(ctx context.Context, signals ...os.Signal) (execute func() error, interrupt func(error)) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SignalError is returned by the signal handler's execute function
// when it terminates due to a received signal.
type SignalError struct {
	Signal os.Signal
}

// Error implements the error interface.
func (e SignalError) Error() string { _ = "STUB: not implemented"; return "" }
