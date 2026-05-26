// Package cflag Wraps and extends go `flag.FlagSet` to build simple command line applications
//
//   - Support auto render a pretty help panel
//   - Allow to add shortcuts for flag option
//   - Allow binding named arguments
//   - Allow set required for argument or option
//   - Allow set validator for argument or option
package cflag

import (
	"errors"
	"flag"
	"strings"

	"github.com/gookit/goutil/strutil"
)

// CFlags wrap and extends the go flag.FlagSet
//
// eg:
//
//	// Can be set required and shorts on desc:
//	// format1: desc;required
//	cmd.IntVar(&age, "age", 0, "your age;true")
//	// format2: desc;required;shorts
//	cmd.IntVar(&age, "age", 0, "your age;true;a")
type CFlags struct {
	*flag.FlagSet
	prepared bool
	// bound options.
	bindOpts map[string]*FlagOpt
	// shortcuts map for options. eg: n -> name
	shortcuts map[string]string

	// argWidth max width value
	argWidth int
	// bound arguments.
	bindArgs []*FlagArg
	// the argument name to index map.
	argNames map[string]int
	// remainArgs after binding args
	remainArgs []string

	// Desc command description
	Desc  string
	Usage string // command usage contents
	// Version command version number
	Version string
	// Example command usage examples
	Example string
	// LongHelp custom help
	LongHelp string
	// HelpOnEmptyArgs show help when not input args. default: false
	HelpOnEmptyArgs bool
	// HelpFunc custom help render func
	HelpFunc func(c *CFlags)

	// AfterFlagParse handler. return false to stop continue run Func.
	//  - fire on flag options parsed, before binding arguments
	AfterFlagParse func(c *CFlags) bool
	// BeforeRun handler for the command. return false to stop run Func.
	//
	// TIP: You can do some processing or intercept command execution before running.
	BeforeRun func(c *CFlags) bool
	// Func handler for the command
	Func func(c *CFlags) error
}

// New create new instance.
//
// Usage:
//
//	cmd := cflag.New(func(c *cflag.CFlags) {
//		c.Version = "0.1.2"
//		c.Desc = "this is my cli tool"
//	})
//
//	// binding opts and args
//
//	cmd.Parse(nil)
func New(fns ...func(c *CFlags)) *CFlags { _ = "STUB: not implemented"; return nil }

// NewWith create new instance.
//
// Usage:
//
//	cmd := cflag.NewWith("0.1.2", "this is my cli tool")
//
//	// binding opts and args
//
//	cmd.Parse(nil)
func NewWith(name, version, desc string, fns ...func(c *CFlags)) *CFlags {
	_ = "STUB: not implemented"
	return nil
}

// NewEmpty instance.
func NewEmpty(fns ...func(c *CFlags)) *CFlags { _ = "STUB: not implemented"; return nil }

/*************************************************************
 * config command flags
 *************************************************************/

// WithDesc for command
func WithDesc(desc string) func(c *CFlags) { _ = "STUB: not implemented"; return nil }

// WithVersion for command
func WithVersion(version string) func(c *CFlags) { _ = "STUB: not implemented"; return nil }

// WithConfigFn for command
func (c *CFlags) WithConfigFn(fns ...func(c *CFlags)) *CFlags {
	_ = "STUB: not implemented"
	return nil
}

// AddValidator for a flag option
func (c *CFlags) AddValidator(name string, fn OptCheckFn) { _ = "STUB: not implemented"; return }

// ConfigOpt for a flag option
func (c *CFlags) ConfigOpt(name string, fn func(opt *FlagOpt)) { _ = "STUB: not implemented"; return }

// init on not exist

// AddShortcuts for option flag
func (c *CFlags) AddShortcuts(name string, shorts ...string) { _ = "STUB: not implemented"; return }

// addShortcuts for option flag
func (c *CFlags) addShortcuts(name string, shorts []string) { _ = "STUB: not implemented"; return }

