package main

func main() {
	var a, b int = 10, 11
	println(a, b)

	// ^：位操作，异或
	a = a ^ b
	b = b ^ a
	a = a ^ b
	println(a, b)

	// 从左到右顺序赋值
	a, b = b, a
	println(a, b)

	var arr = []int{1, 2, 3, 4, 5}
	swap(arr, 0, 1)
	println(arr[0])
}

func swap(p []int, i, j int) {
	p[i], p[j] = p[j], p[i]
}
