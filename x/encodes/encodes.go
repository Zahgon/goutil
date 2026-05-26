// Package encodes provide some util for encode/decode data
package encodes

import (
	"encoding/base32"
	"encoding/base64"
)

// BaseEncoder interface
type BaseEncoder interface {
	Encode(dst []byte, src []byte)
	EncodeToString(src []byte) string
	Decode(dst []byte, src []byte) (n int, err error)
	DecodeString(s string) ([]byte, error)
}

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

// base64 encoding with no padding
var (
	B64Std = base64.StdEncoding.WithPadding(base64.NoPadding)
	B64URL = base64.URLEncoding.WithPadding(base64.NoPadding)
)

// B64Encode base64 encode
func B64Encode(str string) string { _ = "STUB: not implemented"; return "" }

// B64EncodeBytes base64 encode
func B64EncodeBytes(src []byte) []byte { _ = "STUB: not implemented"; return nil }

// B64Decode base64 decode
func B64Decode(str string) string { _ = "STUB: not implemented"; return "" }

// B64DecodeBytes base64 decode
func B64DecodeBytes(str []byte) []byte { _ = "STUB: not implemented"; return nil }
