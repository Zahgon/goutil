// Package goinfo provide some standard util functions for go.
package goinfo

import (
	"regexp"
)

// GoVersion get go runtime version. eg: "1.18.2"
func GoVersion() string { _ = "STUB: not implemented"; return "" }

// GoInfo define
//
// On os by:
//
//	go env GOVERSION GOOS GOARCH
//	go version // "go version go1.19 darwin/amd64"
type GoInfo struct {
	Version string
	GoOS    string
	Arch    string
}

// match "go version go1.19 darwin/amd64"
var goVerRegex = regexp.MustCompile(`\sgo([\d.]+)\s(\w+)/(\w+)`)

// ParseGoVersion get info by parse `go version` results.
//
// Examples:
//
//		line, err := sysutil.ExecLine("go version")
//		if err != nil {
//			return err
//		}
//
//		info, err := goinfo.ParseGoVersion()
//	 	dump.P(info)
func ParseGoVersion(line string) (*GoInfo, error) {
	_ = "STUB: not implemented"
	// eg: [" go1.19 darwin/amd64", "1.19", "darwin", "amd64"]
	return nil, nil
}

// OsGoInfo fetch and parse
func OsGoInfo() (*GoInfo, error) { _ = "STUB: not implemented"; return nil, nil }
