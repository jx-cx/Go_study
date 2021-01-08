package main

import "fmt"

/*
常量


**/
const pi = 3.14159 //常量在定义的时候必须赋值

const (
	p1 = 1
	p2
	p3
) // 同时声明多个常量，如果省略了后面的值，则表示和上一行的值相同，即 p2 = p1 , p3=p2

//iota在const关键字出现时将被重置为0，const中每新增一行常量声明将使iota计数一次
const (
	a1 = iota //0
	a2        //1
	a3        //2
	a4        //3
)

const (
	c1 = iota //0
	c2 = 100  //100  iota 1
	c3 = iota //2
	c4        //3
)

const (
	d1, d2 = iota + 1, iota + 2 //d1=1 d2=2  iota=0
	d3, d4 = iota + 1, iota + 2 //d3=2,d4=3  iota=1
)

const (
	_  = iota
	kb = 1 << (10 * iota) // 1 往左移动10位转换为2进制=1024
	mb = 1 << (10 * iota) //
	gb = 1 << (10 * iota)
	tb = 1 << (10 * iota)
	pb = 1 << (10 * iota)
)

func main() {
	//pi = 3.14   常量的值在声明之后无法改变
	fmt.Println("p1=", p1)
	fmt.Println("p3=", p3)
	fmt.Println("a2=", a2)
	fmt.Println("a4=", a4)
	fmt.Println("c4=", a4)
}
