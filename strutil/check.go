package strutil

import (
	"regexp"
	"strings"

	"github.com/gookit/goutil/internal/checkfn"
)

// Equal check, alias of strings.EqualFold
var Equal = strings.EqualFold
var IsHttpURL = checkfn.IsHttpURL

// IsNumChar returns true if the given character is a numeric, otherwise false.
func IsNumChar(c byte) bool { _ = "STUB: not implemented"; return false }

var (
	uintReg = regexp.MustCompile(`^\d+$`)
	intReg  = regexp.MustCompile(`^[-+]?\d+$`)

	floatReg = regexp.MustCompile(`^[-+]?\d*\.?\d+$`)
)

// IsInt check the string is an integer number
func IsInt(s string) bool { _ = "STUB: not implemented"; return false }

// IsUint check the string is an unsigned integer number
func IsUint(s string) bool { _ = "STUB: not implemented"; return false }

// IsFloat check the string is a float number
func IsFloat(s string) bool { _ = "STUB: not implemented"; return false }

// IsNumeric returns true if the given string is a numeric(int/float), otherwise false.
func IsNumeric(s string) bool { _ = "STUB: not implemented"; return false }

// IsPositiveNum check the string is a positive number
func IsPositiveNum(s string) bool { _ = "STUB: not implemented"; return false }

// IsAlphabet char
func IsAlphabet(char uint8) bool {
	_ = "STUB: not implemented"
	// A 65 -> Z 90
	return false
}

// a 97 -> z 122

// IsAlphaNum reports whether the byte is an ASCII letter, number, or underscore
func IsAlphaNum(c uint8) bool { _ = "STUB: not implemented"; return false }

// IsUpper returns true if the given string is an uppercase, otherwise false.
func IsUpper(s string) bool { _ = "STUB: not implemented"; return false }

// IsLower returns true if the given string is a lowercase, otherwise false.
func IsLower(s string) bool { _ = "STUB: not implemented"; return false }

// IsAllASCII 判断字符串是否全为可打印 ASCII（无中文等多字节字符）
func IsAllASCII(s string) bool { _ = "STUB: not implemented"; return false }

// StrPos alias of the strings.Index
func StrPos(s, sub string) int { _ = "STUB: not implemented"; return 0 }

// BytePos alias of the strings.IndexByte
func BytePos(s string, bt byte) int { _ = "STUB: not implemented"; return 0 }

// IEqual ignore case check given two strings are equals.
func IEqual(s1, s2 string) bool { _ = "STUB: not implemented"; return false }

// NoCaseEq check two strings is equals and case-insensitivity
func NoCaseEq(s, t string) bool { _ = "STUB: not implemented"; return false }

// IContains ignore case check substr in the given string.
func IContains(s, sub string) bool { _ = "STUB: not implemented"; return false }

// ContainsByte in given string.
func ContainsByte(s string, c byte) bool { _ = "STUB: not implemented"; return false }

// ContainsByteOne in given string.
func ContainsByteOne(s string, bs []byte) bool { _ = "STUB: not implemented"; return false }

// InArray alias of HasOneSub()
var InArray = HasOneSub

// ContainsOne substr(s) in the given string. alias of HasOneSub()
func ContainsOne(s string, subs []string) bool { _ = "STUB: not implemented"; return false }

// HasOneSub substr(s) in the given string.
func HasOneSub(s string, subs []string) bool { _ = "STUB: not implemented"; return false }

// IContainsOne ignore case check has one substr(s) in the given string.
func IContainsOne(s string, subs []string) bool { _ = "STUB: not implemented"; return false }

// ContainsAll given string should contain all substrings. alias of HasAllSubs()
func ContainsAll(s string, subs []string) bool { _ = "STUB: not implemented"; return false }

// HasAllSubs given string should contain all substrings
func HasAllSubs(s string, subs []string) bool { _ = "STUB: not implemented"; return false }

// IContainsAll like ContainsAll(), but ignore case
func IContainsAll(s string, subs []string) bool { _ = "STUB: not implemented"; return false }

// StartsWithAny alias of the HasOnePrefix
var StartsWithAny = HasOneSuffix

// IsStartsOf alias of the HasOnePrefix
func IsStartsOf(s string, prefixes []string) bool { _ = "STUB: not implemented"; return false }

// HasOnePrefix the string starts with one of the subs
func HasOnePrefix(s string, prefixes []string) bool { _ = "STUB: not implemented"; return false }

// StartsWith alias func for HasPrefix
var StartsWith = strings.HasPrefix

// HasPrefix substr in the given string.
func HasPrefix(s string, prefix string) bool { _ = "STUB: not implemented"; return false }

// IsStartOf alias of the strings.HasPrefix
func IsStartOf(s, prefix string) bool { _ = "STUB: not implemented"; return false }

// HasSuffix substr in the given string.
func HasSuffix(s string, suffix string) bool { _ = "STUB: not implemented"; return false }

