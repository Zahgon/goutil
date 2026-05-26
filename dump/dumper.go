package dump

import (
	"io"
	"reflect"
	"sync"
)

// printValue must keep track of already-printed pointer values to avoid
// infinite recursion. refer the pkg: github.com/kr/pretty
type visit struct {
	v   uintptr
	typ reflect.Type
}

// Dumper struct definition
type Dumper struct {
	*Options
	// locker for visited
	mu sync.RWMutex
	// visited struct records
	visited map[visit]int
	// is value in the slice, map, struct. will not apply indent.
	msValue bool
	// current depth
	curDepth int
	// current indent string bytes
	indentBytes []byte
	// prevDepth, nextDepth int
	// indentStr, indentPrev, lineEnd string
}

// NewDumper create
func NewDumper(out io.Writer, skip int) *Dumper { _ = "STUB: not implemented"; return nil }

// init map

// NewWithOptions create
func NewWithOptions(fns ...OptionFunc) *Dumper { _ = "STUB: not implemented"; return nil }

// WithSkip for dumper
func (d *Dumper) WithSkip(skip int) *Dumper { _ = "STUB: not implemented"; return nil }

// WithoutColor for dumper
func (d *Dumper) WithoutColor() *Dumper { _ = "STUB: not implemented"; return nil }

// WithOptions for dumper
func (d *Dumper) WithOptions(fns ...OptionFunc) *Dumper { _ = "STUB: not implemented"; return nil }

// ResetOptions for dumper
func (d *Dumper) ResetOptions() { _ = "STUB: not implemented"; return }

// Dump vars
func (d *Dumper) Dump(vs ...any) {
	_ = "STUB: not implemented"

	// Print vars. alias of Dump()
	return
}

func (d *Dumper) Print(vs ...any) {
	_ = "STUB: not implemented"

	// Println vars. alias of Dump()
	return
}

func (d *Dumper) Println(vs ...any) {
	_ = "STUB: not implemented"

	// Fprint print vars to io.Writer
	return
}

func (d *Dumper) Fprint(w io.Writer, vs ...any) {
	_ = "STUB: not implemented"
	// backup
	return
}

// restore

// dump go vars
func (d *Dumper) dump(vs ...any) {
	_ = "STUB: not implemented"
	// reset some settings.
	return
}

// clear all theme settings.

// show print position

// get the print position

// print var data

// d.advance(1)

// d.advance(-1)

func (d *Dumper) printCaller(pc uintptr, file string, line int) {
	_ = "STUB: not implemented"
	// eg: github.com/gookit/goutil/dump.ExamplePrint
	return
}

// eg:
// "PRINT AT github.com/gookit/goutil/dump.ExamplePrint(goutil/dump/dump_test.go:23)"
// "PRINT AT github.com/gookit/goutil/dump.ExamplePrint(dump_test.go:23)"
// "PRINT AT github.com/gookit/goutil/dump.ExamplePrint(:23)"

// has a flag

// full func name

// full file path

// only file name
// file name

// Fline

// fallback. eg: "PRINT AT goutil/dump/dump_test.go:23"

// has func, add ")"

func (d *Dumper) advance(step int) {
	_ = "STUB: not implemented"

	// d.nextDepth = d.curDepth + step
	return
}

func (d *Dumper) printOne(v any) { _ = "STUB: not implemented"; return }

// print reflect value

// print reflect value
func (d *Dumper) printRValue(t reflect.Type, v reflect.Value) {
	_ = "STUB: not implemented"
	// if is a ptr, get real type and value
	return
}

// add prefix

// if v.CanAddr() && !d.checkCyclicRef(t, v) {
// 	return // don't print v again
// }

// if !v.CanInterface() {
// 	d.printf("%s,\n", v.String())
// } else {
// 	// v.Interface() will stack overflow on cyclic refer
// 	d.printf("%#v,\n", v.Interface())
// }

// don't print v again

// d.msValue = true

// d.msValue = false

// d.printf("%v,\n", v.Index(i).Interface())

// don't print v again

// up: special handel time.Time struct

// up: if is type alias of time.Time, use a datetime format

// print field name

// print key name

// d.printf("<cyan>%s</>: ", key.String())

// don't print mv again

// print field value

// don't print v again

// d.advance(1)

// case reflect.Ptr:

// don't print v again

func (d *Dumper) checkCyclicRef(t reflect.Type, v reflect.Value) (goon bool) {
	_ = "STUB: not implemented"
	return false
}

// don't print v again

// record visited

func (d *Dumper) rvStringer(rt reflect.Type, rv reflect.Value) string {
	_ = "STUB: not implemented"
	// fmt.Println("Implements fmt.Stringer:", t.Implements(stringerType))
	return ""
}

func (d *Dumper) fmtTimeValue(v reflect.Value) string { _ = "STUB: not implemented"; return "" }

func (d *Dumper) print(v ...any) { _ = "STUB: not implemented"; return }

func (d *Dumper) printf(f string, v ...any) { _ = "STUB: not implemented"; return }

func (d *Dumper) write(indent bool, v ...any) { _ = "STUB: not implemented"; return }

func (d *Dumper) indentPrint(v ...any) { _ = "STUB: not implemented"; return }
