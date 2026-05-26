package structs

import (
	"sync"

	"github.com/gookit/goutil/maputil"
)

// LiteData simple map[string]any struct. no lock
type LiteData = Data

// NewLiteData create, not locked
func NewLiteData(data map[string]any) *Data { _ = "STUB: not implemented"; return nil }

/*************************************************************
 * data struct and allow enable lock
 *************************************************************/

// Data struct, allow enable lock
type Data struct {
	sync.RWMutex
	lock bool
	data map[string]any
}

// NewData create new data instance
func NewData() *Data { _ = "STUB: not implemented"; return nil }

// WithLock for operate data
func (d *Data) WithLock() *Data { _ = "STUB: not implemented"; return nil }

// EnableLock for operate data
func (d *Data) EnableLock() *Data { _ = "STUB: not implemented"; return nil }

// Data get all
func (d *Data) Data() map[string]any {
	_ = "STUB: not implemented"

	// SetData set all data
	return nil
}

func (d *Data) SetData(data map[string]any) { _ = "STUB: not implemented"; return }

// DataLen of data
func (d *Data) DataLen() int {
	_ = "STUB: not implemented"

	// ResetData all data
	return 0
}

func (d *Data) ResetData() { _ = "STUB: not implemented"; return }

// Merge load new data
func (d *Data) Merge(mp map[string]any) { _ = "STUB: not implemented"; return }

// Set value to data
func (d *Data) Set(key string, val any) { _ = "STUB: not implemented"; return }

// SetValue to data
func (d *Data) SetValue(key string, val any) { _ = "STUB: not implemented"; return }

// Value get from data
func (d *Data) Value(key string) (val any, ok bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

// Get val from data
func (d *Data) Get(key string) any {
	_ = "STUB: not implemented"
	return *

	// GetVal get from data
	new(any)
}

func (d *Data) GetVal(key string) any { _ = "STUB: not implemented"; return *new(any) }

// StrVal get from data
func (d *Data) StrVal(key string) string { _ = "STUB: not implemented"; return "" }

// IntVal get from data
func (d *Data) IntVal(key string) int { _ = "STUB: not implemented"; return 0 }

// BoolVal get from data
func (d *Data) BoolVal(key string) bool { _ = "STUB: not implemented"; return false }

// String format data
func (d *Data) String() string { _ = "STUB: not implemented"; return "" }

// OrderedData data TODO
type OrderedData struct {
	maputil.Data
	cap  int
	keys []string
	// vals []any
}

// NewOrderedData instance.
func NewOrderedData(cap int) *OrderedData { _ = "STUB: not implemented"; return nil }

// Load data
func (om *OrderedData) Load(data map[string]any) { _ = "STUB: not implemented"; return }

// Set key and value to map
func (om *OrderedData) Set(key string, val any) { _ = "STUB: not implemented"; return }
