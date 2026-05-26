package maputil

import (
	"github.com/gookit/goutil/comdef"
)

// some consts for separators
const (
	Wildcard = "*"
	PathSep  = "."
)

// DeepGet value by key path. eg "top" "top.sub"
func DeepGet(mp map[string]any, path string) (val any) { _ = "STUB: not implemented"; return *new(any) }

// QuietGet value by key path. eg "top" "top.sub"
func QuietGet(mp map[string]any, path string) (val any) {
	_ = "STUB: not implemented"
	return *new(any)
}

// GetFromAny get value by key path from any(map,slice) data. eg "top" "top.sub"
func GetFromAny(path string, data any) (val any, ok bool) {
	_ = "STUB: not implemented"
	// empty data
	return *new(any), false
}

// GetByPath get value by key path from a map(map[string]any). eg "top" "top.sub"
func GetByPath(path string, mp map[string]any) (val any, ok bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

// no sub key

// key is path. eg: "top.sub"

// GetByPathKeys get value by path keys from a map(map[string]any). eg "top" "top.sub"
//
// Example:
//
//	mp := map[string]any{
//		"top": map[string]any{
//			"sub": "value",
//		},
//	}
//	val, ok := GetByPathKeys(mp, []string{"top", "sub"}) // return "value", true
func GetByPathKeys(mp map[string]any, keys []string) (val any, ok bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

// find top item data use top key

// find sub item data use sub key

func getByPathKeys(item any, keys []string) (val any, ok bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

// is string map

// is map(decode from toml/json/yaml)

// is map(decode from yaml.v2)

// is an any-map slice

// * is last key

// * is not last key, find sub item data

// k is index number

// * is last key

// check is slice

// * is not last key, find sub item data

// el is map value.

// check k is index number

// as error

// next is last key and it is *

// Keys get all keys of the given map.
func Keys(mp any) (keys []string) { _ = "STUB: not implemented"; return nil }

// TypedKeys get all keys of the given typed map.
func TypedKeys[K comdef.SimpleType, V any](mp map[K]V) (keys []K) {
	_ = "STUB: not implemented"
	return nil
}

// FirstKey returns the first key of the given map.
func FirstKey[T any](mp map[string]T) string { _ = "STUB: not implemented"; return "" }

// Values get all values from the given map.
func Values(mp any) (values []any) { _ = "STUB: not implemented"; return nil }

// TypedValues get all values from the given typed map.
func TypedValues[K comdef.SimpleType, V any](mp map[K]V) (values []V) {
	_ = "STUB: not implemented"
	return nil
}

// EachAnyMap iterates the given map and calls the given function for each item.
func EachAnyMap(mp any, fn func(key string, val any)) { _ = "STUB: not implemented"; return }

// EachTypedMap iterates the given map and calls the given function for each item.
func EachTypedMap[K comdef.SimpleType, V any](mp map[K]V, fn func(key K, val V)) {
	_ = "STUB: not implemented"
	return
}
