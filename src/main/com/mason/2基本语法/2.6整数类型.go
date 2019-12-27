package main

import "math"

// 基本数据类型--整形
// 用来表示 Unicode 字符的 rune 类型和 int32 类型是等价的
// byte 和 uint8 也是等价类
// 无符号的整数类型 uintptr，它没有指定具体的 bit 大小但是足以容纳指针
func main() {
	var maxUInt64 uint64 = math.MaxUint64
	println(maxUInt64)
}
