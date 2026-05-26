package finder

// commonly dot file and dirs
var (
	CommonlyDotDirs  = []string{".git", ".idea", ".vscode", ".svn", ".hg"}
	CommonlyDotFiles = []string{".gitignore", ".dockerignore", ".npmignore", ".DS_Store", ".env"}
)

// FindFlag type for find result.
type FindFlag uint8

// String get string name
func (f FindFlag) String() string { _ = "STUB: not implemented"; return "" }

// flags for find result.
const (
	FlagFile FindFlag = iota + 1 // only find files(default)
	FlagDir
	FlagBoth = FlagFile | FlagDir
)

// ToFlag convert flag string to FindFlag
func ToFlag(s string) FindFlag { _ = "STUB: not implemented"; return *new(FindFlag) }

// Config for finder
type Config struct {
	init bool
	// DebugMode enable debug mode
	DebugMode bool `json:"debug_mode"`

	// ScanDirs scan dir paths for find.
	ScanDirs []string `json:"scan_dirs"`
	// FindFlags type for find result. default is FlagFile
	FindFlags FindFlag `json:"find_flags"`
	// MaxDepth for find result. default is 0 - not limit
	MaxDepth int `json:"max_depth"`
	// Concurrency goroutine number for find result. default is 1
	Concurrency int `json:"concurrency"`
	// UseAbsPath use abs path for find result. default is false
	UseAbsPath bool `json:"use_abs_path"`
	// CacheResult cache result for find result. default is false
	CacheResult bool `json:"cache_result"`
	// ExcludeDotDir exclude dot dir. default is true
	ExcludeDotDir bool `json:"exclude_dot_dir"`
	// ExcludeDotFile exclude dot dir. default is false
	ExcludeDotFile bool `json:"exclude_dot_file"`

	// Matchers generic include matchers for file/dir elems
	Matchers []Matcher
	// ExMatchers generic exclude matchers for file/dir elems
	ExMatchers []Matcher
	// DirMatchers include matchers for dir elems
	DirMatchers []Matcher
	// DirExMatchers exclude matchers for dir elems
	DirExMatchers []Matcher
	// FileMatchers include matchers for file elems
	FileMatchers []Matcher
	// FileExMatchers exclude matchers for file elems
	FileExMatchers []Matcher

	// commonly settings for build matchers

	// IncludeDirs include dir name list. eg: {"model"}
	IncludeDirs []string `json:"include_dirs"`
	// IncludeExts include file ext name list. eg: {".go", ".md"}
	IncludeExts []string `json:"include_exts"`
	// IncludeFiles include file name list. eg: {"go.mod"}
	IncludeFiles []string `json:"include_files"`
	// IncludePaths include file/dir path list. eg: {"path/to"}
	IncludePaths []string `json:"include_paths"`
	// IncludeNames include file/dir name list. eg: {"test", "some.go", "*_test.go"}
	IncludeNames []string `json:"include_names"`

	// ExcludeDirs exclude dir name list. eg: {"test"}
	ExcludeDirs []string `json:"exclude_dirs"`
	// ExcludeExts exclude file ext name list. eg: {".go", ".md"}
	ExcludeExts []string `json:"exclude_exts"`
	// ExcludeFiles exclude file name list. eg: {"go.mod"}
	ExcludeFiles []string `json:"exclude_files"`
	// ExcludePaths exclude file/dir path list. eg: {"path/to"}
	ExcludePaths []string `json:"exclude_paths"`
	// ExcludeNames exclude file/dir name list. eg: {"test", "some.go", "*_test.go"}
	ExcludeNames []string `json:"exclude_names"`
}

// NewConfig create a new Config
func NewConfig(dirs ...string) *Config { _ = "STUB: not implemented"; return nil }

// with default setting.

// NewEmptyConfig create a new Config
func NewEmptyConfig() *Config { _ = "STUB: not implemented"; return nil }

// LoadRules load rules and parse to config
//
//   - addOrExclude: true - include, false - exclude
//   - rule NAME allow: ext, name(names), file(files), path, dir(dirs), size, time(mtime)
//
// Rule Format:
//
//	NAME:pattern1,pattern2
//
// Examples:
//
//	ext:.go,.yaml
//	name:*_test.go,go.mod
func (c *Config) LoadRules(addOrExclude bool, rules []string) error {
	_ = "STUB: not implemented"
	return nil
}

// ext:.go,.yaml

// names:*_test.go,go.mod

// size:>=1M,<=10M

