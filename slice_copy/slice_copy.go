package main

import "fmt"

/*
由于切片是引用类型，所以a和b其实都指向了同一块内存地址。修改b的同时a的值也会发生变化。

Go语言内建的copy()函数可以迅速地将一个切片的数据复制到另外一个切片空间中，copy()函数的使用格式如下：

copy(destSlice, srcSlice []T)
**/
func main() {
	a1 := []int{1, 2, 3}
	a2 := a1 // 赋值，把a1赋值给a2
	var a3 = make([]int, 3)
	copy(a3, a1) //copy 把 a1 copy给a3
	fmt.Println(a1, a2, a3)
	a1[1] = 100
	fmt.Println(a1, a2, a3)
	/*
		Go语言中并没有删除切片元素的专用方法，我们可以使用切片本身的特性来删除元素 (使用append)
		**/
	//删除索引为1的元素
	a1 = append(a1[:1], a1[2:]...) //将 a1 第2个索引到最后的元素追加到 a1 从0至1的索引后面  1 3
	fmt.Println(a1)
	fmt.Println(cap(a1))
}
