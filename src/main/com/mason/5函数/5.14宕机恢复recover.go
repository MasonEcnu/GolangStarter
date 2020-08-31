package main

import (
	"fmt"
	"runtime"
)

// Recover 是一个Go语言的内建函数
// 可以让进入宕机流程中的 goroutine 恢复过来，
// recover 仅在延迟函数 defer 中有效，在正常的执行过程中，调用 recover 会返回 nil 并且没有其他任何效果，
// 如果当前的 goroutine 陷入恐慌，调用 recover 可以捕获到 panic 的输入值，并且恢复正常的执行。

// 崩溃时需要传递的上下文信息
type panicContext struct {
	function string // 所在函数
}

func protectRun(entry func()) {
	// 延迟处理的函数
	defer func() {
		// 发生宕机时，获取panic传递的上下文并打印
		err := recover()
		switch err.(type) {
		case runtime.Error: // 运行时错误
			fmt.Println("runtime error:", err)
		default: // 非运行时错误
			fmt.Println("error:", err)
		}
	}()

	entry()
}

func main() {
	fmt.Println("运行前：")
	// 允许一段手动触发的错误
	protectRun(func() {
		fmt.Println("手动宕机前：")
		// 使用panic传递上下文
		panic(&panicContext{
			"手动触发panic",
		})

		fmt.Println("手动触发宕机后")
	})

	// 故意造成空指针访问错误
	protectRun(func() {
		fmt.Println("赋值宕机前：")

		var a *int
		*a = 1

		fmt.Println("赋值宕机后")
	})

	fmt.Println("运行后")
}
