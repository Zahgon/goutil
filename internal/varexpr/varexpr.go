// Package varexpr provides some commonly ENV var parse functions.
//
// parse env value, allow expressions:
//
//	${VAR_NAME}            Only var name
//	${VAR_NAME | default}  With default value, if value is empty.
//	${VAR_NAME | ?error}   With error on value is empty.
//
// Examples:
//
//	only key     - "${SHELL}"
//	with default - "${NotExist | defValue}"
//	multi key    - "${GOPATH}/${APP_ENV | prod}/dir"
package varexpr

import (
	"regexp"
)

const (
	// SepChar separator char split var name and default value
	SepChar  = "|"
	VarLeft  = "${" // default var left format chars
	VarRight = "}"  // default var right format chars

	mustPrefix = '?' // must prefix char
)

// ParseOptFn option func
type ParseOptFn func(o *ParseOpts)

// ParseOpts parse options for ParseValue
type ParseOpts struct {
	// Getter Env value provider func.
	Getter func(string) string
	// ParseFn custom parse expr func. expr like "${SHELL}" "${NotExist|defValue}"
	ParseFn func(string) (string, error)
	// Regexp custom expression regex.
	Regexp *regexp.Regexp
	// var format chars for expression.
	// default left="${", right="}"
	VarLeft, VarRight string
}

func (opt *ParseOpts) useDefaultRegex() { _ = "STUB: not implemented"; return }

// must add "?" - To ensure that there is no greedy match
var envRegex = regexp.MustCompile(`\${.+?}`)
var std = New()

// Parse parse ENV var value from input string, support default value.
//
// Format:
//
//	${var_name}            Only var name
//	${var_name | default}  With default value
//	${var_name | ?error}   With error on value is empty.
//
// see Parser.Parse
func Parse(val string) (string, error) {
	_ = "STUB: not implemented"
	return "",

		// SafeParse parse ENV var value from input string, support default value.
		//
		// see Parser.Parse
		nil
}

func SafeParse(val string) string { _ = "STUB: not implemented"; return "" }

// ParseWith parse ENV var value from input string, support default value.
func ParseWith(val string, optFns ...ParseOptFn) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Parser parse ENV var value from input string, support default value.
type Parser struct {
	ParseOpts
}

// New create a new Parser
func New(optFns ...ParseOptFn) *Parser { _ = "STUB: not implemented"; return nil }

// Parse parse ENV var value from input string, support default value.
//
// Format:
//
//	${var_name}            Only var name
//	${var_name | default}  With default value
//	${var_name | ?error}   With error on value is empty.
//	${VAR_NAME1}/path/${VAR_NAME2}  Allow multi var name.
func (p *Parser) Parse(val string) (newVal string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

// enhance: see https://github.com/gookit/goutil/issues/135

// parse expression

// parse one node expression.
func (p *Parser) parseOne(eVar string) (val string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

// like "${NotExist | defValue}". first remove "${" and "}", then split it

// with default value.

// get ENV value by name

// check def is "?error"
