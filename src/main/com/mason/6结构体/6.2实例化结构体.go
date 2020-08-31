package main

import "fmt"

// 结构体的定义只是一种内存布局的描述，只有当结构体实例化时，才会真正地分配内存，
// 因此必须在定义结构体并实例化后才能使用结构体的字段。

type Product struct {
	Id   int    // 商品编号
	Name string // 商品名称
}

type Command struct {
	Name    string // 指令名称
	Var     *int   // 指令绑定的变量
	Comment string // 指令的注释
}

func newCommand(name string, varref *int, comment string) *Command {
	return &Command{
		Name:    name,
		Var:     varref,
		Comment: comment,
	}
}

func main() {
	var firstProduct Product
	firstProduct.Id = 1
	firstProduct.Name = "自行车"

	fmt.Println(firstProduct)

	secondProduct := new(Product)
	secondProduct.Id = 2
	secondProduct.Name = "摩托车"

	fmt.Println(secondProduct)

	thirdProduct := &Product{}
	thirdProduct.Id = 3
	thirdProduct.Name = "小轿车"

	fmt.Println(thirdProduct)

	var version int = 1
	cmd := &Command{}
	cmd.Name = "version"
	cmd.Var = &version
	cmd.Comment = "查询版本号"
	fmt.Println(cmd)

	newCmd := newCommand("version", &version, "查询版本号")
	fmt.Println(newCmd)

}
