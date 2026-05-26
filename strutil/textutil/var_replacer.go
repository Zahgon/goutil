package textutil

import (
	"regexp"
)

// DefaultVarFormat var template
const DefaultVarFormat = "{{,}}"

// FallbackFn type
type FallbackFn = func(name string) (val string, ok bool)

// VarReplacer struct
type VarReplacer struct {
	init bool

	Left, Right string
	lLen, rLen  int

	varReg *regexp.Regexp
	// flatten sub map in vars. default: true
	//
	// eg: {name: {a: 1, b: 2}} => {name.a: 1, name.b: 2}
	flatSubs bool
	// do parse env value in var-value and tpl var-name. default: false
	parseEnv bool
	// do parse default value. default: false
	//
	// eg: {{ name | inhere }}
	parseDef bool
	// keepMissVars list.
	//
	// default: False - will clear on each replacement
	keepMissVars bool
	// missing vars list
	missVars []string
	// NotFound hook func. on var-name not found
	NotFound FallbackFn
	// RenderFn custom render func
	RenderFn func(s string, vs map[string]string) string
}

// NewVarReplacer instance. default format is: DefaultVarFormat
//
// Usage:
//
//	rpl := NewVarReplacer("{{,}}") // access var: {{ var }}, {{ top.sub }}
//	// or
//	rpl := NewVarReplacer("$") // access var: $var, $top.sub
func NewVarReplacer(format string, opFns ...func(vp *VarReplacer)) *VarReplacer {
	_ = "STUB: not implemented"
	return nil
}

// NewFullReplacer instance. will enable parse env and parse default.
//
// Usage:
//
//	rpl := NewFullReplacer("{{,}}")
func NewFullReplacer(format string) *VarReplacer { _ = "STUB: not implemented"; return nil }

// DisableFlatten on the input vars map
func (r *VarReplacer) DisableFlatten() *VarReplacer { _ = "STUB: not implemented"; return nil }

// KeepMissingVars on the replacement handle
func (r *VarReplacer) KeepMissingVars() *VarReplacer { _ = "STUB: not implemented"; return nil }

// WithParseDefault value on the input template contents. eg: {{ name | inhere }}
func (r *VarReplacer) WithParseDefault() *VarReplacer { _ = "STUB: not implemented"; return nil }

// WithParseEnv on the input vars value
func (r *VarReplacer) WithParseEnv() *VarReplacer { _ = "STUB: not implemented"; return nil }

// OnNotFound var handle func
func (r *VarReplacer) OnNotFound(fn FallbackFn) *VarReplacer { _ = "STUB: not implemented"; return nil }

// WithFormat custom var template
func (r *VarReplacer) WithFormat(format string) *VarReplacer { _ = "STUB: not implemented"; return nil }

// Init var replacer
func (r *VarReplacer) Init() { _ = "STUB: not implemented"; return }

// no right tag. eg: $name, $user.age

// ParseVars parse the text contents and collect vars
func (r *VarReplacer) ParseVars(s string) []string { _ = "STUB: not implemented"; return nil }

// Replace any-map vars in the text contents
func (r *VarReplacer) Replace(s string, tplVars map[string]any) string {
	_ = "STUB: not implemented"
	return ""
}

// Render any-map vars in the text contents
func (r *VarReplacer) Render(s string, tplVars map[string]any) string {
	_ = "STUB: not implemented"
	return ""
}

// ReplaceSMap string-map vars in the text contents
func (r *VarReplacer) ReplaceSMap(s string, varMap map[string]string) string {
	_ = "STUB: not implemented"
	return ""
}

// RenderSimple string-map vars in the text contents. alias of ReplaceSMap()
func (r *VarReplacer) RenderSimple(s string, varMap map[string]string) string {
	_ = "STUB: not implemented"
	return ""
}

// MissVars list
func (r *VarReplacer) MissVars() []string {
	_ = "STUB: not implemented"

	// ResetMissVars list
	return nil
}

func (r *VarReplacer) ResetMissVars() { _ = "STUB: not implemented"; return }

// Replace string-map vars in the text contents
func (r *VarReplacer) doReplace(s string, varMap map[string]string) string {
	_ = "STUB: not implemented"
	return ""
}

// clear on each replacement

// use custom render func

// has custom not found handle func
