package main

import "fmt"

func main() {
	//声明一个二维切片
	var slice [][]int
	//为二维切片赋值
	slice = [][]int{{10}, {100, 200}}
	fmt.Println(len(slice))
	// slice := [][]int{{10}, {100, 200}}
	slice[0] = append(slice[0], 20)
}
