package reflects

import (
	"reflect"
)

// TryAnyMap convert map[TYPE1]TYPE2 to map[string]any
func TryAnyMap(mp reflect.Value) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// EachMap process any map data
func EachMap(mp reflect.Value, fn func(key, val reflect.Value)) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// EachStrAnyMap process any map data as string key and any value
func EachStrAnyMap(mp reflect.Value, fn func(key string, val any)) error {
	_ = "STUB: not implemented"
	return nil
}

// FlatFunc custom collect handle func
type FlatFunc func(path string, val reflect.Value)

// FlatMap process tree map to flat key-value map.
//
// Examples:
//
//	{"top": {"sub": "value", "sub2": "value2"} }
//	->
//	{"top.sub": "value", "top.sub2": "value2" }
func FlatMap(rv reflect.Value, fn FlatFunc) { _ = "STUB: not implemented"; return }

func flatMap(rv reflect.Value, fn FlatFunc, parent string) { _ = "STUB: not implemented"; return }

func flatSlice(rv reflect.Value, fn FlatFunc, parent string) { _ = "STUB: not implemented"; return }
