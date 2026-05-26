package errorx

import (
	"fmt"
	"io"
	"runtime"
)

// stack represents a stack of program counters.
type stack []uintptr

// Format stack trace
func (s *stack) Format(fs fmt.State, verb rune) {
	_ = "STUB: not implemented"

	// case 'v', 's':
	return
}

// StackLen for error
func (s *stack) StackLen() int {
	_ = "STUB: not implemented"

	// WriteTo for error
	return 0
}

func (s *stack) WriteTo(w io.Writer) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// For historical reasons if pc is interpreted as a uintptr
// its value represents the program counter + 1.

// file eg: workspace/godev/gookit/goutil/errorx/errorx_test.go

// f.Name() eg: github.com/gookit/goutil/errorx_test.TestWithPrev()

// String format to string
func (s *stack) String() string { _ = "STUB: not implemented"; return "" }

// StackFrames stack frame list
func (s *stack) StackFrames() *runtime.Frames { _ = "STUB: not implemented"; return nil }

// CallerPC the caller PC value in the stack. it is first frame.
func (s *stack) CallerPC() uintptr { _ = "STUB: not implemented"; return 0 }

// For historical reasons if pc is interpreted as a uintptr
// its value represents the program counter + 1.

/*************************************************************
 * For error caller func
 *************************************************************/

// Func struct
type Func struct {
	*runtime.Func
	pc uintptr
}

// FuncForPC create.
func FuncForPC(pc uintptr) *Func { _ = "STUB: not implemented"; return nil }

// FileLine returns the file name and line number of the source code
func (f *Func) FileLine() (file string, line int) { _ = "STUB: not implemented"; return "", 0 }

// Location simple location info for the func
//
// Returns eg:
//
//	"github.com/gookit/goutil/errorx_test.TestWithPrev(), errorx_test.go:34"
func (f *Func) Location() string { _ = "STUB: not implemented"; return "" }

// String of the func
//
// Returns eg:
//
//	github.com/gookit/goutil/errorx_test.TestWithPrev()
//	  At /path/to/github.com/gookit/goutil/errorx_test.go:34
func (f *Func) String() string { _ = "STUB: not implemented"; return "" }

// MarshalText handle
func (f *Func) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

/*************************************************************
 * helper func for callers stacks
 *************************************************************/

// ErrStackOpt struct
type ErrStackOpt struct {
	SkipDepth  int
	TraceDepth int
}

// default option
var stdOpt = newErrOpt()

// ResetStdOpt config
func ResetStdOpt() { _ = "STUB: not implemented"; return }

func newErrOpt() *ErrStackOpt { _ = "STUB: not implemented"; return nil }

// Config the stdOpt setting
func Config(fns ...func(opt *ErrStackOpt)) { _ = "STUB: not implemented"; return }

// SkipDepth setting
func SkipDepth(skipDepth int) func(opt *ErrStackOpt) { _ = "STUB: not implemented"; return nil }

// TraceDepth setting
func TraceDepth(traceDepth int) func(opt *ErrStackOpt) { _ = "STUB: not implemented"; return nil }

func callersStack(skip, depth int) *stack { _ = "STUB: not implemented"; return nil }
