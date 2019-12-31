package main

func main() {
	// 数组默认初始化，int->0，string->“”
	var intArr [10]int
	intArr = [10]int{1, 2, 3}
	println(intArr[0])
	println(intArr[len(intArr)-1])

	// 有初始化的元素数量，运行时决定长度
	stringArr := [...]string{"a", "b", "c"}
	println(cap(stringArr))

	// 比较两个数组是否相等
	// 类型相同+长度相同：才有比较的意义
	//println(intArr == stringArr)

	// 遍历数据
	println("传统for-i循环")
	for i := 0; i < len(stringArr); i++ {
		println(i, stringArr[i])
	}

	println("新型foreach循环")
	for k, v := range stringArr {
		println(k, v)
	}
}