// mtime:>=1d,<=10d

// NewFinder create a new Finder by config
func (c *Config) NewFinder() *Finder { _ = "STUB: not implemented"; return nil }

// Init build matchers by config and append to Matchers.
func (c *Config) Init() *Config { _ = "STUB: not implemented"; return nil }

// generic matchers

// dir matchers

// file matchers

//
// --------- config finder by rules ---------
//

// IncludeRule include rules for finder
func (f *Finder) IncludeRule(rules ...string) *Finder { _ = "STUB: not implemented"; return nil }

// IncludeRules include rules for finder
func (f *Finder) IncludeRules(rules []string) *Finder { _ = "STUB: not implemented"; return nil }

// ExcludeRule exclude rules for finder
func (f *Finder) ExcludeRule(rules ...string) *Finder { _ = "STUB: not implemented"; return nil }

// ExcludeRules exclude rules for finder
func (f *Finder) ExcludeRules(rules []string) *Finder { _ = "STUB: not implemented"; return nil }

// WithRules on the finder
//
// Rule Format:
//
//	NAME:pattern1,pattern2
//
// Examples:
//
//	ext:.go,.yaml
//	name:*_test.go,go.mod
func (f *Finder) WithRules(addOrExclude bool, rules []string) *Finder {
	_ = "STUB: not implemented"
	return nil
}

//
// --------- config for finder ---------
//

// WithDebug enable debug mode
func (f *Finder) WithDebug(enable ...bool) *Finder { _ = "STUB: not implemented"; return nil }

// WithConfig on the finder
func (f *Finder) WithConfig(c *Config) *Finder {
	_ = "STUB: not implemented"

	// ConfigFn the finder. alias of WithConfigFn()
	return nil
}

func (f *Finder) ConfigFn(fns ...func(c *Config)) *Finder { _ = "STUB: not implemented"; return nil }

// WithConfigFn the finder
func (f *Finder) WithConfigFn(fns ...func(c *Config)) *Finder {
	_ = "STUB: not implemented"
	return nil
}

// AddScanDirs add source dir for find
func (f *Finder) AddScanDirs(dirPaths []string) *Finder { _ = "STUB: not implemented"; return nil }

// AddScanDir add source dir for find. alias of AddScanDirs()
func (f *Finder) AddScanDir(dirPaths ...string) *Finder { _ = "STUB: not implemented"; return nil }

// AddScan add source dir for find. alias of AddScanDirs()
func (f *Finder) AddScan(dirPaths ...string) *Finder { _ = "STUB: not implemented"; return nil }

// ScanDir add source dir for find. alias of AddScanDirs()
func (f *Finder) ScanDir(dirPaths ...string) *Finder { _ = "STUB: not implemented"; return nil }

// CacheResult cache result for find result.
func (f *Finder) CacheResult(enable ...bool) *Finder { _ = "STUB: not implemented"; return nil }

// WithFlags set find flags.
func (f *Finder) WithFlags(flags FindFlag) *Finder { _ = "STUB: not implemented"; return nil }

// WithStrFlag set find flags by string.
func (f *Finder) WithStrFlag(s string) *Finder { _ = "STUB: not implemented"; return nil }

// TypeFile only find file.
func (f *Finder) TypeFile() *Finder { _ = "STUB: not implemented"; return nil }

// TypeDir only find dir.
func (f *Finder) TypeDir() *Finder { _ = "STUB: not implemented"; return nil }

// OnlyFindDir only find dir.
func (f *Finder) OnlyFindDir() *Finder { _ = "STUB: not implemented"; return nil }

// FileAndDir both find file and dir.
func (f *Finder) FileAndDir() *Finder { _ = "STUB: not implemented"; return nil }

// UseAbsPath use absolute path for find result. alias of WithUseAbsPath()
func (f *Finder) UseAbsPath(enable ...bool) *Finder { _ = "STUB: not implemented"; return nil }

// WithUseAbsPath use absolute path for find result.
func (f *Finder) WithUseAbsPath(enable ...bool) *Finder { _ = "STUB: not implemented"; return nil }

// WithMaxDepth set max depth for find.
func (f *Finder) WithMaxDepth(i int) *Finder { _ = "STUB: not implemented"; return nil }

// WithConcurrency set goroutine number for find.
func (f *Finder) WithConcurrency(i int) *Finder { _ = "STUB: not implemented"; return nil }

// IncludeDir include dir names.
func (f *Finder) IncludeDir(dirs ...string) *Finder { _ = "STUB: not implemented"; return nil }

