package main

import "fmt"

// 函数的基本组成为：关键字 func、函数名、参数列表、返回值、函数体和返回语句
// Go语言是编译型语言
// Go语言里面拥三种类型的函数：
// 1.普通的带有名字的函数
// 2.匿名函数或者 lambda 函数
// 3.方法

// Go语言没有默认参数值

// 在函数中，实参通过值传递的方式进行传递，因此函数的形参是实参的拷贝，对形参进行修改不会影响实参，
// 但是，如果实参包括引用类型，如指针、slice(切片)、map、function、channel 等类型，实参可能会由于函数的间接引用被修改。
func main() {
	fmt.Println(add(1, 2))
	fmt.Println(multiReturnValue())
	fmt.Println(multiReturnValueWithName())
}

// 普通函数
func add(a, b int) int {
	return a + b
}

// 多返回值
func multiReturnValue() (int, int) {
	return 1, 1
}

// 带有变量名的返回值
// 同一种类型返回值和命名返回值两种形式只能二选一
func multiReturnValueWithName() (a int, b int) {
	a = 1
	b = 2
	return
}
