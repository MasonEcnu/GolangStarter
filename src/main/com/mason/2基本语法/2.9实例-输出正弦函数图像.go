package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

const SIN_IMAGE_SIZE int = 300

func main() {
	// 根据指定大小创建灰度图
	grayPic := image.NewGray(image.Rect(0, 0, SIN_IMAGE_SIZE, SIN_IMAGE_SIZE))
	// 遍历每个像素点
	for i := 0; i < SIN_IMAGE_SIZE; i++ {
		for j := 0; i < SIN_IMAGE_SIZE; j++ {
			// 填充为白色
			grayPic.SetGray(i, j, color.Gray{Y: 255})
		}
	}

	// 从0到最大像素生成x坐标
	for i := 0; i < SIN_IMAGE_SIZE; i++ {
		// 让sin的值的范围在0^2Pi之间
		s := float64(i*2) * math.Pi / float64(SIN_IMAGE_SIZE)
		// sin的幅度为一半的像素，向下偏移一半像素并翻转
		y := float64(SIN_IMAGE_SIZE/2) - math.Sin(s)*float64(SIN_IMAGE_SIZE)/2
		// 用黑色绘制sin曲线
		grayPic.SetGray(i, int(y), color.Gray{Y: 0})
	}

	// 创建文件
	file, err := os.Create("sin.png")
	if err != nil {
		log.Fatal(err)
	}

	// 使用png格式将数据写入文件
	// 将image信息写入文件中
	_ = png.Encode(file, grayPic)

	// 关闭文件
	_ = file.Close()
}
