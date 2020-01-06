package main

import "fmt"

// goto 语句通过标签进行代码间的无条件跳转
func main() {
	// 跳出双重循环--001
outer:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 设置退出条件
				break outer
			}
		}
	}
	// 跳出双重循环--002
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				goto breakOuter
			}
		}
	}
breakOuter:
	fmt.Println("done")
}
