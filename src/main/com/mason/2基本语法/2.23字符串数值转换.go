package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 整形转字符串
	num := 123
	str := strconv.Itoa(num)
	println(str)
	// 字符串转整形
	str = "123123a"
	num, err := strconv.Atoi(str)

	if err != nil {
		println(err.Error())
	} else {
		println(num)
	}

	// Parse 系列函数
	// Parse 系列函数用于将字符串转换为指定类型的值
	// 1、0、t、f、T、F、true、false、True、False、TRUE、FALSE
	boo, err := strconv.ParseBool("f")
	if err != nil {
		fmt.Printf("str1: %v\n", err)
	} else {
		fmt.Printf("boo: %v\n", boo)
	}
	// s string, base int, bitSize int
	// ErrSyntax:语法错误，ErrRange:超出范围
	num64, err := strconv.ParseInt("-110", 10, 0)
	println(num64)

	// Format 系列函数实现了将给定类型数据格式化为字符串类型的功能
	boo = false
	booStr := strconv.FormatBool(boo)
	fmt.Printf("type:%T,value:%v\n ", booStr, booStr)

	// Append 系列函数用于将指定类型转换成字符串后追加到一个切片中
	// 声明一个slice
	b10 := []byte("int (base 10):")
	// 将转换为10进制的string，追加到slice中
	b10 = strconv.AppendInt(b10, -42, 10)
	fmt.Println(string(b10))
	b16 := []byte("int (base 16):")
	b16 = strconv.AppendInt(b16, -42, 16)
	fmt.Println(string(b16))
}
