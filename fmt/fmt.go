package main

import "fmt"

func main() {
	// fmt 占位符
	var (
		a = 100
		b = "ha"
		c = "true"
		d = 1.234
	)
	//查看变量类型
	fmt.Printf("%T\n", a)
	//查看值
	fmt.Printf("%v\n", a)
	//查看二进制对应的数
	fmt.Printf("%b\n", a)
	//查看十进制的值
	fmt.Printf("%d\n", a)
	//查看八进制的值
	fmt.Printf("%o\n", a)
	//查看十六进制的值
	fmt.Printf("%x\n", a)
	//查看字符串类型的值
	fmt.Printf("%s\n", b)
	// 打印具体的类型
	fmt.Printf("%#v\n", b)
	fmt.Printf("%#v\n", c)
	fmt.Printf("%#v\n", d)
}