// WithDirName include dir names. alias of IncludeDir()
func (f *Finder) WithDirName(dirs ...string) *Finder { _ = "STUB: not implemented"; return nil }

// IncludeFile include file names.
func (f *Finder) IncludeFile(files ...string) *Finder { _ = "STUB: not implemented"; return nil }

// WithFileName include file names. alias of IncludeFile()
func (f *Finder) WithFileName(files ...string) *Finder { _ = "STUB: not implemented"; return nil }

// IncludeName include file or dir names.
func (f *Finder) IncludeName(names ...string) *Finder { _ = "STUB: not implemented"; return nil }

// WithNames include file or dir names. alias of IncludeName()
func (f *Finder) WithNames(names []string) *Finder { _ = "STUB: not implemented"; return nil }

// IncludeExt include file exts.
func (f *Finder) IncludeExt(exts ...string) *Finder { _ = "STUB: not implemented"; return nil }

// WithExts include file exts. alias of IncludeExt()
func (f *Finder) WithExts(exts []string) *Finder { _ = "STUB: not implemented"; return nil }

// WithFileExt include file exts. alias of IncludeExt()
func (f *Finder) WithFileExt(exts ...string) *Finder { _ = "STUB: not implemented"; return nil }

// IncludePath include file or dir paths.
func (f *Finder) IncludePath(paths ...string) *Finder { _ = "STUB: not implemented"; return nil }

// WithPaths include file or dir paths. alias of IncludePath()
func (f *Finder) WithPaths(paths []string) *Finder { _ = "STUB: not implemented"; return nil }

// WithSubPath include file or dir paths. alias of IncludePath()
func (f *Finder) WithSubPath(paths ...string) *Finder { _ = "STUB: not implemented"; return nil }

// ExcludeDir exclude dir names.
func (f *Finder) ExcludeDir(dirs ...string) *Finder { _ = "STUB: not implemented"; return nil }

// WithoutDir exclude dir names. alias of ExcludeDir()
func (f *Finder) WithoutDir(dirs ...string) *Finder { _ = "STUB: not implemented"; return nil }

// WithoutNames exclude file or dir names. see Config.ExcludeNames
func (f *Finder) WithoutNames(names []string) *Finder { _ = "STUB: not implemented"; return nil }

// ExcludeName exclude file names. alias of WithoutNames()
func (f *Finder) ExcludeName(names ...string) *Finder { _ = "STUB: not implemented"; return nil }

// ExcludeFile exclude file names.
func (f *Finder) ExcludeFile(files ...string) *Finder { _ = "STUB: not implemented"; return nil }

// WithoutFile exclude file names. alias of ExcludeFile()
func (f *Finder) WithoutFile(files ...string) *Finder { _ = "STUB: not implemented"; return nil }

// ExcludeExt exclude file exts.
//
// eg: ExcludeExt(".go", ".java")
func (f *Finder) ExcludeExt(exts ...string) *Finder { _ = "STUB: not implemented"; return nil }

// WithoutExt exclude file exts. alias of ExcludeExt()
func (f *Finder) WithoutExt(exts ...string) *Finder { _ = "STUB: not implemented"; return nil }

// WithoutExts exclude file exts. alias of ExcludeExt()
func (f *Finder) WithoutExts(exts []string) *Finder { _ = "STUB: not implemented"; return nil }

// ExcludePath exclude file paths.
func (f *Finder) ExcludePath(paths ...string) *Finder { _ = "STUB: not implemented"; return nil }

// WithoutPath exclude file paths. alias of ExcludePath()
func (f *Finder) WithoutPath(paths ...string) *Finder { _ = "STUB: not implemented"; return nil }

// WithoutPaths exclude file paths. alias of ExcludePath()
func (f *Finder) WithoutPaths(paths []string) *Finder { _ = "STUB: not implemented"; return nil }

// ExcludeDotDir exclude dot dir names. eg: ".idea"
func (f *Finder) ExcludeDotDir(exclude ...bool) *Finder { _ = "STUB: not implemented"; return nil }

// WithoutDotDir exclude dot dir names. alias of ExcludeDotDir().
func (f *Finder) WithoutDotDir(exclude ...bool) *Finder { _ = "STUB: not implemented"; return nil }

// NoDotDir exclude dot dir names. alias of ExcludeDotDir().
func (f *Finder) NoDotDir(exclude ...bool) *Finder { _ = "STUB: not implemented"; return nil }

