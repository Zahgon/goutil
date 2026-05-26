//go:build windows

package termenv

import (
	"syscall"

	"golang.org/x/sys/windows"
)

// Get the Windows Version and Build Number
var majorVersion, _, buildNumber = windows.RtlGetNtVersionNumbers()

// refer
//
//	https://github.com/Delta456/box-cli-maker/blob/7b5a1ad8a016ce181e7d8b05e24b54ff60b4b38a/detect_windows.go#L30-L57
//	https://github.com/gookit/color/issues/25#issuecomment-738727917
//
// detects the color level supported on Windows: CMD, PowerShell
func detectSpecialTermColor(_ string) (tl ColorLevel, needVTP bool) {
	_ = "STUB: not implemented"
	return *new(ColorLevel), false
}

// ConEmuANSI is "ON" for generic ANSI support
// but True Color option is enabled by default
// I am just assuming that people wouldn't have disabled it
// Even if it is not enabled then ConEmu will auto round off accordingly

// Before Windows 10 Build Number 10586, console never supported ANSI Colors

// Detect if using ANSICON on older systems

// 8-bit Colors were only supported after v1.81 release

// True Color is not available before build 14931 so fallback to 8-bit color.

// Windows 10 build 14931 is the first release that supports 16m/TrueColor

// TryEnableVTP try force enables colors on Windows terminal
func TryEnableVTP(enable bool) bool { _ = "STUB: not implemented"; return false }

// enable colors on Windows terminal

// initKernel32Proc()

func tryEnableOnCONOUT() bool { _ = "STUB: not implemented"; return false }

func tryEnableOnStdout() bool {
	_ = "STUB: not implemented"
	// try direct open syscall.Stdout
	return false
}

// related docs
// https://docs.microsoft.com/zh-cn/windows/console/console-virtual-terminal-sequences
// https://docs.microsoft.com/zh-cn/windows/console/console-virtual-terminal-sequences#samples
var (
	// isMSys bool
	kernel32 *syscall.LazyDLL

	procGetConsoleMode *syscall.LazyProc
	procSetConsoleMode *syscall.LazyProc
)

func initKernel32Proc() { _ = "STUB: not implemented"; return }

// load related Windows dll
// https://docs.microsoft.com/en-us/windows/console/setconsolemode

/*************************************************************
 * render full color code on Windows(8,16,24bit color)
 *************************************************************/

// EnableVTProcessing Enable virtual terminal processing on Windows
//
// ref from github.com/konsorten/go-windows-terminal-sequences
// doc https://docs.microsoft.com/zh-cn/windows/console/console-virtual-terminal-sequences#samples
//
// Usage:
//
//	err := EnableVTProcessing(syscall.Stdout, true)
//	// support print color text
//	err = EnableVTProcessing(syscall.Stdout, false)
func EnableVTProcessing(stream syscall.Handle, enable bool) error {
	_ = "STUB: not implemented"

	// Check if it is currently in the terminal
	// err := syscall.GetConsoleMode(syscall.Stdout, &mode)
	return nil
}

// docs https://docs.microsoft.com/zh-cn/windows/console/getconsolemode#parameters

// ret, _, err := procSetConsoleMode.Call(uintptr(stream), uintptr(mode))
// if ret == 0 {
// 	return err
// }

// on Windows, must convert 'syscall.Stdin' to int
func syscallStdinFd() int { _ = "STUB: not implemented"; return 0 }

// on Windows, must convert to int
func syscallStdoutFd() int { _ = "STUB: not implemented"; return 0 }
