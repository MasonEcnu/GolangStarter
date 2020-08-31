package main

import "fmt"

func main() {
	testConflict()
}

type SA struct {
	a int
}
type SB struct {
	a int
}
type SC struct {
	SA
	SB
}

func testConflict() {
	c := &SC{}
	c.SA.a = 1
	fmt.Printf("c is: %+v\n", *c)
}
