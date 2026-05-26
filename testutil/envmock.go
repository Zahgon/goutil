package testutil

import (
	"os"
)

// Env mocking

// MockEnvValue will store old env value, set new val. will restore old value on end.
func MockEnvValue(key, val string, fn func(nv string)) { _ = "STUB: not implemented"; return }

// if old is empty, unset key.

// MockEnvValues will store old env value, set new val. will restore old value on end.
func MockEnvValues(kvMap map[string]string, fn func()) { _ = "STUB: not implemented"; return }

// MockOsEnvByText by env multi line text string.
// Will **CLEAR** all old ENV data, use given data map,
// and will recover old ENV after fn run. see MockCleanOsEnv
//
// Usage:
//
//	testutil.MockOsEnvByText(`
//		APP_COMMAND = login
//		APP_ENV = dev
//		APP_DEBUG = true
//
//	`, func() {
//			// do something ...
//	})
func MockOsEnvByText(envText string, fn func()) { _ = "STUB: not implemented"; return }

var envGroupSet = make(map[string]map[string]string)

// SetOsEnvs by map data with a group key. should call RemoveTmpEnvs after tested.
//
// Usage:
//
//	tmpKey := testutil.SetOsEnvs(map[string]string{
//		"APP_COMMAND": "login",
//		"APP_ENV":     "dev",
//		"APP_DEBUG":   "true",
//	})
//	defer testutil.RemoveTmpEnvs(tmpKey)
func SetOsEnvs(mp map[string]string) string { _ = "STUB: not implemented"; return "" }

// RemoveTmpEnvs remove test set envs by SetOsEnvs
func RemoveTmpEnvs(tmpKey string) { _ = "STUB: not implemented"; return }

// delete group key

// backup os ENV
var envBak = os.Environ()

// ClearOSEnv info for some testing cases.
//
// Usage:
//
//	testutil.ClearOSEnv()
//	defer testutil.RevertOSEnv()
//	// do something ...
func ClearOSEnv() {
	_ = "STUB: not implemented"

	// RevertOSEnv info
	return
}

func RevertOSEnv() { _ = "STUB: not implemented"; return }

// RunOnCleanEnv will CLEAR all old ENV, then run given func.
// will RECOVER old ENV after fn run.
func RunOnCleanEnv(runFn func()) { _ = "STUB: not implemented"; return }

// MockOsEnv by input map data. alias of MockCleanOsEnv
func MockOsEnv(mp map[string]string, fn func()) { _ = "STUB: not implemented"; return }

// MockCleanOsEnv by input env map data.
//
// will CLEAR all old ENV data, use given a data map.
// will RECOVER old ENV after fn run.
func MockCleanOsEnv(mp map[string]string, fn func()) { _ = "STUB: not implemented"; return }
