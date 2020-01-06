package main

import "fmt"

// continue 仅限在 for 循环内使用
func main() {
outer:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				fmt.Println(i, j)
				continue outer
			}
		}
	}
}
