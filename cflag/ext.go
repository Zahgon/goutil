package cflag

import (
	"github.com/gookit/goutil/comdef"
	"github.com/gookit/goutil/maputil"
)

// RepeatableFlag interface.
type RepeatableFlag interface {
	// IsRepeatable mark option flag can be set multi times
	IsRepeatable() bool
}

// ExtendedFlagType interface.
type ExtendedFlagType interface {
	// FlagTypeDesc flag type description. use for enum and more custom types.
	FlagTypeDesc() string
}

/*************************************************************************
 * options: some special flag vars
 * - implemented flag.Value interface
 *************************************************************************/

// LimitInt limit int value range
func LimitInt(min, max int) comdef.IntCheckFunc {
	_ = "STUB: not implemented"
	return *new(comdef.IntCheckFunc)
}

// IntVar int value can with a check func
//
// Limit min and max value:
//
//	iv := cflag.IntValue{CheckFn: cflag.LimitInt(1, 10)}
//	fs.IntVar(&iv, "int", 1, "the int value")
type IntVar struct {
	val int
	str string
	// check func
	CheckFn comdef.IntCheckFunc
}

// NewIntVar create a new IntVar instance with check func
func NewIntVar(checkFn comdef.IntCheckFunc) IntVar { _ = "STUB: not implemented"; return *new(IntVar) }

// Get value
func (o *IntVar) Get() any {
	_ = "STUB: not implemented"

	// Set new value
	return *new(any)
}

func (o *IntVar) Set(value string) error { _ = "STUB: not implemented"; return nil }

// String value get
func (o *IntVar) String() string {
	_ = "STUB: not implemented"

	// String a special string
	//
	// Usage:
	//
	//	// case 1:
	//	var names cflag.String
	//	c.VarOpt(&names, "names", "", "multi name by comma split")
	//
	//	--names "tom,john,joy"
	//	names.Split(",") // -> []string{"tom","john","joy"}
	//
	//	// case 2:
	//	var ids cflag.String
	//	c.VarOpt(&ids, "ids", "", "multi id by comma split")
	//
	//	--names "23,34,56"
	//	names.Ints(",") // -> []int{23,34,56}
	return ""
}

type String string

// Get value
func (s *String) Get() any {
	_ = "STUB: not implemented"

	// Set value
	return *new(any)
}

func (s *String) Set(val string) error { _ = "STUB: not implemented"; return nil }

// String input value to string
func (s *String) String() string {
	_ = "STUB: not implemented"

	// Strings split value to []string by sep ','
	return ""
}

func (s *String) Strings() []string { _ = "STUB: not implemented"; return nil }

// Split value to []string
func (s *String) Split(sep string) []string { _ = "STUB: not implemented"; return nil }

// Ints value to []int
func (s *String) Ints(sep string) []int { _ = "STUB: not implemented"; return nil }

// StrVar string value can with a check func
//
// Usage:
//
//	sv := cflag.StrVar{}
//	fs.Var(&sv, "str", "the string value")
type StrVar struct {
	val string
	// check func
	CheckFn comdef.StrCheckFunc
}

// NewStrVar create a new StrVar with check func
func NewStrVar(checkFn comdef.StrCheckFunc) StrVar { _ = "STUB: not implemented"; return *new(StrVar) }

// Get value string
func (o *StrVar) Get() any {
	_ = "STUB: not implemented"

	// Set new value
	return *new(any)
}

func (o *StrVar) Set(value string) error { _ = "STUB: not implemented"; return nil }

// String value get
func (o *StrVar) String() string {
	_ = "STUB: not implemented"

	// IntsString The ints-string flag. eg: --get 1,2,3
	//
	// Implemented the flag.Value interface
	return ""
}

type IntsString struct {
	ints []int
	// value and size validate
	ValueFn func(val int) error
	SizeFn  func(ln int) error
}

// String input value to string
func (o *IntsString) String() string { _ = "STUB: not implemented"; return "" }

// Get value
func (o *IntsString) Get() any {
	_ = "STUB: not implemented"

	// Ints value
	return *new(any)
}

func (o *IntsString) Ints() []int {
	_ = "STUB: not implemented"

	// Set new value. eg: "12"
	return nil
}

func (o *IntsString) Set(value string) error { _ = "STUB: not implemented"; return nil }

// Ints The int flag list, repeatable
//
// implemented flag.Value interface
type Ints []int

// String to string
func (s *Ints) String() string { _ = "STUB: not implemented"; return "" }

// Get value
func (s *Ints) Get() any {
	_ = "STUB: not implemented"

	// Set new value
	return *new(any)
}

func (s *Ints) Set(value string) error { _ = "STUB: not implemented"; return nil }

// Ints value
func (s *Ints) Ints() []int {
	_ = "STUB: not implemented"

	// IsRepeatable on input
	return nil
}

func (s *Ints) IsRepeatable() bool {
	_ = "STUB: not implemented"

	// Strings The string flag list, repeatable.
	// eg: --names tom --names john
	return false
}

