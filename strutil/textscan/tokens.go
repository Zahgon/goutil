package textscan

// Kind type
type Kind uint8

// String name for kind
func (k Kind) String() string { _ = "STUB: not implemented"; return "" }

// builtin defined kinds
const (
	TokInvalid Kind = iota
	TokKey
	TokValue
	TokComments
)

// global kinds
var kinds = map[Kind]string{
	TokInvalid:  "Invalid",
	TokKey:      "Key",
	TokValue:    "Value",
	TokComments: "Comments",
}

// AddKind add global kind to kinds
func AddKind(k Kind, name string) { _ = "STUB: not implemented"; return }

// HasKind check
func HasKind(k Kind) bool { _ = "STUB: not implemented"; return false }

// KindString name
func KindString(k Kind) string { _ = "STUB: not implemented"; return "" }

// IsKindToken check
func IsKindToken(k Kind, tok Token) bool { _ = "STUB: not implemented"; return false }

// LiteToken interface
type LiteToken interface {
	Kind() Kind
	Value() string
	IsValid() bool
}

// Token parser
type Token interface {
	LiteToken
	String() string
	// HasMore is multi line values
	HasMore() bool
	// ScanMore scan multi line values
	ScanMore(ts *TextScanner) error
	MergeSame(tok Token) error
}

// BaseToken struct
type BaseToken struct {
	kind  Kind
	value string

	// Offset int // byte offset, starting at 0
	// Line   int // line number, starting at 1
	// Column int // column number, starting at 1 (character count per line)
}

// Kind type
func (t *BaseToken) Kind() Kind {
	_ = "STUB: not implemented"

	// IsValid token
	return *new(Kind)
}

func (t *BaseToken) IsValid() bool { _ = "STUB: not implemented"; return false }

// Value of token
func (t *BaseToken) Value() string {
	_ = "STUB: not implemented"

	// String of token
	return ""
}

func (t *BaseToken) String() string { _ = "STUB: not implemented"; return "" }

// StringToken struct
type StringToken struct {
	BaseToken
}

// NewEmptyToken instance.
// Can use for want skip parse some contents
func NewEmptyToken() *StringToken { _ = "STUB: not implemented"; return nil }

// NewStringToken instance.
func NewStringToken(k Kind, val string) *StringToken { _ = "STUB: not implemented"; return nil }

// HasMore is multi line values
func (t *StringToken) HasMore() bool {
	_ = "STUB: not implemented"

	// ScanMore implements
	return false
}

func (t *StringToken) ScanMore(_ *TextScanner) error {
	_ = "STUB: not implemented"

	// MergeSame implements
	return nil
}

func (t *StringToken) MergeSame(_ Token) error { _ = "STUB: not implemented"; return nil }
