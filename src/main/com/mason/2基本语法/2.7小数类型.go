package main

import (
	"fmt"
	"math"
)

// 可以直接写科学计数法
const Avogadro = 6.02214129e23 // 阿伏伽德罗常数
const Planck = 6.62606957e-34  // 普朗克常数
const PI = 3.1415926           // 圆周率

// 基本数据类型--小数--float
func main() {
	var maxFloat32 float32 = math.MaxFloat32
	println(maxFloat32)
	var maxFloat64 float64 = math.MaxFloat64
	println(maxFloat64)

	var f float32 = 16777216 // 1 << 24
	fmt.Println(f)
	fmt.Println(f + 1)
	fmt.Println(f == f+1) // "true"!

	println(math.Pi)
}