// ExcludeDotFile exclude dot dir names. eg: ".gitignore"
func (f *Finder) ExcludeDotFile(exclude ...bool) *Finder { _ = "STUB: not implemented"; return nil }

// WithoutDotFile exclude dot dir names. alias of ExcludeDotFile().
func (f *Finder) WithoutDotFile(exclude ...bool) *Finder { _ = "STUB: not implemented"; return nil }

// NoDotFile exclude dot dir names. alias of ExcludeDotFile().
func (f *Finder) NoDotFile(exclude ...bool) *Finder { _ = "STUB: not implemented"; return nil }

//
// --------- add matchers to finder ---------
//

// Includes add include match matchers
func (f *Finder) Includes(fls []Matcher) *Finder { _ = "STUB: not implemented"; return nil }

// Collect add include match matchers. alias of Includes()
func (f *Finder) Collect(fls ...Matcher) *Finder { _ = "STUB: not implemented"; return nil }

// Include add include match matchers. alias of Includes()
func (f *Finder) Include(fls ...Matcher) *Finder { _ = "STUB: not implemented"; return nil }

// With add include match matchers. alias of Includes()
func (f *Finder) With(fls ...Matcher) *Finder { _ = "STUB: not implemented"; return nil }

// Adds include match matchers. alias of Includes()
func (f *Finder) Adds(fls []Matcher) *Finder { _ = "STUB: not implemented"; return nil }

// Add include match matchers. alias of Includes()
func (f *Finder) Add(fls ...Matcher) *Finder { _ = "STUB: not implemented"; return nil }

// Excludes add exclude match matchers
func (f *Finder) Excludes(fls []Matcher) *Finder { _ = "STUB: not implemented"; return nil }

// Exclude add exclude match matchers. alias of Excludes()
func (f *Finder) Exclude(fls ...Matcher) *Finder { _ = "STUB: not implemented"; return nil }

// Without add exclude match matchers. alias of Excludes()
func (f *Finder) Without(fls ...Matcher) *Finder { _ = "STUB: not implemented"; return nil }

// Nots add exclude match matchers. alias of Excludes()
func (f *Finder) Nots(fls []Matcher) *Finder { _ = "STUB: not implemented"; return nil }

// Not add exclude match matchers. alias of Excludes()
func (f *Finder) Not(fls ...Matcher) *Finder { _ = "STUB: not implemented"; return nil }

// WithMatchers add include matchers
func (f *Finder) WithMatchers(fls []Matcher) *Finder { _ = "STUB: not implemented"; return nil }

// WithFilter add include matchers
func (f *Finder) WithFilter(fls ...Matcher) *Finder { _ = "STUB: not implemented"; return nil }

// MatchFiles add include file matchers
func (f *Finder) MatchFiles(fls []Matcher) *Finder { _ = "STUB: not implemented"; return nil }

// MatchFile add include file matchers
func (f *Finder) MatchFile(fls ...Matcher) *Finder { _ = "STUB: not implemented"; return nil }

// AddFiles add include file matchers
func (f *Finder) AddFiles(fls []Matcher) *Finder { _ = "STUB: not implemented"; return nil }

// AddFile add include file matchers
func (f *Finder) AddFile(fls ...Matcher) *Finder { _ = "STUB: not implemented"; return nil }

// NotFiles add exclude file matchers
func (f *Finder) NotFiles(fls []Matcher) *Finder { _ = "STUB: not implemented"; return nil }

// NotFile add exclude file matchers
func (f *Finder) NotFile(fls ...Matcher) *Finder { _ = "STUB: not implemented"; return nil }

// MatchDirs add exclude dir matchers
func (f *Finder) MatchDirs(fls []Matcher) *Finder { _ = "STUB: not implemented"; return nil }

// MatchDir add exclude dir matchers
func (f *Finder) MatchDir(fls ...Matcher) *Finder { _ = "STUB: not implemented"; return nil }

// WithDirs add exclude dir matchers
func (f *Finder) WithDirs(fls []Matcher) *Finder { _ = "STUB: not implemented"; return nil }

// WithDir add exclude dir matchers
func (f *Finder) WithDir(fls ...Matcher) *Finder { _ = "STUB: not implemented"; return nil }

// NotDirs add exclude dir matchers
func (f *Finder) NotDirs(fls []Matcher) *Finder { _ = "STUB: not implemented"; return nil }

// NotDir add exclude dir matchers
func (f *Finder) NotDir(fls ...Matcher) *Finder { _ = "STUB: not implemented"; return nil }
