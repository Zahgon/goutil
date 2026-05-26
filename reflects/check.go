package reflects

import (
	"reflect"
)

// IsTimeType check is or alias of time.Time type
func IsTimeType(t reflect.Type) bool { _ = "STUB: not implemented"; return false }

// t == timeType - 无法判断自定义类型

// IsDurationType check is or alias of time.Duration type
func IsDurationType(t reflect.Type) bool { _ = "STUB: not implemented"; return false }

// t == durationType - 无法判断自定义类型

// HasChild type check. eg: array, slice, map, struct
func HasChild(v reflect.Value) bool { _ = "STUB: not implemented"; return false }

// IsArrayOrSlice type check. eg: array, slice
func IsArrayOrSlice(k reflect.Kind) bool { _ = "STUB: not implemented"; return false }

// IsSimpleKind kind in: string, bool, intX, uintX, floatX
func IsSimpleKind(k reflect.Kind) bool { _ = "STUB: not implemented"; return false }

// IsAnyInt check is intX or uintX type. alias of the IsIntLike()
func IsAnyInt(k reflect.Kind) bool { _ = "STUB: not implemented"; return false }

// IsIntLike reports whether the type is int-like(intX, uintX).
func IsIntLike(k reflect.Kind) bool { _ = "STUB: not implemented"; return false }

// IsIntx check is intX type
func IsIntx(k reflect.Kind) bool { _ = "STUB: not implemented"; return false }

// IsUintX check is uintX type
func IsUintX(k reflect.Kind) bool { _ = "STUB: not implemented"; return false }

// IsNil reflect value
func IsNil(v reflect.Value) bool { _ = "STUB: not implemented"; return false }

// IsValidPtr check variable is a valid pointer.
func IsValidPtr(v reflect.Value) bool { _ = "STUB: not implemented"; return false }

// CanBeNil reports whether an untyped nil can be assigned to the type. See reflect.Zero.
func CanBeNil(typ reflect.Type) bool { _ = "STUB: not implemented"; return false }

// IsFunc value
func IsFunc(val any) bool { _ = "STUB: not implemented"; return false }

// IsEqual determines if two objects are considered equal.
//
// TIP: cannot compare a function type
func IsEqual(src, dst any) bool { _ = "STUB: not implemented"; return false }

// IsZero reflect value check, alias of the IsEmpty()
var IsZero = IsEmpty

// IsEmpty reflect value check. if is ptr, check if is nil
func IsEmpty(v reflect.Value) bool { _ = "STUB: not implemented"; return false }

// IsEmptyValue reflect value check, alias of the IsEmptyReal()
var IsEmptyValue = IsEmptyReal

// IsEmptyReal reflect value check.
//
// Note:
//
// Difference the IsEmpty(), if value is ptr or interface, will check real elem.
//
// From src/pkg/encoding/json/encode.go.
func IsEmptyReal(v reflect.Value) bool { _ = "STUB: not implemented"; return false }
