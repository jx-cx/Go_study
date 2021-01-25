package main

import "fmt"

// 闭包是一个函数，这个函数包含了他外部作用域的一个变量
//底层的原理：
// 1.函数可以作为返回值
// 2.函数内部查找变量的顺序，先在内部找，找不到再去外层找
func f1(f func()) {
	fmt.Println("f")
	f()
}

func f2(x, y int) {
	fmt.Println("f2")
	fmt.Println(x + y)
}

//将f2当做参数传到f1中，因为两个函数的类型不匹配，所有新建一个函数进行包装一下（套娃）
func f3(f func(int, int), n, m int) func() {
	tmp := func() {
		f(n, m)
	}
	return tmp

}

func adder(x int) func(int) int {
	return func(y int) int {
		x = x + y
		return x
	}
}

func main() {
	f1(f3(f2, 100, 200))
	ret := f3(f2, 100, 200)
	f1(ret)
	var f = adder(1)   // x=1
	fmt.Println(f(10)) //f =10 传到 adder函数中当做y使用
	fmt.Println(f(20))
	fmt.Println(f(30))

	f1 := adder(1)
	fmt.Println(f1(40))
	fmt.Println(f1(50))
}
