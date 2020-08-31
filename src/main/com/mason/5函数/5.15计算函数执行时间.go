package main

import (
	"fmt"
	"time"
)

func test() {
	start := time.Now()
	sum := 0

	for i := 0; i < 100000000; i++ {
		sum++
	}
	elapsed := time.Now().Sub(start)
	//elapsed := time.Since(start)
	fmt.Println("该函数执行耗时：", elapsed)
}

func main() {
	test()
}
