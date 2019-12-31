package main

import (
	"fmt"
	"time"
)

// 常量是在编译时被创建的，即使定义在函数内部也是如此
// 只能是布尔型、数字型（整数型、浮点型和复数）和字符串型
// 定义常量的表达式必须为能被编译器求值的常量表达式

const (
	e  = 2.7182818
	pi = 3.1415926
)

var (
	a = 10
	b = 20
)

const noDelay time.Duration = 0
const timeout = 5 * time.Minute

// 命名类型
type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	fmt.Printf("%T %[1]v\n", noDelay)     // "time.Duration 0"
	fmt.Printf("%T %[1]v\n", timeout)     // "time.Duration 5m0s"
	fmt.Printf("%T %[1]v\n", time.Minute) // "time.Duration 1m0s"

	println(Sunday)
	println(Saturday)
}
