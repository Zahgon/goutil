// Package timex provides an enhanced time.Time implementation.
// Add more commonly used functional methods.
//
// such as: DayStart(), DayAfter(), DayAgo(), DateFormat() and more.
package timex

import (
	"time"
)

// provide some common time constants
const (
	OneSecond   = 1
	OneMinSec   = 60
	OneHourSec  = 3600
	OneDaySec   = 86400
	OneWeekSec  = 7 * 86400
	OneMonthSec = 30 * 86400

	MinSec   = OneMinSec
	HourSec  = OneHourSec
	DaySec   = OneDaySec
	WeekSec  = OneWeekSec
	MonthSec = OneMonthSec

	Microsecond = time.Microsecond
	Millisecond = time.Millisecond

	Second  = time.Second
	OneMin  = time.Minute
	Minute  = time.Minute
	OneHour = time.Hour
	Hour    = time.Hour
	OneDay  = 24 * time.Hour
	Day     = OneDay
	OneWeek = 7 * 24 * time.Hour
	Week    = OneWeek
	Month   = 30 * 24 * time.Hour
)

// TimeX alias of Time
// Deprecated: use Time instead
type TimeX = Time

// Time an enhanced time.Time implementation.
type Time struct {
	time.Time
	// Layout set the default date format layout. default use DefaultLayout
	Layout string
}

/*************************************************************
 * Create timex instance
 *************************************************************/

// Now time instance
func Now() *Time { _ = "STUB: not implemented"; return nil }

// New instance form given time
func New(t time.Time) *Time { _ = "STUB: not implemented"; return nil }

// Wrap the go time instance. alias of the New()
func Wrap(t time.Time) *Time { _ = "STUB: not implemented"; return nil }

// FromTime new instance form given time.Time. alias of the New()
func FromTime(t time.Time) *Time { _ = "STUB: not implemented"; return nil }

// Local time for now
func Local() *Time { _ = "STUB: not implemented"; return nil }

// FromUnix create from unix time
func FromUnix(sec int64) *Time { _ = "STUB: not implemented"; return nil }

