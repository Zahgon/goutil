package byteutil

import (
	"encoding/base64"
	"encoding/hex"
)

// BytesEncodeFunc type
type BytesEncodeFunc func(src []byte) []byte

// BytesDecodeFunc type
type BytesDecodeFunc func(src []byte) ([]byte, error)

// BytesEncoder interface
type BytesEncoder interface {
	Encode(src []byte) []byte
	Decode(src []byte) ([]byte, error)
}

// StdEncoder implement the BytesEncoder
type StdEncoder struct {
	encodeFn BytesEncodeFunc
	decodeFn BytesDecodeFunc
}

// NewStdEncoder instance
func NewStdEncoder(encFn BytesEncodeFunc, decFn BytesDecodeFunc) *StdEncoder {
	_ = "STUB: not implemented"
	return nil
}

// Encode input
func (e *StdEncoder) Encode(src []byte) []byte { _ = "STUB: not implemented"; return nil }

// Decode input
func (e *StdEncoder) Decode(src []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

var (
	// HexEncoder instance
	HexEncoder = NewStdEncoder(func(src []byte) []byte {
		dst := make([]byte, hex.EncodedLen(len(src)))
		hex.Encode(dst, src)
		return dst
	}, func(src []byte) ([]byte, error) {
		n, err := hex.Decode(src, src)
		return src[:n], err
	})

	// B64Encoder instance
	B64Encoder = NewStdEncoder(func(src []byte) []byte {
		b64Dst := make([]byte, base64.StdEncoding.EncodedLen(len(src)))
		base64.StdEncoding.Encode(b64Dst, src)
		return b64Dst
	}, func(src []byte) ([]byte, error) {
		dBuf := make([]byte, base64.StdEncoding.DecodedLen(len(src)))
		n, err := base64.StdEncoding.Decode(dBuf, src)
		if err != nil {
			return nil, err
		}

		return dBuf[:n], err
	})
)
