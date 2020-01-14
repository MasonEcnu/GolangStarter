package main

import (
	"fmt"
	"strings"
)

func main() {
	// 待处理的字符串列表
	list := []string{
		"go scanner",
		"go parser",
		"go compiler",
		"go printer",
		"go formater",
	}

	// 处理函数链
	chain := []func(string) string{
		removePrefix,
		strings.TrimSpace,
		strings.ToUpper,
	}

	// 处理字符串
	stringProcess(list, chain)

	for _, str := range list {
		fmt.Println(str)
	}
}

func stringProcess(list []string, chain []func(string) string) {
	// 遍历每一个字符串
	for index, str := range list {
		// 第一个需要处理的字符串
		result := str
		// 遍历每一个处理链
		for _, proc := range chain {
			// 输入一个字符串，并返回处理结果
			result = proc(result)
		}

		// 将结果放回切片
		list[index] = result
	}
}

// 自定义的移除前缀函数
func removePrefix(str string) string {
	return strings.TrimPrefix(str, "go")
}
