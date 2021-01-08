/*
整型分为以下两个大类： 按长度分为：int8、int16、int32、int64 对应的无符号整型：uint8、uint16、uint32、uint64

其中，uint8就是我们熟知的byte型，int16对应C语言中的short型，int64对应C语言中的long型。


类型	描述
uint8	无符号 8位整型 (0 到 255)
uint16	无符号 16位整型 (0 到 65535)
uint32	无符号 32位整型 (0 到 4294967295)
uint64	无符号 64位整型 (0 到 18446744073709551615)
int8	有符号 8位整型 (-128 到 127)
int16	有符号 16位整型 (-32768 到 32767)
int32	有符号 32位整型 (-2147483648 到 2147483647)
int64	有符号 64位整型 (-9223372036854775808 到 9223372036854775807)

特殊整型
类型	描述
uint	32位操作系统上就是uint32，64位操作系统上就是uint64
int	32位操作系统上就是int32，64位操作系统上就是int64
uintptr	无符号整型，用于存放一个指针

**/
package main

import "fmt"

func main() {
	var a int = 10 //十进制   0-9
	fmt.Printf("%d \n", a)
	fmt.Printf("%b \n", a) // %b 把十进制转换成二进制  0-1
	fmt.Printf("%o \n", a) // %b 把十进制转换成八进制
	fmt.Printf("%x \n", a) // %b 把十进制转换成十六进制
	var b int = 077        // 0开头代表八进制数  0-7
	fmt.Printf("%d \n", b)
	d := 0x123456 // 十六进制数   0- f
	fmt.Printf("%d \n", d)

	fmt.Printf("%T\n", a) //查看变量类型

	e := int8(9) //指定类型，否则默认int10
	fmt.Printf("%T\n", e)

}
