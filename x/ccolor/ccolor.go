// Package ccolor is a simple color render library for terminal.
// Its main code is the code that is extracted and simplified from gookit/color,
//
// TIP:
//
//	If you want to render with richer colors, use the https://github.com/gookit/color package.
package ccolor

import (
	"io"
	"log"
	"os"
	"regexp"

	"github.com/gookit/goutil/x/termenv"
)

// color render templates
//
// ESC 操作的表示:
//
//	"\033"(Octal 8进制) = "\x1b"(Hexadecimal 16进制) = 27 (10进制)
const (
	// StartSet chars
	StartSet = "\x1b["
	// ResetSet close all properties.
	ResetSet = "\x1b[0m"
	// SettingTpl string.
	SettingTpl = "\x1b[%sm"
	// FullColorTpl for build color code
	FullColorTpl = "\x1b[%sm%s\x1b[0m"
	// CodeSuffix string for color code.
	CodeSuffix = "[0m"
)

// CodeExpr regex to clear color codes eg "\033[1;36mText\x1b[0m"
const CodeExpr = `\033\[[\d;?]+m`

var (
	// last error
	lastErr error
	// output the default io.Writer message print
	output io.Writer = os.Stdout
	// match color codes
	codeRegex = regexp.MustCompile(CodeExpr)
)

// SetOutput set output writer
func SetOutput(w io.Writer) {
	_ = "STUB: not implemented"

	// LastErr info
	return
}

func LastErr() error { _ = "STUB: not implemented"; return nil }

//
// ---------------- support detect from termenv ----------------
//

// Disable color of current terminal.
func Disable() { _ = "STUB: not implemented"; return }

// Level value of current terminal.
func Level() termenv.ColorLevel { _ = "STUB: not implemented"; return *new(termenv.ColorLevel) }

// IsSupportColor returns true if the terminal supports color.
func IsSupportColor() bool { _ = "STUB: not implemented"; return false }

// IsSupport256Color returns true if the terminal supports 256 colors.
func IsSupport256Color() bool { _ = "STUB: not implemented"; return false }

// IsSupportTrueColor returns true if the terminal supports true color.
func IsSupportTrueColor() bool { _ = "STUB: not implemented"; return false }

//
// ---------------- for testing ----------------
//

// ForceEnableColor setting value. TIP: use for unit testing.
//
// Usage:
//
//	ccolor.ForceEnableColor()
//	defer ccolor.RevertColorSupport()
func ForceEnableColor() { _ = "STUB: not implemented"; return }

// RevertColorSupport value
func RevertColorSupport() { _ = "STUB: not implemented"; return }

//
// ---------------- print with color tag style ----------------
//

// Print parse color tag and print messages
func Print(v ...any) { _ = "STUB: not implemented"; return }

// Printf format and print messages
func Printf(format string, v ...any) { _ = "STUB: not implemented"; return }

// Println messages with new line
func Println(v ...any) { _ = "STUB: not implemented"; return }

// Sprint parse color tags, return rendered string
func Sprint(v ...any) string { _ = "STUB: not implemented"; return "" }

// Sprintf format and return rendered string
func Sprintf(format string, a ...any) string { _ = "STUB: not implemented"; return "" }

// Fprint auto parse color-tag, print rendered messages to the writer
func Fprint(w io.Writer, v ...any) { _ = "STUB: not implemented"; return }

// Fprintf auto parse color-tag, print rendered messages to the writer.
func Fprintf(w io.Writer, format string, v ...any) { _ = "STUB: not implemented"; return }

// Fprintln auto parse color-tag, print rendered messages to the writer
func Fprintln(w io.Writer, v ...any) { _ = "STUB: not implemented"; return }

// Lprint passes colored messages to a log.Logger for printing.
func Lprint(l *log.Logger, v ...any) { _ = "STUB: not implemented"; return }
