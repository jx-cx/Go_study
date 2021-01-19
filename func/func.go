package main

import "fmt"

/*
函数定义
Go语言中定义函数使用func关键字，具体格式如下：

func 函数名(参数)(返回值){
    函数体
}
函数名：由字母、数字、下划线组成。但函数名的第一个字母不能是数字。在同一个包内，函数名也称不能重名（包的概念详见后文）。
参数：参数由参数变量和参数变量的类型组成，多个参数之间使用,分隔。
返回值：返回值由返回值变量和其变量类型组成，也可以只写返回值的类型，多个返回值必须用()包裹，并用,分隔。
函数体：实现指定功能的代码块。

//Go语言中函数没有默认参数的概念
**/
// 函数的定义
//命名的返回值就相当于在函数中声明了一个变量，如下：声明了一个 int类型的ret变量
func sum(x int, y int) (ret int) {
	return x + y
	// ret = x+y
	// return ret   使用命名返回值可以省略后面的ret
}

// 没有返回值
func f1(x int, y int) {
	fmt.Println(x + y)
}

//没有参数没有返回值
func f2() {
	fmt.Print("f2")
}

//返回值可以命名也可以不命名
//没有参数有返回值的
func f3() int {
	ret := 3
	return ret
}
func ()  {
	
}

// 多个返回值
func f4() (int, string) {
	return 1, "hah"
}

//参数类型的简写，当参数类型多个且类型一致时，可以将最个一个的参数类型省略
func f5(x, y, z int, a, b, c string, i, o bool) int {
	return x + y
}

// 可变参数,必须放在函数参数的最后
func f6(x string, y ...int) {
	fmt.Println(x)
	fmt.Println(y)
}
func main() {
	fmt.Println(sum(1, 2))
	_, n := f4()
	fmt.Println(n)

	f6("hah1")
	f6("hah1", 1, 2, 3, 4, 5, 6, 7, 8, 9)
}
