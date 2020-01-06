package main

import "fmt"

// Go语言中的循环语句只支持 for 关键字
func main() {
	// 有限循环
	sum := 0
	for i := 1; i < 100; i++ {
		if i%2 == 1 {
			sum += i
		}
	}
	fmt.Println(sum)

	sum = 0
	// 无限循环
	for {
		sum++
		if sum > 100 {
			break
		}
	}
	fmt.Println(sum)

outer:
	for j := 0; j < 5; j++ {
	inner:
		for i := 0; i < 10; i++ {
			if i > 5 {
				break outer
			}
			if i > 2 {
				break inner
			}
			fmt.Println(i)
		}
	}

	step := 2
	for ; step < 10; step++ {
		fmt.Println(step)
	}

	var i int
	for i < 10 {
		i++
	}

}
