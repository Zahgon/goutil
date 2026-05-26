package stdio

import (
	"io"
)

// Fprint to writer, will ignore error
func Fprint(w io.Writer, a ...any) { _ = "STUB: not implemented"; return }

// Fprintf to writer, will ignore error
func Fprintf(w io.Writer, tpl string, vs ...any) { _ = "STUB: not implemented"; return }

// Fprintln to writer, will ignore error
func Fprintln(w io.Writer, a ...any) { _ = "STUB: not implemented"; return }

// WriteStringTo a writer, will ignore error
func WriteStringTo(w io.Writer, ss ...string) { _ = "STUB: not implemented"; return }
