package clog

import (
	"io"
	"os"
)

const (
	SimpleTemplate  = `{emoji} [{level}] | {message}`
	DefaultTemplate = `{time} [{level}] | {emoji} {message}`
)

const (
	DebugLevel   = "debug"
	InfoLevel    = "info"
	WarnLevel    = "warn"
	ErrorLevel   = "error"
	FatalLevel   = "fatal"
	TraceLevel   = "trace"
	SuccessLevel = "success"
)

// LevelColorMap 定义日志级别对应的颜色
var LevelColorMap = map[string]string{
	DebugLevel:   "cyan",
	InfoLevel:    "blue",
	WarnLevel:    "yellow",
	ErrorLevel:   "red",
	FatalLevel:   "red",
	TraceLevel:   "gray",
	SuccessLevel: "green",
}

// LevelEmojiMap 定义日志级别对应的 emoji ⚠️💡
var LevelEmojiMap = map[string]string{
	DebugLevel:   "🐛",
	InfoLevel:    "ℹ️",
	WarnLevel:    "💡",
	ErrorLevel:   "❌",
	FatalLevel:   "🚨",
	TraceLevel:   "🔍",
	SuccessLevel: "🎉",
}

func wrapColor(level, s string) string { _ = "STUB: not implemented"; return "" }

// formatLevel formats the level string
func formatLevel(level string) string { _ = "STUB: not implemented"; return "" }

// getLevelEmoji returns the emoji for the given level
func getLevelEmoji(level string) string { _ = "STUB: not implemented"; return "" }

// default emoji

var std = NewPrinter(os.Stdout)

// Configure the standard log printer
func Configure(optFns ...func(p *Printer)) { _ = "STUB: not implemented"; return }

// SetOnWrite sets a custom write function for the standard logger
func SetOnWrite(fn WriteFn) {
	_ = "STUB: not implemented"

	// SetOutput sets the output for the standard logger
	return
}

func SetOutput(w io.Writer) {
	_ = "STUB: not implemented"

	// SetTemplate sets a custom template for the standard logger
	return
}

func SetTemplate(template string) { _ = "STUB: not implemented"; return }

// Print logs a message with the specified level using the standard logger
func Print(level string, v ...any) { _ = "STUB: not implemented"; return }

// Println logs a message with the specified level using the standard logger
func Println(level string, v ...any) { _ = "STUB: not implemented"; return }

// Printf logs a message with the specified level and format using the standard logger
func Printf(level, format string, v ...any) { _ = "STUB: not implemented"; return }

// Debug logs a debug message using the standard logger
func Debug(v ...any) {
	_ = "STUB: not implemented"

	// Debugf logs a debug message with format using the standard logger
	return
}

func Debugf(format string, v ...any) { _ = "STUB: not implemented"; return }

// Info logs an info message using the standard logger
func Info(v ...any) {
	_ = "STUB: not implemented"

	// Infof logs an info message with format using the standard logger
	return
}

func Infof(format string, v ...any) { _ = "STUB: not implemented"; return }

// Warn logs a warning message using the standard logger
func Warn(v ...any) {
	_ = "STUB: not implemented"

	// Warnf logs a warning message with format using the standard logger
	return
}

func Warnf(format string, v ...any) { _ = "STUB: not implemented"; return }

// Error logs a message using the standard logger
func Error(v ...any) {
	_ = "STUB: not implemented"

	// Errorf logs a message with format using the standard logger
	return
}

func Errorf(format string, v ...any) { _ = "STUB: not implemented"; return }

// Fatal logs a fatal message using the standard logger
func Fatal(v ...any) {
	_ = "STUB: not implemented"

	// Fatalf logs a fatal message with format using the standard logger
	return
}

func Fatalf(format string, v ...any) { _ = "STUB: not implemented"; return }

// Trace logs a trace message using the standard logger
func Trace(v ...any) {
	_ = "STUB: not implemented"

	// Tracef logs a trace message with format using the standard logger
	return
}

func Tracef(format string, v ...any) { _ = "STUB: not implemented"; return }

// Success logs a success message using the standard logger
func Success(v ...any) {
	_ = "STUB: not implemented"

	// Successf logs a success message with format using the standard logger
	return
}

func Successf(format string, v ...any) { _ = "STUB: not implemented"; return }
