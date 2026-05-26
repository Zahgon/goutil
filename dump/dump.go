// Package dump like fmt.Println but more pretty and beautiful print Go values.
package dump

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"time"
)

// These flags define which print caller information
const (
	Fnopos = 1 << iota // no position
	Ffunc
	Ffile
	Ffname
	Fline
)

const defaultSkip = 3

var (
	// valid flag for print caller info
	callerFlags = []int{Ffunc, Ffile, Ffname, Fline}
	// default theme
	defaultTheme = Theme{
		"caller": "magenta",
		"field":  "green", // field name color of the map, struct.
		"value":  "normal",
		// special type
		"msType":  "green", // for keywords map, struct type
		"valTip":  "gray",  // tips comments for string, slice, map len
		"string":  "green",
		"integer": "lightBlue",
	}

	// std dumper
	std = NewDumper(os.Stdout, defaultSkip)
	// no location dumper.
	std2 = NewWithOptions(func(opts *Options) {
		opts.Output = os.Stdout
		opts.ShowFlag = Fnopos
	})

	// some type init
	stringerType = reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
	// time.Time type
	timeType = reflect.TypeOf(time.Time{})
)

// Theme color code/tag map for dump
type Theme map[string]string

func (ct Theme) caller(s string) string  { _ = "STUB: not implemented"; return "" }
func (ct Theme) field(s string) string   { _ = "STUB: not implemented"; return "" }
func (ct Theme) value(s string) string   { _ = "STUB: not implemented"; return "" }
func (ct Theme) msType(s string) string  { _ = "STUB: not implemented"; return "" }
func (ct Theme) valTip(s string) string  { _ = "STUB: not implemented"; return "" }
func (ct Theme) string(s string) string  { _ = "STUB: not implemented"; return "" }
func (ct Theme) integer(s string) string { _ = "STUB: not implemented"; return "" }

// wrap color tag.
func (ct Theme) wrap(key string, s string) string { _ = "STUB: not implemented"; return "" }

// Std dumper
func Std() *Dumper {
	_ = "STUB: not implemented"

	// Reset std dumper
	return nil
}

func Reset() { _ = "STUB: not implemented"; return }

// Config std dumper
func Config(fns ...OptionFunc) { _ = "STUB: not implemented"; return }

// V like fmt.Println, but the output is clearer and more beautiful
func V(vs ...any) {
	_ = "STUB: not implemented"

	// P like fmt.Println, but the output is clearer and more beautiful
	return
}

func P(vs ...any) {
	_ = "STUB: not implemented"

	// Print like fmt.Println, but the output is clearer and more beautiful
	return
}

func Print(vs ...any) {
	_ = "STUB: not implemented"

	// Println like fmt.Println, but the output is clearer and more beautiful
	return
}

func Println(vs ...any) {
	_ = "STUB: not implemented"

	// Fprint like fmt.Println, but the output is clearer and more beautiful
	return
}

func Fprint(w io.Writer, vs ...any) { _ = "STUB: not implemented"; return }

// Std2 dumper
func Std2() *Dumper {
	_ = "STUB: not implemented"

	// Reset2 reset std2 dumper
	return nil
}

func Reset2() { _ = "STUB: not implemented"; return }

// Format like fmt.Println, but the output is clearer and more beautiful
func Format(vs ...any) string { _ = "STUB: not implemented"; return "" }

// NoLoc dump vars data, without location.
func NoLoc(vs ...any) { _ = "STUB: not implemented"; return }

// Clear dump clear data, without location.
func Clear(vs ...any) { _ = "STUB: not implemented"; return }

// is unexported field name on struct
func isUnexported(fieldName string) bool { _ = "STUB: not implemented"; return false }

func isNilOrInvalid(v reflect.Value) bool { _ = "STUB: not implemented"; return false }
