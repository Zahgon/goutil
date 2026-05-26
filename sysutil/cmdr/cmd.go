package cmdr

import (
	"context"
	"io"
	"os/exec"
)

// Cmd struct
type Cmd struct {
	*exec.Cmd
	// Name of the command
	Name string
	// DryRun setting. if True, not really execute command
	DryRun bool
	// Vars mapping TODO
	Vars map[string]string
	// PrintLine placeholder for print command cline.
	PrintLine string

	// BeforeRun hook
	BeforeRun func(c *Cmd)
	// AfterRun hook
	AfterRun func(c *Cmd, err error)
}

// NewGitCmd instance
func NewGitCmd(subCmd string, args ...string) *Cmd { _ = "STUB: not implemented"; return nil }

// NewCmdline instance
//
// see exec.Command
func NewCmdline(line string) *Cmd { _ = "STUB: not implemented"; return nil }

// NewCmd instance
//
// see exec.Command
func NewCmd(bin string, args ...string) *Cmd { _ = "STUB: not implemented"; return nil }

// CmdWithCtx create new instance with context.
//
// see exec.CommandContext
func CmdWithCtx(ctx context.Context, bin string, args ...string) *Cmd {
	_ = "STUB: not implemented"
	return nil
}

// WrapGoCmd instance
func WrapGoCmd(cmd *exec.Cmd) *Cmd { _ = "STUB: not implemented"; return nil }

// -------------------------------------------------
// config the command
// -------------------------------------------------

// Config the command
func (c *Cmd) Config(fn func(c *Cmd)) *Cmd {
	_ = "STUB: not implemented"

	// WithDryRun on exec command
	return nil
}

func (c *Cmd) WithDryRun(dryRun bool) *Cmd { _ = "STUB: not implemented"; return nil }

// PrintCmdline on exec command
func (c *Cmd) PrintCmdline() *Cmd { _ = "STUB: not implemented"; return nil }

// PrintCmdline2 on exec command
func (c *Cmd) PrintCmdline2() *Cmd { _ = "STUB: not implemented"; return nil }

// OnBefore exec add hook
func (c *Cmd) OnBefore(fn func(c *Cmd)) *Cmd { _ = "STUB: not implemented"; return nil }

// OnAfter exec add hook
func (c *Cmd) OnAfter(fn func(c *Cmd, err error)) *Cmd { _ = "STUB: not implemented"; return nil }

// WithBin name returns the current object
func (c *Cmd) WithBin(name string) *Cmd { _ = "STUB: not implemented"; return nil }

func (c *Cmd) lookPath(name string) { _ = "STUB: not implemented"; return }

// Update cmd.Path even if err is non-nil.
// If err is ErrDot (especially on Windows), lp may include a resolved
// extension (like .exe or .bat) that should be preserved.

// WithGoCmd and returns the current instance.
func (c *Cmd) WithGoCmd(ec *exec.Cmd) *Cmd { _ = "STUB: not implemented"; return nil }

// WithWorkDir returns the current object
func (c *Cmd) WithWorkDir(dir string) *Cmd { _ = "STUB: not implemented"; return nil }

// WorkDirOnNE set workdir on input is not empty
func (c *Cmd) WorkDirOnNE(dir string) *Cmd { _ = "STUB: not implemented"; return nil }

// WithEnvMap override set new ENV for run
func (c *Cmd) WithEnvMap(mp map[string]string) *Cmd { _ = "STUB: not implemented"; return nil }

// AppendEnv to the os ENV for run command
func (c *Cmd) AppendEnv(mp map[string]string) *Cmd {
	_ = "STUB: not implemented"

	// init env data
	return nil
}

// OutputToOS output to OS stdout and error
func (c *Cmd) OutputToOS() *Cmd { _ = "STUB: not implemented"; return nil }

// ToOSStdoutStderr output to OS stdout and error
func (c *Cmd) ToOSStdoutStderr() *Cmd { _ = "STUB: not implemented"; return nil }

// ToOSStdout output to OS stdout
func (c *Cmd) ToOSStdout() *Cmd { _ = "STUB: not implemented"; return nil }

// WithStdin returns the current argument
func (c *Cmd) WithStdin(in io.Reader) *Cmd { _ = "STUB: not implemented"; return nil }

// WithOutput returns the current instance
func (c *Cmd) WithOutput(out, errOut io.Writer) *Cmd { _ = "STUB: not implemented"; return nil }

// WithAnyArgs add args and returns the current object.
func (c *Cmd) WithAnyArgs(args ...any) *Cmd { _ = "STUB: not implemented"; return nil }

// AddArg add args and return the current object
func (c *Cmd) AddArg(args ...string) *Cmd { _ = "STUB: not implemented"; return nil }

// WithArg add args and return the current object. alias of the WithArg()
func (c *Cmd) WithArg(args ...string) *Cmd { _ = "STUB: not implemented"; return nil }

// AddArgf add args and return the current object. alias of the WithArgf()
func (c *Cmd) AddArgf(format string, args ...any) *Cmd { _ = "STUB: not implemented"; return nil }

