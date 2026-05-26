package process

// PidFile struct
type PidFile struct {
	pid  int
	file string
	// body string
}

// NewPidFile instance
func NewPidFile(file string) *PidFile { _ = "STUB: not implemented"; return nil }

// Exists of th pid file
func (pf *PidFile) Exists() bool { _ = "STUB: not implemented"; return false }

// File path
func (pf *PidFile) File() string {
	_ = "STUB: not implemented"

	// PID value
	return ""
}

func (pf *PidFile) PID() int { _ = "STUB: not implemented"; return 0 }

// String PID value string
func (pf *PidFile) String() string { _ = "STUB: not implemented"; return "" }

// SetPID value
func (pf *PidFile) SetPID(val int) int { _ = "STUB: not implemented"; return 0 }

// Save PID value to file
func (pf *PidFile) Save() error { _ = "STUB: not implemented"; return nil }
