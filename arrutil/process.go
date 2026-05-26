package arrutil

import (
	"github.com/gookit/goutil/comdef"
)

// Reverse any T slice.
//
// eg: []string{"site", "user", "info", "0"} -> []string{"0", "info", "user", "site"}
func Reverse[T any](ls []T) { _ = "STUB: not implemented"; return }

// Remove give element from slice []T.
//
// eg: []string{"site", "user", "info", "0"} -> []string{"site", "user", "info"}
func Remove[T comdef.Compared](ls []T, val T) []T { _ = "STUB: not implemented"; return nil }

// Filter given slice, default will filter zero value.
//
// Usage:
//
//	// output: [a, b]
//	ss := arrutil.Filter([]string{"a", "", "b", ""})
func Filter[T any](ls []T, filter ...comdef.MatchFunc[T]) []T {
	_ = "STUB: not implemented"
	return nil
}

// if el == nil { // Filter nil value
// 	return false
// }

// Map a list to new list with map filter function.
//
// eg: mapping [object0{},object1{},...] to flatten list [object0.someKey, object1.someKey, ...]
func Map[T, V any](list []T, mapFilter func(input T) (target V, ok bool)) []V {
	_ = "STUB: not implemented"
	return nil
}

// Map1 a list to new list with map function.
//
// eg: mapping [object0{},object1{},...] to flatten list [object0.someKey, object1.someKey, ...]
func Map1[T, R any](list []T, mapFn func(t T) R) []R { _ = "STUB: not implemented"; return nil }

// Column collect sub elements from list. alias of Map func
//
// Example:
//
//	list := []map[string]any{
//	  {"id": 1, "name": "one", "age": 23},
//	  {"id": 2, "name": "two", "age": 23},
//	  {"id": 3, "name": "three", "age": 23},
//	}
//	names := arrutil.Column(list, func(el map[string]any) string {
//	  return el["name"].(string)
//	})
func Column[T any, V any](list []T, mapFn func(obj T) (val V, find bool)) []V {
	_ = "STUB: not implemented"
	return nil

	// Unique value in the given slice data.
}

func Unique[T comdef.NumberOrString](list []T) []T { _ = "STUB: not implemented"; return nil }

// IndexOf value in given slice.
func IndexOf[T comdef.NumberOrString](val T, list []T) int { _ = "STUB: not implemented"; return 0 }

// FirstOr get first value of slice, if slice is empty, return the default value.
func FirstOr[T any](list []T, defVal ...T) T { _ = "STUB: not implemented"; return *new(T) }

// Chunk split slice to chunks by size.
//
// eg: [1,2,3,4,5,6,7,8,9,10] -> [[1,2,3,4], [5,6,7,8], [9,10]]
func Chunk[T any](list []T, size int) [][]T { _ = "STUB: not implemented"; return nil }

// ChunkBy split slice to chunks by size, and with custom chunk function.
//
// Example:
//
//	  list := []map[string]any{
//	    {"id": 1, "name": "one", "age": 23},
//	    {"id": 2, "name": "two", "age": 23},
//	    {"id": 3, "name": "three", "age": 23},
//	  }
//	  chunks := arrutil.ChunkBy(list, 2, func(el map[string]any) map[string]any {
//	    return map[string]any{
//	      "id": el["id"],
//	      "name": el["name"],
//	    }
//	  })
//		Output: [
//			[{"id": 1, "name": "one"}, {"id": 2, "name": "two"}],
//			[{"id": 3, "name": "three"}]
//		]
func ChunkBy[T, R any](list []T, size int, mapFn func(el T) R) [][]R {
	_ = "STUB: not implemented"
	return nil
}

// 计算需要的块数量

// 创建当前块的切片
