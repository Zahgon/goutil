package arrutil

import (
	"github.com/gookit/goutil/comdef"
)

// Ints type
type Ints[T comdef.Integer] []T

// String to string
func (is Ints[T]) String() string { _ = "STUB: not implemented"; return "" }

// Has given element
func (is Ints[T]) Has(i T) bool { _ = "STUB: not implemented"; return false }

// First element value.
func (is Ints[T]) First(defVal ...T) T { _ = "STUB: not implemented"; return *new(T) }

// Last element value.
func (is Ints[T]) Last(defVal ...T) T { _ = "STUB: not implemented"; return *new(T) }

// Sort the int slice
func (is Ints[T]) Sort() {
	_ = "STUB: not implemented"

	// Len get length
	return
}

func (is Ints[T]) Len() int {
	_ = "STUB: not implemented"

	// Less compare two elements
	return 0
}

func (is Ints[T]) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// Swap elements by indexes
func (is Ints[T]) Swap(i, j int) { _ = "STUB: not implemented"; return }

// Strings type
type Strings []string

// String to string
func (ss Strings) String() string { _ = "STUB: not implemented"; return "" }

// Join to string
func (ss Strings) Join(sep string) string { _ = "STUB: not implemented"; return "" }

// Has given element
func (ss Strings) Has(sub string) bool { _ = "STUB: not implemented"; return false }

// Contains given element
func (ss Strings) Contains(sub string) bool { _ = "STUB: not implemented"; return false }

// First element value.
func (ss Strings) First(defVal ...string) string { _ = "STUB: not implemented"; return "" }

// Last element value.
func (ss Strings) Last(defVal ...string) string { _ = "STUB: not implemented"; return "" }

// Sort the string slice
func (ss Strings) Sort() {
	_ = "STUB: not implemented"

	// SortedList definition for compared type
	return
}

type SortedList[T comdef.Compared] []T

// Len get length
func (ls SortedList[T]) Len() int {
	_ = "STUB: not implemented"

	// Less compare two elements
	return 0
}

func (ls SortedList[T]) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// Swap elements by indexes
func (ls SortedList[T]) Swap(i, j int) { _ = "STUB: not implemented"; return }

// IsEmpty check
func (ls SortedList[T]) IsEmpty() bool { _ = "STUB: not implemented"; return false }

// String to string
func (ls SortedList[T]) String() string { _ = "STUB: not implemented"; return "" }

// Has given element
func (ls SortedList[T]) Has(el T) bool { _ = "STUB: not implemented"; return false }

// Contains given element
func (ls SortedList[T]) Contains(el T) bool { _ = "STUB: not implemented"; return false }

// First element value.
func (ls SortedList[T]) First(defVal ...T) T { _ = "STUB: not implemented"; return *new(T) }

// Last element value.
func (ls SortedList[T]) Last(defVal ...T) T { _ = "STUB: not implemented"; return *new(T) }

// Remove given element
func (ls SortedList[T]) Remove(el T) SortedList[T] { _ = "STUB: not implemented"; return nil }

// Filter the slice, default will filter zero value.
func (ls SortedList[T]) Filter(filter ...comdef.MatchFunc[T]) SortedList[T] {
	_ = "STUB: not implemented"
	return nil
}

// Map the slice to new slice. TODO syntax ERROR: Method cannot have type parameters
// func (ls SortedList[T]) Map[V any](mapFn MapFn[T, V]) SortedList[V] {
// 	return Map(ls, mapFn)
// }

// Sort the slice
func (ls SortedList[T]) Sort() { _ = "STUB: not implemented"; return }
