package main

import (
	"fmt"
	"math/rand"
)

// if else
func main() {
	for i := 0; i < 10; i++ {
		var a = rand.Int31()
		fmt.Println(a)
		if a%2 == 0 {
			fmt.Println("Hello World!")
		} else {
			fmt.Println("Hello Hell")
		}
	}

	// 特殊语法
	// err:局部变量，只在if范围内生效
	// 返回值的作用范围被限制在 if、else 语句组合中
	if err := Connect(); err != nil {
		fmt.Println(err)
		return
	}
}

func Connect() error {
	return nil
}
