//go:build windows

package strutil

import (
	"math"
	"math/rand"
	"time"
)

const MaximumCapacity = math.MaxInt>>1 + 1

var rn = rand.NewSource(time.Now().UnixNano())

// nearestPowerOfTwo 返回一个大于等于cap的最近的2的整数次幂，参考java8的hashmap的tableSizeFor函数
//   - cap 输入参数
//
// 返回一个大于等于cap的最近的2的整数次幂
func nearestPowerOfTwo(cap int) int { _ = "STUB: not implemented"; return 0 }

// buildRandomString 生成随机字符串
//   - letters 字符串模板
//   - length 生成长度
//
// 返回一个指定长度的随机字符串
func buildRandomString(letters string, length int) string {
	_ = "STUB: not implemented"
	// 仿照strings.Builder
	// 创建一个长度为 length 的字节切片
	return ""
}

// letters的字符需要使用多少个比特位数才能表示完
// letterIdBits := int(math.Ceil(math.Log2(strLength))),下面比上面的代码快

// 最大的字母id掩码

// 可用次数的最大值

// UnixNano: 1607400451937462000
// 循环生成随机字符串

// 检查随机数生成器是否用尽所有随机数

// 从可用字符的字符串中随机选择一个字符

// 右移比特位数，为下次选择字符做准备

// 仿照strings.Builder用unsafe包返回一个字符串，避免拷贝
// 将字节切片转换为字符串并返回
