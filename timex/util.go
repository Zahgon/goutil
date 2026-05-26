package timex

import (
	"time"
)

// NowUnix is short of time.Now().Unix()
func NowUnix() int64 { _ = "STUB: not implemented"; return 0 }

// NowDate quick get current date string. if template is empty, will use DefaultTemplate.
func NowDate(template ...string) string { _ = "STUB: not implemented"; return "" }

// Format convert time to string use default layout
func Format(t time.Time) string { _ = "STUB: not implemented"; return "" }

// FormatBy given default layout
func FormatBy(t time.Time, layout string) string { _ = "STUB: not implemented"; return "" }

// Date format time by given date template. see ToLayout() for template parse.
func Date(t time.Time, template ...string) string { _ = "STUB: not implemented"; return "" }

// Datetime convert time to string use template. see ToLayout() for template parse.
func Datetime(t time.Time, template ...string) string { _ = "STUB: not implemented"; return "" }

// DateFormat format time by given date template. see ToLayout()
func DateFormat(t time.Time, template string) string { _ = "STUB: not implemented"; return "" }

// FormatByTpl format time by given date template. see ToLayout()
func FormatByTpl(t time.Time, template string) string { _ = "STUB: not implemented"; return "" }

// FormatUnix time seconds use default layout
func FormatUnix(sec int64, layout ...string) string { _ = "STUB: not implemented"; return "" }

// FormatUnixBy format time seconds use given layout
func FormatUnixBy(sec int64, layout string) string { _ = "STUB: not implemented"; return "" }

// FormatUnixByTpl format time seconds use given date template.
// see ToLayout()
func FormatUnixByTpl(sec int64, template ...string) string { _ = "STUB: not implemented"; return "" }
