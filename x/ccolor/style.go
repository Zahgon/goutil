package ccolor

import (
	"io"
)

// String color style string. TODO
// eg:
//
//	s := String("red,bold")
//	s.Println("some text message")
type String string

// Style for color render.
type Style struct {
	Fg   Color
	Bg   Color
	Opts []Color
}

// NewStyle fg, bg and options
func NewStyle(fg Color, bg Color, opts ...Color) *Style { _ = "STUB: not implemented"; return nil }

// Print like fmt.Print, but with color
func (s *Style) Print(v ...any) { _ = "STUB: not implemented"; return }

// Println like fmt.Println, but with color
func (s *Style) Println(v ...any) { _ = "STUB: not implemented"; return }

// Printf render and print text
func (s *Style) Printf(format string, v ...any) { _ = "STUB: not implemented"; return }

// Sprint like fmt.Sprint, but with color
func (s *Style) Sprint(v ...any) string { _ = "STUB: not implemented"; return "" }

// Sprintln like fmt.Sprintln, but with color
func (s *Style) Sprintln(v ...any) string { _ = "STUB: not implemented"; return "" }

// Sprintf format and render message.
func (s *Style) Sprintf(format string, v ...any) string { _ = "STUB: not implemented"; return "" }

// Fprint like fmt.Fprint, but with color
func (s *Style) Fprint(w io.Writer, v ...any) { _ = "STUB: not implemented"; return }

// String convert style setting to color code string.
func (s *Style) String() string { _ = "STUB: not implemented"; return "" }

var (
	// Info color style
	Info = &Style{Fg: FgGreen}
	// Warn color style
	Warn = &Style{Fg: FgYellow}
	// Error color style
	Error = NewStyle(FgLightWhite, BgRed)
	// Debug color style
	Debug = &Style{Fg: FgCyan}
	// Success color style
	Success = &Style{Fg: FgGreen, Opts: []Color{OpBold}}
)
