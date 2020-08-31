package main

import (
	"log"
	"runtime"
	"time"
)

func main() {
	entry()
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		runtime.GC()
	}
}

type Road int

func findRoad(r *Road) {
	log.Println("road:", *r)
}
func entry() {
	var rd Road = Road(999)
	r := &rd
	// finalizer函数设置
	runtime.SetFinalizer(r, findRoad)
}
