package main

import (
	"fmt"
)

// for range 可以遍历数组、切片、字符串、map 及通道（channel）
// 我们总结一下 for 的功能：
// 1. Go语言的 for 包含初始化语句、条件表达式、结束语句，这 3 个部分均可缺省。
// 2.for range 支持对数组、切片、字符串、map、通道进行遍历操作。
// 3.在需要时，可以使用匿名变量对 for range 的变量进行选取。
func main() {
	var str = "Hello!"
	// value:始终为集合中对应索引的值拷贝
	for index, value := range str {
		fmt.Println(index, "-->", string(value))
	}

	m := map[string]int{
		"hello": 100,
		"world": 200,
	}
	for key, value := range m {
		fmt.Println(key, value)
	}

	// 遍历通道
	c := make(chan int)
	go func() {
		c <- 1
		c <- 2
		c <- 3
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}
}
