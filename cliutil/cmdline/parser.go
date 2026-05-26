package cmdline

import (
	"bytes"
	"os/exec"
)

// LineParser struct
// parse input command line to []string, such as cli os.Args
type LineParser struct {
	parsed bool
	// Line the full input command line text
	// eg `kite top sub -a "this is a message" --foo val1 --bar "val 2"`
	Line string
	// ParseEnv parse ENV var on the line.
	ParseEnv bool
	// the exploded nodes by space.
	nodes []string
	// the parsed args
	args []string

	// temp value
	quoteChar  byte
	quoteIndex int // if > 0, mark is not on start
	tempNode   bytes.Buffer
}

// NewParser create
func NewParser(line string) *LineParser { _ = "STUB: not implemented"; return nil }

// WithParseEnv with parse ENV var
func (p *LineParser) WithParseEnv() *LineParser { _ = "STUB: not implemented"; return nil }

// AlsoEnvParse input command line text to os.Args, will parse ENV var
func (p *LineParser) AlsoEnvParse() []string { _ = "STUB: not implemented"; return nil }

// NewExecCmd quick create exec.Cmd by cmdline string
func (p *LineParser) NewExecCmd() *exec.Cmd {
	_ = "STUB: not implemented"
	// parse get bin and args
	return nil
}

// create a new Cmd instance

// BinAndArgs get binName and args
func (p *LineParser) BinAndArgs() (bin string, args []string) {
	_ = "STUB: not implemented"
	// ensure parsed.
	return "", nil
}

// Parse input command line text to os.Args
func (p *LineParser) Parse() []string { _ = "STUB: not implemented"; return nil }

// enable parse Env var

func (p *LineParser) parseNode(node string) { _ = "STUB: not implemented"; return }

// in quotes

// end quotes

// eg: node="--pretty=format:'one two'"

// remove last quote

// goon ... write to temp node

// quote start

// only one words. eg: `-m "msg"`

// only one node: `msg"`

// eg: --pretty=format:'one two three'

// mark is not on start

// in quote, append to temp-node

func (p *LineParser) appendTempNode() { _ = "STUB: not implemented"; return }

// reset context value
