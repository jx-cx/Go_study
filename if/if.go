package main

import "fmt"

func main() {
	age := 60
	if age > 60 { // if age:=40;age>60  这种写法只在if条件作用域中生效
		fmt.Println("老骨头，蹦跶不动要死了")
	} else if age > 30 {
		fmt.Println("有钱的老男人")
	} else {
		fmt.Println("没钱的屌丝")
	}
}
