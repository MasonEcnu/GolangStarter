package main

import (
	"fmt"
	"net"
)

/*
每个变量会初始化其类型的默认值
1.整型和浮点型变量的默认值为 0 和 0.0。
2.字符串变量的默认值为空字符串。
3.布尔型变量默认为 bool。
4.切片、函数、指针变量的默认为 nil。
*/
func main() {
	var attack = 40
	var defence = 20
	var damageRate float32 = 0.23
	var damage = float32(attack-defence) * damageRate
	println(fmt.Sprintf("%.2f", damage))

	var hp int
	hp = 10
	println(hp)

	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	conn2, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		println(err.Error())
	}
	println("conn=", conn, " conn2:", conn2)

}
