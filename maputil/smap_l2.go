package maputil

// L2StrMap is alias of map[string]map[string]string
type L2StrMap map[string]map[string]string

// Load data, merge new data to old
func (m L2StrMap) Load(mp map[string]map[string]string) { _ = "STUB: not implemented"; return }

// Value get by key path. eg: "top.sub"
func (m L2StrMap) Value(key string) (val string, ok bool) {
	_ = "STUB: not implemented"
	return "", false
}

// Get value by key path. eg: "top.sub"
func (m L2StrMap) Get(key string) string { _ = "STUB: not implemented"; return "" }

// Exists check key path exists. eg: "top.sub"
func (m L2StrMap) Exists(key string) bool { _ = "STUB: not implemented"; return false }

// StrMap get by top key. eg: "top"
func (m L2StrMap) StrMap(top string) StrMap { _ = "STUB: not implemented"; return *new(StrMap) }
