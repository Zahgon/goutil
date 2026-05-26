//go:build windows

package sysutil

// ChangeUserByName change work user by new username.
func ChangeUserByName(newUname string) error { _ = "STUB: not implemented"; return nil }

// ChangeUserUidGid change work user by new username uid,gid
//
// Deprecated: use ChangeUserUIDGid instead
func ChangeUserUidGid(newUid int, newGid int) error { _ = "STUB: not implemented"; return nil }

// ChangeUserUIDGid change work user by new username uid,gid
func ChangeUserUIDGid(newUid int, newGid int) (err error) {
	_ = "STUB: not implemented"

	// IsAdmin Determine whether the current user is an administrator
	return nil
}

func IsAdmin() bool {
	_ = "STUB: not implemented"
	// 执行 net session 判断
	return false
}
