package sysutil

import (
	"github.com/gookit/goutil/x/goinfo"
)

// GoVersion get go runtime version. eg: "1.18.2"
func GoVersion() string { _ = "STUB: not implemented"; return "" }

// GoInfo define. alias of goinfo.GoInfo
type GoInfo = goinfo.GoInfo

// CallerInfo define. alias of goinfo.CallerInfo
type CallerInfo = goinfo.CallerInfo

// ParseGoVersion get info by parse `go version` results. alias of goinfo.ParseGoVersion()
//
// Examples:
//
//		line, err := sysutil.ExecLine("go version")
//		if err != nil {
//			return err
//		}
//
//		info, err := sysutil.ParseGoVersion()
//	 	dump.P(info)
func ParseGoVersion(line string) (*GoInfo, error) { _ = "STUB: not implemented"; return nil, nil }

// OsGoInfo fetch and parse. alias of goinfo.OsGoInfo()
func OsGoInfo() (*GoInfo, error) {
	_ = "STUB: not implemented"
	return nil,

		// CallersInfos returns an array of the CallerInfo. can with filters
		nil
}

func CallersInfos(skip, num int, filters ...goinfo.CallerFilterFunc) []*CallerInfo {
	_ = "STUB: not implemented"
	return nil
}
