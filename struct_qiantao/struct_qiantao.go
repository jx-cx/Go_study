package main

import "fmt"

type a struct {
	id   int64
	name string
	c    //匿名嵌套结构体,可以直接使用访问
}
type b struct {
	card int64
	user string
	c
}

type c struct {
	addr string
}

func main() {
	p1 := a{
		1,
		"小红",
		c{
			"哈哈",
		},
	}
	p2 := b{
		2,
		"小明",
		c{
			"嘿嘿",
		},
	}
	fmt.Println(p1.addr) //直接访问匿名嵌套结构体
	fmt.Println(p2.c.addr)

}
