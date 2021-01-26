package main

import "fmt"

// 定义结构体 person
type person struct {
	name, gender string
	age          int
	hobby        []string
}

func main() {
	var xm person
	//结构体实例化
	xm.name = "小明"
	xm.gender = "男"
	xm.age = 18
	xm.hobby = []string{"爱钱", "爱江山", "爱美人"}
	fmt.Println(xm)
	fmt.Printf("%T\n", xm)
	//匿名结构体
	var a struct {
		a int
		b string
	}
	a.a = 100
	fmt.Printf("%T %v\n", a, a)
}
