package goinfo

// FullFcName struct.
type FullFcName struct {
	// FullName eg: "github.com/gookit/goutil/x/goinfo.PanicIf"
	FullName string
	pkgPath  string // "github.com/gookit/goutil/x/goinfo"
	pkgName  string // "goinfo"
	funcName string // "PanicIf"
}

// Parse the full func name.
func (ffn *FullFcName) Parse() { _ = "STUB: not implemented"; return }

// spilt get pkg and func name

// PkgPath string get. eg: github.com/gookit/goutil/x/goinfo
func (ffn *FullFcName) PkgPath() string { _ = "STUB: not implemented"; return "" }

// PkgName string get. eg: goinfo
func (ffn *FullFcName) PkgName() string { _ = "STUB: not implemented"; return "" }

// FuncName get short func name. eg: PanicIf
func (ffn *FullFcName) FuncName() string { _ = "STUB: not implemented"; return "" }

// String get full func name string, pkg path and func name.
func (ffn *FullFcName) String() string { _ = "STUB: not implemented"; return "" }

// FuncName get full func name, contains pkg path.
//
// eg:
//
//	// OUTPUT: github.com/gookit/goutil/x/goinfo.PkgName
//	goinfo.FuncName(goinfo.PkgName)
func FuncName(fn any) string { _ = "STUB: not implemented"; return "" }

// CutFuncName get pkg path and short func name
// eg:
//
//	"github.com/gookit/goutil/x/goinfo.FuncName" => [github.com/gookit/goutil/x/goinfo, FuncName]
func CutFuncName(fullFcName string) (pkgPath, shortFnName string) {
	_ = "STUB: not implemented"
	return "", ""
}

// PkgName get current package name
//
// Usage:
//
//	fullFcName := goinfo.FuncName(fn)
//	pgkName := goinfo.PkgName(fullFcName)
func PkgName(fullFcName string) string { _ = "STUB: not implemented"; return "" }

// GoodFuncName reports whether the function name is a valid identifier.
func GoodFuncName(name string) bool { _ = "STUB: not implemented"; return false }
