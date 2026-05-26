package secutil

import (
	"github.com/gookit/goutil/byteutil"
)

// aes methods
const (
	CryptAes128CBC = "aes-128-cbc"
	CryptAes256CBC = "aes-256-cbc"
)

// padding type for encoding contents.
const (
	PadTypeNone uint8 = iota
	PadTypeZeros
	PadTypePKCS5
	PadTypePKCS7
)

// type BlockModeFunc func(b cipher.Block, iv []byte, isEnc bool) cipher.BlockMode

// CryptConfig struct
type CryptConfig struct {
	Key string
	// IV string, length must be equals to cipher.Block.BlockSize()
	IV string
	// Method name. eg: CryptAes256CBC
	Method string
	// PadType for padding string.
	PadType uint8
	Encoder byteutil.BytesEncoder

	// BlockModeFn for create encrypter or decrypter blockMode
	// BlockModeFn BlockModeFunc
}

// AesCrypt struct
type AesCrypt struct {
	CryptConfig

	// iv length must is 16 = aes.BlockSize
	iv  []byte
	key []byte

	init   bool
	keyLen int
	// padding string

	encryptType string
}

// NewAesCrypt instance
func NewAesCrypt() *AesCrypt { _ = "STUB: not implemented"; return nil }

// Config crypt instance
func (p *AesCrypt) Config(fn func(c *CryptConfig)) *AesCrypt { _ = "STUB: not implemented"; return nil }

// Init crypt instance
func (p *AesCrypt) Init() error { _ = "STUB: not implemented"; return nil }

// iv length must is 16 = aes.BlockSize

// padding ASCII 0(NUL)

// Encrypt input source bytes. return error on fail.
func (p *AesCrypt) Encrypt(src []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// TODO 优化: 在 Init() 创建好 block 和 blockMode

// EncryptString to encoded string. return error on fail.
func (p *AesCrypt) EncryptString(src string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Decrypt an encrypt to source data
func (p *AesCrypt) Decrypt(enc []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// TODO 优化: 在 Init() 创建好 block 和 blockMode

// DecryptString to source string. return error on fail.
func (p *AesCrypt) DecryptString(enc string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
