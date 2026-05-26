package structs

import (
	"reflect"
)

// Wrapper struct for read or set field value
type Wrapper struct {
	// src any // source struct

	// raw reflect.Value of source struct
	rv reflect.Value

	// FieldTagName field name for read/write value. default tag: json
	FieldTagName string

	// caches for field rv and name and tag name TODO
	fieldNames []string                 //lint:ignore U1000 for unused
	fvCacheMap map[string]reflect.Value //lint:ignore U1000 for unused
}

// Wrap quick create a struct wrapper
func Wrap(src any) *Wrapper { _ = "STUB: not implemented"; return nil }

// NewWrapper create a struct wrapper
func NewWrapper(src any) *Wrapper { _ = "STUB: not implemented"; return nil }

// WrapValue create a struct wrapper
func WrapValue(rv reflect.Value) *Wrapper { _ = "STUB: not implemented"; return nil }

// Get field value by name, name allows to use dot syntax.
func (r *Wrapper) Get(name string) any { _ = "STUB: not implemented"; return *new(any) }

// Lookup field value by name, name allows to use dot syntax.
func (r *Wrapper) Lookup(name string) (val any, ok bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

// Set field value by name. name allows using dot syntax.
func (r *Wrapper) Set(name string, val any) error { _ = "STUB: not implemented"; return nil }
