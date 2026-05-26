package cmdr

import (
	"github.com/gookit/goutil/errorx"
	"github.com/gookit/goutil/maputil"
	"github.com/gookit/goutil/strutil/textutil"
)

// Task struct
type Task struct {
	err   error
	index int

	// ID for task
	ID  string
	Cmd *Cmd

	// BeforeRun hook
	BeforeRun func(t *Task)
	PrevCond  func(prev *Task) bool
}

// NewTask instance
func NewTask(cmd *Cmd) *Task { _ = "STUB: not implemented"; return nil }

// get task id by cmd.Name
func (t *Task) ensureID(idx int) { _ = "STUB: not implemented"; return }

var rpl = textutil.NewVarReplacer("$").DisableFlatten()

// RunWith command
func (t *Task) RunWith(ctx maputil.Data) error { _ = "STUB: not implemented"; return nil }

// rpl := strutil.NewReplacer(cmdVars)

// Run command
func (t *Task) Run() error { _ = "STUB: not implemented"; return nil }

// Err get
func (t *Task) Err() error {
	_ = "STUB: not implemented"

	// Index get
	return nil
}

func (t *Task) Index() int {
	_ = "STUB: not implemented"

	// Cmdline get
	return 0
}

func (t *Task) Cmdline() string { _ = "STUB: not implemented"; return "" }

// IsSuccess of task
func (t *Task) IsSuccess() bool { _ = "STUB: not implemented"; return false }

// RunnerHookFn func
type RunnerHookFn func(r *Runner, t *Task) bool

// Runner use for batch run multi task commands
type Runner struct {
	prev *Task
	// task name to index
	idMap map[string]int
	tasks []*Task
	// Errs on run tasks, key is Task.ID
	Errs errorx.ErrMap

	// TODO Concurrent run

	// Workdir common workdir
	Workdir string
	// EnvMap will append to task.Cmd on run
	EnvMap map[string]string

	// Params for add custom params
	Params maputil.Map

	// DryRun dry run all commands
	DryRun bool
	// OutToStd stdout and stderr
	OutToStd bool
	// IgnoreErr continue on error
	IgnoreErr bool
	// BeforeRun hooks on each task. return false to skip current task.
	BeforeRun func(r *Runner, t *Task) bool
	// AfterRun hook on each task. return false to stop running.
	AfterRun func(r *Runner, t *Task) bool
}

// NewRunner instance with config func
func NewRunner(fns ...func(rr *Runner)) *Runner { _ = "STUB: not implemented"; return nil }

// WithOutToStd set
func (r *Runner) WithOutToStd() *Runner { _ = "STUB: not implemented"; return nil }

// Add multitask at once
func (r *Runner) Add(tasks ...*Task) *Runner { _ = "STUB: not implemented"; return nil }

// AddTask add one task
func (r *Runner) AddTask(task *Task) *Runner { _ = "STUB: not implemented"; return nil }

// TODO check id repeat

// AddCmd commands
func (r *Runner) AddCmd(cmds ...*Cmd) *Runner { _ = "STUB: not implemented"; return nil }

// GitCmd quick a git command task
func (r *Runner) GitCmd(subCmd string, args ...string) *Runner {
	_ = "STUB: not implemented"
	return nil
}

// CmdWithArgs a command task
func (r *Runner) CmdWithArgs(cmdName string, args ...string) *Runner {
	_ = "STUB: not implemented"
	return nil
}

// CmdWithAnys a command task
func (r *Runner) CmdWithAnys(cmdName string, args ...any) *Runner {
	_ = "STUB: not implemented"
	return nil
}

// AddCmdline as a command task
func (r *Runner) AddCmdline(line string) *Runner { _ = "STUB: not implemented"; return nil }

// Run all tasks
func (r *Runner) Run() error {
	_ = "STUB: not implemented"
	// do run tasks
	return nil
}

// with newline.

// StepRun one command
func (r *Runner) StepRun() error {
	_ = "STUB: not implemented"
	// TODO

	// RunTask command
	return nil
}

func (r *Runner) RunTask(task *Task) (goon bool) { _ = "STUB: not implemented"; return false }

// common workdir

// do running

// not ignore error, stop.

// store prev

// Len of tasks
func (r *Runner) Len() int { _ = "STUB: not implemented"; return 0 }

// Reset instance
func (r *Runner) Reset() *Runner { _ = "STUB: not implemented"; return nil }

// TaskIDs get
func (r *Runner) TaskIDs() []string { _ = "STUB: not implemented"; return nil }

// Prev task instance after running
func (r *Runner) Prev() *Task {
	_ = "STUB: not implemented"

	// Task get by id name
	return nil
}

func (r *Runner) Task(id string) (*Task, error) { _ = "STUB: not implemented"; return nil, nil }
