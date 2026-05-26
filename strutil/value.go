package strutil

// Value string
type Value string

// StrVal string. alias of Value
type StrVal = Value

// Set value
func (s *Value) Set(val string) error { _ = "STUB: not implemented"; return nil }

// IsEmpty check
func (s Value) IsEmpty() bool { _ = "STUB: not implemented"; return false }

// IsBlank check
func (s Value) IsBlank() bool { _ = "STUB: not implemented"; return false }

// IsStartWith prefix
func (s Value) IsStartWith(sub string) bool { _ = "STUB: not implemented"; return false }

// HasPrefix prefix
func (s Value) HasPrefix(sub string) bool { _ = "STUB: not implemented"; return false }

// IsEndWith suffix
func (s Value) IsEndWith(sub string) bool { _ = "STUB: not implemented"; return false }

// HasSuffix suffix
func (s Value) HasSuffix(sub string) bool { _ = "STUB: not implemented"; return false }

// Bytes string to bytes
func (s Value) Bytes() []byte {
	_ = "STUB: not implemented"

	// Val string
	return nil
}

func (s Value) Val() string {
	_ = "STUB: not implemented"

	// Int convert
	return ""
}

func (s Value) Int() int { _ = "STUB: not implemented"; return 0 }

// Int64 convert
func (s Value) Int64() int64 { _ = "STUB: not implemented"; return 0 }

// Bool convert
func (s Value) Bool() bool { _ = "STUB: not implemented"; return false }

// Value string
func (s Value) String() string {
	_ = "STUB: not implemented"

	// OrElse string
	return ""
}

func (s Value) OrElse(or string) string { _ = "STUB: not implemented"; return "" }

// Split string
func (s Value) Split(sep string) []string { _ = "STUB: not implemented"; return nil }

// SplitN string
func (s Value) SplitN(sep string, n int) []string { _ = "STUB: not implemented"; return nil }

// WithTrimSpace string and return new
func (s Value) WithTrimSpace() Value { _ = "STUB: not implemented"; return *new(Value) }
