package structs

import (
	"reflect"
)

const defaultInitTag = "default"
const defaultEnvPrefixTag = "defaultenvprefix"

// InitOptFunc define
type InitOptFunc func(opt *InitOptions)

// InitOptions struct
type InitOptions struct {
	// TagName default value tag name. tag: default
	TagName string
	// EnvPrefix default ENV prefix name.
	EnvPrefix string
	// EnvPrefixTagName default value tag name. tag: defaultenvprefix
	EnvPrefixTagName string
	// ParseEnv var name on default value. eg: `default:"${APP_ENV}"`
	//
	// default: false
	ParseEnv bool
	// ParseTime parse string to `time.Duration`, `time.Time`. default: false
	//
	// eg: default:"10s", default:"2025-04-23 15:04:05"
	ParseTime bool
	// ValueHook before set value hook TODO
	ValueHook func(val string) any
}

// WithParseTime set parse time string on default value.
func (opt *InitOptions) WithParseTime(val bool) *InitOptions { _ = "STUB: not implemented"; return nil }

// WithParseEnv set parse env var on default value.
func (opt *InitOptions) WithParseEnv(val bool) *InitOptions { _ = "STUB: not implemented"; return nil }

// Init struct default value by field "default" tag.
func Init(ptr any, optFns ...InitOptFunc) error { _ = "STUB: not implemented"; return nil }

// InitDefaults init struct default value by field "default" tag.
//
// TIPS:
//
//	Support init field types: string, bool, intX, uintX, floatX, array, slice
//
// Example:
//
//	type User1 struct {
//		Name string `default:"inhere"`
//		Age  int32  `default:"30"`
//	}
//
//	u1 := &User1{}
//	err = structs.InitDefaults(u1)
//	fmt.Printf("%+v\n", u1) // Output: {Name:inhere Age:30}
func InitDefaults(ptr any, optFns ...InitOptFunc) error { _ = "STUB: not implemented"; return nil }

// type InitBuilder struct {
// 	opt InitOptions
// }

func initDefaults(rv reflect.Value, opt *InitOptions, envPrefix string) error {
	_ = "STUB: not implemented"
	return nil
}

// skip doesn't exported field

// opt.EnvPrefix = childPrefixVar

// special: struct is time.Time type

// Skip init on field has value. but will check slice and pointer field

// special: handle for pointer struct field

// init sub struct in slice. like `[]SubStruct` or `[]*SubStruct`

// handle for pointer field

// init sub struct in slice. like `[]SubStruct` or `[]*SubStruct`

// up: if slice elem is struct and slice len=0, will be skip init default value

// make sub-struct and init. like: `SubStruct`

// make new slice and set value.

func enhanceDefaultVar(val string, envPrefix string) string { _ = "STUB: not implemented"; return "" }

func initDefaultValue(fv reflect.Value, val string, opt *InitOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// parse env var

// enhance: parse special value type. eg: time.Duration, time.Time

// simple slice: convert simple kind(string,intX,uintX,...) to slice. eg: "1,2,3" => []int{1,2,3}

// set value
