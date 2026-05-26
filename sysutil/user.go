package sysutil

import (
	"os/user"
)

// MustFindUser must find a system user by name
func MustFindUser(uname string) *user.User { _ = "STUB: not implemented"; return nil }

// LoginUser must get current user, will panic if error
func LoginUser() *user.User { _ = "STUB: not implemented"; return nil }

// CurrentUser must get current user, will panic if error
func CurrentUser() *user.User { _ = "STUB: not implemented"; return nil }

// UHomeDir get user home dir path, ignore error. (by user.Current)
func UHomeDir() string { _ = "STUB: not implemented"; return "" }

// homeDir cache
var _homeDir string

// UserHomeDir is alias of os.UserHomeDir, but ignore error.(by os.UserHomeDir)
func UserHomeDir() string { _ = "STUB: not implemented"; return "" }

// HomeDir get user home dir path.
func HomeDir() string { _ = "STUB: not implemented"; return "" }

// UserDir will prepend user home dir to subPaths
func UserDir(subPaths ...string) string { _ = "STUB: not implemented"; return "" }

// UserCacheDir will prepend user `$HOME/.cache` to subPaths
func UserCacheDir(subPaths ...string) string { _ = "STUB: not implemented"; return "" }

// UserConfigDir will prepend user `$HOME/.config` to subPath
func UserConfigDir(subPaths ...string) string { _ = "STUB: not implemented"; return "" }

// ExpandPath will parse `~` as user home dir path.
func ExpandPath(path string) string { _ = "STUB: not implemented"; return "" }

// ExpandHome will parse `~` as user home dir path.
func ExpandHome(path string) string { _ = "STUB: not implemented"; return "" }
