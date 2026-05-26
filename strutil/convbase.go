package strutil

//
// -------------------- convert base --------------------
//

const (
	Base10Chars = "0123456789"
	Base16Chars = "0123456789abcdef"
	Base32Chars = "0123456789abcdefghjkmnpqrstvwxyz"
	Base36Chars = "0123456789abcdefghijklmnopqrstuvwxyz"
	Base48Chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKL"
	Base62Chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Base64Chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ+/"
)

// Base10Conv convert base10 string to new base string.
func Base10Conv(src string, to int) string { _ = "STUB: not implemented"; return "" }

// BaseConv convert base string by from and to base.
//
// Note: from and to base must be in [2, 64]
//
// Usage:
//
//	BaseConv("123", 10, 16) // Output: "7b"
//	BaseConv("7b", 16, 10) // Output: "123"
func BaseConv(src string, from, to int) string { _ = "STUB: not implemented"; return "" }

// BaseConvInt convert base int to new base string.
//
// Usage:
//
//	BaseConv(123, 16) // Output: "7b"
func BaseConvInt(src uint64, toBase int) string { _ = "STUB: not implemented"; return "" }

// bigInt 支持 2-62 进制转换处理 TODO

// BaseConvByTpl convert base string by template.
//
// Usage:
//
//	BaseConvert("123", Base62Chars, Base16Chars) // Output: "1e"
//	BaseConvert("1e", Base16Chars, Base62Chars) // Output: "123"
func BaseConvByTpl(src string, fromBase, toBase string) string {
	_ = "STUB: not implemented"
	return ""
}

// convert to base 10

// convert to new base

// BaseConvIntByTpl convert base int to new base string.
func BaseConvIntByTpl(dec uint64, toBase string) string {
	_ = "STUB: not implemented"
	// convert to new base
	return ""
}
