package main

import "fmt"

/*
构造函数
当结构体比较大的时候尽量使用结构结构体指针，健身程序的内存开销
*/

type person struct {
	name string
	age  int
}

func newperson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}

func main() {
	p1 := newperson("xx", 1)
	p2 := newperson("yy", 2)
	fmt.Println(p1, p2)

}
