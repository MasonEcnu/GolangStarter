package main

import "fmt"

func main() {
	var intSlice = []int{}
	intSlice = append(intSlice, 1)
	intSlice = append(intSlice, 1, 2)
	// 这冒号居然是必须的
	intSlice = append(intSlice, []int{1, 2, 3}...) // 追加一个切片, 切片需要解包
	// 容量的扩展规律是按容量的2倍数
	fmt.Println(intSlice)

	// 在切片开头添加元素一般都会导致内存的重新分配，而且会导致已有元素全部被复制1次
	var a = []int{1, 2, 3}
	a = append([]int{0}, a...)          // 在开头添加1个元素
	a = append([]int{-3, -2, -1}, a...) // 在开头添加1个切片

	//var b []int
	//b = append(b[:i], append([]int{x}, b[i:]...)...) // 在第i个位置插入x
	//b = append(b[:i], append([]int{1,2,3}, b[i:]...)...) // 在第i个位置插入切片
}
