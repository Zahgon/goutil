package clipboard

// clipboard writer, reader program names
const (
	// WriterOnMac driver
	//
	// Example:
	//   echo hello | pbcopy
	//   pbcopy < tempfile.txt
	WriterOnMac = "pbcopy"

	// WriterOnWin driver on Windows
	//
	// TIP: clip only support write contents to clipboard.
	WriterOnWin = "clip"

	// WriterOnLin driver name
	//
	// linux:
	//   echo "hello-c" | xclip -selection c
	WriterOnLin = "xclip -selection clipboard"

	// ReaderOnMac driver
	//
	// Example:
	// 	Mac: pbpaste >> tasklist.txt
	ReaderOnMac = "pbpaste"

	// ReaderOnWin driver on Windows
	//
	// read clipboard should use: powershell get-clipboard
	ReaderOnWin = "powershell -NoProfile get-clipboard"

	// ReaderOnLin driver name
	//
	// Usage:
	// 	xclip -o -selection clipboard
	// 	xclip -o -selection c // can use shorts
	ReaderOnLin = "xclip -o -selection clipboard"
)

var (
	// TODO select an valid driver on New()
	writerOnLin = []string{"xclip", "xsel"} //lint:ignore U1000 for TODO

	// std instance
	std = New()
)

// Std get
func Std() *Clipboard {
	_ = "STUB: not implemented"

	// Reset clipboard data
	return nil
}

func Reset() error {
	_ = "STUB: not implemented"

	// Available clipboard available check
	return nil
}

func Available() bool { _ = "STUB: not implemented"; return false }

// ReadString contents from clipboard
func ReadString() (string, error) {
	_ = "STUB: not implemented"
	return "",

		// WriteString contents to clipboard and flush
		nil
}

func WriteString(s string) error { _ = "STUB: not implemented"; return nil }

// special handle on with args
func parseLine(line string) (bin string, args []string) { _ = "STUB: not implemented"; return "", nil }
