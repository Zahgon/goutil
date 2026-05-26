package strutil

import (

	// TODO use v2 on 1.22+

	"sync"
)

// global id:
//
//	https://github.com/rs/xid
//	https://github.com/satori/go.uuid
var (
	DefMinInt = 1000
	DefMaxInt = 9999
)

// MicroTimeID generate.
//   - return like: 16074145697981929446(len: 20)
//
// Conv Base:
//
//	mtId := MicroTimeID() // eg: 16935349145643425047 len: 20
//	b16id := Base10Conv(mtId, 16) // eg: eb067252154a9d17 len: 16
//	b32id := Base10Conv(mtId, 32) // eg: em1jia8akl78n len: 13
//	b36id := Base10Conv(mtId, 36) // eg: 3ko088phiuoev len: 13
//	b62id := Base10Conv(mtId, 62) // eg: kb24SKgsQ9V len: 11
func MicroTimeID() string { _ = "STUB: not implemented"; return "" }

// MicroTimeHexID micro time HEX ID generate.
//
// return like: 643d4cec7db9e(len: 13)
func MicroTimeHexID() string {
	_ = "STUB: not implemented"

	// MTimeHexID micro time HEX ID generate.
	//
	// return like: 643d4cec7db9e(len: 13)
	return ""
}

func MTimeHexID() string { _ = "STUB: not implemented"; return "" }

// MTimeBase36 micro time BASE36 id generate.
func MTimeBase36() string { _ = "STUB: not implemented"; return "" }

// MTimeBaseID micro time BASE id generate. toBase: 2-36
//
// Examples:
//   - toBase=16: 643d4cec7db9e(len: 13)
//   - toBase=36: hd312z9ka2(len: 10)
func MTimeBaseID(toBase int) string {
	_ = "STUB: not implemented"
	// eg: 1763431181849557
	return ""
}

// rand 1000 - 9999
// ri := mathutil.RandomInt(DefMinInt, DefMaxInt)

// DatetimeNo generate. can use for order-no.
//
//   - No prefix, return like: 2023041410484904074285478388(len: 28)
//   - With prefix, return like: prefix2023041410484904074285478388(len: 28 + len(prefix))
func DatetimeNo(prefix string) string { _ = "STUB: not implemented"; return "" }

// DateSN generate date serial number. PREFIX + yyyyMMddHHmmss + ext(微秒+随机数)
func DateSN(prefix string) string { _ = "STUB: not implemented"; return "" }

// micro datetime

// host

// eg: 4006367001

// rand 1000 - 9999
// rs := rand.New(rand.NewSource(nt.UnixNano()))

// DateSNOpt 基于时间生成唯一编号
type DateSNOpt struct {
	Layout string // time layout
	// RandMax   int    // rand max
	DateLen  int // 时间格式长度，后面部分将会进行进制转换 默认 8(yyyyMMdd)
	ConvBase int // DateLen 之后的转换 base 2-64. default 36
	// EnableSeq bool  // 需要高并发生成时可以启用自增序号。默认不启用
	SeqMaxVal int   // 自增序号最大值，之后后自动重置
	globalSeq int64 // 自增，确保同一时刻生成的编号不重复. EnableSeq=true 时启用
}

// default setting: {时间年到秒14位}
var defOpt = NewDateSNOpt()

// ConfigSNOpt config default date sn option
func ConfigSNOpt(fn func(opt *DateSNOpt)) {
	_ = "STUB: not implemented"

	// NewDateSNOpt create a new DateSNOpt instance.
	return
}

func NewDateSNOpt() *DateSNOpt { _ = "STUB: not implemented"; return nil }

// RandMax:   8999,

// EnableSeq: true,

// prepare for generate
func (do *DateSNOpt) prepare() { _ = "STUB: not implemented"; return }

// default 8 for yyyyMMdd

// get sequence max value (default 9999)

func (do *DateSNOpt) getSeqValue() int64 {
	_ = "STUB: not implemented"
	// use atomic sequence for guaranteed uniqueness (even without EnableSeq)
	// this ensures no collisions in tight loops
	return 0
}

// auto reset when seq exceeds SeqMaxVal (thread-safe using CAS)

// try to reset to 1, other goroutines may have already done it

// GenSN generate date serial number.
func (do *DateSNOpt) GenSN(prefix string) string { _ = "STUB: not implemented"; return "" }

// get time and format

// remove the dot separator if exists

// determine date length (default 8 for yyyyMMdd)

// high concurrency mode: sequence is the main differentiator

// convert extension to target base

// 确保同一时刻生成的编号不重复 max: 89999
var globalSeqSnV2 int64 = 0

func getSeqValue() int64 {
	_ = "STUB: not implemented"
	// use atomic sequence for guaranteed uniqueness (even without EnableSeq)
	// this ensures no collisions in tight loops
	return 0
}

// auto reset when seq exceeds SeqMaxVal (thread-safe using CAS)

// try to reset to 1, other goroutines may have already done it

// DateSNv2 generate date serial number.
//   - 2 < extBase <= 36
//   - return: PREFIX + yyyyMMddHHmmss + extBase(6bit micro + 5bit random number)
//
// Example:
//   - prefix=P, extBase=16, return: P2023091414361354b4490(len=22)
//   - prefix=P, extBase=36, return: P202309141436131gw3jg(len=21)
func DateSNv2(prefix string, extBase ...int) string { _ = "STUB: not implemented"; return "" }

// micro datetime

// 6bit micro + 5bit rand 10000 - 99999
// ext := strconv.AppendInt(bs[16+pl:], 10000+rand.Int63n(89999), 10)

// prefix + yyyyMMddHHmmss + ext(convert to base)

var (
	// key is dateLen + extBase
	dsoMap   = make(map[string]*DateSNOpt)
	dsoMutex sync.RWMutex
)

// DateSNv3 generate date serial number.
//   - 2 < extBase <= 64
//   - return: PREFIX + DATETIME(yyyyMMddHHmmss).dateLen + extBase(DATETIME.after+6bit micro + 5bit random number)
//   - dateLen: 为 DATETIME(yyyyMMddHHmmss) 保留的长度，默认为 8(yyyyMMdd) 后面的给 extBase 使用
//
// Example:
//   - prefix=P, dateLen=8, extBase=16, return: P202511139vs99gbifnj len: 20
//   - prefix=P, dateLen=6, extBase=36, return: P2025119yn52qhefati len: 19
//   - prefix=P, dateLen=6, extBase=48, return: P202511k9ksgD1fe6x len: 18
//   - prefix=P, dateLen=4, extBase=62, return: P2025aZl8N0y58M7 len: 16
func DateSNv3(prefix string, dateLen int, extBase ...int) string {
	_ = "STUB: not implemented"
	return ""
}

// double check，防止在加锁期间其他 goroutine 已经创建

// dso.EnableSeq = true
