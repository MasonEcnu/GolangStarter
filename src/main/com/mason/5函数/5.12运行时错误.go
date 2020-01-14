package main

import (
	"errors"
	"fmt"
)

// Go语言的错误处理思想及设计包含以下特征：
// 1.一个可能造成错误的函数，需要返回值中返回一个错误接口（error），
// 如果调用是成功的，错误接口将返回 nil，否则返回错误。
// 2.在函数调用后需要检查错误，如果发生错误，则进行必要的错误处理。

// 自定义一个错误
var errDivisionByZero = errors.New("division by zero")

// 声明一个解析错误
type ParseError struct {
	Filename string // 文件名
	Line     int    // 行号
}

// 实现error接口
func (e *ParseError) Error() string {
	return fmt.Sprintf("%s:%d", e.Filename, e.Line)
}

// 创建一些解析错误
func newParseError(filename string, line int) error {
	return &ParseError{filename, line}
}

func main() {
	fmt.Println(div(1, 0))

	// 示例：在解析中使用自定义错误
	var e error
	// 创建一个错误实例
	e = newParseError("main.go", 1)
	// 通过error结构查看错误信息
	fmt.Println(e.Error())

	// 根据错误接口的具体类型，获取详细错误信息
	switch detail := e.(type) {
	case *ParseError:
		fmt.Printf("Filename: %s Line: %d\n", detail.Filename, detail.Line)
	default: // 其他类型的错误
		fmt.Println("other error")
	}
}

func div(dividend, divisor int) (int, error) {
	// 判断除数为0
	if divisor == 0 {
		return 0, errDivisionByZero
	}
	return dividend / divisor, nil
}
