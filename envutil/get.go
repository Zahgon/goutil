package envutil

// HasEnv check ENV key exists
func HasEnv(name string) bool { _ = "STUB: not implemented"; return false }

// Getenv get ENV value by key name, can with default value
func Getenv(name string, def ...string) string { _ = "STUB: not implemented"; return "" }

// MustGet get ENV value by key name, if not exists or empty, will panic
func MustGet(name string) string { _ = "STUB: not implemented"; return "" }

// GetInt get int ENV value by key name, can with default value
func GetInt(name string, def ...int) int { _ = "STUB: not implemented"; return 0 }

// GetBool get bool ENV value by key name, can with default value
func GetBool(name string, def ...bool) bool { _ = "STUB: not implemented"; return false }

// GetOne get one not empty ENV value by input names.
func GetOne(names []string, defVal ...string) string { _ = "STUB: not implemented"; return "" }

// GetMulti ENV values by input names.
func GetMulti(names ...string) map[string]string { _ = "STUB: not implemented"; return nil }

// OnExist check ENV value by key name, will call fn on value exists(not-empty)
func OnExist(name string, fn func(val string)) bool { _ = "STUB: not implemented"; return false }

// EnvPaths get and split $PATH to []string
func EnvPaths() []string { _ = "STUB: not implemented"; return nil }

// EnvMap like os.Environ, but will returns key-value map[string]string data.
func EnvMap() map[string]string { _ = "STUB: not implemented"; return nil }

// Environ like os.Environ, but will returns key-value map[string]string data.
func Environ() map[string]string { _ = "STUB: not implemented"; return nil }

// SearchEnvKeys values by given keywords
func SearchEnvKeys(keywords string) map[string]string { _ = "STUB: not implemented"; return nil }

// SearchEnv values by given keywords
func SearchEnv(keywords string, matchValue bool) map[string]string {
	_ = "STUB: not implemented"
	return nil
}