// IsEndOf alias of the strings.HasSuffix
func IsEndOf(s, suffix string) bool { _ = "STUB: not implemented"; return false }

// HasOneSuffix the string end with one of the subs
func HasOneSuffix(s string, suffixes []string) bool { _ = "STUB: not implemented"; return false }

// IsValidUtf8 valid utf8 string check
func IsValidUtf8(s string) bool { _ = "STUB: not implemented"; return false }

// ----- refer from github.com/yuin/goldmark/util

// refer from github.com/yuin/goldmark/util
var spaceTable = [256]int8{
	0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
}

// IsSpace returns true if the given character is a space, otherwise false.
func IsSpace(c byte) bool { _ = "STUB: not implemented"; return false }

// IsEmpty returns true if the given string is empty.
func IsEmpty(s string) bool {
	_ = "STUB: not implemented"

	// IsBlank returns true if the given string is all space characters.
	return false
}

func IsBlank(s string) bool { _ = "STUB: not implemented"; return false }

// IsNotBlank returns true if the given string is not blank.
func IsNotBlank(s string) bool { _ = "STUB: not implemented"; return false }

// IsBlankBytes returns true if the given []byte is all space characters.
func IsBlankBytes(bs []byte) bool { _ = "STUB: not implemented"; return false }

// IsSymbol reports whether the rune is a symbolic character.
func IsSymbol(r rune) bool { _ = "STUB: not implemented"; return false }

// HasEmpty value for input strings
func HasEmpty(ss ...string) bool { _ = "STUB: not implemented"; return false }

// IsAllEmpty for input strings
func IsAllEmpty(ss ...string) bool { _ = "STUB: not implemented"; return false }

var (
	// regex for check version number
	verRegex = regexp.MustCompile(`^[0-9][\d.]+(-\w+)?$`)
	// regex for check variable name
	varRegex = regexp.MustCompile(`^[a-zA-Z][\w-]*$`)
	// regex for check env var name
	envRegex = regexp.MustCompile(`^[A-Z][A-Z0-9_]*$`)
	// IsVariableName alias for IsVarName
	IsVariableName = IsVarName
	// regex for check uuid string. format: 8-4-4-4-12
	uuidPattern = regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`)
)

// IsVersion number. eg: 1.2.0
func IsVersion(s string) bool { _ = "STUB: not implemented"; return false }

// IsVarName is valid variable name.
func IsVarName(s string) bool { _ = "STUB: not implemented"; return false }

// IsEnvName is valid ENV var name. eg: APP_NAME
func IsEnvName(s string) bool { _ = "STUB: not implemented"; return false }

// IsUUID check if the string is a valid UUID format.
func IsUUID(s string) bool { _ = "STUB: not implemented"; return false }

// Compare for two strings.
func Compare(s1, s2, op string) bool { _ = "STUB: not implemented"; return false }

// eq

// VersionCompare for two version strings. eg: 1.2.0 > 1.1.0
func VersionCompare(v1, v2, op string) bool { _ = "STUB: not implemented"; return false }

// parseVersion 将版本号字符串解析为整数数组
func parseVersion(version string) []int { _ = "STUB: not implemented"; return nil }

// compareVersions 比较两个版本号数组
// 返回: -1 表示 v1 < v2, 0 表示 v1 = v2, 1 表示 v1 > v2
func compareVersions(v1, v2 []int) int { _ = "STUB: not implemented"; return 0 }

// SimpleMatch all substring in the give text string.
//
// Difference the ContainsAll:
//
//   - start with ^ for exclude contains check.
//   - end with $ for the check end with keyword.
func SimpleMatch(s string, keywords []string) bool { _ = "STUB: not implemented"; return false }

// exclude

// end with

// include

// QuickMatch check for a string. pattern can be a substring.
func QuickMatch(pattern, s string) bool { _ = "STUB: not implemented"; return false }

// PathMatch check for a string match the pattern. alias of the path.Match()
//
// TIP: `*` can match any char, not contain `/`.
func PathMatch(pattern, s string) bool { _ = "STUB: not implemented"; return false }

// GlobMatch check for a string match the pattern.
//
// Difference with PathMatch() is: `*` can match any char, contain `/`.
func GlobMatch(pattern, s string) bool {
	_ = "STUB: not implemented"
	// replace `/` to `S` for path.Match
	return false
}

// LikeMatch simple check for a string match the pattern. pattern like the SQL LIKE.
func LikeMatch(pattern, s string) bool { _ = "STUB: not implemented"; return false }

// eg `%abc` `%abc%`

// eg `abc%`

// MatchNodePath check for a string match the pattern.
//
// Use on a pattern:
//   - `*` match any to sep
//   - `**` match any to end. only allow at start or end on pattern.
//
// Example:
//
//	strutil.MatchNodePath()
func MatchNodePath(pattern, s string, sep string) bool { _ = "STUB: not implemented"; return false }

// at start
