package cflag

import (
	"github.com/gookit/goutil/structs"
)

// OptCheckFn define
type OptCheckFn func(val any) error

// FlagOpt struct
type FlagOpt struct {
	// Shortcuts short names. eg: ["o", "a"]
	Shortcuts []string
	// Required option
	Required bool
	// Validator for check option value
	Validator OptCheckFn
}

// HelpName string
func (o *FlagOpt) HelpName(name string) string { _ = "STUB: not implemented"; return "" }

// FlagArg struct
type FlagArg struct {
	// Value for the flag argument
	*structs.Value
	// default value
	defVal any
	// default value string, use of help
	defStr string
	// Name of the argument
	Name string
	// Desc arg description
	Desc string
	// Index of the argument
	Index int
	// Required argument
	Required bool
	// Arrayed argument. MUST on the last
	Arrayed bool // support arrayed argument
	// Validator for check value
	Validator func(val string) error
}

// NewArg create instance
func NewArg(name, desc string, required bool) *FlagArg { _ = "STUB: not implemented"; return nil }

// check arg config and init
func (a *FlagArg) check() error { _ = "STUB: not implemented"; return nil }

// HelpDesc string build
func (a *FlagArg) HelpDesc() string { _ = "STUB: not implemented"; return "" }
