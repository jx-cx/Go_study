package main

import "fmt"

/** Go语言中的defer语句会将其后面跟随的语句进行延迟处理。
在defer归属的函数即将返回时，将延迟处理的语句按defer定义的逆序进行执行，也就是说，先被defer的语句最后被执行，最后被defer的语句，最先被执行。
*/
func deferdemo() {
	fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3") // 多个defer存在的时候,先进后出的延迟执行 先把2放下面，再把3放2上面。先执行3再执行2
	fmt.Println("4")

}

// go语言中函数的return不是原子操作，在底层分两步来执行
// 第一步：返回值赋值
// defer在这两步中间
// 第二步：真的的ret返回
func f1() int {
	x := 5
	defer func() {
		x++ //修改的是x
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5 //返回值是x
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

func main() {
	deferdemo()
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}
