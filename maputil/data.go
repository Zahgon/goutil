package maputil

// Map alias of Data
type Map = Data

// Data alias of map[string]any
type Data map[string]any

// Has value on the data map
func (d Data) Has(key string) bool { _ = "STUB: not implemented"; return false }

// IsEmpty if the data map
func (d Data) IsEmpty() bool {
	_ = "STUB: not implemented"

	// endregion
	// region T: set value(s)
	return false
}

// Set value to the data map
func (d Data) Set(key string, val any) {
	_ = "STUB: not implemented"

	// SetByPath sets a value in the map.
	// Supports dot syntax to set deep values.
	//
	// Example:
	//
	//	d.SetByPath("name.first", "Mat")
	return
}

func (d Data) SetByPath(path string, value any) error { _ = "STUB: not implemented"; return nil }

// SetByKeys sets a value in the map by path keys.
// Supports dot syntax to set deep values.
//
// Example:
//
//	d.SetByKeys([]string{"name", "first"}, "Mat")
func (d Data) SetByKeys(keys []string, value any) error { _ = "STUB: not implemented"; return nil }

// special handle d is empty.

// It's ok, but use `func (d *Data)`
// return SetByKeys((*map[string]any)(d), keys, value)

//
// endregion
// region T: read value(s)
//

// Value get from the data map
func (d Data) Value(key string) (any, bool) { _ = "STUB: not implemented"; return *new(any), false }

// Get value from the data map.
// Supports dot syntax to get deep values. eg: top.sub
func (d Data) Get(key string) any { _ = "STUB: not implemented"; return *new(any) }

// One get value from the data by multi paths. will return first founded value
func (d Data) One(keys ...string) any { _ = "STUB: not implemented"; return *new(any) }

// TryOne get value from the data by multi paths. will return first founded value
func (d Data) TryOne(keys ...string) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

// GetByPath get value from the data map by path. eg: top.sub
// Supports dot syntax to get deep values.
func (d Data) GetByPath(path string) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

// is a key path.

// Default get value from the data map with default value
func (d Data) Default(key string, def any) any { _ = "STUB: not implemented"; return *new(any) }

// Int value get, or default value
func (d Data) Int(key string, defVal ...int) int { _ = "STUB: not implemented"; return 0 }

// Int64 value get, or default value
func (d Data) Int64(key string, defVal ...int64) int64 { _ = "STUB: not implemented"; return 0 }

// Uint value get, or default value
func (d Data) Uint(key string, defVal ...uint) uint { _ = "STUB: not implemented"; return 0 }

// Uint16 value get, or default value
func (d Data) Uint16(key string, defVal ...uint16) uint16 { _ = "STUB: not implemented"; return 0 }

// Uint64 value get, or default value
func (d Data) Uint64(key string, defVal ...uint64) uint64 { _ = "STUB: not implemented"; return 0 }

// Str value gets by key, or default value
func (d Data) Str(key string, defVal ...string) string { _ = "STUB: not implemented"; return "" }

// StrOne value gets by multi keys, will return first value
func (d Data) StrOne(keys ...string) string { _ = "STUB: not implemented"; return "" }

// Bool value get
func (d Data) Bool(key string) bool { _ = "STUB: not implemented"; return false }

// BoolOne value gets from multi keys, return first value
func (d Data) BoolOne(keys ...string) bool { _ = "STUB: not implemented"; return false }

// StringsOne get []string value by multi keys, return first founded value
func (d Data) StringsOne(keys ...string) []string { _ = "STUB: not implemented"; return nil }

// Strings get []string value by key
func (d Data) Strings(key string) []string { _ = "STUB: not implemented"; return nil }

// StrSplit get strings by split string value
func (d Data) StrSplit(key, sep string) []string { _ = "STUB: not implemented"; return nil }

// StringsByStr value gets by key, will split string value by ","
func (d Data) StringsByStr(key string) []string { _ = "STUB: not implemented"; return nil }

// StrMap get map[string]string value
func (d Data) StrMap(key string) map[string]string { _ = "STUB: not implemented"; return nil }

// StringMap get map[string]string value
func (d Data) StringMap(key string) map[string]string { _ = "STUB: not implemented"; return nil }

// Sub get sub value(map[string]any) as new Data
func (d Data) Sub(key string) Data { _ = "STUB: not implemented"; return *new(Data) }

// AnyMap get sub value as map[string]any
func (d Data) AnyMap(key string) map[string]any { _ = "STUB: not implemented"; return nil }

// AnyMap get sub value as map[string]any
func (d Data) toAnyMap(val any) map[string]any { _ = "STUB: not implemented"; return nil }

// Slice get []any value from data map
func (d Data) Slice(key string) ([]any, error) { _ = "STUB: not implemented"; return nil, nil }

// Keys of the data map
func (d Data) Keys() []string { _ = "STUB: not implemented"; return nil }

// ToStringMap convert to map[string]string
func (d Data) ToStringMap() map[string]string { _ = "STUB: not implemented"; return nil }

// String data to string
func (d Data) String() string {
	_ = "STUB: not implemented"

	// Load other data to current data map
	return ""
}

func (d Data) Load(sub map[string]any) { _ = "STUB: not implemented"; return }

// LoadSMap to data
func (d Data) LoadSMap(smp map[string]string) { _ = "STUB: not implemented"; return }
