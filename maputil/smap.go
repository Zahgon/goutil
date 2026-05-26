package maputil

// SM is alias of map[string]string
type SM = StrMap

// SMap and StrMap is alias of map[string]string
type SMap = StrMap
type StrMap map[string]string

// IsEmpty of the data map
func (m StrMap) IsEmpty() bool {
	_ = "STUB: not implemented"

	// Has key on the data map
	return false
}

func (m StrMap) Has(key string) bool { _ = "STUB: not implemented"; return false }

// HasValue on the data map
func (m StrMap) HasValue(val string) bool { _ = "STUB: not implemented"; return false }

// Load data to the map
func (m StrMap) Load(data map[string]string) { _ = "STUB: not implemented"; return }

// Set value to the data map
func (m StrMap) Set(key string, val any) { _ = "STUB: not implemented"; return }

// Value get from the data map
func (m StrMap) Value(key string) (string, bool) { _ = "STUB: not implemented"; return "", false }

// Default get value by key. if not found, return defVal
func (m StrMap) Default(key, defVal string) string { _ = "STUB: not implemented"; return "" }

// Get value by key
func (m StrMap) Get(key string) string {
	_ = "STUB: not implemented"

	// Int value get
	return ""
}

func (m StrMap) Int(key string) int { _ = "STUB: not implemented"; return 0 }

// Int64 value get
func (m StrMap) Int64(key string) int64 { _ = "STUB: not implemented"; return 0 }

// Str value get
func (m StrMap) Str(key string) string {
	_ = "STUB: not implemented"

	// StrOne get first founded value by keys
	return ""
}

func (m SMap) StrOne(keys ...string) string { _ = "STUB: not implemented"; return "" }

// Bool value get
func (m StrMap) Bool(key string) bool { _ = "STUB: not implemented"; return false }

// Ints value to []int
func (m StrMap) Ints(key string) []int { _ = "STUB: not implemented"; return nil }

// Strings value to []string
func (m StrMap) Strings(key string) (ss []string) { _ = "STUB: not implemented"; return nil }

// IfExist key, then call the fn with value.
func (m StrMap) IfExist(key string, fn func(val string)) { _ = "STUB: not implemented"; return }

// IfValid value is not empty, then call the fn
func (m StrMap) IfValid(key string, fn func(val string)) { _ = "STUB: not implemented"; return }

// Keys of the string-map
func (m StrMap) Keys() []string { _ = "STUB: not implemented"; return nil }

// Values of the string-map
func (m StrMap) Values() []string { _ = "STUB: not implemented"; return nil }

// ToKVPairs slice convert. eg: {k1:v1,k2:v2} => {k1,v1,k2,v2}
func (m StrMap) ToKVPairs() []string { _ = "STUB: not implemented"; return nil }

// String data to string
func (m StrMap) String() string { _ = "STUB: not implemented"; return "" }
