package main

import "fmt"

/*
切片（Slice）是一个拥有相同类型元素的可变长度的序列。它是基于数组类型做的一层封装。它非常灵活，支持自动扩容。
切片是一个引用类型，它的内部结构包含地址、长度和容量。切片一般用于快速地操作一块数据集合。
语法：var name []T
**/
func main() {
	var a []int // 定义切片
	var b []string
	var c []int
	fmt.Println(a, b)
	//初始化
	a = []int{1, 2}
	b = []string{"hah", "heihei", "呵呵"}
	c = []int{1, 3}
	fmt.Println(a, len(a), c)

}
