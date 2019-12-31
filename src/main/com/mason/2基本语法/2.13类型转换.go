package main

import (
	"fmt"
	"math"
)

// Go语言不存在隐式类型转换，因此所有的类型转换都必须显式的声明
func main() {
	// valueOfTypeB = typeB(valueOfTypeA)
	a := 1.9
	b := int(a)
	println(b)

	// 输出各种数值范围
	fmt.Println("int8 range:", math.MinInt8, math.MaxInt8)
	fmt.Println("int16 range:", math.MinInt16, math.MaxInt16)
	fmt.Println("int32 range:", math.MinInt32, math.MaxInt32)
	fmt.Println("int64 range:", math.MinInt64, math.MaxInt64)

	var x int32 = 1047483647
	fmt.Printf("int32: 0x%x %d\n", x, x)

	y := int16(x)
	fmt.Printf("int16: 0x%x %d\n", y, y)

	var z float32 = math.Pi
	fmt.Println(int(z))
}
