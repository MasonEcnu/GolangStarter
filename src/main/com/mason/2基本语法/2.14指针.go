package main

import (
	"flag"
	"fmt"
)

// 指针（pointer）在Go语言中可以被拆分为两个核心概念：
// 1.类型指针，允许对这个指针类型的数据进行修改，传递数据可以直接使用指针，而无须拷贝数据，类型指针不能进行偏移和运算。
// 2.切片，由指向起始元素的原始指针、元素数量和容量组成。
func main() {
	var cat int = 1
	var str string = "Hello World"
	// 指针的值是带有0x十六进制前缀的一组数据
	// 变量、指针和地址三者的关系是，每个变量都拥有地址，指针的值就是地址
	fmt.Printf("%p, %p\n", &cat, &str)

	var house = "Shanghai, China, 10086"
	ptr := &house
	// 打印ptr的类型
	fmt.Printf("ptr type:%T\n", ptr)
	// 打印ptr的指针地址
	fmt.Printf("ptr address:%p\n", ptr)
	// 对指针进行取值曹组
	value := *ptr
	// 取值后的类型
	fmt.Printf("value type:%T\n", value)
	// 指针取值后就是指向变量的值
	fmt.Printf("value: %s\n", value)

	// 变量、指针地址、指针变量、取地址、取值的相互关系和特性如下：
	// 1.对变量进行取地址操作使用&操作符，可以获得这个变量的指针变量。
	// 2.指针变量的值是指针地址。
	// 3.对指针变量进行取值操作使用*操作符，可以获得指针变量指向的原变量的值。

	// 使用指针变量进行数值交换
	x, y := 1, 2
	swapByPtr(&x, &y)
	fmt.Printf("%d, %d\n", x, y)

	parseCommanderLine()

	createPointByNew()
}

// *int:指向int类型变量的指针类型
func swapByPtr(a, b *int) {
	// *a：取指针a指向地址的值
	// *操作符作为右值时，意义是取指针的值，作为左值时，也就是放在赋值操作符的左边时，表示 a 指针指向的变量。
	// 其实归纳起来，*操作符的根本意义就是操作指针指向的变量。
	// 当操作在右值时，就是取指向变量的值，当操作在左值时，就是将值设置给指向的变量。
	t := *a
	*a = *b
	*b = t
}

var mode = flag.String("mode", "", "process mode")

func parseCommanderLine() {
	// 解析命令行参数
	flag.Parse()
	// 输出命令行参数
	fmt.Println(*mode)
}

func createPointByNew() {
	// 通过new函数创建指针
	// new() 函数可以创建一个对应类型的指针，创建过程会分配内存，被创建的指针指向默认值
	str := new(string)
	println(*str)
	*str = "Hello World"
	println(*str)
}
