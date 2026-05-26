package strutil

import (
	"github.com/gookit/goutil/byteutil"
)

// Buffer wrap and extends the bytes.Buffer
type Buffer = byteutil.Buffer

// NewBuffer instance, can set init size
func NewBuffer(initSize ...int) *Buffer { _ = "STUB: not implemented"; return nil }

// ByteChanPool struct
//
// Usage:
//
//	bp := strutil.NewByteChanPool(500, 1024, 1024)
//	buf:=bp.Get()
//	defer bp.Put(buf)
//	// use buf do something ...
type ByteChanPool = byteutil.ChanPool

// NewByteChanPool instance
func NewByteChanPool(maxSize, width, capWidth int) *ByteChanPool {
	_ = "STUB: not implemented"
	return nil
}
