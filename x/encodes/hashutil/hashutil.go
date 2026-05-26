// Package hashutil provide some util for quickly generate hash
package hashutil

import (
	"hash"
)

// hash algorithm names
const (
	AlgoCRC32  = "crc32"
	AlgoCRC64  = "crc64"
	AlgoMD5    = "md5"
	AlgoSHA1   = "sha1"
	AlgoSHA224 = "sha224"
	AlgoSHA256 = "sha256"
	AlgoSHA384 = "sha384"
	AlgoSHA512 = "sha512"
)

// MD5 generate md5 string by given src
func MD5(src any) string { _ = "STUB: not implemented"; return "" }

// ShortMD5 Generate a 16-bit md5 bytes.
// remove first 8 and last 8 bytes from 32-bit md5.
func ShortMD5(src any) string { _ = "STUB: not implemented"; return "" }

// Hash generate hex hash string by given algorithm
func Hash(algo string, src any) string { _ = "STUB: not implemented"; return "" }

// HexBytes generate hex hash bytes by given algorithm
func HexBytes(algo string, src any) []byte { _ = "STUB: not implemented"; return nil }

// Hash32 generate hash by given algorithm, then use base32 encode.
func Hash32(algo string, src any) string { _ = "STUB: not implemented"; return "" }

// Base32Bytes generate base32 hash bytes by given algorithm
func Base32Bytes(algo string, src any) []byte { _ = "STUB: not implemented"; return nil }

// Hash64 generate hash by given algorithm, then use base64 encode.
func Hash64(algo string, src any) string { _ = "STUB: not implemented"; return "" }

// Base64Bytes generate base64 hash bytes by given algorithm
func Base64Bytes(algo string, src any) []byte { _ = "STUB: not implemented"; return nil }

// HashSum generate hash sum bytes by given algorithm
func HashSum(algo string, src any) []byte { _ = "STUB: not implemented"; return nil }

// NewHash create hash.Hash instance
//
// algo: crc32, crc64, md5, sha1, sha224, sha256, sha384, sha512, sha512_224, sha512_256
func NewHash(algo string) hash.Hash { _ = "STUB: not implemented"; return *new(hash.Hash) }

// ShortHash 生成 8 位十六进制 FNV-32a hash
func ShortHash(s string) string { _ = "STUB: not implemented"; return "" }
