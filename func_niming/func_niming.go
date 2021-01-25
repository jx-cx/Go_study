package main

import "fmt"

//匿名函数

var q = func(x, y int) {
	fmt.Println(x + y)
}

func main() {
	q(1, 2)
	// 函数内部无法声明带名字的函数
	f1 := func(x, y int) {
		fmt.Println(x + y)
	}
	f1(1, 2)

	// 如果只是调用一次的函数，可以简写成立即执行函数
	func(x, y int) {
		fmt.Println(x + y)
	}(1, 2)

}
