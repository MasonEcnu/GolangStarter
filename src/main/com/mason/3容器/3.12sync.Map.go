package main

import (
	"fmt"
	"sync"
	"time"
)

// Go语言中的 map 在并发情况下，只读是线程安全的，同时读写是线程不安全的
func main() {
	// 创建一个int到int的映射
	//m := make(map[int]int)
	var m sync.Map
	// 开启一段并发代码
	go func() {
		// 不停地对map进行写入
		for {
			//m[1] = 1
			m.Store(1, 1)
		}
	}()
	// 开启一段并发代码
	go func() {
		// 不停地对map进行读取
		for {
			//_ = m[1]
			m.Load(1)
		}
	}()
	// 无限循环, 让并发程序在后台执行
	// 警告---代码还是可以执行的
	//for {
	//}
	// fatal error: concurrent map read and map write
	time.Sleep(1000)

	// sync.Map 有以下特性：
	// 1.无须初始化，直接声明即可。
	// 2.sync.Map 不能使用 map 的方式进行取值和设置等操作，而是使用 sync.Map 的方法进行调用，Store 表示存储，Load 表示获取，Delete 表示删除。
	// 3.使用 Range 配合一个回调函数进行遍历操作，通过回调函数返回内部遍历出来的值，
	// Range 参数中回调函数的返回值在需要继续迭代遍历时，返回 true，终止迭代遍历时，返回 false。

	var scene sync.Map
	// 将键值对保存到sync.Map
	scene.Store("greece", 97)
	scene.Store("london", 100)
	scene.Store("egypt", 200)

	// 取值
	value, err := scene.Load("greece")
	fmt.Println(value, err)

	// 删除
	scene.Delete("london")

	// 遍历
	scene.Range(func(key, value interface{}) bool {
		fmt.Println(key, "-->", value)
		return true
	})
}
