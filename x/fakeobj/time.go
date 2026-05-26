package fakeobj

import (
	"time"
)

// Clock mock time clock for test
type Clock struct {
	tt time.Time
}

// NewClock create a mock clock instance from layout "2006-01-02 15:04:05"
//
// Example:
//
//	tc := NewClock("2023-01-01 12:00:00")
//	tc.Add(time.Second * 15)
//	ds := tc.Datetime() // "2023-01-01 12:00:15"
func NewClock(value string) *Clock { _ = "STUB: not implemented"; return nil }

// Now get current time.
func (mt *Clock) Now() time.Time {
	_ = "STUB: not implemented"

	// Add progresses time by the given duration.
	return *new(time.Time)
}

func (mt *Clock) Add(d time.Duration) { _ = "STUB: not implemented"; return }

// Datetime returns the current time in the format "2006-01-02 15:04:05".
func (mt *Clock) Datetime() string { _ = "STUB: not implemented"; return "" }
