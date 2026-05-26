package structs

import (
	"reflect"
)

// ToMap quickly convert structs to map by reflection
func ToMap(st any, optFns ...MapOptFunc) map[string]any { _ = "STUB: not implemented"; return nil }

// MustToMap alis of TryToMap, but will panic on error
func MustToMap(st any, optFns ...MapOptFunc) map[string]any { _ = "STUB: not implemented"; return nil }

// TryToMap simple convert structs to map by reflect
func TryToMap(st any, optFns ...MapOptFunc) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ToSMap quickly and safe convert structs to map[string]string by reflection
func ToSMap(st any, optFns ...MapOptFunc) map[string]string { _ = "STUB: not implemented"; return nil }

// TryToSMap quickly convert structs to map[string]string by reflection
func TryToSMap(st any, optFns ...MapOptFunc) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MustToSMap alias of ToStringMap(), but will panic on error
func MustToSMap(st any, optFns ...MapOptFunc) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// ToString quickly format struct to string
func ToString(st any, optFns ...MapOptFunc) string { _ = "STUB: not implemented"; return "" }

const defaultFieldTag = "json"

// CustomUserFunc for map convert
//   - fName: raw field name in struct
//
// Returns:
//   - ok: return true to collect field, otherwise excluded.
//   - newVal: `newVal != nil` return new value to collect, otherwise collect original value.
type CustomUserFunc func(fName string, fv reflect.Value) (ok bool, newVal any)

// MapOptions for convert struct to map
type MapOptions struct {
	// TagName for map filed. default is "json"
	TagName string
	// ParseDepth for parse and collect. TODO support depth
	ParseDepth int
	// MergeAnonymous struct fields to parent map. default is true
	MergeAnonymous bool
	// ExportPrivate export private fields. default is false
	ExportPrivate bool
	// IgnoreEmpty ignore empty value item. default: false
	IgnoreEmpty bool
	// UserFunc custom interceptor for filter or handle field value.
	UserFunc CustomUserFunc
}

// MapOptFunc define
type MapOptFunc func(opt *MapOptions)

// WithMapTagName set tag name for map field
func WithMapTagName(tagName string) MapOptFunc { _ = "STUB: not implemented"; return *new(MapOptFunc) }

// WithUserFunc custom user func
func WithUserFunc(fn CustomUserFunc) MapOptFunc { _ = "STUB: not implemented"; return *new(MapOptFunc) }

// MergeAnonymous merge anonymous struct fields to parent map
func MergeAnonymous(opt *MapOptions) { _ = "STUB: not implemented"; return }

// ExportPrivate merge anonymous struct fields to parent map
func ExportPrivate(opt *MapOptions) { _ = "STUB: not implemented"; return }

// WithIgnoreEmpty ignore on field value is empty
func WithIgnoreEmpty(opt *MapOptions) { _ = "STUB: not implemented"; return }

// StructToMap quickly convert structs to map[string]any by reflection.
//
// Can custom export field name by tag `json` or custom tag. see MapOptions
func StructToMap(st any, optFns ...MapOptFunc) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func structToMap(obj reflect.Value, opt *MapOptions, mp map[string]any) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// skip un-exported field

// un-exported field

// opt: ignore empty field

// up: special handle time.Time field value

// collect anonymous struct values to parent.

// collect struct values to submap

// TODO support struct slice field.

// up: support custom user func

// ok1=true, newVal != nil

// for unexported field
