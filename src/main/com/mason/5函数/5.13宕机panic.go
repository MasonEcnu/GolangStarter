package main

import (
	"regexp"
	"strconv"
)

func main() {
	mustCompile("123")
}

func mustCompile(str string) *regexp.Regexp {
	reg, err := regexp.Compile(str)
	if err != nil {
		panic(`regexp: Compile(` + strconv.Quote(str) + `): ` + err.Error())
	}
	return reg
}
