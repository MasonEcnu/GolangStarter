package main

import (
	"encoding/json"
	"fmt"
)

// 首字母小写表示字段对结构体外不可见，即就是private
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	p := &Person{}
	p.Name = "Mason"
	p.Age = 200
	data, _ := json.Marshal(p)
	fmt.Println(string(data))

	str := `{"Name":"Mason","Age":200}`
	fmt.Println(str)
}
