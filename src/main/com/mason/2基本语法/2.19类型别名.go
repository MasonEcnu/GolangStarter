package main

import (
	"fmt"
	"reflect"
)

type byteInt8 int8
type byteAlias = int8

// 类型别名与类型定义
// 类型别名
// type TypeAlias = Type
// TypeAlias：类型只会在代码中存在，编译完成时，不会有TypeAlias类型

func testMain() {
	var bb byteInt8 = 1
	var bbb byteAlias = 1
	fmt.Printf("%T\n", bb) // main.byteInt8
	fmt.Printf("%T", bbb)  // int8
}

// 不能为不在一个包中的类型定义方法

// 定义商标结构
type Brand struct {
}

// 为商标结构定义show方法
func (brand Brand) show() {

}

// 为Brand定义一个别名
type FakeBrand = Brand

// 定义车辆结构
type Vehicle struct {
	// 嵌入两个结构
	FakeBrand
	Brand
}

func main() {
	// 声明变量a为车辆类型
	var a Vehicle
	// 指定调用FakeBrand的show
	a.FakeBrand.show()

	// 取a的类型反射对象
	ta := reflect.TypeOf(a)

	// 遍历a的所有成员
	for i := 0; i < ta.NumField(); i++ {
		// a的成员变量
		f := ta.Field(i)
		// 打印成员的字段名和类型
		fmt.Printf("FieldName: %v, FieldType: %v\n", f.Name, f.Type.Name())
	}
}
