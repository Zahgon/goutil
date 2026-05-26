// Package finder Provides a simple and convenient filedir lookup function,
// supports filtering, excluding, matching, ignoring, etc.
// and with some commonly built-in matchers.
package finder

import (
	"os"
	"sync"
)

type scanDir struct {
	path  string // dir path to scan
	depth int    // current depth
}

// FileFinder type alias.
type FileFinder = Finder

// Finder struct
type Finder struct {
	// config for finder
	c *Config
	// last error
	err error
	// num - founded fs elem number
	num uint32
	// ch - founded fs elem chan
	ch chan Elem
	// 等待组,跟踪任务完成
	wg sync.WaitGroup
	// dir queue channel, used for concurrency mode
	dirQueue chan scanDir
	// caches - cache found fs elem. if config.CacheResult is true
	caches []Elem
}

// New instance with source dir paths.
func New(dirs []string) *Finder { _ = "STUB: not implemented"; return nil }

// NewFinder new instance with source dir paths.
func NewFinder(dirPaths ...string) *Finder { _ = "STUB: not implemented"; return nil }

// NewWithConfig new instance with config.
func NewWithConfig(c *Config) *Finder { _ = "STUB: not implemented"; return nil }

// NewEmpty new empty Finder instance
func NewEmpty() *Finder { _ = "STUB: not implemented"; return nil }

// EmptyFinder new empty Finder instance. alias of NewEmpty()
func EmptyFinder() *Finder {
	_ = "STUB: not implemented"

	// --------- do finding ---------
	return nil
}

// Find files in given dir paths. will return a channel, you can use it to get the result.
//
// Usage:
//
//	f := NewFinder("/path/to/dir")
//	for el := range f.Find() {
//		fmt.Println(el.Path())
//	}
func (f *Finder) Find() <-chan Elem {
	_ = "STUB: not implemented"

	// Elems find and return founded file Elem. alias of Find()
	return nil
}

func (f *Finder) Elems() <-chan Elem {
	_ = "STUB: not implemented"

	// Results find and return founded file Elem. alias of Find()
	return nil
}

func (f *Finder) Results() <-chan Elem {
	_ = "STUB: not implemented"

	// FindNames find and return founded file/dir names.
	return nil
}

func (f *Finder) FindNames() []string { _ = "STUB: not implemented"; return nil }

// FindPaths find and return founded file/dir paths.
func (f *Finder) FindPaths() []string { _ = "STUB: not implemented"; return nil }

// Each founded file or dir Elem.
func (f *Finder) Each(fn func(el Elem)) {
	_ = "STUB: not implemented"

	// EachElem founded file or dir Elem.
	return
}

func (f *Finder) EachElem(fn func(el Elem)) { _ = "STUB: not implemented"; return }

// EachPath founded file paths.
func (f *Finder) EachPath(fn func(filePath string)) { _ = "STUB: not implemented"; return }

// EachFile each file os.File
func (f *Finder) EachFile(fn func(file *os.File)) { _ = "STUB: not implemented"; return }

// EachStat each file os.FileInfo
func (f *Finder) EachStat(fn func(fi os.FileInfo, filePath string)) {
	_ = "STUB: not implemented"
	return
}

// EachContents handle each found file contents
func (f *Finder) EachContents(fn func(contents, filePath string)) {
	_ = "STUB: not implemented"
	return
}

// prepare for find.
func (f *Finder) prepare() { _ = "STUB: not implemented"; return }

// ensure config

// 创建队列

// Do finding
//
// Usage:
//
//	for el := range f.find() {
//		fmt.Println(el.Path())
//	}
func (f *Finder) find() <-chan Elem {
	_ = "STUB: not implemented"
	// has caches, return it
	return nil
}

// 添加初始任务

// 启动工作goroutine

// 等待所有任务完成并关闭通道

// reset wg
// f.wg = sync.WaitGroup{}

// worker 处理目录的工作goroutine
func (f *Finder) worker(index int) { _ = "STUB: not implemented"; return }

func (f *Finder) safeFindDir(index int, dirPath string, depth int) {
	_ = "STUB: not implemented"
	return
}

// recover error and always call wg.Done()

func (f *Finder) addRootDirs() { _ = "STUB: not implemented"; return }

// add task

// code refer filepath.glob()
func (f *Finder) findDir(dirPath string, depth int) { _ = "STUB: not implemented"; return }

// ignore I/O error

// apply generic filters

// --- dir: apply dir filters

// match ok, send to consumer

// if cfg.FindFlags == FlagDir {
// 	continue // only find sub-dir on ok=false
// }

// find in sub dir. 添加子目录任务

// fix: 创建一个 goroutine 添加子目录任务，不然会造成阻塞

// --- type: file

// apply file filters

// write to consumer

func applyMatchers(el Elem, fls []Matcher) bool { _ = "STUB: not implemented"; return false }

func applyExMatchers(el Elem, fls []Matcher) bool { _ = "STUB: not implemented"; return false }

// Reset filters config setting and results info.
func (f *Finder) Reset() { _ = "STUB: not implemented"; return }

// ResetResult reset result info.
func (f *Finder) ResetResult() { _ = "STUB: not implemented"; return }

// Num get found elem num. only valid after finding.
func (f *Finder) Num() uint {
	_ = "STUB: not implemented"

	// Err get last error
	return 0
}

func (f *Finder) Err() error {
	_ = "STUB: not implemented"

	// Caches get cached results. only valid after finding.
	return nil
}

func (f *Finder) Caches() []Elem {
	_ = "STUB: not implemented"

	// CacheNum get
	return nil
}

func (f *Finder) CacheNum() int { _ = "STUB: not implemented"; return 0 }

// Config get, NOTE: it's a copy of config.
func (f *Finder) Config() Config {
	_ = "STUB: not implemented"

	// String all dir paths
	return *new(Config)
}

func (f *Finder) String() string { _ = "STUB: not implemented"; return "" }

func (f *Finder) debugf(tpl string, vs ...any) { _ = "STUB: not implemented"; return }

func (f *Finder) setError(err error) { _ = "STUB: not implemented"; return }
