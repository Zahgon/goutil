package goutil

import (
	"context"

	"github.com/gookit/goutil/structs"
	"github.com/gookit/goutil/syncs"
)

// ErrGroup is a collection of goroutines working on subtasks that
// are part of the same overall task.
type ErrGroup = syncs.ErrGroup

// NewCtxErrGroup instance. use for batch run tasks, can with context.
//
// Deprecated: use syncs.NewCtxErrGroup instead
func NewCtxErrGroup(ctx context.Context, limit ...int) (*ErrGroup, context.Context) {
	_ = "STUB: not implemented"
	return nil, *new(context.Context)
}

// NewErrGroup instance. use for batch run tasks
//
// Deprecated: use syncs.NewErrGroup instead
func NewErrGroup(limit ...int) *ErrGroup { _ = "STUB: not implemented"; return nil }

// RunFn func
type RunFn func(ctx *structs.Data) error

// QuickRun struct
type QuickRun struct {
	ctx *structs.Data
	// err error
	fns []RunFn
}

// NewQuickRun instance
func NewQuickRun() *QuickRun { _ = "STUB: not implemented"; return nil }

// Add func for run
func (p *QuickRun) Add(fns ...RunFn) *QuickRun { _ = "STUB: not implemented"; return nil }

// Run all func
func (p *QuickRun) Run() error { _ = "STUB: not implemented"; return nil }
