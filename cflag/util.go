package cflag

import (
	"flag"
	"regexp"

	"github.com/gookit/goutil/envutil"
)

const (
	// RegGoodName match a good option, argument name
	RegGoodName = `^[a-zA-Z][\w-]*$`
)

var (
	// GoodName good name for option and argument
	goodName = regexp.MustCompile(RegGoodName)
)

// Debug mode
var Debug = envutil.GetBool("CFLAG_DEBUG")

// SetDebug mode
func SetDebug(open bool) {
	_ = "STUB: not implemented"

	// DebugMsg print debug message
	return
}

func DebugMsg(format string, args ...any) { _ = "STUB: not implemented"; return }

// IsGoodName check
func IsGoodName(name string) bool { _ = "STUB: not implemented"; return false }

// IsZeroValue determines whether the string represents the zero
// value for a flag.
//
// from flag.isZeroValue() and more return the second arg for check is string.
func IsZeroValue(opt *flag.Flag, value string) (bool, bool) {
	_ = "STUB: not implemented"
	// Build a zero value of the flag's Value type, and see if the
	// result of calling its String method equals the value passed in.
	// This works unless the Value type is itself an interface type.
	return false, false
}

// AddPrefix for render flag options help
func AddPrefix(name string) string { _ = "STUB: not implemented"; return "" }

// AddPrefixes for render flag options help, name will first add.
func AddPrefixes(name string, shorts []string) string { _ = "STUB: not implemented"; return "" }

// AddPrefixes2 for render flag options help, can custom name add position.
func AddPrefixes2(name string, shorts []string, nameAtEnd bool) string {
	_ = "STUB: not implemented"
	return ""
}

// append shorts

// SplitShortcut string to []string
func SplitShortcut(shortcut string) []string { _ = "STUB: not implemented"; return nil }

// FilterNames for option names, will clear there are: "-+= "
func FilterNames(names []string) []string { _ = "STUB: not implemented"; return nil }

// IsFlagHelpErr check
func IsFlagHelpErr(err error) bool { _ = "STUB: not implemented"; return false }

// regex: "`[\w ]+`"
// regex: "`.+`"
var codeReg = regexp.MustCompile("`" + `.+` + "`")

// WrapColorForCode convert "hello `keywords`" to "hello <mga>keywords</>"
func WrapColorForCode(s string) string { _ = "STUB: not implemented"; return "" }

// ParseStopMark string
const ParseStopMark = "--"

// ReplaceShorts replace shorts to full option. will stop on ParseStopMark
//
// For example:
//
//	eg: '-f' -> '--file'.
//	eg: '-n=tom' -> '--name=tom'.
func ReplaceShorts(args []string, shortsMap map[string]string) []string {
	_ = "STUB: not implemented"
	return nil
}

// is short name, replace to full opt. eg: '-f' -> '--file'

// special, use '=' split value. eg: '-n=tom' -> '--name=tom'
