package main

// main包为go程序的入口
import (
	"fmt"
	"net/http"
)

func main() {
	// 使用 http.FileServer 文件服务器将当前目录作为根目录（/目录）的处理器，访问根目录，就会进入当前目录
	http.Handle("/", http.FileServer(http.Dir(".")))
	err := http.ListenAndServe(":8080", nil)
	fmt.Println(err)
}
