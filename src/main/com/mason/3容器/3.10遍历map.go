package main

import (
	"fmt"
	"sort"
)

func main() {
	scene := make(map[string]int)
	scene["route"] = 66
	scene["brazil"] = 4
	scene["china"] = 960

	for key, value := range scene {
		fmt.Println(key, "-->", value)
	}

	// 顺序输出
	// 先排序
	var sceneList []string
	for key := range scene {
		sceneList = append(sceneList, key)
	}

	sort.Strings(sceneList)

	fmt.Println("key排序后：")
	for _, value := range sceneList {
		fmt.Println(value, "-->", scene[value])
	}
}
