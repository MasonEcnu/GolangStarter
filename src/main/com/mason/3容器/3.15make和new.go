package main

import (
	"fmt"
)

// new 和 make 是两个内置函数，主要用来创建并分配类型的内存
// new 只分配内存，而 make 只能用于 slice、map 和 channel 的初始化
type Student struct {
	name string
	age  int
}

func main() {
	// The new built-in function allocates memory. The first argument is a type,
	// not a value, and the value returned is a pointer to a newly
	// allocated zero value of that type.
	var sum *int      // 指针类型
	sum = new(int)    // 分配空间
	fmt.Println(*sum) // 只分配空间时，赋默认零值0
	*sum = 98         // 赋值
	fmt.Println(*sum)

	var s *Student
	s = new(Student) // 分配空间，如果不执行这步，会报错
	// panic: runtime error: invalid memory address or nil pointer dereference
	s.name = "mason"
	fmt.Println(s)

	// make 也是用于内存分配的，但是和 new 不同，它只用于 chan、map 以及 slice 的内存创建，而且它返回的类型就是这三个类型本身
	// The make built-in function allocates and initializes an object of type
	// slice, map, or chan (only). Like new, the first argument is a type, not a
	// value. Unlike new, make's return type is the same as the type of its
	// argument, not a pointer to it. The specification of the result depends on
	// the type:
	// Slice: The size specifies the length. The capacity of the slice is
	// equal to its length. A second integer argument may be provided to
	// specify a different capacity; it must be no smaller than the
	// length, so make([]int, 0, 10) allocates a slice of length 0 and
	// capacity 10.
	// Map: An empty map is allocated with enough space to hold the
	// specified number of elements. The size may be omitted, in which case
	// a small starting size is allocated.
	// Channel: The channel's buffer is initialized with the specified
	// buffer capacity. If zero, or the size is omitted, the channel is
	// unbuffered.
	// func make(t Type, size ...IntegerType) Type

	// Go语言中的 new 和 make 主要区别如下：
	// 1.make 只能用来分配及初始化类型为 slice、map、chan 的数据。new 可以分配任意类型的数据；
	// 2.new 分配返回的是指针，即类型 *Type。make 返回引用，即 Type；
	// 3.new 分配的空间被清零。make 分配空间后，会进行初始化；

	// 总结
	// make 关键字的主要作用是创建 slice、map 和 Channel 等内置的数据结构，而 new 的主要作用是为类型申请一片内存空间，并返回指向这片内存的指针。
}
