package fmtutil

// data size
const (
	OneKByte = 1024
	OneMByte = 1024 * 1024
	OneGByte = 1024 * 1024 * 1024
)

// DataSize format bytes number friendly.
//
// Usage:
//
//	file, err := os.Open(path)
//	fl, err := file.Stat()
//	fmtSize := DataSize(fl.Size())
func DataSize(size uint64) string { _ = "STUB: not implemented"; return "" }

// HumanSize alias of the DataSize
func HumanSize(size uint64) string { _ = "STUB: not implemented"; return "" }

// SizeToString alias of the DataSize
func SizeToString(size uint64) string { _ = "STUB: not implemented"; return "" }

// StringToByte alias of the ParseByte
func StringToByte(sizeStr string) uint64 { _ = "STUB: not implemented"; return 0 }

// ParseByte converts size string like 1GB/1g or 12mb/12M into an unsigned integer number of bytes
func ParseByte(sizeStr string) uint64 { _ = "STUB: not implemented"; return 0 }

// PrettyJSON get pretty Json string
func PrettyJSON(v any) (string, error) { _ = "STUB: not implemented"; return "", nil }

// ArgsWithSpaces it like Println, will add spaces for each argument
func ArgsWithSpaces(vs []any) (message string) { _ = "STUB: not implemented"; return "" }

// add space
