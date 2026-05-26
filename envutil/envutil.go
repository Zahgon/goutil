// Package envutil provide some commonly ENV util functions.
package envutil

import (
	"os"
)

// ValueGetter Env value provider func.
//
// TIPS: you can custom provide data.
var ValueGetter = os.Getenv

// VarReplace replaces ${var} or $var in the string according to the values.
//
// is alias of the os.ExpandEnv()
func VarReplace(s string) string { _ = "STUB: not implemented"; return "" }

// ParseOrErr parse ENV var value from input string, support default value.
//
// Diff with the ParseValue, this support return error.
//
// With error format: ${VAR_NAME | ?error}
func ParseOrErr(val string) (string, error) {
	_ = "STUB: not implemented"
	return "",

		// ParseValue parse ENV var value from input string, support default value.
		//
		// Format:
		//
		//	${var_name}            Only var name
		//	${var_name | default}  With default value
		//
		// Usage:
		//
		//	envutil.ParseValue("${ APP_NAME }")
		//	envutil.ParseValue("${ APP_ENV | dev }")
		nil
}

func ParseValue(val string) string { _ = "STUB: not implemented"; return "" }

// VarParse alias of the ParseValue
func VarParse(val string) string { _ = "STUB: not implemented"; return "" }

// ParseEnvValue alias of the ParseValue
func ParseEnvValue(val string) string { _ = "STUB: not implemented"; return "" }

// SplitText2map parse ENV text to map. Can use to parse .env file contents.
func SplitText2map(text string) map[string]string { _ = "STUB: not implemented"; return nil }

// SplitLineToKv parse ENV line to k-v. eg: 'DEBUG=true' => ['DEBUG', 'true']
func SplitLineToKv(line string) (string, string) { _ = "STUB: not implemented"; return "", "" }
