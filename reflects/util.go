package reflects

import (
	"reflect"
)

// loopIndirect returns the item at the end of indirection, and a bool to indicate
// if it's nil. If the returned bool is true, the returned value's kind will be
// either a pointer or interface.
func loopIndirect(v reflect.Value) (rv reflect.Value, isNil bool) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), false
}

// indirectInterface returns the concrete value in an interface value,
// or else the zero reflect.Value.
// That is, if v represents the interface value x, the result is the same as reflect.ValueOf(x):
// the fact that x was an interface value is forgotten.
func indirectInterface(v reflect.Value) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}

// Elem returns the value that the interface v contains
// or that the pointer v points to. otherwise, will return self
func Elem(v reflect.Value) reflect.Value { _ = "STUB: not implemented"; return *new(reflect.Value) }

// Indirect like reflect.Indirect(), but can also indirect reflect.Interface. otherwise, will return self
func Indirect(v reflect.Value) reflect.Value { _ = "STUB: not implemented"; return *new(reflect.Value) }

// UnwrapAny unwrap reflect.Interface value. otherwise, will return self
func UnwrapAny(v reflect.Value) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}

// TypeReal returns a ptr type's real type. otherwise, will return self.
func TypeReal(t reflect.Type) reflect.Type { _ = "STUB: not implemented"; return *new(reflect.Type) }

// TypeElem returns the array, slice, chan, map type's element type. otherwise, will return self.
func TypeElem(t reflect.Type) reflect.Type { _ = "STUB: not implemented"; return *new(reflect.Type) }

// Len get reflect value length. allow: intX, uintX, floatX, string, map, array, chan, slice.
//
// Note: (u)intX use width. float to string then calc len.
func Len(v reflect.Value) int { _ = "STUB: not implemented"; return 0 }

// (u)int use width.

// cannot get length

// SliceSubKind get sub-elem kind of the array, slice, variadic-var. alias SliceElemKind()
func SliceSubKind(typ reflect.Type) reflect.Kind {
	_ = "STUB: not implemented"
	return *new(reflect.Kind)
}

// SliceElemKind get sub-elem kind of the array, slice, variadic-var.
//
// Usage:
//
//	SliceElemKind(reflect.TypeOf([]string{"abc"})) // reflect.String
func SliceElemKind(typ reflect.Type) reflect.Kind {
	_ = "STUB: not implemented"
	return *new(reflect.Kind)
}

// UnexportedValue quickly get unexported value by reflect.Value
//
// NOTE: this method is unsafe, use it carefully.
// should ensure rv is addressable by field.CanAddr()
//
// refer: https://stackoverflow.com/questions/42664837/how-to-access-unexported-struct-fields
func UnexportedValue(rv reflect.Value) any {
	_ = "STUB: not implemented"

	// create new value from addr, now can be read and set.
	return *new(any)
}

// If the rv is not addressable this trick won't work, but you can create an addressable copy like this

// Now rv can be read. TIP: Setting will succeed but only affects the temporary copy.

// SetUnexportedValue quickly set unexported field value by reflect
//
// NOTE: this method is unsafe, use it carefully.
// should ensure rv is addressable by field.CanAddr()
func SetUnexportedValue(rv reflect.Value, value any) { _ = "STUB: not implemented"; return }

// SetValue to a `reflect.Value`. will auto convert type if needed.
func SetValue(rv reflect.Value, val any) error {
	_ = "STUB: not implemented"
	// get a real type of the ptr value
	return nil
}

// use elem for set value

// SetRValue to a `reflect.Value`. will direct set value without a type convert.
func SetRValue(rv, val reflect.Value) { _ = "STUB: not implemented"; return }
