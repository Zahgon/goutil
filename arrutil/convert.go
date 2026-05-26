package arrutil

import (
	"errors"

	"github.com/gookit/goutil/comdef"
)

// ErrInvalidType error
var ErrInvalidType = errors.New("the input param type is invalid")

/*************************************************************
 * Join func for slice
 *************************************************************/

// JoinStrings alias of strings.Join
func JoinStrings(sep string, ss ...string) string { _ = "STUB: not implemented"; return "" }

// StringsJoin alias of strings.Join
func StringsJoin(sep string, ss ...string) string { _ = "STUB: not implemented"; return "" }

// JoinTyped join typed []T slice to string.
//
// Usage:
//
//	JoinTyped(",", 1,2,3) // "1,2,3"
//	JoinTyped(",", "a","b","c") // "a,b,c"
//	JoinTyped[any](",", "a",1,"c") // "a,1,c"
func JoinTyped[T any](sep string, arr ...T) string { _ = "STUB: not implemented"; return "" }

// JoinSlice join []any slice to string.
func JoinSlice(sep string, arr ...any) string { _ = "STUB: not implemented"; return "" }

/*************************************************************
 * convert func for ints
 *************************************************************/

// IntsToString convert []T to string
func IntsToString[T comdef.Integer](ints []T) string { _ = "STUB: not implemented"; return "" }

// ToInt64s convert any(allow: array,slice) to []int64
func ToInt64s(arr any) (ret []int64, err error) { _ = "STUB: not implemented"; return nil, nil }

// MustToInt64s convert any(allow: array,slice) to []int64
func MustToInt64s(arr any) []int64 { _ = "STUB: not implemented"; return nil }

// SliceToInt64s convert []any to []int64
func SliceToInt64s(arr []any) []int64 { _ = "STUB: not implemented"; return nil }

// ToMap convert a list to new map.
//
// Example:
//
//	 type User struct {
//			Name string
//			Age  int
//		}
//		users := []User{{"Tom", 18}, {"Jack", 20}}
//		mp := arrutil.ToMap(users, func(u User) (string, int) {
//			return u.Name, u.Age
//		})
//	 // mp = map[string]int{"Tom":18, "Jack":20}
func ToMap[T any, K comdef.ScalarType, V any](list []T, mapFn func(T) (K, V)) map[K]V {
	_ = "STUB: not implemented"
	return nil
}

/*************************************************************
 * convert func for any-slice
 *************************************************************/

// AnyToSlice convert any(allow: array,slice) to []any
func AnyToSlice(sl any) (ls []any, err error) { _ = "STUB: not implemented"; return nil, nil }

// AnyToStrings convert array or slice to []string
func AnyToStrings(arr any) []string { _ = "STUB: not implemented"; return nil }

// MustToStrings convert array or slice to []string
func MustToStrings(arr any) []string { _ = "STUB: not implemented"; return nil }

// ToStrings convert any(allow: array,slice) to []string
func ToStrings(arr any) (ret []string, err error) {
	_ = "STUB: not implemented"
	// try direct convert
	return nil, nil
}

// try use reflect to convert

// SliceToStrings safe convert []any to []string
func SliceToStrings(arr []any) []string { _ = "STUB: not implemented"; return nil }

// QuietStrings safe convert []any to []string
func QuietStrings(arr []any) []string { _ = "STUB: not implemented"; return nil }

// ConvType convert type of slice elements to new type slice, by the given newElemTyp type.
//
// Supports conversion between []string, []intX, []uintX, []floatX.
//
// Usage:
//
//	ints, _ := arrutil.ConvType([]string{"12", "23"}, 1) // []int{12, 23}
func ConvType[T any, R any](arr []T, newElemTyp R) ([]R, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// type is same.

// need conv type.

// AnyToString simple and quickly convert any array, slice to string
func AnyToString(arr any) string { _ = "STUB: not implemented"; return "" }

// SliceToString convert []any to string
func SliceToString(arr ...any) string { _ = "STUB: not implemented"; return "" }

// ToString simple and quickly convert []T to string
func ToString[T any](arr []T) string {
	_ = "STUB: not implemented"
	// like fmt.Println([]any(nil))
	return ""
}

// CombineToMap combine []K and []V slice to map[K]V.
//
// If keys length is greater than values, the extra keys will be ignored.
func CombineToMap[K comdef.SortedType, V any](keys []K, values []V) map[K]V {
	_ = "STUB: not implemented"
	return nil
}

// CombineToSMap combine two string-slice to map[string]string
func CombineToSMap(keys, values []string) map[string]string { _ = "STUB: not implemented"; return nil }
