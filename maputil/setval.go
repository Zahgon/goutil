package maputil

import (
	"reflect"
)

// SetByPath set sub-map value by key path.
// Supports dot syntax to set deep values.
//
// For example:
//
//	SetByPath("name.first", "Mat")
func SetByPath(mp *map[string]any, path string, val any) error {
	_ = "STUB: not implemented"
	return nil
}

// SetByKeys set sub-map value by path keys.
// Supports dot syntax to set deep values.
//
// For example:
//
//	SetByKeys([]string{"name", "first"}, "Mat")
func SetByKeys(mp *map[string]any, keys []string, val any) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func setMapByKeys(rv reflect.Value, keys []string, nv reflect.Value) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// If the map is nil, make a new map

// slice index key must be ended on the keys.
// eg: "top.arr[2]" -> "arr[2]"

// update value

// deep make map by keys

// get real type: any -> map

// tmpV.Index(idx).Set(elemV)

// last - set value

// set value on last key

// key is slice index

// get real type: any -> map

// sub is slice and is not ptr

// to (E)

// next key is index number.

// rv = tmpV.Index(idx) // TODO

// deep make map by keys

// (E). slice from ptr slice

func getRealVal(rv reflect.Value) (reflect.Value, bool) {
	_ = "STUB: not implemented"
	// get real type: any -> map
	return *new(reflect.Value), false
}

// "arr[2]" => "arr", 2, true
func parseArrKeyIndex(key string) (string, int, bool) {
	_ = "STUB: not implemented"
	return "", 0, false
}
