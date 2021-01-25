package main

import (
	"fmt"
)

func a() {

}
func b() int {
	return 10
}

// 既然函数是一种类型，那么也可以当做参数类型来使用
//x 的类型是一个函数类型，无参数，int返回值的类型
func c(x func() int) int {
	ret := x()
	fmt.Println(ret)
	return x()
}

//函数也可以作为返回值
func f(x func() int) func(int, int) int {
	return d
}
func d(x int, y int) int {
	return x + y
}

func main() {

	q := a
	w := b

	//  q是函数类型，类型是 a 函数 无参数，无返回值
	// w 啊函数类型，类型是 b 函数，无参数，有返回值
	fmt.Printf("%T\n", q)
	fmt.Printf("%T\n", w)
	c(w)
	fmt.Printf("%T\n", d)
	// c(d)  因为d 函数与c函数的类型不一致，无法作为参数传入
	r := f(b)
	fmt.Printf("%T\n", r)

}
