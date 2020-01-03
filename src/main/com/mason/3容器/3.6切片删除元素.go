package main

import "fmt"

func main() {
	var intSlice = []int{1, 2, 3, 4, 5}
	fmt.Println(intSlice)

	// 从开头删除
	var n = 1
	intSlice = intSlice[n:] // 删除开头的n个元素
	// intSlice[:0]：相当于置空原切片
	intSlice = append(intSlice[:0], intSlice[n:]...)   // 删除开头N个元素
	intSlice = intSlice[:copy(intSlice, intSlice[n:])] // 删除开头N个元素
	fmt.Println(intSlice)

	intSlice = []int{1, 2, 3, 4, 5}
	// 从中间位置删除
	var i = 1
	intSlice = append(intSlice[:i], intSlice[i+1:]...)         // 删除中间1个元素
	intSlice = append(intSlice[:i], intSlice[i+n:]...)         // 删除中间N个元素
	intSlice = intSlice[:i+copy(intSlice[i:], intSlice[i+1:])] // 删除中间1个元素
	intSlice = intSlice[:i+copy(intSlice[i:], intSlice[i+n:])] // 删除中间N个元素
	fmt.Println(intSlice)

	// 从尾部删除
	intSlice = []int{1, 2, 3, 4, 5}
	intSlice = intSlice[:len(intSlice)-1] // 删除尾部第一个元素
	intSlice = intSlice[:len(intSlice)-n] // 删除尾部第一个元素
	fmt.Println(intSlice)
}
