package assert

import (
	"github.com/gookit/goutil/comdef"
)

// isEmpty value check
func isEmpty(v any) bool { _ = "STUB: not implemented"; return false }

func checkEqualArgs(expected, actual any) error { _ = "STUB: not implemented"; return nil }

// formatUnequalValues takes two values of arbitrary types and returns string
// representations appropriate to be presented to the user.
//
// If the values are not of like type, the returned strings will be prefixed
// with the type name, and the value will be enclosed in parentheses similar
// to a type conversion in the Go grammar.
func formatUnequalValues(expected, actual any) (e string, a string) {
	_ = "STUB: not implemented"
	return "", ""
}

// return fmt.Sprintf("%T(%s)", expected, truncatingFormat(expected)),
// 	fmt.Sprintf("%T(%s)", actual, truncatingFormat(actual))

// truncatingFormat formats the data and truncates it if it's too long.
//
// This helps keep formatted error messages lines from exceeding the
// bufio.MaxScanTokenSize max line length that the go testing framework imposes.
func truncatingFormat(data any) string { _ = "STUB: not implemented"; return "" }

// Give us some space the type info too if needed.

func callerInfos() []string { _ = "STUB: not implemented"; return nil }

// The breaks below failed to terminate the loop, and we ran off the
// end of the call stack.

// This is a huge edge case, but it will panic if this is the case

// eg: runtime.goexit

// refers from stretchr/testify/assert
type labeledText struct {
	label   string
	message string
}

func formatLabeledTexts(lts []labeledText) string { _ = "STUB: not implemented"; return "" }

func formatMessage(message string, labelWidth int, buf comdef.StringWriteStringer) string {
	_ = "STUB: not implemented"
	return ""
}

// skip add prefix for first line.

// +3: is len of ":  "
