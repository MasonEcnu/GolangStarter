package main

import "fmt"

func main() {
	black := newBlackCat("Little Black")
	fmt.Println(&black)
}

// 模拟构造函数
type Cat struct {
	Name  string
	Color string
}

func newCatByName(name string) *Cat {
	return &Cat{
		Name: name,
	}
}

func newCatByColor(color string) *Cat {
	return &Cat{
		Color: color,
	}
}

// 父子继承关系通过组合实现
type BlackCat struct {
	Cat // 嵌入Cat，类似于派生
}

func newCat(name string) *Cat {
	return &Cat{
		Name: name,
	}
}

func newBlackCat(name string) *BlackCat {
	cat := &BlackCat{}
	cat.Name = name
	cat.Color = "Black"
	return cat
}
