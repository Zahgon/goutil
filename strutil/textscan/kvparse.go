package textscan

import (
	"errors"
)

// define special chars constants
const (
	MultiLineValMarkS = "'''"
	MultiLineValMarkD = `"""`
	MultiLineValMarkH = "<<<" // heredoc at start. <<<TXT ... TXT
	MultiLineValMarkQ = "\\"  // at end. eg: properties contents
	MultiLineCmtEnd   = "*/"
	// VarRefStartChars  = "${"
)

// KeyValueMatcher match key-value token.
// Support parse `KEY=VALUE` line text contents.
type KeyValueMatcher struct {
	// Separator string for split key and value, default is "="
	Separator string
	// MergeComments collect previous comments token to value token.
	// If set as True, on each s.Scan() please notice skip TokComments
	MergeComments bool
	// InlineComment parse and split inline comment
	InlineComment bool
	// DisableMultiLine value parse
	DisableMultiLine bool
	// KeyCheckFn set func check key string is valid
	KeyCheckFn func(key string) error
}

// Match text line.
func (m *KeyValueMatcher) Match(text string, prev Token) (Token, error) {
	_ = "STUB: not implemented"
	return *new(Token), nil
}

// check key string.

// handle value

// collect prev comments token

// split inline comments and clear quotes

// multi line value ended by \

// multi line value start

// split inline comments and clear quotes

func (m *KeyValueMatcher) inlineCommentsAndUnquote(vt *ValueToken, val string) string {
	_ = "STUB: not implemented"
	return ""

	// split inline comments
}

// merge comments token

// clear quotes

// DetectEnd for multi line value
func (m *KeyValueMatcher) DetectEnd(mark, text string) (ok bool, val string) {
	_ = "STUB: not implemented"
	return false, ""
}

// multi line value

// end

// goon

// goon

// end

// ValueToken contains key and value contents
type ValueToken struct {
	BaseToken
	m *KeyValueMatcher

	more bool
	mark string // end mark for multi line

	// key for token
	key string
	// for multi line value.
	values []string
	// comment for the item
	comment Token
}

// Key name
func (t *ValueToken) Key() string {
	_ = "STUB: not implemented"

	// Mark for multi line values
	return ""
}

func (t *ValueToken) Mark() string {
	_ = "STUB: not implemented"

	// Values for multi line values
	return ""
}

func (t *ValueToken) Values() []string {
	_ = "STUB: not implemented"

	// Comment lines string
	return nil
}

func (t *ValueToken) Comment() string { _ = "STUB: not implemented"; return "" }

// Value text string.
func (t *ValueToken) Value() string { _ = "STUB: not implemented"; return "" }

// HasMore is multi line values
func (t *ValueToken) HasMore() bool {
	_ = "STUB: not implemented"

	// HasComment for the value
	return false
}

func (t *ValueToken) HasComment() bool { _ = "STUB: not implemented"; return false }

// MergeSame comments token
func (t *ValueToken) MergeSame(_ Token) error { _ = "STUB: not implemented"; return nil }

// String of token
func (t *ValueToken) String() string { _ = "STUB: not implemented"; return "" }

// ErrMLineValueNotEnd error
var ErrMLineValueNotEnd = errors.New("not end of multi line value")

// ScanMore scan multi line values
func (t *ValueToken) ScanMore(ts *TextScanner) error { _ = "STUB: not implemented"; return nil }

// detect value end line

// CommentsMatcher match comments lines.
// will auto merge prev comments token
type CommentsMatcher struct {
	// InlineChars for match inline comments. default is: #
	InlineChars []byte
	// MatchFn for comments line
	// - mark 	useful on multi line comments
	MatchFn func(text string) (ok, more bool, err error)
	// DetectEnd for multi line comments
	DetectEnd func(text string) bool
}

// Match comments token
func (m *CommentsMatcher) Match(text string, prev Token) (Token, error) {
	_ = "STUB: not implemented"
	return *new(Token), nil
}

// skip empty line

// CommentsDetect check.
//
// - inlineChars: #
//
// default match:
//
//   - inline #, //
//   - multi line: /*
func CommentsDetect(str string, inlineChars []byte) (ok, more bool, err error) {
	_ = "STUB: not implemented"
	return false, false, nil
}

// match inline comments by prefix char.

// match start withs // OR /*

// multi line comments start

// end at line

// MatchEnd for multi line comments
func (m *CommentsMatcher) MatchEnd(text string) bool { _ = "STUB: not implemented"; return false }

// CommentToken struct
type CommentToken struct {
	BaseToken
	m *CommentsMatcher

	more bool
	// mark string // end mark for multi line

	// for multi line comments.
	comments []string
}

// NewCommentToken instance.
func NewCommentToken(val string) *CommentToken { _ = "STUB: not implemented"; return nil }

// Value fo token
func (t *CommentToken) Value() string { _ = "STUB: not implemented"; return "" }

// String for token
func (t *CommentToken) String() string {
	_ = "STUB: not implemented"

	// MergeSame comments token
	return ""
}

func (t *CommentToken) MergeSame(tok Token) error { _ = "STUB: not implemented"; return nil }

// HasMore is multi line values
func (t *CommentToken) HasMore() bool {
	_ = "STUB: not implemented"

	// ErrCommentsNotEnd error
	return false
}

var ErrCommentsNotEnd = errors.New("not end of multi-line comments")

// ScanMore scan multi line values
func (t *CommentToken) ScanMore(ts *TextScanner) error { _ = "STUB: not implemented"; return nil }

// detect comments end line

// CommentsDetectEnd multi line comments end
func CommentsDetectEnd(line string) bool { _ = "STUB: not implemented"; return false }
