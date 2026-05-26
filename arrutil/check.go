package arrutil

import (
	"github.com/gookit/goutil/comdef"
)

// SliceHas check the slice contains the given value
func SliceHas[T comdef.ScalarType](slice []T, val T) bool { _ = "STUB: not implemented"; return false }

// IntsHas check the []comdef.Integer contains the given value
func IntsHas[T comdef.Integer](ints []T, val T) bool { _ = "STUB: not implemented"; return false }

// Int64sHas check the []int64 contains the given value
func Int64sHas(ints []int64, val int64) bool { _ = "STUB: not implemented"; return false }

// StringsHas check the []string contains the given element
func StringsHas[T ~string](ss []T, val T) bool { _ = "STUB: not implemented"; return false }

// InStrings check elem in the ss. alias of StringsHas()
func InStrings[T ~string](elem T, ss []T) bool { _ = "STUB: not implemented"; return false }

// NotIn check the given value whether not in the list
func NotIn[T comdef.ScalarType](value T, list []T) bool { _ = "STUB: not implemented"; return false }

// In check the given value whether in the list
func In[T comdef.ScalarType](value T, list []T) bool { _ = "STUB: not implemented"; return false }

// ContainsAll check given values is sub-list of sample list.
func ContainsAll[T comdef.ScalarType](list, values []T) bool {
	_ = "STUB: not implemented"
	return false
}

// IsSubList check given values is sub-list of sample list.
func IsSubList[T comdef.ScalarType](values, list []T) bool { _ = "STUB: not implemented"; return false }

// IsParent check given values is parent-list of samples.
func IsParent[T comdef.ScalarType](values, list []T) bool { _ = "STUB: not implemented"; return false }

// HasValue check array(strings, intXs, uintXs) should be contained the given value(int(X),string).
func HasValue(arr, val any) bool { _ = "STUB: not implemented"; return false }

// Contains check slice/array(strings, intXs, uintXs) should be contained the given value(int(X),string).
//
// TIP: Difference the In(), Contains() will try to convert value type,
// and Contains() support array type.
func Contains(arr, val any) bool { _ = "STUB: not implemented"; return false }

// if is string value

// as int value

// NotContains check array(strings, ints, uints) should be not contains the given value.
func NotContains(arr, val any) bool { _ = "STUB: not implemented"; return false }
