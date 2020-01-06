package main

import (
	"fmt"
	"math/rand"
)

// Binary Search
func main() {
	//定义一个数组
	var arr = make([]int, 10)
	for i := 0; i < 10; i++ {
		arr[i] = rand.Intn(10)
	}
	fmt.Println("排序前arr=", arr)
	bubbleSort(arr)
	fmt.Println("排序后arr=", arr)
}

func bubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[i] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}
