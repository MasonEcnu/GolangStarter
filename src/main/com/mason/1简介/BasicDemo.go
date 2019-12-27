package main

func main() {
	for i := 0; i < 10; i++ {
		println(i)
	}

	flag := true
	if flag {
		println("Hello World!")
	}
	i := 10
	// i++：是一条语句，不再是表达式/操作符
	i++
	println(i)
}
