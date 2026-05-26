package main

import (
	"bytes"

	"github.com/gookit/goutil/cflag"
)

var (
	// hidden pacakges.
	hidden = []string{
		"basefn",
		"netutil",
		"comdef",
		"internal",
		"syncs",
	}
	nameMap = map[string]string{
		"arr":     "array and Slice",
		"str":     "string Utils",
		"byte":    "Bytes Utils",
		"sys":     "system Utils",
		"math":    "math/Number",
		"fs":      "file System",
		"fmt":     "format Utils",
		"test":    "testing Utils",
		"dump":    "var Dumper",
		"structs": "struct Utils",
		"json":    "JSON Utils",
		"cli":     "CLI Utils",
		"env":     "ENV/Environment",
	}
	// show details in readme markdown.
	showDetails = []string{
		"jsonutil",
	}

	// allowLang = map[string]int{
	// 	"en":    1,
	// 	"zh-CN": 1,
	// }
	exFileNames = []string{
		"color_print.go",
	}
	exSuffixes = []string{
		"_test.go",
		"_windows.go",
		"_darwin.go",
		// "_linux",
	}
)

type genOptsSt struct {
	lang     string
	baseDir  string
	output   string
	template string
	tplDir   string
}

//lint:ignore U1000 for test
func (o genOptsSt) filePattern() string { _ = "STUB: not implemented"; return "" }

func (o genOptsSt) tplFilename() string { _ = "STUB: not implemented"; return "" }

func (o genOptsSt) tplFilepath(givePath string) string { _ = "STUB: not implemented"; return "" }

var (
	genOpts = genOptsSt{}
	// collected sub package names.
	// short name => full name.
	pkgNames = make(map[string]string, 16)
)

// go run ./internal/gendoc -h
// go run ./internal/gendoc
func main() {
	cflag.SetDebug(true)
	cmd := cflag.New(func(c *cflag.CFlags) {
		c.Version = "0.1.2"
		c.Desc = "Collect and dump all exported functions for goutil"
	})

	cmd.StringVar(&genOpts.lang, "lang", "en", "package desc message language. allow: en, zh-CN;;l")
	cmd.StringVar(&genOpts.baseDir, "dir", "./", "the base dir path for collect;;d")
	cmd.StringVar(&genOpts.output,
		"output",
		"./metadata.log",
		"the result output target.、n if is 'stdout', will direct print it;;o",
	)
	cmd.StringVar(&genOpts.tplDir,
		"tpl",
		"./internal/gendoc/template",
		"template file dir, use for generate, will inject metadata to the template.\nsee ./internal/gendoc/template/*.tpl;;t",
	)
	cmd.StringVar(&genOpts.template, "template", "", "the template file")

	cmd.Func = handle
	cmd.Example = `
  go run ./internal/gendoc -o stdout
  go run ./internal/gendoc -o stdout -l zh-CN
  go run ./internal/gendoc -o README.md
  go run ./internal/gendoc -o README.zh-CN.md
`
	cmd.MustParse(nil)
}

func handle(_ *cflag.CFlags) error { _ = "STUB: not implemented"; return nil }

// auto detect language

// close after a handle

// want output by template file
// var tplFile *os.File

// collect functions

// write to output

// ccolor.Cyanln("Collected packages:")
// dump.Clear(pkgNames)

func collectPgkFunc(ms []string, basePkg string) *bytes.Buffer {
	_ = "STUB: not implemented"
	// name - format dirname, will remove suffix: util
	return nil
}

// match func

// for each go file
// "jsonutil/jsonutil_test.go"
// "sysutil/sysutil_windows.go"

// sub pkg name.

// end of prev package.

// load prev sub-pkg doc file.

// now: name is package name.

// load sub-pkg start doc file.

// 隐藏详情

// read contents

// load last sub-pkg doc file.

func bufWritef(buf *bytes.Buffer, f string, a ...any) { _ = "STUB: not implemented"; return }

func bufWriteln(buf *bytes.Buffer, a ...any) { _ = "STUB: not implemented"; return }

func bufWriteDoc(buf *bytes.Buffer, partType, pkgName string) { _ = "STUB: not implemented"; return }

// fallback use en docs

func doWriteDoc2buf(buf *bytes.Buffer, filename string) bool {
	_ = "STUB: not implemented"
	return false
}

// ccolor.Infoln("- try read part readme from", partFile)
