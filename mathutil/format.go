package mathutil

// DataSize format bytes number friendly. eg: 1024 => 1KB, 1024*1024 => 1MB
//
// Usage:
//
//	file, err := os.Open(path)
//	fl, err := file.Stat()
//	fmtSize := DataSize(fl.Size())
func DataSize(size uint64) string { _ = "STUB: not implemented"; return "" }

// FormatBytes Format the byte size to be a readable string. eg: 1024 => 1 KB
func FormatBytes(bytes int) string { _ = "STUB: not implemented"; return "" }

var timeFormats = [][]int{
	{0},
	{1},
	{2, 1},
	{60},
	{120, 60},
	{3600},
	{7200, 3600},
	{86400},
	{172800, 86400}, // second elem is unit.
	{2592000},
	{2592000 * 2, 2592000},
}

var timeMessages = []string{
	"< 1 sec", "1 sec", "secs", "1 min", "mins", "1 hr", "hrs", "1 day", "days", "1 month", "months",
}

// HowLongAgo format a seconds, get how lang ago. eg: 1 day, 1 week
func HowLongAgo(sec int64) string { _ = "STUB: not implemented"; return "" }

// next exists

// current <= intVal < next

// current is last

// match success

// He should never happen
