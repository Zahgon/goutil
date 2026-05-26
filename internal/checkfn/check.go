package checkfn

import (
	"reflect"
	"regexp"
)

// IsNil value check
func IsNil(v any) bool { _ = "STUB: not implemented"; return false }

// IsSimpleKind kind in: string, bool, intX, uintX, floatX
func IsSimpleKind(k reflect.Kind) bool { _ = "STUB: not implemented"; return false }

// IsEqual determines if two objects are considered equal.
//
// TIP: cannot compare function type
func IsEqual(src, dst any) bool { _ = "STUB: not implemented"; return false }

// Contains try loop over the data check if the data includes the element.
//
// data allow types: string, map, array, slice
//
//	map         - check key exists
//	string      - check sub-string exists
//	array,slice - check sub-element exists
//
// Returns:
//   - valid: data is valid
//   - found: element was found
//
// return (false, false) if impossible.
// return (true, false) if element was not found.
// return (true, true) if element was found.
func Contains(data, elem any) (valid, found bool) { _ = "STUB: not implemented"; return false, false }

// string

// map

// array, slice - other return false

// StringsContains check string slice contains sub-string
func StringsContains(ss []string, sub string) bool { _ = "STUB: not implemented"; return false }

var (
	// check is number: int or float
	numReg = regexp.MustCompile(`^[-+]?\d*\.?\d+$`)
	// is positive number: int or float
	pNumReg = regexp.MustCompile(`^\d*\.?\d+$`)
)

// IsNumeric returns true if the given string is a numeric, otherwise false.
func IsNumeric(s string) bool { _ = "STUB: not implemented"; return false }

// IsPositiveNum check input string is positive number
func IsPositiveNum(s string) bool { _ = "STUB: not implemented"; return false }

// IsHttpURL check input is http/https url
func IsHttpURL(s string) bool { _ = "STUB: not implemented"; return false }

// IndexByteAfter find index of byte after startIndex. return -1 if not found
//
// eg:
//
//	IndexByteAfter("abcabc", 'b', 0) = 1
//	IndexByteAfter("abcabc", 'b', 2) = 4
func IndexByteAfter(s string, b byte, startIndex int) int { _ = "STUB: not implemented"; return 0 }
