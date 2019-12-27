package main

import "math/cmplx"

// 基本数据类型--复数--complex
// 复数是由两个浮点数表示的，其中一个表示实部（real），一个表示虚部（imag）
// complex128（64 位实数和虚数）和 complex64（32 位实数和虚数），其中 complex128 为复数的默认类型
func main() {
	var a float64 = 123.1
	var b float64 = 456.4
	var c128 complex128 = complex(a, b)
	println(c128)
	println(real(c128))
	println(imag(c128))
	println(cmplx.Sqrt(c128))
}

// 复数计算规则
// 加减法
// z1=a+bi，z2=c+di
//  (a+bi)+(c+di)=(a+c)+(b+d)i

// z1=a+bi，z2=c+di
// (a+bi)-(c+di)=(a-c)+(b-d)i

// 乘除法
// z1=a+bi，z2=c+di
// (a+bi)(c+di)=(ac-bd)+(bc+ad)i

// 满足(c+di)(x+yi)=(a+bi)的复数x+yi(x,y∈R)叫复数a+bi除以复数c+di的商
