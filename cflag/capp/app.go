// Package capp provides a simple command line application build.
//
//   - Support add multiple commands
//   - Support add aliases for command
package capp

import (
	"io"
	"os"
	"path/filepath"

	"github.com/gookit/goutil/cflag"
	"github.com/gookit/goutil/maputil"
	"github.com/gookit/goutil/strutil"
)

// App struct
type App struct {
	*cflag.CFlags // save global flags
	// added commands
	names []string
	cmds  map[string]*Cmd
	cmdAs maputil.Aliases

	Name string
	Desc string
	// Version for app
	Version string
	// NameWidth max width for command name
	NameWidth  int
	HelpWriter io.Writer

	// OnAppFlagParsed hook func
	OnAppFlagParsed func(app *App) bool
	// AfterHelpBuild hook
	AfterHelpBuild func(buf *strutil.Buffer)

	// BeforeRun each command hook func
	//  - cmdArgs: input raw args for current command.
	//  - return false to stop run.
	BeforeRun func(c *Cmd, cmdArgs []string) bool
	// AfterRun command hook func
	AfterRun func(c *Cmd, err error)
}

// New App instance
//
// Usage:
//
//	app := capp.New(func(app *cflag.App) {})
//	app.Name = "mycli"
//	app.Version = "0.0.1"
//	app.Desc = "mycli is a command line tool"
func New(fns ...func(app *App)) *App {
	_ = "STUB: not implemented"
	// global flags for app
	return nil
}

// with default version

// NameWidth default value

// NewApp instance. alias of New()
func NewApp(fns ...func(app *App)) *App {
	_ = "STUB: not implemented"

	// NewWith name and desc and option functions
	return nil
}

func NewWith(name, version, desc string, fns ...func(app *App)) *App {
	_ = "STUB: not implemented"
	return nil
}

// WithConfigFn config app
func (a *App) WithConfigFn(fns ...func(app *App)) *App { _ = "STUB: not implemented"; return nil }

// Add command(s) to app. panic if error.
//
// NOTE: command object should create use NewCmd()
//
// Usage:
//
//	app.Add(
//		cflag.NewCmd("cmd1", "desc1"),
//		cflag.NewCmd("cmd2", "desc2"),
//	)
//
// Or:
//
//	app.Add(cflag.NewCmd("cmd1", "desc1"))
//	app.Add(cflag.NewCmd("cmd2", "desc2"))
func (a *App) Add(cmds ...*Cmd) { _ = "STUB: not implemented"; return }

// AddOrErr add command(s) to app.
func (a *App) AddOrErr(cmds ...*Cmd) error { _ = "STUB: not implemented"; return nil }

func (a *App) addCmd(c *Cmd) error { _ = "STUB: not implemented"; return nil }

// add aliases

// attach handle func

//
// region Run with args
// -----------------------------------

// Run app by os.Args
func (a *App) Run() { _ = "STUB: not implemented"; return }

// RunWithArgs run app by input args
func (a *App) RunWithArgs(args []string) error {
	_ = "STUB: not implemented"
	// init for run
	return nil
}

// stop run.

// fire onAppFlagParsed hook

// update args after parse global flags

// first as command name

// parse command flags and execute func.

// fire after run hook

func (a *App) preRun(args []string) (showHelp bool, err error) {
	_ = "STUB: not implemented"
	// prepare
	return false, nil
}

// empty args or help flag

// parse global flags

// ignore help error

// rArgs := a.RemainArgs()
// if len(rArgs) == 0 || isHelp(rArgs[0]) {
// 	return true, nil
// }

func (a *App) init() {
	if a.Name == "" {
		// fix: path.Base not support windows
		a.Name = filepath.Base(os.Args[0])
	}
}

func (a *App) findCmd(name string) (*Cmd, bool) { _ = "STUB: not implemented"; return nil, false }

// resolve alias

//
// region Show help
// -----------------------------------

func (a *App) showHelp() error { _ = "STUB: not implemented"; return nil }

func isHelp(s string) bool { _ = "STUB: not implemented"; return false }
