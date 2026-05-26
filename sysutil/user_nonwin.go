//go:build !windows

package sysutil

// IsAdmin Determine whether the current user is an administrator(root)
func IsAdmin() bool { _ = "STUB: not implemented"; return false }

// ChangeUserByName change work user by new username.
func ChangeUserByName(newUname string) error { _ = "STUB: not implemented"; return nil }

// syscall.Setlogin(newUname)

// ChangeUserUidGid change work user by new username uid,gid
//
// Deprecated: use ChangeUserUIDGid instead
func ChangeUserUidGid(newUID int, newGid int) error { _ = "STUB: not implemented"; return nil }

// ChangeUserUIDGid change work user by new username uid,gid
func ChangeUserUIDGid(newUID int, newGid int) (err error) { _ = "STUB: not implemented"; return nil }

// update group id
