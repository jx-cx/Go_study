package main

import "fmt"

/*
Go语言支持两种浮点型数：float32和float64。
这两种浮点型数据格式遵循IEEE 754标准： float32 的浮点数的最大范围约为 3.4e38，可以使用常量定义：math.MaxFloat32。
 float64 的浮点数的最大范围约为 1.8e308，可以使用一个常量定义：math.MaxFloat64。

**/
func main() {
	a := 1.234
	fmt.Printf("%T\n", a) // Go语言中的小数默认都是float64类型
	b := float32(1.234)   //显示成名float32类型
	fmt.Printf("%T\n", b)
	// b=a   //float64与float32类型不能直接赋值

	var c1 complex64
	c1 = 1 + 2i
	var c2 complex128
	c2 = 2 + 3i
	fmt.Println(c1)
	fmt.Println(c2)
	//复数有实部和虚部，complex64的实部和虚部为32位，complex128的实部和虚部为64位。

}
