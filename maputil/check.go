package maputil

// HasKey check of the given map.
func HasKey(mp, key any) (ok bool) { _ = "STUB: not implemented"; return false }

// HasOneKey check of the given map. return the first exist key
func HasOneKey(mp any, keys ...any) (ok bool, key any) {
	_ = "STUB: not implemented"
	return false, *new(any)
}

// HasAllKeys check of the given map. return the first not exist key
func HasAllKeys(mp any, keys ...any) (ok bool, noKey any) {
	_ = "STUB: not implemented"
	return false, *new(any)
}
