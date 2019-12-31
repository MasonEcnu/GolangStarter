package main

import "fmt"

type Weapon int

const (
	Arrow Weapon = iota
	Shuriken
	SniperRifle
	Rifle
	Blower
)

const (
	FlagNone = 1 << iota
	FlagRed
	FlagGreen
	FlagBlue
)

// 声明芯片类型
type ChipType int

const (
	None ChipType = iota
	CPU
	GPU
)

// 扩展函数--kotlin
func (c ChipType) toString() string {
	switch c {
	case None:
		return "None"
	case CPU:
		return "CPU"
	case GPU:
		return "GPU"
	}
	return "N/A"
}

func main() {
	var weapon Weapon = Rifle
	println(weapon)

	println(FlagNone, FlagRed, FlagGreen, FlagBlue)
	fmt.Printf("%b %b %b %b\n", FlagNone, FlagRed, FlagGreen, FlagBlue)

	println(CPU.toString(), CPU)

	println(ChipType.toString(99))
}
