package capp

import (
	"github.com/gookit/goutil/cflag"
)

// CmdOptionFn for one command
type CmdOptionFn func(c *Cmd)

// Cmd for App
type Cmd struct {
	*cflag.CFlags
	init bool
	Name string
	Desc string // desc for command, will sync set to CFlags.Desc
	// Aliases name for command
	Aliases []string
	// OnAdd hook func. fire on add to App
	//  - you can add some cli options or arguments.
	OnAdd func(c *Cmd)
	// Func for run command, will call after options parsed. will sync set to CFlags.Func
	Func func(c *Cmd) error
}

// WrapRunFunc wrap a no-params func as cmd run func
func WrapRunFunc(fn func() error) func(c *Cmd) error { _ = "STUB: not implemented"; return nil }

// NewCmd create a Cmd instance
func NewCmd(name, desc string, runFunc ...func(c *Cmd) error) *Cmd {
	_ = "STUB: not implemented"
	return nil
}

// WithConfigFn config cmd, alias of ConfigCmd()
func (c *Cmd) WithConfigFn(fns ...CmdOptionFn) *Cmd { _ = "STUB: not implemented"; return nil }

// Config the cmd. eg: bing flags
func (c *Cmd) Config(fns ...CmdOptionFn) *Cmd { _ = "STUB: not implemented"; return nil }

// QuickRun parse OS flags and run command, will auto handle error
func (c *Cmd) QuickRun() {
	_ = "STUB: not implemented"

	// MustRun parse flags and run command. alias of MustParse()
	return
}

func (c *Cmd) MustRun(args []string) {
	_ = "STUB: not implemented"

	// MustParse parse flags and run command, will auto handle error
	return
}

func (c *Cmd) MustParse(args []string) { _ = "STUB: not implemented"; return }

// Parse flags and run command func
//
// If args is nil, will parse os.Args
func (c *Cmd) Parse(args []string) error {
	_ = "STUB: not implemented"
	// fix: cmd.xxRun not exec Cmd.Func
	return nil
}

func (c *Cmd) initCmd() { _ = "STUB: not implemented"; return }

// attach handle func

// fix: init c.CFlags on not exist

func (c *Cmd) getDesc() string { _ = "STUB: not implemented"; return "" }

// WithAliases set aliases for command
func WithAliases(aliases ...string) CmdOptionFn {
	_ = "STUB: not implemented"
	return *new(CmdOptionFn)
}
