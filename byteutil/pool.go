package byteutil

// ChanPool struct
//
// Usage:
//
//	bp := strutil.NewByteChanPool(500, 1024, 1024)
//	buf:=bp.Get()
//	defer bp.Put(buf)
//	// use buf do something ...
//
// refer https://www.flysnow.org/2020/08/21/golang-chan-byte-pool.html
// from https://github.com/minio/minio/blob/master/internal/bpool/bpool.go
type ChanPool struct {
	c    chan []byte
	w    int // init byte width
	wcap int // set byte cap
}

// NewChanPool instance
func NewChanPool(chSize int, width int, capWidth int) *ChanPool {
	_ = "STUB: not implemented"
	return nil
}

// Get gets a []byte from the BytePool, or creates a new one if none are
// available in the pool.
func (bp *ChanPool) Get() (b []byte) { _ = "STUB: not implemented"; return nil }

// reuse existing buffer

// create new buffer

// Put returns the given Buffer to the BytePool.
func (bp *ChanPool) Put(b []byte) { _ = "STUB: not implemented"; return }

// buffer went back into pool

// buffer didn't go back into pool, just discard

// Width returns the width of the byte arrays in this pool.
func (bp *ChanPool) Width() (n int) {
	_ = "STUB: not implemented"

	// WidthCap returns the cap width of the byte arrays in this pool.
	return 0
}

func (bp *ChanPool) WidthCap() (n int) { _ = "STUB: not implemented"; return 0 }
