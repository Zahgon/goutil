// Package secutil provide some security utils
package secutil

import (
	"errors"
)

// ErrUnPadding error
var ErrUnPadding = errors.New("un-padding decrypted data fail")

// PKCS5Padding input data
func PKCS5Padding(ciphertext []byte, blockSize int) []byte { _ = "STUB: not implemented"; return nil }

// PKCS5UnPadding input data
func PKCS5UnPadding(origData []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// fix: 检查删除的填充是否是一样的字符，不一样说明 delLen 值是有问题的，无法解码

// PKCS7Padding input data
func PKCS7Padding(ciphertext []byte, blockSize int) []byte { _ = "STUB: not implemented"; return nil }

// PKCS7UnPadding input data
func PKCS7UnPadding(origData []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