// AddArg binding for command, by position
//
//	c.AddArg("name", "desc ...")
//	c.AddArg("name", "desc ...", required: bool)
//	c.AddArg("name", "desc ...", required: bool, default: any)
//	c.AddArg("name", "desc ...", required: bool, default: any, isArray: bool)
func (c *CFlags) AddArg(name, desc string, requireDefaultArrayed ...any) {
	_ = "STUB: not implemented"
	return
}

// 2th set default value

// 3th set isArrayed

// required

// BindArg for command
func (c *CFlags) BindArg(arg *FlagArg) { _ = "STUB: not implemented"; return }

// check arg info

// register

/*************************************************************
 * parse command flags
 *************************************************************/

// QuickRun parse OS flags and run command, will auto handle error
func (c *CFlags) QuickRun() {
	_ = "STUB: not implemented"

	// MustRun parse flags and run command. alias of MustParse()
	return
}

func (c *CFlags) MustRun(args []string) {
	_ = "STUB: not implemented"

	// MustParse parse flags and run command, will auto handle error
	return
}

func (c *CFlags) MustParse(args []string) { _ = "STUB: not implemented"; return }

// ErrStopRun error
var ErrStopRun = errors.New("stop run")

// Parse flags and run command func.
//
// If args is nil, will parse os.Args
//
//   - will auto handle display help on with --help, -h
func (c *CFlags) Parse(args []string) error { _ = "STUB: not implemented"; return nil }

// prepare

// show help when no args

// do parsing(will handle show help)

// ignore help error

// call before run

// call func

// Prepare for parse. (internal use)
func (c *CFlags) Prepare() error { _ = "STUB: not implemented"; return nil }

// dont use flag output.

// parse flag usage string

// custom something
// c.FlagSet.Usage = c.ShowHelp

// do parse flag.Usage string.
func (c *CFlags) parseFlagUsage(name, usage string) string { _ = "STUB: not implemented"; return "" }

// FORMAT: desc;required;shorts

// required

// shortcuts

// DoParse parse options and validate, collect args. (internal use)
func (c *CFlags) DoParse(args []string) error { _ = "STUB: not implemented"; return nil }

// do parsing

// check option values

// fire hook: after flag parse

// check bind option flags
func (c *CFlags) checkBindOpts() error { _ = "STUB: not implemented"; return nil }

// call validator

// desc for command
func (c *CFlags) bindParsedArgs() error { _ = "STUB: not implemented"; return nil }

// max index

// fix: need reset value to default, on repeat parse

// clean it

// collect remain args

// all args

// Arg get by bind name
//
//	val := c.Arg("name").String()
func (c *CFlags) Arg(name string) *FlagArg { _ = "STUB: not implemented"; return nil }

// RemainArgs get
func (c *CFlags) RemainArgs() []string {
	_ = "STUB: not implemented"

	// Name for command
	return nil
}

func (c *CFlags) Name() string { _ = "STUB: not implemented"; return "" }

// BinFile path for command
func (c *CFlags) BinFile() string { _ = "STUB: not implemented"; return "" }

/*************************************************************
 * render command help
 *************************************************************/

// desc for command
func (c *CFlags) helpDesc() string { _ = "STUB: not implemented"; return "" }

// ShowHelp for command
func (c *CFlags) ShowHelp() {
	_ = "STUB: not implemented"

	// show help for command
	return
}

func (c *CFlags) showHelp(err error) { _ = "STUB: not implemented"; return }

// render options help

var optionIndentSpace = "\n" + strings.Repeat("    ", 7)

// RenderOptionsHelp prints, to standard error unless configured otherwise, the
// default values of all defined command-line flags in the set. See the
// documentation for the global function PrintDefaults for more information.
//
// from flag.PrintDefaults
func (c *CFlags) RenderOptionsHelp(buf *strutil.Buffer) { _ = "STUB: not implemented"; return }

// Boolean flags of one ASCII letter are so common we
// treat them specially, putting their usage on the same line.
// -9: <info></>

// Four spaces before the tab triggers good alignment
// for both 4- and 8-space tab stops.

// put quotes on the string value

// arrayed, repeatable
