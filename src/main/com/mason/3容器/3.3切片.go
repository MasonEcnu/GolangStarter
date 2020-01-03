package main

import (
	"fmt"
)

// 切片（slice）是对数组的一个连续片段的引用，所以切片是一个引用类型
// 就是数组
// slice [开始位置 : 结束位置]
// 语法说明如下：
// 1.slice：表示目标切片对象；
// 2.开始位置：对应目标切片对象的索引；
// 3.结束位置：对应目标切片的结束索引。

// 从数组或切片生成新的切片拥有如下特性：
// 1.取出的元素数量为：结束位置 - 开始位置；
// 2.取出元素不包含结束位置对应的索引，切片最后一个元素使用 slice[len(slice)] 获取；
// 3.当缺省开始位置时，表示从连续区域开头到结束位置；
// 4.当缺省结束位置时，表示从开始位置到整个连续区域末尾；
// 5.两者同时缺省时，与切片本身等效；
// 6.两者同时为 0 时，等效于空切片，一般用于切片复位。
func main() {
	var arr = [3]int{1, 2, 3}
	fmt.Println(arr, arr[1:2])
	var slice = arr[1:2]
	fmt.Printf("slice type:%T\n", slice)

	// 1.从指定范围中生成切片
	// 跟python的list很像啊
	var sliceDemo [20]int
	for i := 0; i < 20; i++ {
		sliceDemo[i] = i + 1
	}
	// 以下生成的都是新的slice

	// 区间
	fmt.Println(sliceDemo[0:15])

	// 中间的到尾部的元素
	fmt.Println(sliceDemo[10:])

	// 从开头到中间的指定元素
	fmt.Println(sliceDemo[:5])

	// 全体
	var newSliceDemo = sliceDemo[:]
	fmt.Println(sliceDemo[:])

	fmt.Printf("%p\n", &newSliceDemo)
	fmt.Printf("%p\n", &sliceDemo)

	// 重置切片
	fmt.Println(sliceDemo[0:0])
	fmt.Println(sliceDemo)

	// 直接声明新的切片
	// 声明字符串切片
	var stringSlice []string
	// 声明整形切片
	var intSlice []int
	// 声明空切片
	var emptySlice = []int{}

	// 输出
	fmt.Println(stringSlice, intSlice, emptySlice)

	// 输出大小
	fmt.Println(len(stringSlice), len(intSlice), len(emptySlice))

	// 切片判空
	fmt.Println(stringSlice == nil) // true
	fmt.Println(intSlice == nil)    // true
	fmt.Println(emptySlice == nil)  // false
	// append：也是生成新的切片
	fmt.Println(append(intSlice, 1))
	fmt.Println(append(intSlice, 2))
	fmt.Println(intSlice)

	// 用make()构造切片
	var makeSlice = make([]int, 2)
	fmt.Println(makeSlice)
	// size 指的是为这个类型分配多少个元素，cap 为预分配的元素数量
	// 使用 make() 函数生成的切片一定发生了内存分配操作
	makeSlice = make([]int, 5, 10)
	fmt.Println(makeSlice)
}
