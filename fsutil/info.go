package fsutil

// DirPath get dir path from filepath, without a last name.
//
//	eg: "/foo/bar/baz.js" => "/foo/bar"
func DirPath(fPath string) string { _ = "STUB: not implemented"; return "" }

// Dir get dir path from filepath, without a last name.
//
//	eg: "/foo/bar/baz.js" => "/foo/bar"
func Dir(fPath string) string { _ = "STUB: not implemented"; return "" }

// PathName get file/dir name from a full path.
//
//	eg: "/foo/bar/baz.js" => "baz.js"
func PathName(fPath string) string { _ = "STUB: not implemented"; return "" }

// PathNoExt get path from full path, without ext.
//
// eg: path/to/main.go => "path/to/main"
func PathNoExt(fPath string) string { _ = "STUB: not implemented"; return "" }

// Name get file/dir name from full path.
//
// eg:
//
//	"path/to/main.go" => "main.go"
//	"/foo/bar/baz" => "baz"
func Name(fPath string) string { _ = "STUB: not implemented"; return "" }

// NameNoExt get file name from a full path, without an ext.
//
// eg: path/to/main.go => "main"
func NameNoExt(fPath string) string { _ = "STUB: not implemented"; return "" }

// FileExt get filename ext. alias of filepath.Ext()
//
// eg: path/to/main.go => ".go"
func FileExt(fPath string) string { _ = "STUB: not implemented"; return "" }

// Extname get filename ext. alias of filepath.Ext()
//
// eg: path/to/main.go => "go"
func Extname(fPath string) string { _ = "STUB: not implemented"; return "" }

// Suffix get filename ext. alias of filepath.Ext()
//
// eg: path/to/main.go => ".go"
func Suffix(fPath string) string { _ = "STUB: not implemented"; return "" }

// Expand will parse first `~` to user home dir path.
func Expand(pathStr string) string { _ = "STUB: not implemented"; return "" }

// ExpandHome will parse first `~` to user home dir path.
func ExpandHome(pathStr string) string { _ = "STUB: not implemented"; return "" }

// ExpandPath will parse `~` to user home dir path.
func ExpandPath(pathStr string) string { _ = "STUB: not implemented"; return "" }

// ResolvePath will parse `~` and ENV var in path
func ResolvePath(pathStr string) string { _ = "STUB: not implemented"; return "" }

// return comfunc.ParseEnvVar()

// SplitPath splits path immediately following the final Separator, separating it into a directory and file name component
func SplitPath(pathStr string) (dir, name string) { _ = "STUB: not implemented"; return "", "" }

// homeDir cache
var _homeDir string

// UserHomeDir is alias of os.UserHomeDir, but ignore error.(by os.UserHomeDir)
func UserHomeDir() string { _ = "STUB: not implemented"; return "" }

// HomeDir get user home dir path.
func HomeDir() string { _ = "STUB: not implemented"; return "" }
