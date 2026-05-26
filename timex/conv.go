package timex

import (
	"time"
)

// Elapsed calc elapsed time from start time to end time.
func Elapsed(start, end time.Time) string { _ = "STUB: not implemented"; return "" }

// ElapsedNow calc elapsed time from start time to now.
func ElapsedNow(start time.Time) string { _ = "STUB: not implemented"; return "" }

// FormatDuration Formatting time consumption is clock format. 格式化时间消耗为时钟格式
//
// eg: 90 * time.Second => "01:30"
func FormatDuration(d time.Duration) string { _ = "STUB: not implemented"; return "" }

//
// -------- parse time diff to string --------
//

// TimeMessage struct for HowLongAgo2(), FromNowWith()
type TimeMessage struct {
	// Message string or format string
	Message string
	// Seconds time range.
	// first elem is max boundary value, second elem is divisor(unit).
	Seconds []int
}

// TimeMessages time message list.
//
// NOTE: last item.Seconds[0] is min boundary value.
var TimeMessages = []TimeMessage{
	{"< 1 sec ago", []int{0}},
	{"1 sec ago", []int{1}},
	{"%d secs ago", []int{45, 1}},
	{"1 min ago", []int{89}},
	{"%d mins ago", []int{44 * 60, 60}},       // 89s - 44min
	{"1 hour ago", []int{89 * 60}},            // 45min - 89min
	{"%d hours ago", []int{21 * 3600, 3600}},  // 90min - 21hr
	{"1 day ago", []int{35 * 3600}},           // 22 - 35 hours
	{"%d days ago", []int{30 * 86400, 86400}}, // 36hr - 30 day
	// {"1 week ago", []int{10 * 86400}},          // 7 - 10 days
	// {"%d weeks ago", []int{30 * 86400, 604800}},  // 10 - 30 days
	{"1 month ago", []int{45 * 86400}},                 // 31 - 45 days
	{"%d months ago", []int{319 * 86400, 2592000}},     // 44 - 319 days
	{"1 year ago", []int{547 * 86400}},                 // 320 - 547 days
	{"%d years ago", []int{547 * 86400, 12 * 2592000}}, // > 547 days, unit is year
}

// FromNow format time from now, returns like: 1 hour ago, 2 days ago
//
// refer: https://gist.github.com/davidrleonard/259fe449b1ec13bf7d87cde567ca0fde
func FromNow(t time.Time) string { _ = "STUB: not implemented"; return "" }

// FromNowWith format time from now with custom TimeMessage list
func FromNowWith(u time.Time, tms []TimeMessage) string { _ = "STUB: not implemented"; return "" }

// HowLongAgo format diff time seconds to string. alias of HowLongAgo2()
func HowLongAgo(diffSec int64) string { _ = "STUB: not implemented"; return "" }

// HowLongAgo2 format diff time seconds with custom TimeMessage list
func HowLongAgo2(diffSec int64, tms []TimeMessage) string { _ = "STUB: not implemented"; return "" }

// match success: is last elem or diffSec <= secs[0]

//
// -------- parse string to time --------
//

// ToTime parse a datetime string. alias of strutil.ToTime()
func ToTime(s string, layouts ...string) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

// ToDur parse a duration string. alias of ToDuration()
func ToDur(s string) (time.Duration, error) {
	_ = "STUB: not implemented"
	return *

	// ParseDuration parse a duration string. alias of ToDuration()
	new(time.Duration), nil
}

func ParseDuration(s string) (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}

// ToDuration parses a duration string. such as "300ms", "-1.5h" or "2h45m".
// Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
//
// it like time.ParseDuration, but supports more unit string. eg: "1d", "2w"
func ToDuration(s string) (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}

// TryToTime parse a date string or duration string to time.Time.
//
// if s is empty, return zero time.
func TryToTime(s string, bt time.Time) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

// if s is a duration string, add it to bt(base time)

// as a date string, parse it to time.Time

//
// -------- parse string to time range --------
//

// ParseRangeOpt is the option for ParseRange
type ParseRangeOpt struct {
	// BaseTime is the base time for relative time string.
	// if is zero, use time.Now() as base time.
	BaseTime time.Time
	// OneAsEnd is the option for one time range.
	//  - False: "-1h" => "-1h,0"; "1h" => "+1h, feature"
	//  - True:  "-1h" => "zero,-1h"; "1h" => "zero,1h"
	OneAsEnd bool
	// AutoSort is the option for sort the time range.
	AutoSort bool
	// SepChar is the separator char for time range string. default is '~'
	SepChar byte
	// BeforeFn hook for before parse time string.
	BeforeFn func(string) string
	// KeywordFn is the function for parse keyword time string.
	KeywordFn func(string) (time.Time, time.Time, error)
}

func ensureOpt(opt *ParseRangeOpt) *ParseRangeOpt { _ = "STUB: not implemented"; return nil }

// ParseRange parse time range expression string to time.Time range.
//
//   - "0" will use opt.BaseTime.
//
// Expression format:
//
//	"-5h~-1h"       	=> 5 hours ago to 1 hour ago
//	"1h~5h"         	=> 1 hour after to 5 hours after
//	"-1h~1h"        	=> 1 hour ago to 1 hour after
//	"-1h"            	=> 1 hour ago to feature. eq "-1h~"
//	"-1h~0"          	=> 1 hour ago to now.
//	"< -1h" OR "~-1h"   => 1 hour ago.
//	"> 1h" OR "1h"     	=> 1 hour after to feature
//	// keyword: now, today, yesterday, tomorrow
//	"today"          => today start to today end
//	"yesterday"      => yesterday start to yesterday end
//	"tomorrow"       => tomorrow start to tomorrow end
//
// Usage:
//
//	start, end, err := ParseRange("-1h~1h", nil)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(start, end)
func ParseRange(expr string, opt *ParseRangeOpt) (start, end time.Time, err error) {
	_ = "STUB: not implemented"
	return *new(time.Time), *new(time.Time), nil
}

// parse time range. eg: "5h~1h"

// auto sort range time

// single time. eg: "5h", "1h", "-1h"

// with compare operator. eg: "<1h", ">1h"

// parse keyword time string

// single datetime. eg: "2019-01-01"
