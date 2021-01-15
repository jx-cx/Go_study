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
	fmt.Println(a, b, c)
	fmt.Println(a == nil) //a 是否为nil (nil表示空，没有开辟内存空间)
	// fmt.Println(a == b) 切片是引用类型，不支持直接比较，只能和nil比较
	fmt.Println(b == nil)

	/*
	   切片的长度和容量
	   切片拥有自己的长度和容量，我们可以通过使用内置的len()函数求长度，使用内置的cap()函数求切片的容量。
	   切片表达式从字符串、数组、指向数组或切片的指针构造子字符串或切片。
	   它有两种变体：一种指定low和high两个索引界限值的简单的形式，另一种是除了low和high索引界限值外还指定容量的完整的形式。
	   切片的底层就是一个数组，所以我们可以基于数组通过切片表达式得到切片。
	   切片表达式中的low和high表示一个索引范围（左包含，右不包含），也就是下面代码中从数组a中选出1<=索引值<4的元素组成切片s，得到的切片长度=high-low，容量等于得到的切片的底层数组的容量。
	   **/
	fmt.Printf("len a:%d cap a:%d\n", len(a), cap(a)) //长度和容量  len表示长度。cap表示容量
	fmt.Printf("len b:%d cap b:%d\n", len(b), cap(b))
	a1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 11, 12, 13}
	a2 := a1[0:7] //基于一个数组切割，左包含又不包含（左闭右开） 不包含第7个元素
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Printf("len a2:%d cap a2:%d\n", len(a2), cap(a2)) //切片的容量是底层数组第一个元素到最后的元素数量
	a3 := a1[2:]                                          // 等同于 a1[2:len(a)]
	fmt.Printf("len a3:%d cap a3:%d\n", len(a3), cap(a3))
	a4 := a1[:3] // 等同于 a1[0:3]
	fmt.Printf("len a4:%d cap a4:%d\n", len(a4), cap(a4))
	a5 := a1[:] // 等同于 a1[0:len(a)]
	fmt.Printf("len a5:%d cap a5:%d\n", len(a5), cap(a5))
	//切片再切片
	a6 := a3[3:]
	fmt.Printf("len a6:%d cap a6:%d\n", len(a6), cap(a6))
	a1[12] = 1000 // 修改了底层数组，所有后面基于这个数组的切片的值也会随之变化
	fmt.Println(a6)

	

}
