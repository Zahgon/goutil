package structs

// Value data store
type Value struct {
	// V value
	V any
}

// NewValue instance.
func NewValue(val any) *Value { _ = "STUB: not implemented"; return nil }

// Set value
func (v *Value) Set(val any) {
	_ = "STUB: not implemented"

	// Reset value
	return
}

func (v *Value) Reset() {
	_ = "STUB: not implemented"

	// Val get
	return
}

func (v *Value) Val() any {
	_ = "STUB: not implemented"

	// Val get
	//
	//	func (v *Value) ValOr[T any](defVal T) T {
	//		return v.V
	//	}
	return *new(any)
}

// Int value get
func (v *Value) Int() int { _ = "STUB: not implemented"; return 0 }

// Int64 value
func (v *Value) Int64() int64 { _ = "STUB: not implemented"; return 0 }

// Bool value
func (v *Value) Bool() bool { _ = "STUB: not implemented"; return false }

// Float64 value
func (v *Value) Float64() float64 { _ = "STUB: not implemented"; return 0 }

// String value
func (v *Value) String() string { _ = "STUB: not implemented"; return "" }

// Strings value
func (v *Value) Strings() (ss []string) { _ = "STUB: not implemented"; return nil }

// SplitToStrings split string value to strings. sep default is comma(,)
func (v *Value) SplitToStrings(sep ...string) (ss []string) { _ = "STUB: not implemented"; return nil }

// SplitToInts split string value to []int. sep default is comma(,)
func (v *Value) SplitToInts(sep ...string) (ss []int) { _ = "STUB: not implemented"; return nil }

// IsEmpty value
func (v *Value) IsEmpty() bool { _ = "STUB: not implemented"; return false }

func sepStr(seps []string) string { _ = "STUB: not implemented"; return "" }
