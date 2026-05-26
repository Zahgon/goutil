package textutil

import (
	"io"
	"strings"
	"text/template"

	"github.com/gookit/goutil"
	"github.com/gookit/goutil/strutil"
)

var builtInFuncs = template.FuncMap{
	// don't escape content
	"raw": func(s string) string {
		return s
	},
	"trim": strings.TrimSpace,
	// lower first char
	"lcFirst": strutil.LowerFirst,
	// upper first char
	"upFirst": strutil.UpperFirst,
	// upper case
	"upper": strings.ToUpper,
	// lower case
	"lower": strings.ToLower,
	// cut sub-string
	"substr": strutil.Substr,
	// default value
	"default": func(v, defVal any) string {
		if goutil.IsEmpty(v) {
			return strutil.SafeString(defVal)
		}
		return strutil.SafeString(v)
	},
	// join strings
	"join": func(ss []string, sep string) string {
		return strings.Join(ss, sep)
	},
}

// TextRenderOpt render text template options
type TextRenderOpt struct {
	// Output use custom output writer
	Output io.Writer
	// Funcs add custom template functions
	Funcs template.FuncMap
}

// RenderOptFn render option func
type RenderOptFn func(opt *TextRenderOpt)

// NewRenderOpt create a new render options
func NewRenderOpt(optFns []RenderOptFn) *TextRenderOpt { _ = "STUB: not implemented"; return nil }

// RenderGoTpl render input text or template file.
func RenderGoTpl(input string, data any, optFns ...RenderOptFn) string {
	_ = "STUB: not implemented"
	return ""
}

// use custom output writer

// return empty string

// use buffer receive rendered content
