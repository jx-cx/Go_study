package main

import "fmt"

/*
Go语言的内建函数append()可以为切片动态添加元素。 可以一次添加一个元素，可以添加多个元素，也可以添加另一个切片中的元素（后面加…）。
每个切片会指向一个底层数组，这个数组的容量够用就添加新增元素。
当底层数组不能容纳新增的元素时，切片就会自动按照一定的策略进行“扩容”，此时该切片指向的底层数组就会更换。“扩容”操作往往发生在append()函数调用时，所以我们通常都需要用原变量接收append函数的返回值。
首先判断，如果新申请容量（cap）大于2倍的旧容量（old.cap），最终容量（newcap）就是新申请的容量（cap）。
否则判断，如果旧切片的长度小于1024，则最终容量(newcap)就是旧容量(old.cap)的两倍，即（newcap=doublecap），
否则判断，如果旧切片长度大于等于1024，则最终容量（newcap）从旧容量（old.cap）开始循环增加原来的1/4，即（newcap=old.cap,for {newcap += newcap/4}）直到最终容量（newcap）大于等于新申请的容量(cap)，即（newcap >= cap）
如果最终容量（cap）计算值溢出，则最终容量（cap）就是新申请容量（cap）。
需要注意的是，切片扩容还会根据切片中元素的类型不同而做不同的处理，比如int和string类型的处理方式就不一样。
注意：通过var声明的零值切片可以在append()函数直接使用，无需初始化。
var s []int
s = append(s, 1, 2, 3)
**/
func main() {
	s1 := []int{1, 2, 3}
	//调用append函数必须用原来的切片变量接收返回值
	//append 追加元素，原来的底层数组放不下的时候，go底层会把底层数组换一个
	s1 = append(s1, 4)
	fmt.Printf("s1:%v s1 len:%d s1 cap:%d\n", s1, len(s1), cap(s1))
	s2 := []int{5, 6}
	s1 = append(s1, s2...) // s2... 表示将s2拆开，全部元素追加到s1中
	fmt.Printf("s1:%v s1 len:%d s1 cap:%d\n", s1, len(s1), cap(s1))
	//

}
