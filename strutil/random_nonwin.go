//go:build !windows

package strutil

import (
	"math/rand"
)

var rn = newRand()

func newRand() *rand.Rand { _ = "STUB: not implemented"; return nil }

// buildRandomString 生成随机字符串
func buildRandomString(letters string, length int) string {
	_ = "STUB: not implemented"
	// rn := newRand()
	return ""
}
