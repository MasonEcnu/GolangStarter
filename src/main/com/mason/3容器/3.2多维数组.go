package main

func main() {
	var mulIntArr [4][2]int = [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
	println("二维数组-行数：", len(mulIntArr))
	println("二维数组-列数：", len(mulIntArr[0]))
	for outerK, outerV := range mulIntArr {
		for innerK, innerV := range outerV {
			println(outerK, ",", innerK, ",", innerV)
		}
	}
}
