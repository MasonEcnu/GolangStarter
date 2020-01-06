package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// 基本写法
	var a = "hello"
	switch a {
	case "hello":
		fmt.Println(a)
	case "world":
		fmt.Println("...")
	}

	// 多值分支
	a = "num"
	switch a {
	case "num", "other":
		fmt.Println(a)
	}

	// 分支表达式
	var r int32 = rand.Int31()
	switch {
	case r%2 == 1:
		fmt.Println(r)
	case r > 10:
		fmt.Println(r)
	}

	// fallthrough
	a = "hello"
	switch {
	case a == "hello":
		fmt.Println(a)
		// 会继续执行下面的case语句
		fallthrough
	case a != "world":
		fmt.Println(a, "world")
	}

}
