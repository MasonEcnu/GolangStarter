package main

func main() {
	var boolVar = 10
	println(boolVar == 10)

	i := 10
	println(i)

	// 布尔值并不会隐式转换为数字值 0 或 1，反之亦然，必须使用 if 语句显式的进行转换：
	var intVar = 10
	if intVar == 10 {
		println("123")
	} else {
		println("456")
	}
}
