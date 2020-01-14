package main

import (
	"fmt"
	"os"
	"sync"
)

// defer 语句会将其后面跟随的语句进行延迟处理
// 在 defer 归属的函数即将返回时，将延迟处理的语句按 defer 的逆序进行执行，
// 也就是说，先被 defer 的语句最后被执行，最后被 defer 的语句，最先被执行
func main() {
	fmt.Println("defer begin")

	// 将defer放入延迟调用栈
	// 栈：先进后出
	defer fmt.Println(1)

	defer fmt.Println(2)

	defer fmt.Println(3)

	fmt.Println("defer end")

	// 使用延迟执行语句在函数退出时释放资源
	// 使用延迟并发解锁

}

var (
	valueByKey = make(map[string]int)

	valueByKeyGuard sync.Mutex
)

func readValue(key string) int {
	// 加锁
	valueByKeyGuard.Lock()
	// 取值
	v := valueByKey[key]
	valueByKeyGuard.Unlock()
	return v
}

func readValueWithDefer(key string) int {
	// 加锁
	valueByKeyGuard.Lock()
	// defer后面的语句不会马上调用, 而是延迟到函数结束时调用
	defer valueByKeyGuard.Unlock()
	return valueByKey[key]
}

// 使用延迟释放文件句柄
func fileSize(filename string) int64 {
	f, err := os.Open(filename)
	if err != nil {
		return 0
	}

	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		// defer机制触发, 调用Close关闭文件
		return 0
	}

	return info.Size()
}
