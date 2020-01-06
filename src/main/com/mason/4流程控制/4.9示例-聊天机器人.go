package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// 从标准输入读取数据
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please input your name:")
	// 读取数据
	input, err := inputReader.ReadString('\n')
	if err != nil {
		goto exitWhenErrorOccurred
	} else {
		// 用切片操作删除最后的\n
		name := input[:len(input)-1]
		fmt.Printf("Hello, %s! What can I do for you?\n", name)
	}

	for {
		input, err = inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("An error occurred:%s\n", err)
			continue
		}

		input = input[:len(input)-1]
		// 转小写
		input = strings.ToLower(input)
		switch input {
		case "":
			continue
		case "nothing", "bye":
			fmt.Println("Bye!")
			os.Exit(0)
		default:
			fmt.Println("Sorry, I didn't catch you.")
		}
	}

	// 发生异常时退出
exitWhenErrorOccurred:
	fmt.Printf("An error occurred:%s\n", err)
	// 异常退出
	os.Exit(1)
}
