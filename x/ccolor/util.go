package ccolor

import (
	"io"
)

// ColorsToCode convert colors to code. return like "32;45;3"
func ColorsToCode(colors ...Color) string { _ = "STUB: not implemented"; return "" }

func shouldCleanColor() bool { _ = "STUB: not implemented"; return false }

/*************************************************************
 * render color code
 *************************************************************/

// RenderCode render message by color code.
//
// Usage:
//
//	msg := RenderCode("3;32;45", "some", "message")
func RenderCode(code string, args ...any) string { _ = "STUB: not implemented"; return "" }

// disabled OR not support color

// return fmt.Sprintf(FullColorTpl, code, message)

// RenderString render a string with color code.
//
// Usage:
//
//	msg := RenderString("3;32;45", "a message")
func RenderString(code string, str string) string { _ = "STUB: not implemented"; return "" }

// disabled OR not support color

// return fmt.Sprintf(FullColorTpl, code, str)

// RenderWithSpaces Render code with spaces.
// If the number of args is > 1, a space will be added between the args
func RenderWithSpaces(code string, args ...any) string { _ = "STUB: not implemented"; return "" }

// disabled OR not support color

// ClearCode clear color codes.
//
// eg:
//
//	"\033[36;1mText\x1b[0m" -> "Text"
func ClearCode(str string) string { _ = "STUB: not implemented"; return "" }

/*************************************************************
 * helper methods for print
 *************************************************************/

// new implementation, support render full color code on pwsh.exe, cmd.exe
func doPrint(code, str string) { _ = "STUB: not implemented"; return }

func doPrintTo(w io.Writer, code, str string) { _ = "STUB: not implemented"; return }

// new implementation, support render full color code on pwsh.exe, cmd.exe
func doPrintln(code string, args []any) { _ = "STUB: not implemented"; return }

// use Println, will add spaces for each arg
func formatLikePrintln(args []any) (message string) { _ = "STUB: not implemented"; return "" }

// clear last "\n"
