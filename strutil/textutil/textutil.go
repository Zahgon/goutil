// Package textutil provide some extensions text handle util functions.
package textutil

import (
	"github.com/gookit/goutil/maputil"
)

// ReplaceVars by regex replace given tpl vars.
//
// If a format is empty, will use {const DefaultVarFormat}
func ReplaceVars(text string, vars map[string]any, format string) string {
	_ = "STUB: not implemented"
	return ""
}

// RenderSMap by regex replacement given tpl vars.
//
// If a format is empty, will use {const DefaultVarFormat}
func RenderSMap(text string, vars map[string]string, format string) string {
	_ = "STUB: not implemented"
	return ""
}

// IsMatchAll keywords in the give text string.
//
// TIP: can use ^ for exclude match.
func IsMatchAll(s string, keywords []string) bool { _ = "STUB: not implemented"; return false }

// ParseInlineINI parse config string to string-map. it's like INI format contents.
//
// Examples:
//
//	eg: "name=val0;shorts=i;required=true;desc=a message"
//	=>
//	{name: val0, shorts: i, required: true, desc: a message}
func ParseInlineINI(tagVal string, keys ...string) (mp maputil.SMap, err error) {
	_ = "STUB: not implemented"
	return *new(maputil.SMap), nil
}

// ParseSimpleINI parse simple multiline config string to a string-map.
// Can use to parse simple INI or dotenv file contents.
//
// NOTE:
//
//   - it's like INI format contents.
//   - support comments line with: "#", ";", "//"
//   - support inline comments with: " #" eg: name=tom # a comments
//   - DON'T support submap parse.
func ParseSimpleINI(text string) (mp maputil.SMap, err error) {
	_ = "STUB: not implemented"
	return *new(maputil.SMap), nil
}
