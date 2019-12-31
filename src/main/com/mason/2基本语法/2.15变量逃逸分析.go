package main

import "fmt"

// 变量逃逸分析
// 通过编译器分析代码的特征和代码的生命周期，决定应该使用堆还是栈来进行内存分配

// 编译器觉得变量应该分配在堆和栈上的原则是：
// 1.变量是否被取地址；
// 2.变量是否发生逃逸。
func main() {
	//var a int
	//void()
	//fmt.Println(a, dummy(0))

	fmt.Println(structDummy())
}

// 测试入口参数和返回值情况
func dummy(b int) int {
	var c int
	c = b
	return c
}

func void() {

}

type Data struct {
}

func structDummy() *Data {
	var c Data
	return &c
}
