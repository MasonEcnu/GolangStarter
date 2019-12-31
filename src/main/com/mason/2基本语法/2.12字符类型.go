package main

import (
	"fmt"
	"unicode"
)

// byte和rune
// Go语言的字符有以下两种：
// 1.一种是 uint8 类型，或者叫 byte 型，代表了 ASCII 码的一个字符。
// 2.另一种是 rune 类型，代表一个 UTF-8 字符，当需要处理中文、日文或者其他复合字符时，则需要用到 rune 类型。
// rune 类型等价于 int32 类型。
func main() {
	// ASCII表中，A=65=十六进制的41
	var chb byte = 65
	var ch16 byte = '\x41'
	var ch8 byte = '\101'
	println(chb == ch16)
	println(chb == ch8)

	var ch int = '\u0041'
	var ch2 int = '\u03B2'
	var ch3 int = '\U00101234'
	fmt.Printf("%d - %d - %d\n", ch, ch2, ch3) // integer
	fmt.Printf("%c - %c - %c\n", ch, ch2, ch3) // character
	fmt.Printf("%X - %X - %X\n", ch, ch2, ch3) // UTF-8 bytes
	fmt.Printf("%U - %U - %U", ch, ch2, ch3)   // UTF-8 code point
	// 判断是否为字母：unicode.IsLetter(ch)
	// 判断是否为数字：unicode.IsDigit(ch)
	// 判断是否为空白符号：unicode.IsSpace(ch)
	var intVar = 'A'
	println("123")
	println(unicode.IsLetter(intVar))
}
