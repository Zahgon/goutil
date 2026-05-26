package clog

import (
	"io"
)

type WriteFn func(level string, data map[string]string)

// Printer log printer for console.
type Printer struct {
	Template   string
	TimeFormat string
	// Output log output. default: os.Stdout
	Output io.Writer
	// OnWriteFn log write callback. see formatOutput
	OnWriteFn WriteFn
}

// NewPrinter create a new printer
func NewPrinter(output io.Writer) *Printer { _ = "STUB: not implemented"; return nil }

// Configure the printer
func (p *Printer) Configure(fns ...func(p *Printer)) { _ = "STUB: not implemented"; return }

// SetTemplate sets a custom template
func (p *Printer) SetTemplate(template string) { _ = "STUB: not implemented"; return }

// Print logs a message with the specified level
func (p *Printer) Print(level string, v ...any) { _ = "STUB: not implemented"; return }

// Println logs a message with the specified level. alias of Print
func (p *Printer) Println(level string, v ...any) { _ = "STUB: not implemented"; return }

// Printf logs a message with the specified level and format
func (p *Printer) Printf(level, format string, v ...any) { _ = "STUB: not implemented"; return }

// formatOutput formats the output based on template
func (p *Printer) formatOutput(level, message string) string { _ = "STUB: not implemented"; return "" }

// fire onWrite hook

// Replace placeholders in template

// Debug logs a debug message
func (p *Printer) Debug(v ...any) { _ = "STUB: not implemented"; return }

// Debugf logs a debug message with format
func (p *Printer) Debugf(format string, v ...any) { _ = "STUB: not implemented"; return }

// Info logs an info message
func (p *Printer) Info(v ...any) { _ = "STUB: not implemented"; return }

// Infof logs an info message with format
func (p *Printer) Infof(format string, v ...any) { _ = "STUB: not implemented"; return }

// Warn logs a warning message
func (p *Printer) Warn(v ...any) { _ = "STUB: not implemented"; return }

// Warnf logs a warning message with format
func (p *Printer) Warnf(format string, v ...any) { _ = "STUB: not implemented"; return }

// Error logs an error message
func (p *Printer) Error(v ...any) { _ = "STUB: not implemented"; return }

// Errorf logs an error message with format
func (p *Printer) Errorf(format string, v ...any) { _ = "STUB: not implemented"; return }

// ErrorT logs an error type value
// func (p *Printer) ErrorT(err error) {
// 	if err != nil {
// 		p.Print(ErrorLevel, fmt.Sprint(err))
// 	}
// }

// Fatal logs a fatal message
func (p *Printer) Fatal(v ...any) { _ = "STUB: not implemented"; return }

// Fatalf logs a fatal message with format
func (p *Printer) Fatalf(format string, v ...any) { _ = "STUB: not implemented"; return }

// Trace logs a trace message
func (p *Printer) Trace(v ...any) { _ = "STUB: not implemented"; return }

// Tracef logs a trace message with format
func (p *Printer) Tracef(format string, v ...any) { _ = "STUB: not implemented"; return }

// Success logs a success message
func (p *Printer) Success(v ...any) { _ = "STUB: not implemented"; return }

// Successf logs a success message with format
func (p *Printer) Successf(format string, v ...any) { _ = "STUB: not implemented"; return }
