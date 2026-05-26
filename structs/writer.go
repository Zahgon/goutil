package structs

import (
	"reflect"
)

// NewWriter create a struct writer
//
// TIP: must be pointer for set field value
func NewWriter(ptr any) *Wrapper { _ = "STUB: not implemented"; return nil }

/*************************************************************
 * set values to a struct
 *************************************************************/

// SetOptFunc define
type SetOptFunc func(opt *SetOptions)

// BeforeSetFunc hook type.
type BeforeSetFunc func(fieldName string, value any, fv reflect.Value) any

// SetOptions for set values to struct
type SetOptions struct {
	// FieldTagName get field name for read value. default tag: json
	FieldTagName string
	// BeforeSetFn hook func. will fire on before set value.
	//  - you can modify value here.
	//  - returns nil will skip set value.
	//  - returns value will be set to field value.
	BeforeSetFn BeforeSetFunc

	// ParseTime parse string to `time.Duration`, `time.Time`. default: false
	//
	// eg: default:"10s", default:"2025-04-23 15:04:05"
	ParseTime bool
	// ParseDefault init default value by DefaultValTag tag value. default: false
	//
	// see InitDefaults()
	ParseDefault bool
	// DefaultValTag name. tag: default
	DefaultValTag string

	// ParseDefaultEnv parse env var on default tag. eg: `default:"${APP_ENV}"` default: false
	ParseDefaultEnv bool
	// DefaultEnvPrefixTag name. tag: defaultenvprefix
	DefaultEnvPrefixTag string

	// StopOnError if true, will stop set value on error happened. default: false
	// StopOnError bool
}

// WithParseDefault value by tag "default"
func WithParseDefault(opt *SetOptions) { _ = "STUB: not implemented"; return }

// WithBeforeSetFn value by tag "default"
func WithBeforeSetFn(fn BeforeSetFunc) SetOptFunc {
	_ = "STUB: not implemented"
	return *new(SetOptFunc)
}

// TODO refactoring SetValues to the struct
type ValuesSetter struct {
	src any // source struct
	// raw reflect.Value of source struct
	rv reflect.Value

	option  *SetOptions
	initOpt *InitOptions
}

// BindData set values to struct ptr from map data.
func BindData(ptr any, data map[string]any, optFns ...SetOptFunc) error {
	_ = "STUB: not implemented"
	return nil
}

// SetValues set values to struct ptr from map data.
//
// TIPS:
//
//	Only support set: string, bool, intX, uintX, floatX
func SetValues(ptr any, data map[string]any, optFns ...SetOptFunc) error {
	_ = "STUB: not implemented"
	return nil
}

func setValues(rv reflect.Value, data map[string]any, opt *SetOptions, envPrefix string) error {
	_ = "STUB: not implemented"
	return nil
}

// skip don't exported field

// get field name

// set field value by default tag.

// handle for pointer field

// hook: call BeforeSet func

// on value is not nil

// field is struct

// up: special handle time.Time field

// maybe val is time.Time

// val is datetime string

// val as map

// recursive processing sub-struct

// set field value
