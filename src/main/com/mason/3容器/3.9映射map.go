package main

import "fmt"

// 一种元素对（pair）的无序集合
// 关联数组或字典
func main() {
	// var mapname map[keytype]valuetype
	// 1.mapname 为 map 的变量名。
	// 2.keytype 为键类型。
	// 3.valuetype 是键对应的值类型。

	// var mapLit map[string]int:这种方式生命的map没有初始化，值是nil
	// make：该方式会分配内存
	var firstMap = make(map[int]int)
	firstMap[1] = 1
	fmt.Println(firstMap)

	var secondMap = map[string]int{"one": 1, "two": 2}
	secondMap["third"] = 3
	fmt.Println(secondMap)

	// 10：容量
	// map 的大小会自动加 1
	var thirdMap = make(map[int]int, 10)
	thirdMap[123] = 456
	fmt.Println(thirdMap)
	fmt.Println(thirdMap[122])
}