type Strings []string

// String input value to string
func (s *Strings) String() string { _ = "STUB: not implemented"; return "" }

// Get value
func (s *Strings) Get() any {
	_ = "STUB: not implemented"

	// Set new value
	return *new(any)
}

func (s *Strings) Set(value string) error { _ = "STUB: not implemented"; return nil }

// Strings value
func (s *Strings) Strings() []string {
	_ = "STUB: not implemented"

	// IsRepeatable on input
	return nil
}

func (s *Strings) IsRepeatable() bool {
	_ = "STUB: not implemented"

	// Booleans The bool flag list, repeatable.
	// eg: -v -v => []bool{true, true}
	return false
}

type Booleans []bool

// String input value to string
func (s *Booleans) String() string { _ = "STUB: not implemented"; return "" }

// Bools value
func (s *Booleans) Bools() []bool {
	_ = "STUB: not implemented"

	// Set new value
	return nil
}

func (s *Booleans) Set(value string) error { _ = "STUB: not implemented"; return nil }

// IsRepeatable on input
func (s *Booleans) IsRepeatable() bool {
	_ = "STUB: not implemented"

	// EnumString limit input value is in the enum list.
	// implemented flag.Value interface
	//
	// Usage:
	//
	//	var enumStr = cflag.NewEnumString("php", "go", "java")
	//	c.VarOpt(&enumStr, "lang", "", "input language name")
	return false
}

type EnumString struct {
	val  string
	enum []string
}

// NewEnumString instance
func NewEnumString(enum ...string) EnumString { _ = "STUB: not implemented"; return *new(EnumString) }

// Get value
func (s *EnumString) Get() any {
	_ = "STUB: not implemented"

	// String input value to string
	return *new(any)
}

func (s *EnumString) String() string {
	_ = "STUB: not implemented"

	// SetEnum values
	return ""
}

func (s *EnumString) SetEnum(enum []string) {
	_ = "STUB: not implemented"

	// WithEnum values
	return
}

func (s *EnumString) WithEnum(enum []string) *EnumString { _ = "STUB: not implemented"; return nil }

// EnumString to string
func (s *EnumString) EnumString() string { _ = "STUB: not implemented"; return "" }

// Set new value, will check value is right
func (s *EnumString) Set(value string) error { _ = "STUB: not implemented"; return nil }

// Enum to string
func (s *EnumString) Enum() []string {
	_ = "STUB: not implemented"

	// FlagTypeDesc message. will display on the flag description end.
	return nil
}

func (s *EnumString) FlagTypeDesc() string { _ = "STUB: not implemented"; return "" }

// KVString The kv-string flag, allow input multi.
//
// Implemented the flag.Value interface.
//
// Usage:
//
//		type myOpts struct {
//			vars cflag.KVString
//		}
//	 var mo &myOpts{ vars: cflag.NewKVString() }
//
// Example:
//
//	--var name=inhere => string map {name:inhere}
//	--var name=inhere --var age=234 => string map {name:inhere, age:234}
type KVString struct {
	maputil.SMap
	Sep string // Default: "="
}

// KVStrMap alias for KVString
type KVStrMap = KVString

// NewKVString instance
func NewKVString() KVString { _ = "STUB: not implemented"; return *new(KVString) }

// Init settings
func (s *KVString) Init() *KVString { _ = "STUB: not implemented"; return nil }

// Get value
func (s *KVString) Get() any {
	_ = "STUB: not implemented"

	// Data map get
	return *new(any)
}

func (s *KVString) Data() maputil.SMap {
	_ = "STUB: not implemented"
	return *

	// Set new value, will check value is right
	new(maputil.SMap)
}

func (s *KVString) Set(value string) error { _ = "STUB: not implemented"; return nil }

// IsRepeatable on input
func (s *KVString) IsRepeatable() bool {
	_ = "STUB: not implemented"

	// ConfString The config-string flag, INI format, like nginx-config.
	//
	// Example:
	//
	//	--config 'k0=val0;k1=val1' => string map {k0:val0, k1:val1}
	return false
}

type ConfString struct {
	maputil.SMap
	val string
}

// String to string
func (s *ConfString) String() string {
	_ = "STUB: not implemented"

	// SetData value
	return ""
}

func (s *ConfString) SetData(mp map[string]string) {
	_ = "STUB: not implemented"

	// Data map get
	return
}

func (s *ConfString) Data() maputil.SMap {
	_ = "STUB: not implemented"

	// Get value
	return *new(maputil.SMap)
}

func (s *ConfString) Get() any {
	_ = "STUB: not implemented"

	// Set new value, will check value is right
	return *new(any)
}

func (s *ConfString) Set(value string) error { _ = "STUB: not implemented"; return nil }

// SafeFuncVar safe func Value
type SafeFuncVar func(string)

// Set value
func (f SafeFuncVar) Set(s string) error {
	_ = "STUB: not implemented"

	// String get
	return nil
}

func (f SafeFuncVar) String() string { _ = "STUB: not implemented"; return "" }
