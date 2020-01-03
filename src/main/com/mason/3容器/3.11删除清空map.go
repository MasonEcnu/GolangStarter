package main

import (
	"fmt"
)

func main() {
	scene := make(map[string]int)
	scene["route"] = 66
	scene["brazil"] = 4
	scene["china"] = 960

	for key, value := range scene {
		fmt.Println(key, "-->", value)
	}

	fmt.Println("删除后：")
	// delete
	delete(scene, "brazil")
	for key, value := range scene {
		fmt.Println(key, "-->", value)
	}
	// 重新赋值一个空的map--->清空
	scene = make(map[string]int)
	fmt.Println(len(scene))
	for key, value := range scene {
		fmt.Println(key, "-->", value)
	}
}