// FromDate create from datetime string.
func FromDate(s string, template ...string) (*Time, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FromString create from datetime string. see strutil.ToTime()
func FromString(s string, layouts ...string) (*Time, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// LocalByName time for now. eg: "Local", "UTC", "Asia/Shanghai"
func LocalByName(tzName string) *Time { _ = "STUB: not implemented"; return nil }

/*************************************************************
 * timex usage
 *************************************************************/

// T returns the t.Time
func (t *Time) T() time.Time {
	_ = "STUB: not implemented"

	// Format returns a textual representation of the time value formatted according to the layout defined by the argument.
	//
	// see time.Time.Format()
	return *new(time.Time)
}

func (t *Time) Format(layout string) string { _ = "STUB: not implemented"; return "" }

// Datetime use layout or DefaultLayout format time to date. see Format()
func (t *Time) Datetime(layout ...string) string { _ = "STUB: not implemented"; return "" }

// TplFormat use input template format time to date.
//
// alias of DateFormat()
func (t *Time) TplFormat(template string) string { _ = "STUB: not implemented"; return "" }

// DateFormat use input template format time to date.
//
// Example:
//
//	tn := timex.Now()
//	tn.DateFormat("Y-m-d H:i:s") // Output: 2019-01-01 12:12:12
//	tn.DateFormat("Y-m-d H:i") // Output: 2019-01-01 12:12
//	tn.DateFormat("Y-m-d") // Output: 2019-01-01
//	tn.DateFormat("Y-m") // Output: 2019-01
//	tn.DateFormat("y-m-d") // Output: 19-01-01
//	tn.DateFormat("ymd") // Output: 190101
//
// see ToLayout() for convert template to layout.
func (t *Time) DateFormat(template string) string { _ = "STUB: not implemented"; return "" }

// Yesterday got day ago time for the time
func (t *Time) Yesterday() *Time { _ = "STUB: not implemented"; return nil }

// DayAgo get some day ago time for the time
func (t *Time) DayAgo(day int) *Time { _ = "STUB: not implemented"; return nil }

// AddDay add some daytime for the time
func (t *Time) AddDay(day int) *Time { _ = "STUB: not implemented"; return nil }

// SubDay add some day time for the time
func (t *Time) SubDay(day int) *Time { _ = "STUB: not implemented"; return nil }

// Tomorrow time. get tomorrow time for the time
func (t *Time) Tomorrow() *Time { _ = "STUB: not implemented"; return nil }

// DayAfter get some day after time for the time.
// alias of Time.AddDay()
func (t *Time) DayAfter(day int) *Time { _ = "STUB: not implemented"; return nil }

// AddDur some duration time
func (t *Time) AddDur(dur time.Duration) *Time { _ = "STUB: not implemented"; return nil }

// AddString add duration time string.
//
// Example:
//
//	tn := timex.Now() // example as "2019-01-01 12:12:12"
//	nt := tn.AddString("1h")
//	nt.Datetime() // Output: 2019-01-01 13:12:12
func (t *Time) AddString(dur string) *Time { _ = "STUB: not implemented"; return nil }

// AddHour add some hour time
func (t *Time) AddHour(hours int) *Time { _ = "STUB: not implemented"; return nil }

// SubHour minus some hour time
func (t *Time) SubHour(hours int) *Time { _ = "STUB: not implemented"; return nil }

// AddMinutes add some minutes time for the time
func (t *Time) AddMinutes(minutes int) *Time { _ = "STUB: not implemented"; return nil }

// SubMinutes minus some minutes time for the time
func (t *Time) SubMinutes(minutes int) *Time { _ = "STUB: not implemented"; return nil }

// AddSeconds add some seconds time the time
func (t *Time) AddSeconds(seconds int) *Time { _ = "STUB: not implemented"; return nil }

// with layout

// SubSeconds minus some seconds time the time
func (t *Time) SubSeconds(seconds int) *Time { _ = "STUB: not implemented"; return nil }

// with layout

// Diff calc diff duration for t - u. alias of time.Time.Sub()
func (t *Time) Diff(u time.Time) time.Duration {
	_ = "STUB: not implemented"

	// DiffSec calc diff seconds for t - u
	return *new(time.Duration)
}

func (t *Time) DiffSec(u time.Time) int { _ = "STUB: not implemented"; return 0 }

// DiffUnix calc diff seconds for t.Unix() - u
func (t *Time) DiffUnix(u int64) int { _ = "STUB: not implemented"; return 0 }

// SubUnix calc diff seconds for t - u
func (t *Time) SubUnix(u time.Time) int { _ = "STUB: not implemented"; return 0 }

// HourStart time
func (t *Time) HourStart() *Time { _ = "STUB: not implemented"; return nil }

// HourEnd time
func (t *Time) HourEnd() *Time { _ = "STUB: not implemented"; return nil }

// DayStart get time at 00:00:00
func (t *Time) DayStart() *Time { _ = "STUB: not implemented"; return nil }

// DayEnd get time at 23:59:59
func (t *Time) DayEnd() *Time { _ = "STUB: not implemented"; return nil }

// CustomHMS custom change the hour, minute, second for create new time.
func (t *Time) CustomHMS(hour, min, sec int) *Time { _ = "STUB: not implemented"; return nil }

// IsEmpty check if time is empty. alias for t.IsZero
func (t *Time) IsEmpty() bool {
	_ = "STUB: not implemented"

	// IsBefore the given time
	return false
}

func (t *Time) IsBefore(u time.Time) bool {
	_ = "STUB: not implemented"

	// IsBeforeUnix the given unix timestamp
	return false
}

func (t *Time) IsBeforeUnix(ux int64) bool { _ = "STUB: not implemented"; return false }

// IsAfter the given time
func (t *Time) IsAfter(u time.Time) bool {
	_ = "STUB: not implemented"

	// IsAfterUnix the given unix timestamp
	return false
}

func (t *Time) IsAfterUnix(ux int64) bool { _ = "STUB: not implemented"; return false }

// Timestamp value. alias of t.Unix()
func (t *Time) Timestamp() int64 {
	_ = "STUB: not implemented"

	// HowLongAgo format diff time to string.
	return 0
}

func (t *Time) HowLongAgo(before time.Time) string { _ = "STUB: not implemented"; return "" }

// UnmarshalJSON implements the json.Unmarshaler interface.
//
// Tip: will auto match a format by strutil.ToTime()
func (t *Time) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	// Ignore null, like in the main JSON package.
	return nil
}

// Fractional seconds are handled implicitly by Parse.

// UnmarshalText implements the encoding.TextUnmarshaler interface.
//
// Tip: will auto match a format by strutil.ToTime()
func (t *Time) UnmarshalText(data []byte) error {
	_ = "STUB: not implemented"
	// Fractional seconds are handled implicitly by Parse.
	return nil
}
