package envutil

// SetEnvMap set multi ENV(string-map) to os
func SetEnvMap(mp map[string]string) { _ = "STUB: not implemented"; return }

// SetEnvs set multi k-v ENV pairs to os
func SetEnvs(kvPairs ...string) { _ = "STUB: not implemented"; return }

// UnsetEnvs from os
func UnsetEnvs(keys ...string) { _ = "STUB: not implemented"; return }

// LoadText parse multiline text to ENV. Can use to load .env file contents.
//
// Usage:
//
//	envutil.LoadText(fsutil.ReadFile(".env"))
func LoadText(text string) { _ = "STUB: not implemented"; return }

// LoadString set line to ENV. e.g.: "KEY=VALUE"
func LoadString(line string) bool { _ = "STUB: not implemented"; return false }
