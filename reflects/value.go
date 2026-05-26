package reflects

import "reflect"

// Value struct
type Value struct {
	reflect.Value
	baseKind BKind
}

// Wrap the give value
func Wrap(rv reflect.Value) Value { _ = "STUB: not implemented"; return *new(Value) }

// ValueOf the give value
func ValueOf(v any) Value { _ = "STUB: not implemented"; return *new(Value) }

// Indirect value. alias of the reflect.Indirect()
func (v Value) Indirect() Value { _ = "STUB: not implemented"; return *new(Value) }

// Elem returns the value that the interface v contains or that the pointer v points to.
//
// TIP: not like reflect.Value.Elem. otherwise, will return self.
func (v Value) Elem() Value { _ = "STUB: not implemented"; return *new(Value) }

// otherwise, will return self

// Type of value.
func (v Value) Type() Type { _ = "STUB: not implemented"; return *new(Type) }

// BKind value
func (v Value) BKind() BKind {
	_ = "STUB: not implemented"

	// BaseKind value
	return *new(BKind)
}

func (v Value) BaseKind() BKind {
	_ = "STUB: not implemented"

	// HasChild check. eg: array, slice, map, struct
	return *new(BKind)
}

func (v Value) HasChild() bool { _ = "STUB: not implemented"; return false }

// Int value. if is uintX will convert to int64
func (v Value) Int() int64 { _ = "STUB: not implemented"; return 0 }

// Uint value. if is intX will convert to uint64
func (v Value) Uint() uint64 { _ = "STUB: not implemented"; return 0 }
