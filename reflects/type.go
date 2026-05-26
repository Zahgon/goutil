package reflects

import "reflect"

// BKind base data kind type, alias of reflect.Kind
//
// Diff with reflect.Kind:
//   - Int contains all intX types
//   - Uint contains all uintX types
//   - Float contains all floatX types
//   - Array for array and slice types
//   - Complex contains all complexX types
type BKind = reflect.Kind

// base kinds
const (
	// Int for all intX types
	Int = reflect.Int
	// Uint for all uintX types
	Uint = reflect.Uint
	// Float for all floatX types
	Float = reflect.Float32
	// Array for array,slice types
	Array = reflect.Array
	// Complex for all complexX types
	Complex = reflect.Complex64
)

// ToBaseKind convert reflect.Kind to base kind
func ToBaseKind(kind reflect.Kind) BKind {
	_ = "STUB: not implemented"
	return *

	// ToBKind convert reflect.Kind to base kind
	new(BKind)
}

func ToBKind(kind reflect.Kind) BKind { _ = "STUB: not implemented"; return *new(BKind) }

// like: string, map, struct, ptr, func, interface ...

// Type struct
type Type interface {
	reflect.Type
	// BaseKind value
	BaseKind() BKind
	// RealType returns a ptr type's real type. otherwise, will return self.
	RealType() reflect.Type
	// SafeElem returns a type's element type. otherwise, will return self.
	SafeElem() reflect.Type
}

type xType struct {
	reflect.Type
	baseKind BKind
}

// TypeOf value
func TypeOf(v any) Type { _ = "STUB: not implemented"; return *new(Type) }

// BaseKind value
func (t *xType) BaseKind() BKind {
	_ = "STUB: not implemented"

	// RealType returns a ptr type's real type. otherwise, will return self.
	return *new(BKind)
}

func (t *xType) RealType() reflect.Type {
	_ = "STUB: not implemented"
	return *

	// SafeElem returns the array, slice, chan, map type's element type. otherwise, will return self.
	new(reflect.Type)
}

func (t *xType) SafeElem() reflect.Type { _ = "STUB: not implemented"; return *new(reflect.Type) }
