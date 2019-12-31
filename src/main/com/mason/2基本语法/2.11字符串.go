package main

// 一个字符串是一个不可改变的字节序列
// 字符串是字节的定长数组
func main() {
	var str = "Hello World"
	var strLen = len(str)
	println(strLen)
	println(string(str[0]))
	println(str[strLen-1])
	// 获取字符串中某个字节的地址属于非法行为，例如 &str[i]
	var linesStr = `123
					456`
	// 多行字符串中，所有的转义字符均无效
	println(len(linesStr))
	for i := 0; i < len(str); i++ {
		println(string(str[i]))
	}
}
