package comfunc

import (
	"regexp"
	"time"
)

var (
	// check is duration string. TIP: extend unit d,w.  eg: "1d", "2w"
	//
	// time.ParseDuration() is max support hour "h".
	durStrReg = regexp.MustCompile(`^-?([0-9]+(?:\.[0-9]*)?(ns|us|µs|ms|s|m|h|d|w))+$`)

	// check long duration string. 验证整体格式是否符合
	//
	// eg: "1hour", "2hours", "3minutes", "4mins", "5days", "1weeks", "1month"
	//
	// time.ParseDuration() is not support long unit.
	durStrRegL = regexp.MustCompile(`^-?([0-9]+(?:\.[0-9]*)?[nuµsmhdw][a-zA-Z]{0,8})+$`)
	// use for parse duration string. see ToDuration()
	//
	// NOTE: 解析时，不能加最后的 `+` 会导致只匹配了最后一组 时间单位
	durStrRegL2 = regexp.MustCompile(`-?([0-9]+(?:\.[0-9]*)?)([nuµsmhdw][a-z]{0,8})`)
)

// IsDuration check the string is a duration string.
func IsDuration(s string) bool { _ = "STUB: not implemented"; return false }

// ToDuration parses a duration string. such as "300ms", "-1.5h" or "2h45m".
// Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
//
// Diff of time.ParseDuration:
//   - support extends unit d, w at the end of string. such as "1d", "2w".
//   - support extends unit: month, week, day
//   - support long string unit at the end. such as "1hour", "2hours", "3minutes", "4mins", "5days", "1weeks".
//
// If the string is not a valid duration string, it will return an error.
func ToDuration(s string) (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}

// check duration string is valid

// if ln < 4 AND end != d|w, directly call time.ParseDuration()

// time.ParseDuration() is not support long unit.

// fmt.Println(ssList)

// only one element. eg: "1day"

// more than one element. eg: "1day2hour3min"

// convert to short unit
func parseLongUnit(ss []string, bts []byte) []byte {
	_ = "STUB: not implemented"
	// eg: "3sec" -> ss=[3sec, -3, sec]
	return nil
}

// time lib max unit is hour, so need convert by 24 * 30*n

// time lib max unit is hour, so need convert by 24 * 7*n

// time lib max unit is hour, so need convert by 24*n

// '-' has been added on ToDuration()

func appendNumToBytes(bts []byte, num string, multiple int) []byte {
	_ = "STUB: not implemented"
	return nil
}

// is float number

// 使用 Float 保留两位小数 -> 会始终有两位小数，即使是N.00
// bts = strconv.AppendFloat(bts, val, 'f', 2, 64)

// 四舍五入到两位小数

// 使用 AppendFloat 自动去除末尾的 .0 或 .00
