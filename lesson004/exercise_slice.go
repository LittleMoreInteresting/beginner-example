package main

import "golang.org/x/tour/pic"

/**

实现 Pic函数.它应该返回一个长度为dy的切片，其中的每个元素都是dx 8位无符号整数的切片. 当你运行该程序时, 它将显示您的图片，将整数解释为灰度值.

图像的选择取决于您。有趣的函数包括（x+y）/2、x*y和x^y
您需要使用一个循环来分配[][]uint8中的每个[]uint8。
(使用uint8（intValue）在类型之间进行转换。)
*/
func Pic(dx, dy int) [][]uint8 {
	total := make([]uint8, dx*dy)
	res := make([][]uint8, dx)
	for x := 0; x < dx; x++ {
		res[x], total = total[0:dy], total[dy:]
		for y := 0; y < dy; y++ {
			res[x][y] = uint8(x * y)
		}
	}
	return res
}

func main() {
	pic.Show(Pic)
}
