package secutil

import (
	"crypto/cipher"
	"errors"
)

var (
	// ErrInvalidAESKeySize means the AES key length is not 16, 24 or 32 bytes.
	ErrInvalidAESKeySize = errors.New("secutil: invalid AES key size")
	// ErrInvalidGCMNonceSize means the GCM nonce length does not match AEAD requirements.
	ErrInvalidGCMNonceSize = errors.New("secutil: invalid GCM nonce size")
)

// EncryptGCM encrypts plaintext with AES-GCM.
//
// The key length must be 16, 24 or 32 bytes. A random nonce is generated and
// returned alongside the ciphertext.
func EncryptGCM(key, plaintext []byte) (ciphertext, nonce []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// DecryptGCM decrypts ciphertext with AES-GCM.
func DecryptGCM(key, nonce, ciphertext []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newGCM(key []byte) (cipher.AEAD, error) {
	_ = "STUB: not implemented"
	return *new(cipher.AEAD), nil
}

func validateAESKeySize(n int) error { _ = "STUB: not implemented"; return nil }
