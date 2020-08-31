package main

import "fmt"

func main() {
	testACar()
}

// 车轮
type AWheel struct {
	Size int
}

// 车
type ACar struct {
	AWheel
	// 引擎
	Engine struct {
		Power int    // 功率
		Type  string // 类型
	}
}

func testACar() {
	c := ACar{
		// 初始化轮子
		AWheel: AWheel{
			Size: 18,
		},
		// 初始化引擎
		Engine: struct {
			Power int
			Type  string
		}{
			Type:  "1.4T",
			Power: 143,
		},
	}
	fmt.Printf("%+v\n", c)
}
