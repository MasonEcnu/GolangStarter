package main

import "fmt"

// copy( destSlice, srcSlice []T) int
// copy() 函数的返回值表示实际发生复制的元素个数
func main() {
	var srcSlice = []int{1, 2, 3}
	var destSlice = make([]int, len(srcSlice))
	copy(destSlice, srcSlice)
	fmt.Println(srcSlice, destSlice)

	// 设置元素数量
	const eleCount = 10

	// 预分配足够多的元素切片
	srcData := make([]int, eleCount)

	// 给切片赋值
	for i := 0; i < eleCount; i++ {
		srcData[i] = i + 1
	}

	// 引用切片数据
	// 引用是指向地址相同
	refData := srcData

	// 预分配足够多的元素切片
	copyData := make([]int, eleCount)
	// 将数据复制到新的切片空间中
	// 而复制，只复制内容
	copy(copyData, srcData)

	// 修改原始数据的第一个元素
	srcData[0] = 999

	// 打印引用切片的第一个元素
	fmt.Println(refData[0])

	// 打印复制切片的第一个和最后一个元素
	fmt.Println(copyData[0], copyData[eleCount-1])

	// 复制原始数据从4到6（不包含）
	copy(copyData, srcData[4:6])

	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", copyData[i])
	}
}
