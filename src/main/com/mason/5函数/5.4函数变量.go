package main

import "fmt"

func main() {
	var f func(string)
	f = hello
	f("hello")

	var fire func()
	fire = fireFunc
	fire()
}

func fireFunc() {
	fmt.Println("fireFunc in the hole")
}

func hello(str string) {
	fmt.Println(str)
}
