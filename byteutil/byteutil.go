// Package byteutil provides some useful functions for byte slice.
package byteutil

// Md5 Generate a 32-bit md5 bytes
func Md5(src any) []byte { _ = "STUB: not implemented"; return nil }

// Md5Sum Generate a md5 bytes
func Md5Sum(src any) []byte { _ = "STUB: not implemented"; return nil }

// cap(bs) == 16

// ShortMd5 Generate a 16-bit md5 bytes. remove the first 8 and last 8 bytes from 32-bit md5.
func ShortMd5(src any) []byte { _ = "STUB: not implemented"; return nil }

// Random bytes generate
func Random(length int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// Note that err == nil only if we read len(b) bytes.
		nil
}

// FirstLine from command output
func FirstLine(bs []byte) []byte { _ = "STUB: not implemented"; return nil }

// AppendAny append any value to byte slice
func AppendAny(dst []byte, v any) []byte { _ = "STUB: not implemented"; return nil }

// Cut bytes by one byte char. like bytes.Cut(), but sep is byte.
func Cut(bs []byte, sep byte) (before, after []byte, found bool) {
	_ = "STUB: not implemented"
	return nil, nil, false

	// SafeCut bytes by one byte char. always return before and after
}

func SafeCut(bs []byte, sep byte) (before, after []byte) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SafeCuts bytes by sub bytes. like the bytes.Cut(), but always return before and after
func SafeCuts(bs []byte, sep []byte) (before, after []byte) {
	_ = "STUB: not implemented"
	return nil, nil
}
