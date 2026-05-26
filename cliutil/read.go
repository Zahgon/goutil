package cliutil

import (
	"io"
	"os"
)

// the global input output stream
var (
	// Input global input stream
	Input io.Reader = os.Stdin
	// Output global output stream
	Output io.Writer = os.Stdout
)

// ReadInput read user input form Stdin
func ReadInput(question string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// reading

// ReadLine read first line from user input.
//
// Usage:
//
//	in := cliutil.ReadLine("")
//	ans, _ := cliutil.ReadLine("your name?")
func ReadLine(question string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// ReadInt read input value as int
func ReadInt(question string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// ReadFirst read first char
//
// Usage:
//
//	ans, _ := cliutil.ReadFirst("continue?[y/n] ")
func ReadFirst(question string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// ReadFirstByte read first byte char
//
// Usage:
//
//	ans, _ := cliutil.ReadFirstByte("continue?[y/n] ")
func ReadFirstByte(question string) (byte, error) { _ = "STUB: not implemented"; return 0, nil }

// ReadFirstRune read first rune char
func ReadFirstRune(question string) (rune, error) { _ = "STUB: not implemented"; return 0, nil }

// ReadAsBool check user inputted answer is right
//
// Usage:
//
//	ok := ReadAsBool("are you OK? [y/N]", false)
func ReadAsBool(tip string, defVal bool) bool { _ = "STUB: not implemented"; return false }

// Confirm with user input
func Confirm(tip string, defVal ...bool) bool { _ = "STUB: not implemented"; return false }

// InputIsYes answer: yes, y, Yes, Y
func InputIsYes(ans string) bool { _ = "STUB: not implemented"; return false }

// ByteIsYes answer: yes, y, Yes, Y
func ByteIsYes(ans byte) bool { _ = "STUB: not implemented"; return false }

// ReadPassword from console terminal
func ReadPassword(question ...string) string { _ = "STUB: not implemented"; return "" }
