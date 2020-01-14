package main

import (
	"bytes"
	"fmt"
)

func main() {
	varParamFunc(1)
	varParamFunc(1, 2)

	var ii64 int64 = 3
	myPrintf(1, "2", ii64)

	fmt.Println(joinStrings("pig ", "and", " rat"))
}

// 可变参数函数
// 最后一个参数
// interface{}:任意类型
func varParamFunc(args ...interface{}) {
	// []interface：本质就是数组
	//fmt.Printf("%T\n", args)
	fmt.Println(args)
}

func myPrintf(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is an int value.")
		case string:
			fmt.Println(arg, "is a string value.")
		case int64:
			fmt.Println(arg, "is an int64 value.")
		default:
			fmt.Println(arg, "is an unknown type.")
		}
	}
}

// 拼接字符串
func joinStrings(slist ...string) string {
	// 定义一个字节缓冲，快速拼接字符串
	var b bytes.Buffer
	for _, s := range slist {
		b.WriteString(s)
	}

	return b.String()
}
