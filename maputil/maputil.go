// Package maputil provide map data util functions. eg: convert, sub-value get, simple merge
package maputil

// Key, value sep char consts
const (
	ValSepStr  = ","
	ValSepChar = ','
	KeySepStr  = "."
	KeySepChar = '.'
)

// Copy copies all key/value pairs in src adding them to dst.
// When a key in src is already present in dst,
// the value in dst will be overwritten by the value associated
// with the key in src.
func Copy[M1 ~map[K]V, M2 ~map[K]V, K comparable, V any](dst M1, src M2) {
	_ = "STUB: not implemented"
	return
}

// DeleteFunc deletes any key/value pairs from m for which del returns true.
func DeleteFunc[M ~map[K]V, K comparable, V any](m M, del func(K, V) bool) {
	_ = "STUB: not implemented"
	return
}

// SimpleMerge simple merge two data map by string key. will merge the src to dst map
func SimpleMerge(src, dst map[string]any) map[string]any { _ = "STUB: not implemented"; return nil }

// simple merge

// Merge1level merge multi any map[string]any data. only merge one level data.
func Merge1level(mps ...map[string]any) map[string]any { _ = "STUB: not implemented"; return nil }

// func DeepMerge(src, dst map[string]any, deep int) map[string]any { TODO
// }

// MergeSMap simple merge two string map. merge src to dst map
func MergeSMap(src, dst map[string]string, ignoreCase bool) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// MergeStrMap simple merge two string map. merge src to dst map
func MergeStrMap(src, dst map[string]string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// AppendSMap append string map data to dst map.
func AppendSMap(dst, src map[string]string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// MergeStringMap simple merge two string map. merge src to dst map
func MergeStringMap(src, dst map[string]string, ignoreCase bool) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// MergeMultiSMap quick merge multi string-map data.
func MergeMultiSMap(mps ...map[string]string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// MergeL2StrMap merge multi level2 string-map data. The back map covers the front.
func MergeL2StrMap(mps ...map[string]map[string]string) map[string]map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// merge level 2 value

// FilterSMap filter empty elem for the string map.
func FilterSMap(sm map[string]string) map[string]string { _ = "STUB: not implemented"; return nil }

// MakeByPath build new value by key names
//
// Example:
//
//	"site.info"
//	->
//	map[string]any {
//		site: {info: val}
//	}
//
//	// case 2, last key is slice:
//	"site.tags[1]"
//	->
//	map[string]any {
//		site: {tags: [val]}
//	}
func MakeByPath(path string, val any) (mp map[string]any) { _ = "STUB: not implemented"; return nil }

// MakeByKeys build new value by key names
//
// Example:
//
//	// case 1:
//	[]string{"site", "info"}
//	->
//	map[string]any {
//		site: {info: val}
//	}
//
//	// case 2, last key is slice:
//	[]string{"site", "tags[1]"}
//	->
//	map[string]any {
//		site: {tags: [val]}
//	}
func MakeByKeys(keys []string, val any) (mp map[string]any) {
	_ = "STUB: not implemented"

	// if last key contains slice index, make slice wrap the val
	return nil
}

// valTyp := reflect.TypeOf(val)

// update val and last key

// multi nodes
