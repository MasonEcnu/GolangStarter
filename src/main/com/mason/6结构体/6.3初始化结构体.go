package main

import "fmt"

type People struct {
	name  string
	child *People
}

func main() {
	p := createPeopleByKeyValue()
	fmt.Println(p.name)
	anonymousStruct()
}

// 键值对
func createPeopleByKeyValue() *People {
	return &People{
		name: "爷爷",
		child: &People{
			name: "爸爸",
			child: &People{
				name: "我",
			},
		},
	}
}

type Address struct {
	Province    string
	City        string
	ZipCode     int
	PhoneNumber string
}

// 值列表
func createAddressByValueList() *Address {
	return &Address{
		"四川",
		"成都",
		610000,
		"0",
	}
}

// 匿名结构体
func anonymousStruct() {
	s := &struct {
		id   int
		data string
	}{
		1024,
		"hello",
	}
	printAnonymousStruct(s)
}

func printAnonymousStruct(msg *struct {
	id   int
	data string
}) {
	// 使用动词%T打印msg的类型
	fmt.Printf("%T\n", msg)
}
