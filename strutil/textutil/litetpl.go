package textutil

import (
	"io"
	"text/template"

	"github.com/gookit/goutil/reflects"
	"github.com/gookit/goutil/structs"
)

// LTemplateOptFn lite template option func
type LTemplateOptFn func(opt *LiteTemplateOpt)

// LiteTemplateOpt template options for LiteTemplate
type LiteTemplateOpt struct {
	// func name alias map. eg: {"up_first": "upFirst"}
	nameMp structs.Aliases
	Funcs  template.FuncMap

	Left, Right string

	ParseDef bool
	ParseEnv bool
}

// SetVarFmt custom sets the variable format in template
func (o *LiteTemplateOpt) SetVarFmt(varFmt string) { _ = "STUB: not implemented"; return }

// LiteTemplate implement a simple text template engine.
//
//   - support parse template vars
//   - support access multi-level map field. eg: {{ user.name }}
//   - support parse default value
//   - support parse env vars
//   - support custom pipeline func handle. eg: {{ name | upper }} {{ name | def:guest }}
//
// NOTE: not support control flow, eg: if/else/for/with
type LiteTemplate struct {
	LiteTemplateOpt
	vr VarReplacer
	// template func map. refer the text/template
	//
	// Func allow return 1 or 2 values, if return 2 values, the second value is error.
	fxs map[string]*reflects.FuncX
}

// NewLiteTemplate instance
func NewLiteTemplate(opFns ...LTemplateOptFn) *LiteTemplate { _ = "STUB: not implemented"; return nil }

// with default options

// Init LiteTemplate
func (t *LiteTemplate) Init() { _ = "STUB: not implemented"; return }

// init var replacer

// add built-in funcs

// add custom funcs

func (t *LiteTemplate) initReplacer(vr *VarReplacer) { _ = "STUB: not implemented"; return }

// 排除匹配，防止匹配到类似 "{} adb ddf {var}"

// eg: \{(?s:([^\}]+?))\}
// (?s:...) - 让 "." 匹配换行
// (?s:(.+?)) - 第二个 "?" 非贪婪匹配

// AddFuncs add custom template functions
func (t *LiteTemplate) AddFuncs(fns map[string]any) { _ = "STUB: not implemented"; return }

// RenderString render template string with vars
func (t *LiteTemplate) RenderString(s string, vars map[string]any) string {
	_ = "STUB: not implemented"
	return ""
}

// RenderFile render template file with vars
func (t *LiteTemplate) RenderFile(filePath string, vars map[string]any) (string, error) {
	_ = "STUB: not implemented"
	// read file contents
	return "", nil
}

// RenderWrite render template string with vars, and write result to writer
func (t *LiteTemplate) RenderWrite(wr io.Writer, s string, vars map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *LiteTemplate) renderVars(s string, varMap map[string]string) string {
	_ = "STUB: not implemented"
	return ""
}

// var name or pipe expression.

// compatible default value. eg: {{ name | inhere }}

// clear pipes
// collect pipe functions

// var not found

// check is default func. eg: {{ name | def:guest }}

func (t *LiteTemplate) applyPipes(val any, pipes []string) (string, error) {
	_ = "STUB: not implemented"

	// pipe expr: "trim|upper|substr:1,2"
	// =>
	// pipes: ["trim", "upper", "substr:1,2"]
	return "", nil
}

// has custom args. eg: "substr:1,2"

// call pipe func

func (t *LiteTemplate) isFunc(name string) bool { _ = "STUB: not implemented"; return false }

// check name alias

func (t *LiteTemplate) isDefaultFunc(name string) bool { _ = "STUB: not implemented"; return false }

var stdTpl = NewLiteTemplate()

// RenderFile render template file with vars
func RenderFile(filePath string, vars map[string]any) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// RenderString render str template string or file.
func RenderString(input string, data map[string]any) string { _ = "STUB: not implemented"; return "" }

// RenderWrite render template string with vars, and write result to writer
func RenderWrite(wr io.Writer, s string, vars map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func parseArgStr(argStr string) (ss []any) {
	_ = "STUB: not implemented"
	// no arg
	return nil
}

// one char
