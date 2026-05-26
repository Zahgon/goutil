package assert

// Assertions provide assertion methods around the TestingT interface.
type Assertions struct {
	t  TestingT
	ok bool // last assert result
	// prefix message for each assert TODO
	Msg string
}

// New makes a new Assertions object for the specified TestingT
//
// Usage:
//
//	as := assert.New(t)
//	as.True(true)
func New(t TestingT) *Assertions { _ = "STUB: not implemented"; return nil }

// WithMsg set with prefix message.
func (as *Assertions) WithMsg(msg string) *Assertions { _ = "STUB: not implemented"; return nil }

// IsOk for last check
func (as *Assertions) IsOk() bool {
	_ = "STUB: not implemented"

	// IsFail for last check
	return false
}

func (as *Assertions) IsFail() bool { _ = "STUB: not implemented"; return false }
