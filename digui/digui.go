package main

import "fmt"

/*
递归：函数自己调用自己

*/
//阶乘
func f(n uint64) uint64 {
	if n <= 1 {
		return 1
	}
	return n * f(n-1)
}

func t(n uint64) uint64 {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return t(n-1) + t(n-2)
}

func main() {
	fmt.Println(f(5))
	fmt.Println(t(50))
}
