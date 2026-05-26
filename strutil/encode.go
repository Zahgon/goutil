package strutil

import (
	"encoding/base32"
	"encoding/base64"
)

//
// -------------------- escape --------------------
//

// EscapeJS escape javascript string
func EscapeJS(s string) string { _ = "STUB: not implemented"; return "" }

// EscapeHTML escape html string
func EscapeHTML(s string) string { _ = "STUB: not implemented"; return "" }

// AddSlashes add slashes for the string.
func AddSlashes(s string) string { _ = "STUB: not implemented"; return "" }

// StripSlashes strip slashes for the string.
func StripSlashes(s string) string { _ = "STUB: not implemented"; return "" }

//
// -------------------- encode --------------------
//

// URLEncode encode url string.
func URLEncode(s string) string { _ = "STUB: not implemented"; return "" }

// escape query data

// URLDecode decode url string.
func URLDecode(s string) string { _ = "STUB: not implemented"; return "" }

// un-escape query data

//
// -------------------- base encode --------------------
//

// base32 encoding with no padding
var (
	B32Std = base32.StdEncoding.WithPadding(base32.NoPadding)
	B32Hex = base32.HexEncoding.WithPadding(base32.NoPadding)
)

// B32Encode base32 encode
func B32Encode(str string) string { _ = "STUB: not implemented"; return "" }

// B32Decode base32 decode
func B32Decode(str string) string { _ = "STUB: not implemented"; return "" }

// B64Std base64 encoding with no padding
var B64Std = base64.StdEncoding.WithPadding(base64.NoPadding)

// B64Encode base64 encode
func B64Encode(str string) string { _ = "STUB: not implemented"; return "" }

// B64EncodeBytes base64 encode
func B64EncodeBytes(src []byte) []byte { _ = "STUB: not implemented"; return nil }

// B64Decode base64 decode
func B64Decode(str string) string { _ = "STUB: not implemented"; return "" }

// B64DecodeBytes base64 decode
func B64DecodeBytes(str []byte) []byte { _ = "STUB: not implemented"; return nil }
