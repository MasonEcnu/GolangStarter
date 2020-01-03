package main

import (
	"fmt"
	"unsafe"
)

// 在Go语言中，布尔类型的零值（初始值）为 false，数值类型的零值为 0，字符串类型的零值为空字符串""，
// 而指针、切片、映射、通道、函数和接口的零值则是 nil。
func main() {
	// nil
	// nil 标识符是不能比较的
	// fmt.Println(nil == nil)	// invalid operation: nil == nil (operator == not defined on nil)
	// nil 不是关键字或保留字
	//var nil = "我是nil"
	//fmt.Println(nil)
	// nil 没有默认类型
	fmt.Printf("%T\n", nil) // <nil>
	// 不同类型 nil 的指针是一样的
	var arr []int
	var num *int
	fmt.Printf("%p\n", arr) // 0x0
	fmt.Printf("%p\n", num)
	// 不同类型的 nil 是不能比较的
	//var m map[int]int
	//var ptr *int
	//fmt.Println(m == ptr)	// invalid operation: m == ptr (mismatched types map[int]int and *int)

	// 两个相同类型的 nil 值也可能无法比较
	//var s1 []int
	//var s2 []int
	//fmt.Println(s1 == s2)	// invalid operation: s1 == s2 (slice can only be compared to nil)
	// fmt.Println(s1 == nil)	// true

	// nil 是 map、slice、pointer、channel、func、interface 的零值
	var m map[int]string
	var ptr *int
	var c chan int
	var sl []int
	var f func()
	var i interface{}
	fmt.Printf("%#v\n", m)
	fmt.Printf("%#v\n", ptr)
	fmt.Printf("%#v\n", c)
	fmt.Printf("%#v\n", sl)
	fmt.Printf("%#v\n", f)
	fmt.Printf("%#v\n", i)
	// map[int]string(nil)
	// (*int)(nil)
	// (chan int)(nil)
	// []int(nil)
	// (func())(nil)
	// <nil>

	// 不同类型的 nil 值占用的内存大小可能是不一样的
	var p *struct{}
	fmt.Println(unsafe.Sizeof(p)) // 8
	var s []int
	fmt.Println(unsafe.Sizeof(s)) // 24
	var mm map[int]bool
	fmt.Println(unsafe.Sizeof(mm)) // 8
	var cc chan string
	fmt.Println(unsafe.Sizeof(cc)) // 8
	var ff func()
	fmt.Println(unsafe.Sizeof(ff)) // 8
	var ii interface{}
	fmt.Println(unsafe.Sizeof(ii)) // 16
}
