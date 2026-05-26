package strutil

// Md5 Generate a 32-bit md5 string
func Md5(src any) string { _ = "STUB: not implemented"; return "" }

// MD5 Generate a 32-bit md5 string
func MD5(src any) string {
	_ = "STUB: not implemented"

	// GenMd5 Generate a 32-bit md5 string
	return ""
}

func GenMd5(src any) string {
	_ = "STUB: not implemented"

	// Md5Simple md5 加密原始二进制的每个byte转为 base62，缩短长度(16)
	return ""
}

func Md5Simple(src any) string { _ = "STUB: not implemented"; return "" }

// 直接将每个 byte 转为 base62 数字然后拼接

// Md5Base62 md5 加密原始二进制的转为 base62 字符串，缩短长度(21~22)
func Md5Base62(src any) string { _ = "STUB: not implemented"; return "" }

// Step 2: 将字节数组转为 big.Int

// Step 3: 转换为 Base62 字符串

// base62.Grow(22)

// 反转字符串以得到正确的顺序

// Md5Bytes Generate a 32-bit md5 bytes
func Md5Bytes(src any) []byte { _ = "STUB: not implemented"; return nil }

// ShortMd5 Generate a 16-bit md5 string. remove the first 8 and last 8 bytes from 32-bit md5 string.
func ShortMd5(src any) string { _ = "STUB: not implemented"; return "" }

//
// ----------------------- Simple UUID -----------------------------
//

// UUIDv4 Generate a simple UUIDv4 string
func UUIDv4() (string, error) { _ = "STUB: not implemented"; return "", nil }

// 设置版本（4）和变体
// Version 4
// Variant 10

// ShortUUID Generate a short UUID (8 characters)
func ShortUUID() string { _ = "STUB: not implemented"; return "" }
