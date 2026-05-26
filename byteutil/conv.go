package byteutil

// StrOrErr convert to string, return empty string on error.
func StrOrErr(bs []byte, err error) (string, error) { _ = "STUB: not implemented"; return "", nil }

// SafeString convert to string, return empty string on error.
func SafeString(bs []byte, err error) string { _ = "STUB: not implemented"; return "" }

// String unsafe convert bytes to string
func String(b []byte) string { _ = "STUB: not implemented"; return "" }

// ToString convert bytes to string
func ToString(b []byte) string { _ = "STUB: not implemented"; return "" }

// ToBytes convert any value to []byte. return error on convert failed.
func ToBytes(v any) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// SafeBytes convert any value to []byte. use fmt.Sprint() on convert failed.
func SafeBytes(v any) []byte { _ = "STUB: not implemented"; return nil }

// ToBytesFunc convert any value to []byte
type ToBytesFunc = func(v any) ([]byte, error)

// ToBytesWithFunc convert any value to []byte with custom fallback func.
//
// refer the strutil.ToStringWithFunc
//
// On not convert:
//   - If usrFn is nil, will return comdef.ErrConvType.
//   - If usrFn is not nil, will call it to convert.
func ToBytesWithFunc(v any, usrFn ToBytesFunc) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// same as `rune`

// Reverse 反转字节数组 eg: ABCD -> DCBA
func Reverse(arr []byte) { _ = "STUB: not implemented"; return }
