package maputil

import (
	"github.com/gookit/goutil/comdef"
	"github.com/gookit/goutil/reflects"
)

// alias functions
var (
	// ToStrMap convert map[string]any to map[string]string
	ToStrMap = ToStringMap
	// ToL2StrMap convert map[string]any to map[string]map[string]string
	ToL2StrMap = ToL2StringMap
)

// KeyToLower convert keys to lower case.
func KeyToLower(src map[string]string) map[string]string { _ = "STUB: not implemented"; return nil }

// AnyToStrMap try convert any(map[string]any, map[string]string) to map[string]string
func AnyToStrMap(src any) map[string]string { _ = "STUB: not implemented"; return nil }

// ToStringMap simple convert map[string]any to map[string]string
func ToStringMap(src map[string]any) map[string]string { _ = "STUB: not implemented"; return nil }

// ToL2StringMap convert map[string]any to map[string]map[string]string
func ToL2StringMap(groupsMap map[string]any) map[string]map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// CombineToSMap combine two string-slices to SMap(map[string]string)
func CombineToSMap(keys, values []string) SMap { _ = "STUB: not implemented"; return *new(SMap) }

// CombineToMap combine two any slice to map[K]V. alias of arrutil.CombineToMap
func CombineToMap[K comdef.SortedType, V any](keys []K, values []V) map[K]V {
	_ = "STUB: not implemented"
	return nil
}

// SliceToSMap convert string k-v pairs slice to map[string]string
//   - eg: []string{k1,v1,k2,v2} -> map[string]string{k1:v1, k2:v2}
func SliceToSMap(kvPairs ...string) map[string]string {
	_ = "STUB: not implemented"

	// check kvPairs length must be even
	return nil
}

// SliceToMap convert any k-v pairs slice to map[string]any
func SliceToMap(kvPairs ...any) map[string]any {
	_ = "STUB: not implemented"

	// check kvPairs length must be even
	return nil
}

// SliceToTypeMap convert k-v pairs slice to map[string]T
func SliceToTypeMap[T any](valFunc func(any) T, kvPairs ...any) map[string]T {
	_ = "STUB: not implemented"

	// check kvPairs length must be even
	return nil
}

// ToAnyMap convert map[TYPE1]TYPE2 to map[string]any
func ToAnyMap(mp any) map[string]any { _ = "STUB: not implemented"; return nil }

// TryAnyMap convert map[TYPE1]TYPE2 to map[string]any
func TryAnyMap(mp any) (map[string]any, error) { _ = "STUB: not implemented"; return nil, nil }

// HTTPQueryString convert map[string]any data to http query string.
func HTTPQueryString(data map[string]any) string { _ = "STUB: not implemented"; return "" }

// StringsMapToAnyMap convert map[string][]string to map[string]any
//
//	Example:
//	{"k1": []string{"v1", "v2"}, "k2": []string{"v3"}}
//	=>
//	{"k": []string{"v1", "v2"}, "k2": "v3"}
//
//	mp := StringsMapToAnyMap(httpReq.Header)
func StringsMapToAnyMap(ssMp map[string][]string) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

// ToString simple and quickly convert map[string]any to string.
func ToString(mp map[string]any) string { _ = "STUB: not implemented"; return "" }

// remove last ', '

// ToString2 simple and quickly convert a map to string.
func ToString2(mp any) string { _ = "STUB: not implemented"; return "" }

// FormatIndent format map data to string with newline and indent.
func FormatIndent(mp any, indent string) string { _ = "STUB: not implemented"; return "" }

// StrMapToText 将 map[string]string 转换为多行 key=value 格式文本
func StrMapToText(m map[string]string) string { _ = "STUB: not implemented"; return "" }

/*************************************************************
 * Flat convert tree map to flatten key-value map.
 *************************************************************/

// Flatten convert tree map to flat key-value map.
//
// Examples:
//
//	{"top": {"sub": "value", "sub2": "value2"} }
//	->
//	{"top.sub": "value", "top.sub2": "value2" }
func Flatten(mp map[string]any) map[string]any { _ = "STUB: not implemented"; return nil }

// FlatWithFunc flat a tree-map with custom collect handle func
func FlatWithFunc(mp map[string]any, fn reflects.FlatFunc) { _ = "STUB: not implemented"; return }
