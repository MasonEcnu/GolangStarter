package main

import "fmt"

// 结构体可以包含一个或多个匿名（或内嵌）字段，即这些字段没有显式的名字
// 只有字段的类型是必须的，此时类型也就是字段的名字。
func main() {
	testInnerS()

	b := B{A{1, 2}, 3.0, 4.0}
	fmt.Printf("b is: %+v\n", b)
	fmt.Printf("b.A is: %+v\n", b.A)
}

// 匿名（内嵌）字段
type innerS struct {
	in1 int
	in2 int
}
type outerS struct {
	b      int
	c      float32
	int    // anonymous field
	innerS //anonymous field
}

func testInnerS() {
	outer := new(outerS)
	outer.b = 6
	outer.c = 7.5
	outer.int = 60
	outer.in1 = 5
	outer.in2 = 10
	fmt.Printf("outer.b is: %d\n", outer.b)
	fmt.Printf("outer.c is: %f\n", outer.c)
	fmt.Printf("outer.int is: %d\n", outer.int)
	fmt.Printf("outer.in1 is: %d\n", outer.in1)
	fmt.Printf("outer.in2 is: %d\n", outer.in2)
	// 使用结构体字面量
	outer2 := outerS{6, 7.5, 60, innerS{5, 10}}
	fmt.Printf("outer2 is:%+v\n", outer2)
}

// 内嵌结构体--可以实现类似继承的功能
type A struct {
	ax, ay int
}
type B struct {
	A
	bx, by float32
}
