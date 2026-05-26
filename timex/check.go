package timex

import (
	"time"
)

// IsDuration check the string is a valid duration string.
// alias of comfunc.IsDuration()
func IsDuration(s string) bool { _ = "STUB: not implemented"; return false }

// InRange check the dst time is in the range of start and end.
//
// if start is zero, only check dst < end,
// if end is zero, only check dst > start.
func InRange(dst, start, end time.Time) bool { _ = "STUB: not implemented"; return false }
