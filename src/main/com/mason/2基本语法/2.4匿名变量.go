package main

// 匿名变量不占用内存空间，不会分配内存。匿名变量与匿名变量之间也不会因为多次声明而无法使用
func main() {
	a, _ := getData()
	println(a)
}

func getData() (int, int) {
	return 100, 200
}
