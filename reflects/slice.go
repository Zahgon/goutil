package reflects

import (
	"reflect"
)

// MakeSliceByElem create a new slice by the element type.
//
// - elType: the type of the element.
// - returns: the new slice.
//
// Usage:
//
//	sl := MakeSliceByElem(reflect.TypeOf(1), 10, 20)
//	sl.Index(0).SetInt(10)
//
//	// Or use reflect.AppendSlice() merge two slice
//	// Or use `for` with `reflect.Append()` add elements
func MakeSliceByElem(elTyp reflect.Type, len, cap int) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}

// FlatSlice flatten multi-level slice to given depth-level slice.
//
// Example:
//
//	FlatSlice([]any{ []any{3, 4}, []any{5, 6} }, 1) // Output: []any{3, 4, 5, 6}
//
// always return reflect.Value of []any. note: maybe flatSl.Cap != flatSl.Len
func FlatSlice(sl reflect.Value, depth int) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}

func addSliceItem(sl reflect.Value, depth int, collector func(item reflect.Value)) (c int) {
	_ = "STUB: not implemented"
	return 0
}
