package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	// 第一个值是当前迭代到的索引位置，第二个值是该位置对应元素值的一份副本
	for index, value := range slice {
		fmt.Println(index, "-->", value)
	}

	// range 关键字不仅仅可以用来遍历切片，它还可以用来遍历数组、字符串、map 或者通道等
	str := "abcdefg"
	for index, value := range str {
		fmt.Printf("%T\n", value) // int32
		fmt.Println(index, "-->", string(value))
	}
}
