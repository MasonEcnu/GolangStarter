package main

import "fmt"

// 构成递归需要具备以下条件：
// 1.一个问题可以被拆分成多个子问题；
// 2.拆分前的原问题与拆分后的子问题除了数据规模不同，但处理问题的思路是一样的；
// 3.不能无限制的调用本身，子问题需要有退出递归状态的条件。
// 注意：编写递归函数时，一定要有终止条件，否则就会无限调用下去，直到内存溢出。
func main() {
	for i := 1; i <= 5; i++ {
		fmt.Println(fibonacciWithRecursion(i))
	}

	fmt.Println(factorialWithRecursion(10))

	// 多个函数组成递归
}

func fibonacciWithRecursion(n int) (res int) {
	if n <= 2 {
		res = 1
	} else {
		res = fibonacciWithRecursion(n-1) + fibonacciWithRecursion(n-2)
	}
	return
}

func factorialWithRecursion(n uint64) (res uint64) {
	if n > 0 {
		res = n * factorialWithRecursion(n-1)
		return
	}
	return 1
}

// 多此一举
// 完全是为了示例啊
func even(nr int) bool {
	if nr == 0 {
		return true
	}
	return odd(revSign(nr) - 1)
}

func odd(nr int) bool {
	if nr == 0 {
		return false
	}
	return even(revSign(nr) - 1)
}

func revSign(nr int) int {
	if nr < 0 {
		return -nr
	}
	return nr
}
