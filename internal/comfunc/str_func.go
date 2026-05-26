package comfunc

var commentsPrefixes = []string{"#", ";", "//"}

// ParseEnvLineOption parse env line options
type ParseEnvLineOption struct {
	// NotInlineComments dont parse inline comments.
	//  - default: false. will parse inline comments
	NotInlineComments bool
	// SkipOnErrorLine skip error line, continue parse next line
	//  - False: return error, clear parsed map
	SkipOnErrorLine bool
}

// ParseEnvLines parse simple multiline k-v string to a string-map.
// Can use to parse simple INI or DOTENV file contents.
//
// NOTE:
//
//   - It's like INI/ENV format contents.
//   - Support comments line starts with: "#", ";", "//"
//   - Support inline comments split with: " #" eg: "name=tom # a comments"
//   - DON'T support submap parse.
func ParseEnvLines(text string, opt ParseEnvLineOption) (mp map[string]string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// skip comments line

// invalid line

// SplitLineToKv parse string line to k-v, not support comments.
//
// Example:
//
//	'DEBUG=true' => ['DEBUG', 'true']
//
// NOTE: line must contain '=', allow: 'ENV_KEY='
func SplitLineToKv(line, sep string) (string, string) { _ = "STUB: not implemented"; return "", "" }

// SplitKvBySep parse string line to k-v, support parse comments.
//   - rmInlineComments: check and remove inline comments by ' #'
func SplitKvBySep(line, sep string, rmInlineComments bool) (key, val string) {
	_ = "STUB: not implemented"
	return "", ""
}

func splitLineByChar(line string, sep byte, rmInlineComments bool) (key, val string) {
	_ = "STUB: not implemented"
	return "", ""
}

func splitKvBySepPos(line string, sepPos, sepLen int, rmInlineComments bool) (key, val string) {
	_ = "STUB: not implemented"
	// key cannot be empty
	return "", ""
}

// check quotes if present

// remove quotes

// value is empty, only inline comments

// remove inline comments

// remove quotes
