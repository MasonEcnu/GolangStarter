package main

import (
	"fmt"
	"math/rand"
)

// Go语言中闭包是引用了自由变量的函数
func main() {
	var str = "hello world"

	var foo = func() {
		str = "hello dude"
	}

	foo()

	fmt.Println(str)

	// 示例：闭包的记忆效应
	// 创建一个累加器, 初始值为1
	accumulator := accumulate(1)
	// 累加1并打印
	fmt.Println(accumulator())
	fmt.Println(accumulator())
	// 打印累加器的函数地址
	fmt.Printf("%p\n", accumulator)
	// 创建一个累加器, 初始值为1
	accumulator2 := accumulate(10)
	// 累加1并打印
	fmt.Println(accumulator2())
	// 打印累加器的函数地址
	fmt.Printf("%p\n", accumulator2)

	// 示例：闭包实现生成器
	var fib = fibonacciWithClosure()
	fmt.Println(fib())
	fmt.Println(fib())
	fmt.Println(fib())

	// 玩家生成器
	var pg = playerGen("Mason")
	for i := 0; i < 5; i++ {
		fmt.Println(pg())
	}
}

func accumulate(value int) func() int {
	return func() int {
		value++
		return value
	}
}

// 闭包的环境关联的循环变量a,b是引用方式而不是值传递方式
func fibonacciWithClosure() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func playerGen(name string) func() (string, int) {
	hp := 150

	return func() (string, int) {
		hp += rand.Intn(50)
		return name, hp
	}
}
