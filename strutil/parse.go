package strutil

import (
	"errors"
	"regexp"
	"time"
)

var regNumVersion = regexp.MustCompile(`[0-9][\d.]+([_-]\d+)?`)

// NumVersion parse input string, get valid number version. eg: go-1.22.3 -> 1.22.3
func NumVersion(s string) string { _ = "STUB: not implemented"; return "" }

// MustToTime convert date string to time.Time
func MustToTime(s string, layouts ...string) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

// auto match uses some common layouts.
// key is layout length.
var layoutMap = map[int][]string{
	6:  {"200601", "060102", time.Kitchen},
	8:  {"20060102", "06-01-02"},
	10: {"2006-01-02"},
	13: {"2006-01-02 15"},
	15: {time.Stamp},
	16: {"2006-01-02 15:04"},
	19: {"2006-01-02 15:04:05", time.RFC822, time.StampMilli},
	20: {"2006-01-02 15:04:05Z"},
	21: {time.RFC822Z},
	22: {time.StampMicro},
	23: {"2006-01-02 15:04:05.000", "2006-01-02 15:04:05.999"},
	24: {time.ANSIC},
	25: {time.RFC3339, time.StampNano},
	// time.Layout}, // must go >= 1.19
	26: {"2006-01-02 15:04:05.000000"},
	28: {time.UnixDate},
	29: {time.RFC1123, "2006-01-02 15:04:05.000000000"},
	30: {time.RFC850},
	31: {time.RFC1123Z},
	35: {time.RFC3339Nano},
}

// ToTime convert date string to time.Time
//
// NOTE: always use local timezone.
func ToTime(s string, layouts ...string) (t time.Time, err error) {
	_ = "STUB: not implemented"
	// custom layout
	return *new(time.Time), nil
}

// auto match use some commonly layouts.

// date string has "T". eg: "2006-01-02T15:04:05"

// date string has "/". eg: "2006/01/02 15:04:05"

// t, err = time.ParseInLocation(layout, s, time.Local)

// ParseSizeOpt parse size expression options
type ParseSizeOpt struct {
	// OneAsMax if only one size value, use it as max size. default is false
	OneAsMax bool
	// SepChar is the separator char for time range string. default is '~'
	SepChar byte
	// KeywordFn is the function for parse keyword time string.
	KeywordFn func(string) (min, max uint64, err error)
}

func ensureOpt(opt *ParseSizeOpt) *ParseSizeOpt { _ = "STUB: not implemented"; return nil }

// ErrInvalidSizeExpr invalid size expression error
var ErrInvalidSizeExpr = errors.New("invalid size expr")

// ParseSizeRange parse range size expression to min and max size.
//
// Expression format:
//
//	"1KB~2MB"       => 1KB to 2MB
//	"-1KB"          => <1KB
//	"~1MB"          => <1MB
//	"< 1KB"         => <1KB
//	"1KB"           => >1KB
//	"1KB~"          => >1KB
//	">1KB"          => >1KB
//	"+1KB"          => >1KB
func ParseSizeRange(expr string, opt *ParseSizeOpt) (min, max uint64, err error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

// parse size range. eg: "1KB~2MB"

// parse single size. eg: "1KB"

// parse with prefix. eg: "<1KB", ">= 1KB", "-1KB", "+1KB"

// parse keyword. eg: "small", "large"

// SafeByteSize converts size string like 1GB/1g or 12mb/12M into an unsigned integer number of bytes
func SafeByteSize(sizeStr string) uint64 { _ = "STUB: not implemented"; return 0 }

// ToByteSize converts size string like 1GB/1g or 12mb/12M into an unsigned integer number of bytes
func ToByteSize(sizeStr string) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

// last second char is k,m,g,t

// not unit suffix. eg: 346

// b
