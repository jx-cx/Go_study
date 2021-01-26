package main

import "fmt"

//将mint定义为int类型 - 自定义类型
type mint int

//将 uint 设置为int类型的别名 - 类型别名
type youint = int

func main() {
	var a mint
	a = 100
	var y youint
	y = 200
	fmt.Println(a, y)
	fmt.Printf("%T\n%T\n", a, y)
}
