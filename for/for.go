package main

/*
for(循环结构)
Go 语言中的所有循环类型均可以使用for关键字来完成。

for循环的基本格式如下：

for 初始语句;条件表达式;结束语句{
    循环体语句
}
条件表达式返回true时循环体不停地进行循环，直到条件表达式返回false时自动退出循环。

无限循环
for {
    循环体语句
}

for循环可以通过break、goto、return、panic语句强制退出循环。

**/

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
	fmt.Println()

	for a1 := 0; a1 < 10; a1++ {
		if a1 == 5 {
			break // 当a1=5时，break跳出循环
		}
		fmt.Println(a1)
	}
	fmt.Println("over")
	fmt.Println()

	for a1 := 0; a1 < 10; a1++ {
		if a1 == 5 {
			continue // 当a1=5时，continue跳出本次循环
		}
		fmt.Println(a1)
	}
	fmt.Println("over")

}
