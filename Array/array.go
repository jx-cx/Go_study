package main

import "fmt"

/*
数组是同一种数据类型元素的集合
在Go语言中，数组从声明时就确定，使用时可以修改数组成员，但是数组大小不可变化
比如：var a [5]int， 数组的长度必须是常量，并且长度是数组类型的一部分。
一旦定义，长度不能变。 [5]int和[10]int是不同的类型
数组可以通过下标进行访问，下标是从0开始，最后一个元素下标是：len-1
访问越界（下标在合法范围之外），则触发访问越界，会panic。
数组是值类型，赋值和传参会复制整个数组。因此改变副本的值，不会改变本身的值。
注意：

数组支持 “==“、”!=” 操作符，因为内存总是被初始化过的。
[n]*T表示指针数组，*[n]T表示数组指针 。
**/

func main() {
	//数组的定义：var 数组变量名 [元素数量]类型
	var a [3]int    //声明 一个int类型数组 a  长度为3
	var b [4]bool   //不初始化，默认都是零值，（bool:false,int and float :0 ,string:""
	var c [2]string //
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	//数组初始化
	c = [2]string{"哈哈", "嘿嘿"} //方式1
	fmt.Printf("%T\n", c)
	fmt.Println(c)
	d := [...]int{1, 2, 3} //方法2 根据初始值自动推断数组的长度
	fmt.Println(d)
	g := [...]int{1: 1, 3: 5} //方法3 使用索引来初始化数组
	fmt.Println(g)
	for i := 0; i < len(g); i++ { //方法1：for循环遍历，根据索引遍历 len是数组的长度
		fmt.Println(g[i])
	}
	// 方法2：for range遍历
	for index, value := range g {
		fmt.Println(index, value)
	}

	//多维数组
	//多维数组只有第一层可以使用...来让编译器推导数组长度
	//[[1 2][3 4][5 6]]
	var a1 [3][2]int
	a1 = [3][2]int{
		// [2]int{1, 2},  初始化多维数组
		// [2]int{3, 4},
		// [2]int{5, 6},
		{1, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Println(a1)

	// 多维数组的遍历
	for _, aa := range a1 {
		for _, bb := range aa {
			fmt.Println(bb)
		}
		fmt.Println()

	}
	q1 := [3]int{1, 2, 3}
	q2 := q1    //把q1 copy给q2
	q2[0] = 100 //修改的是q2的值，q1不变
	fmt.Println(q1, q2)

	var (
		s2 int
		s3 int
	)
	s1 := [...]int{1, 3, 5, 7, 8}
	for i := 0; i < len(s1); i++ {
		s2 += s1[i]
	}
	fmt.Println(s2)		

	for _, v1 := range s1 {
		s3 += v1
	}
	fmt.Println(s3)

	for i := 0; i < len(s1); i++ {
		for j := i; j < len(s1); j++ {
			if s1[i]+s1[j] == 8 {
				fmt.Println(i, j)
			}
		}
	}

}
