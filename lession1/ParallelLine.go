package main

import "fmt"

func main() {
	fmt.Println("平行线判断")
	var zero float64 = 0.0
	var line1Point1x float64
	var line1Point1y float64
	var line1Point2x float64
	var line1Point2y float64
	fmt.Print("请输入第一条线第一点x坐标: ")
	fmt.Scanln(&line1Point1x)
	fmt.Print("请输入第一条线第一点y坐标: ")
	fmt.Scanln(&line1Point1y)
	fmt.Print("请输入第一条线第二点x坐标: ")
	fmt.Scanln(&line1Point2x)
	fmt.Print("请输入第一条线第二点y坐标: ")
	fmt.Scanln(&line1Point2y)
	fmt.Println("第一条线坐标为(", line1Point1x, ", ", line1Point1y, ")  (", line1Point2x, ", ", line1Point2y, ") ")
	var line1DiffX = line1Point2x - line1Point1x
	var line1DiffY = line1Point2y - line1Point1y
	fmt.Println("第一条线 x距离:", line1DiffX, " y距离:", line1DiffY)

	var line2Point1x float64
	var line2Point1y float64
	var line2Point2x float64
	var line2Point2y float64
	fmt.Print("请输入第二条线第一点x坐标: ")
	fmt.Scanln(&line2Point1x)
	fmt.Print("请输入第二条线第一点y坐标: ")
	fmt.Scanln(&line2Point1y)
	fmt.Print("请输入第二条线第二点x坐标: ")
	fmt.Scanln(&line2Point2x)
	fmt.Print("请输入第二条线第二点y坐标: ")
	fmt.Scanln(&line2Point2y)
	fmt.Println("第二条线坐标为(", line2Point1x, ", ", line2Point1y, ")  (", line2Point2x, ", ", line2Point2y, ") ")
	var line2DiffX = line2Point2x - line2Point1x
	var line2DiffY = line2Point2y - line2Point1y
	fmt.Println("第二条线 x距离:", line2DiffX, " y距离:", line2DiffY)

	var result string
	if line1DiffX == zero && line2DiffX == zero {
		result = "Yes"
	} else {
		var k1 float64
		var k2 float64
		if line1DiffX != zero {
			k1 = line1DiffY / line1DiffX
		}
		if line2DiffX != zero {
			k2 = line2DiffY / line2DiffX
		}
		if k1 == k1 {
			result = "Yes"
		} else {
			result = "No"
		}
		fmt.Println("第一条线斜率:", k1, "  第二条线斜率:", k2)

	}

	fmt.Println("平行线判断结果: ", result)
}
