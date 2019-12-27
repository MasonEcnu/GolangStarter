package main

func main() {
	testLocalVar()
	println(testFormVar(1, 2))
}

// 全局变量只需要在一个源文件中定义，就可以在所有源文件中使用
var globalA int = 789

// Go语言程序中全局变量与局部变量名称可以相同，但是函数体内的局部变量会被优先考虑

// 局部变量
func testLocalVar() {
	var localA int = 123
	var localB int = 456
	println(localA + localB + globalA)
}

// 形式参数
func testFormVar(i, j int) int {
	return i + j
}
