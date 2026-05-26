package jsonutil

import (
	"io"
)

// MustString encode data to json string, will panic on error
func MustString(v any) string { _ = "STUB: not implemented"; return "" }

// Encode data to json bytes. alias of json.Marshal
func Encode(v any) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// EncodePretty encode data to pretty JSON bytes.
		nil
}

func EncodePretty(v any) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// EncodeString encode data to JSON string.
func EncodeString(v any) (string, error) { _ = "STUB: not implemented"; return "", nil }

// EncodeToWriter encode data to json and write to writer.
func EncodeToWriter(v any, w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeUnescapeHTML data to json bytes. will close escape HTML
func EncodeUnescapeHTML(v any) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Decode json bytes to data ptr. alias of json.Unmarshal
func Decode(bts []byte, ptr any) error { _ = "STUB: not implemented"; return nil }

// DecodeString json string to data ptr.
func DecodeString(str string, ptr any) error { _ = "STUB: not implemented"; return nil }

// DecodeReader decode JSON from io reader.
func DecodeReader(r io.Reader, ptr any) error { _ = "STUB: not implemented"; return nil }

// DecodeFile decode JSON from file, bind data to ptr.
func DecodeFile(file string, ptr any) error { _ = "STUB: not implemented"; return nil }
