package main

import "fmt"

func main() {
	var (
		a = 5
		b = 2
	)

	//算术运算符
	fmt.Println(a + b) // 相加
	fmt.Println(a - b) //相减
	fmt.Println(a * b) //相乘
	fmt.Println(a / b) //相除
	fmt.Println(a % b) //取余 5/3 =1 余 2
	a++                //++（自增）和--（自减）在Go语言中是单独的语句，并不是运算符
	b--
	fmt.Println(a, b)
	//关系运算符
	fmt.Println(a == b) //检查两个值是否相等，如果相等返回 True 否则返回 False  GO语言强类型，相同类型的变量才能比较.
	fmt.Println(a != b) //检查两个值是否不相等，如果不相等返回 True 否则返回 False。
	fmt.Println(a <= b) //检查左边值是否小于等于右边值，如果是返回 True 否则返回 False。
	fmt.Println(a >= b) //检查左边值是否大于等于右边值，如果是返回 True 否则返回 False。
	fmt.Println(a > b)  //检查左边值是否大于右边值，如果是返回 True 否则返回 False。
	fmt.Println(a < b)  //检查左边值是否小于右边值，如果是返回 True 否则返回 False。

	//逻辑运算符
	c := 25
	if c > 20 && c < 60 { //&& 如果两边的操作数都是 True，则为 True，否则为 False
		fmt.Println("上班的穷逼")
	} else {
		fmt.Println("没上班的穷学生")
	}
	if c > 20 || c < 60 { //||如果两边的操作数有一个 True，则为 True，否则为 False
		fmt.Println("上班的穷逼")
	} else {
		fmt.Println("没上班的穷学生")
	}
	isbool := false
	fmt.Println(!isbool) // ! 如果条件为 True，则为 False，否则为 True
	fmt.Println(isbool)

	//位运算符：针对二进制数
	//5的二进制是：101
	//2的二进制是：010

	fmt.Println(5 & 2)  // 参与运算的两数各对应的二进位相与（两位均为1才为1） 000
	fmt.Println(5 | 2)  //参与运算的两数各对应的二进位相或（两位有一个为1就为1） 111
	fmt.Println(5 ^ 2)  //参与运算的两数各对应的二进位相异或，当两对应的二进位相异时，结果为1（两位不一样则为1） 111
	fmt.Println(5 << 1) // 把 5（101）往左移1位 1010
	fmt.Println(5 >> 1) // 把 5（101）往右移1位 10

	fmt.Println()
	//赋值运算符
	var f int
	f = 1
	fmt.Println(f)
	f += 1 // f= f+1
	fmt.Println(f)
	f -= 1
	fmt.Println(f)
	f *= 1
	fmt.Println(f)
	f /= 1
	fmt.Println(f)
	f %= 1
	fmt.Println(f)
	f <<= 2
	fmt.Println(f)
	f >>= 2
	fmt.Println(f)
	f |= 2
	fmt.Println(f)
	f ^= 2
	fmt.Println(f)
	f &= 2

}