// WithArgf add arg and return the current object
func (c *Cmd) WithArgf(format string, args ...any) *Cmd { _ = "STUB: not implemented"; return nil }

// ArgIf add arg and return the current object
func (c *Cmd) ArgIf(arg string, exprOk bool) *Cmd { _ = "STUB: not implemented"; return nil }

// WithArgIf add arg and return the current object
func (c *Cmd) WithArgIf(arg string, exprOk bool) *Cmd { _ = "STUB: not implemented"; return nil }

// AddArgs for the git. alias of WithArgs()
func (c *Cmd) AddArgs(args []string) *Cmd { _ = "STUB: not implemented"; return nil }

// WithArgs for the git
func (c *Cmd) WithArgs(args []string) *Cmd { _ = "STUB: not implemented"; return nil }

// WithArgsIf add arg and return the current object
func (c *Cmd) WithArgsIf(args []string, exprOk bool) *Cmd { _ = "STUB: not implemented"; return nil }

// WithVars add vars and return the current object
func (c *Cmd) WithVars(vs map[string]string) *Cmd { _ = "STUB: not implemented"; return nil }

// SetVar add var and return the current object
func (c *Cmd) SetVar(name, val string) *Cmd { _ = "STUB: not implemented"; return nil }

// -------------------------------------------------
// helper command
// -------------------------------------------------

// IDString of the command
func (c *Cmd) IDString() string { _ = "STUB: not implemented"; return "" }

// BinName of the command
func (c *Cmd) BinName() string { _ = "STUB: not implemented"; return "" }

// BinOrPath of the command
func (c *Cmd) BinOrPath() string { _ = "STUB: not implemented"; return "" }

// OnlyArgs of the command, not contains bin name.
func (c *Cmd) OnlyArgs() (ss []string) { _ = "STUB: not implemented"; return nil }

// ResetArgs for command, but will keep bin name.
func (c *Cmd) ResetArgs() { _ = "STUB: not implemented"; return }

// Workdir of the command
func (c *Cmd) Workdir() string {
	_ = "STUB: not implemented"

	// Cmdline to command line
	return ""
}

func (c *Cmd) Cmdline() string { _ = "STUB: not implemented"; return "" }

// RawLine raw command line for print show.
func (c *Cmd) RawLine() string { _ = "STUB: not implemented"; return "" }

// Copy new instance from current command, with new args.
func (c *Cmd) Copy(args ...string) *Cmd {
	_ = "STUB: not implemented"

	// copy bin name.
	return nil
}

// GoCmd get exec.Cmd
func (c *Cmd) GoCmd() *exec.Cmd {
	_ = "STUB: not implemented"

	// -------------------------------------------------
	// run command
	// -------------------------------------------------
	return nil
}

// Success run and return whether success
func (c *Cmd) Success() bool { _ = "STUB: not implemented"; return false }

// HasStdout output setting.
func (c *Cmd) HasStdout() bool { _ = "STUB: not implemented"; return false }

// SafeLines run and return output as lines
func (c *Cmd) SafeLines() []string { _ = "STUB: not implemented"; return nil }

// OutputLines run and return output as lines
func (c *Cmd) OutputLines() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

// SafeOutput run and return output
func (c *Cmd) SafeOutput() string { _ = "STUB: not implemented"; return "" }

// Output run and return output
func (c *Cmd) Output() (string, error) { _ = "STUB: not implemented"; return "", nil }

// AllOutput run and return output, will combine stderr and stdout output
func (c *Cmd) AllOutput() (string, error) {
	_ = "STUB: not implemented"
	return "",

		// CombinedOutput run and return output, will combine stderr and stdout output
		nil
}

func (c *Cmd) CombinedOutput() (string, error) { _ = "STUB: not implemented"; return "", nil }

// MustRun a command. will panic on error
func (c *Cmd) MustRun() { _ = "STUB: not implemented"; return }

// FlushRun runs command and flush output to stdout
func (c *Cmd) FlushRun() error { _ = "STUB: not implemented"; return nil }

// Run runs command
func (c *Cmd) Run() error { _ = "STUB: not implemented"; return nil }

// do running

// if IsWindows() {
// 	return c.Spawn()
// }
// return c.Exec()

// Spawn runs command with spawn(3)
// func (c *Cmd) Spawn() error {
// 	return c.Cmd.Run()
// }
//
// // Exec runs command with exec(3)
// // Note that Windows doesn't support exec(3): http://golang.org/src/pkg/syscall/exec_windows.go#L339
// func (c *Cmd) Exec() error {
// 	binary, err := exec.LookPath(c.Path)
// 	if err != nil {
// 		return &exec.Error{
// 			Name: c.Path,
// 			Err:  errorx.Newf("%s not found in the system", c.Path),
// 		}
// 	}
//
// 	args := []string{binary}
// 	args = append(args, c.Args...)
//
// 	return syscall.Exec(binary, args, os.Environ())
// }
