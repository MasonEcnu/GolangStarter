package main

import (
	"fmt"
	"math"
)

/*
Go语言的基本类型有：
    bool
    string
    int、int8、int16、int32、int64
    uint、uint8、uint16、uint32、uint64、uintptr
    byte // uint8 的别名
    rune // int32 的别名 代表一个 Unicode 码
    float32、float64
    complex64、complex128

*/
func main() {
	var str string = "Hello World"
	println(str)
	var testInt8 int64 = math.MaxInt64
	println(testInt8)

	var (
		a int
		b string = "b"
		c []float32
		d func() bool
		e struct {
			x int
		}
	)
	e.x = 10
	d = testBoo
	fmt.Println(a, b, c, d, e)
}

func testBoo() bool {
	return true
}
