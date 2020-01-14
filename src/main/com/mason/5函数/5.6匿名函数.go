package main

import (
	"flag"
	"fmt"
)

// 匿名函数没有函数名只有函数体
func main() {
	var anon = func(data int) {
		fmt.Println("Hello,", data)
	}

	anon(100)

	// 匿名函数用作回调
	var data = []int{1, 2, 3}
	visit(data, func(value int) {
		fmt.Println(value)
	})

	// 使用匿名函数实现操作封装
	fmt.Printf("%T\n", skillParam)

	flag.Parse()
	var skill = map[string]func(){
		"fire": func() {
			fmt.Println("chicken fire")
		},
		"run": func() {
			fmt.Println("soldier run")
		},
		"fly": func() {
			fmt.Println("angel fly")
		},
	}

	if f, ok := skill[*skillParam]; ok {
		f()
	} else {
		fmt.Println("skill not found")
	}
}

var skillParam = flag.String("skill", "", "skill to perform")

func visit(list []int, f func(int)) {
	for _, value := range list {
		f(value)
	}
}
