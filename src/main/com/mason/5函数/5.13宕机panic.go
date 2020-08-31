package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	defer fmt.Println("Hello World!")
	mustCompile("^{_+}@![]")
}

func mustCompile(str string) *regexp.Regexp {
	reg, err := regexp.Compile(str)
	if err != nil {
		// 手动触发宕机
		panic("regexp: Compile(" + strconv.Quote(str) + "): " + err.Error())
	}
	return reg
}
