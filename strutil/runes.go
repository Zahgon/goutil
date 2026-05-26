package strutil

// RuneIsWord char: a-zA-Z
func RuneIsWord(c rune) bool { _ = "STUB: not implemented"; return false }

// RuneIsLower char
func RuneIsLower(c rune) bool { _ = "STUB: not implemented"; return false }

// RuneIsUpper char
func RuneIsUpper(c rune) bool { _ = "STUB: not implemented"; return false }

// RunePos alias of the strings.IndexRune
func RunePos(s string, ru rune) int { _ = "STUB: not implemented"; return 0 }

// IsSpaceRune returns true if the given rune is a space, otherwise false.
func IsSpaceRune(r rune) bool { _ = "STUB: not implemented"; return false }

// Utf8Len count rune of the string.
//
// Examples:
//
//	str := "hi,你好"
//
//	len(str) // 9
//	strutil.RunesWidth(str) // 7 一个中文字占两个字符
//	RuneCount(str) = Utf8Len(s) // 5 按字算
func Utf8Len(s string) int { _ = "STUB: not implemented"; return 0 }

// Utf8len count rune of the string.
func Utf8len(s string) int { _ = "STUB: not implemented"; return 0 }

// RuneCount of the string.
//
// Examples:
//
//	str := "hi,你好"
//
//	len(str) // 9
//	strutil.RunesWidth(str) // 7 一个中文字占两个字符
//	RuneCount(str) = utf8.RuneCountInString(s) // 5 按字算
func RuneCount(s string) int { _ = "STUB: not implemented"; return 0 }

// RuneWidth of the rune.
//
// Example:
//
//	RuneWidth('你') // 2
//	RuneWidth('a') // 1
//	RuneWidth('\n') // 0
func RuneWidth(r rune) int { _ = "STUB: not implemented"; return 0 }

// eg: "\n"

// TextWidth utf8 string width. alias of RunesWidth()
func TextWidth(s string) int {
	_ = "STUB: not implemented"

	// Utf8Width utf8 string width. alias of RunesWidth
	return 0
}

func Utf8Width(s string) int { _ = "STUB: not implemented"; return 0 }

// RunesWidth utf8 runes string width.
//
// Examples:
//
//	str := "hi,你好"
//
//	len(str) // 9
//	strutil.RunesWidth(str) // 7 一个中文字占两个字符
//	len([]rune(str)) = utf8.RuneCountInString(s) // 5 按字算
func RunesWidth(rs []rune) (w int) { _ = "STUB: not implemented"; return 0 }

// Truncate alias of the Utf8Truncate()
func Truncate(s string, w int, tail string) string { _ = "STUB: not implemented"; return "" }

// TextTruncate alias of the Utf8Truncate()
func TextTruncate(s string, w int, tail string) string { _ = "STUB: not implemented"; return "" }

// Utf8Truncate a string with given width.
func Utf8Truncate(s string, w int, tail string) string { _ = "STUB: not implemented"; return "" }

// utf8Truncate a string with given width.
func utf8Truncate(s string, sw, w int, tail string) string { _ = "STUB: not implemented"; return "" }

// Chunk split string to chunks by size.
// func Chunk[T ~string](s T, size int) []T {
// }

// TextSplit alias of the Utf8Split()
func TextSplit(s string, w int) []string { _ = "STUB: not implemented"; return nil }

// Utf8Split split a string by width.
func Utf8Split(s string, w int) []string { _ = "STUB: not implemented"; return nil }

// reset

// append to next line.

// TextWrap a string by "\n". alias of the WidthWrap()
func TextWrap(s string, w int) string { _ = "STUB: not implemented"; return "" }

// WidthWrap a string by "\n"
//
// Example:
//
//		s := "hello 你好, world 世界"
//	 s1 := strutil.TextWrap(s, 6) // "hello \n你好, \nworld \n世界"
func WidthWrap(s string, w int) string { _ = "STUB: not implemented"; return "" }

// WordWrap text string and limit width.
func WordWrap(s string, w int) string { _ = "STUB: not implemented"; return "" }

// Runes data slice
type Runes []rune

// Padding a rune to want length and with position
func (rs Runes) Padding(pad rune, length int, pos PosFlag) []rune {
	_ = "STUB: not implemented"
	return nil
}

// PadLeft a rune to want length
func (rs Runes) PadLeft(pad rune, length int) []rune { _ = "STUB: not implemented"; return nil }

// PadRight a rune to want length
func (rs Runes) PadRight(pad rune, length int) []rune { _ = "STUB: not implemented"; return nil }
