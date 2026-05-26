package strutil

import (
	"github.com/gookit/goutil/comdef"
)

// PosFlag type
type PosFlag = comdef.Position

// Position for padding/resize string
const (
	PosLeft PosFlag = iota
	PosMiddle
	PosRight
	PosAuto
)

/*************************************************************
 * String padding operation
 *************************************************************/

// Padding Fill the string to the specified length.
//
// params:
//   - s: 原始字符串
//   - pad: 用于填充的字符或字符串
//   - length: 目标长度
//   - padPos: 填充位置标志（左填充或右填充）
func Padding(s, pad string, length int, padPos PosFlag) string {
	_ = "STUB: not implemented"
	return ""
}

// Utf8Padding Fill the string to the specified length. use utf8 width.
func Utf8Padding(s, pad string, wantLen int, padPos PosFlag) string {
	_ = "STUB: not implemented"
	return ""
}

func padding(s, pad string, sLen, wantLen int, padPos PosFlag) string {
	_ = "STUB: not implemented"
	return ""
}

// do not need padding.

// pad space.

// Sprintf: 是按字数来填充的，不管中英文都是一个字符 - 有问题
// to right

// other character.
// to right

// PadLeft a string.
func PadLeft(s, pad string, length int) string { _ = "STUB: not implemented"; return "" }

// PadRight a string.
func PadRight(s, pad string, length int) string { _ = "STUB: not implemented"; return "" }

// PadChars padding a rune/byte to want length and with position flag
func PadChars[T byte | rune](cs []T, pad T, length int, pos PosFlag) []T {
	_ = "STUB: not implemented"
	return nil
}

// to left

// PadBytes padding a byte to want length and with position flag
func PadBytes(bs []byte, pad byte, length int, pos PosFlag) []byte {
	_ = "STUB: not implemented"
	return nil
}

// PadBytesLeft a byte to want length
func PadBytesLeft(bs []byte, pad byte, length int) []byte { _ = "STUB: not implemented"; return nil }

// PadBytesRight a byte to want length
func PadBytesRight(bs []byte, pad byte, length int) []byte { _ = "STUB: not implemented"; return nil }

// PadRunes padding a rune to want length and with position flag
func PadRunes(rs []rune, pad rune, length int, pos PosFlag) []rune {
	_ = "STUB: not implemented"
	return nil
}

// PadRunesLeft a rune to want length
func PadRunesLeft(rs []rune, pad rune, length int) []rune { _ = "STUB: not implemented"; return nil }

// PadRunesRight a rune to want length
func PadRunesRight(rs []rune, pad rune, length int) []rune { _ = "STUB: not implemented"; return nil }

// Align a string by given length and align settings. alias of Resize
func Align(s string, length int, align comdef.Align) string { _ = "STUB: not implemented"; return "" }

// Utf8Align a string by given length and align settings. alias of Utf8Resize
func Utf8Align(s string, length int, align comdef.Align) string {
	_ = "STUB: not implemented"
	return ""
}

// Resize a string by given length and align settings. use padding space.
// If len(s) > wantLen, will truncate it.
func Resize(s string, length int, align comdef.Align) string { _ = "STUB: not implemented"; return "" }

// Utf8Resize a string by given length and align settings. use padding space.
// If width(s) > wantLen, will truncate it.
func Utf8Resize(s string, length int, align comdef.Align) string {
	_ = "STUB: not implemented"
	return ""
}

// resize a string by given length and align settings. use padding space.
func resize(s string, sLen, wantLen int, align comdef.Align, cutOverflow bool) string {
	_ = "STUB: not implemented"
	return ""
}

// do not need padding.
// cutOverflow: truncate on sLen > wantLen

// tip: 左对齐 - 使用空白填充右边

/*************************************************************
 * String repeat operation
 *************************************************************/

// Repeat a string by given times.
func Repeat(s string, times int) string { _ = "STUB: not implemented"; return "" }

// RepeatRune repeat a rune char.
func RepeatRune(char rune, times int) []rune { _ = "STUB: not implemented"; return nil }

// RepeatBytes repeat a byte char.
func RepeatBytes(char byte, times int) []byte { _ = "STUB: not implemented"; return nil }

// RepeatChars repeat a byte char.
func RepeatChars[T byte | rune](char T, times int) []T { _ = "STUB: not implemented"; return nil }
