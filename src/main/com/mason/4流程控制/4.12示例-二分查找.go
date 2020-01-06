package main

import (
	"fmt"
	"sort"
)

// Binary Search
func main() {
	//定义一个数组
	arr := []int{1, 2, 5, 7, 15, 25, 30, 36, 39, 51, 67, 78, 80, 82, 85, 91, 92, 97}
	sort.Ints(arr)
	binarySearch(arr, 0, len(arr)-1, 30)
	fmt.Println("main arr=", arr)
}

func binarySearch(arr []int, low int, high int, value int) {
	if low > high {
		fmt.Println("Can't find...")
		return
	}
	var middle = (low + high) / 2
	var midValue = arr[middle]
	if midValue == value {
		fmt.Println("找到了，下标为：", middle)
	} else if midValue > value {
		binarySearch(arr, low, middle, value)
	} else {
		binarySearch(arr, middle, high, value)
	}
}
