package checkfn

var detectedWSL bool
var detectedWSLContents string

// IsWSL system env
// https://github.com/Microsoft/WSL/issues/423#issuecomment-221627364
func IsWSL() bool {
	_ = "STUB: not implemented"
	// ENV:
	//
	//	WSL_DISTRO_NAME=Debian
	return false
}

// `cat /proc/version`
// on mac:
// 	!not the file!
// on linux(debian,ubuntu,alpine):
//	Linux version 4.19.121-linuxkit (root@18b3f92ade35) (gcc version 9.2.0 (Alpine 9.2.0)) #1 SMP Thu Jan 21 15:36:34 UTC 2021
// on win git bash, conEmu:
// 	MINGW64_NT-10.0-19042 version 3.1.7-340.x86_64 (@WIN-N0G619FD3UK) (gcc version 9.3.0 (GCC) ) 2020-10-23 13:08 UTC
// on WSL:
//  Linux version 4.4.0-19041-Microsoft (Microsoft@Microsoft.com) (gcc version 5.4.0 (GCC) ) #488-Microsoft Mon Sep 01 13:43:00 PST 2020

// ignore error
