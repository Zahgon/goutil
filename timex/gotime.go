package timex

import (
	"time"
)

// some time layout or time
const (
	DatetimeLayout = "2006-01-02 15:04:05"
	LayoutWithMs3  = "2006-01-02 15:04:05.000"
	LayoutWithMs6  = "2006-01-02 15:04:05.000000"
	DateOnlyLayout = "2006-01-02"
	TimeOnlyLayout = "15:04:05"

	// ZeroUnix zero unix timestamp
	ZeroUnix int64 = -62135596800
)

var (
	// DefaultLayout template for format time
	DefaultLayout = DatetimeLayout
	// ZeroTime zero time instance
	ZeroTime = time.Time{}
)

// SetLocalByName set local by tz name. eg: UTC, PRC
func SetLocalByName(tzName string) error { _ = "STUB: not implemented"; return nil }

// NowAddDay add some day time from now
func NowAddDay(day int) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// NowAddHour add some hour time from now
func NowAddHour(hour int) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// NowAddMinutes add some minutes time from now
func NowAddMinutes(minutes int) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// NowAddSec add some seconds time from now. alias of NowAddSeconds()
func NowAddSec(seconds int) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// NowAddSeconds add some seconds time from now
func NowAddSeconds(seconds int) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// NowHourStart time
func NowHourStart() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// NowHourEnd time
func NowHourEnd() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// AddDay add some day time for given time
func AddDay(t time.Time, day int) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// AddHour add some hour time for given time
func AddHour(t time.Time, hour int) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// AddMinutes add some minutes time for given time
func AddMinutes(t time.Time, minutes int) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

// AddSeconds add some seconds time for given time
func AddSeconds(t time.Time, seconds int) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

// AddSec add some seconds time for given time. alias of AddSeconds()
func AddSec(t time.Time, seconds int) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// HourStart time for given time
func HourStart(t time.Time) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// HourEnd time for given time
func HourEnd(t time.Time) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// DayStart time for given time
func DayStart(t time.Time) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// DayEnd time for given time
func DayEnd(t time.Time) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// TodayStart time
func TodayStart() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// TodayEnd time
func TodayEnd() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }
