// Package termenv provides detect color support of the current terminal.
// And with some utils for terminal env.
package termenv

import (
	"os"
)

var (
	lastErr error
	// debug mode
	debugMode bool
	// support color of current terminal
	supportColor bool

	// value of os color render and display
	//
	// NOTICE:
	// if ENV: NO_COLOR is not empty, will disable color render.
	noColor = os.Getenv("NO_COLOR") != ""

	// the color support level for current terminal
	// needVTP - need enable VTP, only for Windows OS
	colorLevel, needVTP = detectTermColorLevel()
)

// SetDebugMode sets debug mode.
func SetDebugMode(enable bool) {
	_ = "STUB: not implemented"

	// LastErr returns the last error.
	return
}

func LastErr() error { _ = "STUB: not implemented"; return nil }

// reset on get

func debugf(tpl string, v ...any) { _ = "STUB: not implemented"; return }

func setLastErr(err error) { _ = "STUB: not implemented"; return }

// exec: `stty -a 2>&1`
// const (
// mac: speed 9600 baud; 97 rows; 362 columns;
// macSttyMsgPattern = `(\d+)\s+rows;\s*(\d+)\s+columns;`
// linux: speed 38400 baud; rows 97; columns 362; line = 0;
// linuxSttyMsgPattern = `rows\s+(\d+);\s*columns\s+(\d+);`
// )
var terminalWidth, terminalHeight int

// GetTermSize for current console terminal. will first try to get from environment variables COLUMNS and LINES.
func GetTermSize(refresh ...bool) (w int, h int) { _ = "STUB: not implemented"; return 0, 0 }

// 首先尝试从环境变量获取

// cache result

// ReadPassword from console terminal
func ReadPassword(question ...string) string { _ = "STUB: not implemented"; return "" }

// new line
