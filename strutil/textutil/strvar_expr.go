package textutil

import (
	"regexp"

	"github.com/gookit/goutil/maputil"
	"github.com/gookit/goutil/reflects"
)

// type SimpleAnyFunc func(args ...any) any

// StrVarRenderer implements like shell vars renderer
// 简单的实现类似 php, kotlin, shell 插值变量渲染，表达式解析处理。
//
//   - var format: $var_name, ${some_var}, ${top.sub_var}
//   - func call: ${func($var_name, 'const string')}
type StrVarRenderer struct {
	// global variables
	vars map[string]any
	// fallback func for var not exists
	getter FallbackFn
	// funcMap map[string]any TODO use any, add reflect value to rfs
	funcMap map[string]func(...any) any
	// var func map. refer the text/template TODO
	//
	// Func allow return 1 or 2 values, if return 2 values, the second value is error.
	rfs map[string]*reflects.FuncX
}

// NewStrVarRenderer create a new StrVarRenderer
func NewStrVarRenderer() *StrVarRenderer { _ = "STUB: not implemented"; return nil }

// funcMap: make(map[string]any),

// SetVars set variables
func (r *StrVarRenderer) SetVars(vars map[string]any) *StrVarRenderer {
	_ = "STUB: not implemented"
	return nil
}

// SetVar set a variable
func (r *StrVarRenderer) SetVar(name string, value any) *StrVarRenderer {
	_ = "STUB: not implemented"
	return nil
}

// SetFuncMap set function map
func (r *StrVarRenderer) SetFuncMap(funcMap map[string]func(...any) any) *StrVarRenderer {
	_ = "STUB: not implemented"
	return nil
}

// SetFunc set a function
func (r *StrVarRenderer) SetFunc(name string, fn func(...any) any) *StrVarRenderer {
	_ = "STUB: not implemented"
	return nil
}

// SetGetter set variable getter
func (r *StrVarRenderer) SetGetter(getter FallbackFn) *StrVarRenderer {
	_ = "STUB: not implemented"
	return nil
}

var (
	// 处理 $var_name 格式
	// - 允许：$1..$N 这样的变量
	// - 也支持 $@, $* 变量
	reS = regexp.MustCompile(`\$(\w[a-zA-Z0-9_]*|[@|*])`)
	// 处理 ${var_name} ${top.sub} 格式
	reQ = regexp.MustCompile(`\$\{([a-zA-Z][a-zA-Z0-9_.]*)\}`)
	// 处理 ${func(...)} 格式
	reFn = regexp.MustCompile(`\$\{([a-zA-Z][a-zA-Z0-9_]*)\(([^}]*)\)\}`)
)

// Render rendering input string with variables
func (r *StrVarRenderer) Render(input string, vars map[string]any) string {
	_ = "STUB: not implemented"
	return ""
}

// 处理 $var_name 格式

// 处理 ${var.name} 格式

// 处理 ${func(...)} 格式

func (r *StrVarRenderer) handleFuncCalls(input string, re *regexp.Regexp, vars maputil.Map) string {
	_ = "STUB: not implemented"
	return ""
}

// 解析参数 并 调用函数

func (r *StrVarRenderer) parseArgs(argsStr string, data maputil.Map) []any {
	_ = "STUB: not implemented"
	return nil
}

// 简单参数解析，按逗号分割

// is var name

// 常量值：去掉引号

// strconv.Unquote( arg)

func (r *StrVarRenderer) replaceVars(input string, re *regexp.Regexp, data maputil.Map) string {
	_ = "STUB: not implemented"
	return ""
}

// format: ${var_name} 提取变量名

// format: $var_name

// 从 vars map 获取值，支持嵌套变量名

// fallback: 使用 getter fn 获取值
