package main

import "fmt"

func main() {
	//基本格式
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println()

	// 写法 1
	i := 5
	for ; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println()

	// 写法 2
	b := 8
	for b < 10 {
		fmt.Println(b)
		b++
	}
	fmt.Println()

	//编写代码打印9*9乘法表
	// 遍历, 决定处理第几行
	for y := 1; y <= 9; y++ {
		// 遍历, 决定这一行有多少列
		for x := 1; x <= y; x++ {
			fmt.Printf("%d*%d=%d ", x, y, x*y)
		}
		// 没执行一次换行等于手动回车
		fmt.Println()
	}
	fmt.Println()

	//倒着打印乘法表
	for jj := 9; jj >= 1; jj-- {
		for ii := 1; ii <= jj; ii++ {
			fmt.Printf("%d*%d=%d  ", ii, jj, ii*jj)
		}
		fmt.Println()
	}
	return

}
