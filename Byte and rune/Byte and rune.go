package main

import (
	"fmt"
	"unicode"
)

/*
byte和rune类型
组成每个字符串的元素叫做“字符”，可以通过遍历或者单个获取字符串元素获得字符。 字符用单引号（’）包裹起来，如：

var a := '中'
var b := 'x'
Go 语言的字符有以下两种：

uint8类型，或者叫 byte 型，代表了ASCII码的一个字符。
rune类型，代表一个 UTF-8字符。
当需要处理中文、日文或者其他复合字符时，则需要用到rune类型。rune类型实际是一个int32。

Go 使用了特殊的 rune 类型来处理 Unicode，让基于 Unicode 的文本处理更为方便，也可以使用 byte 型进行默认字符串处理，性能和扩展性都有照顾。

**/

func main() {
	aa := "hello小红"
	for i := 0; i < len(aa); i++ {
		fmt.Printf("%v(%c)", aa[i], aa[i])
	}
	fmt.Println()
	for _, r := range aa {
		fmt.Printf("%v(%c)", r, r)
		// 		因为UTF8编码下一个中文汉字由3~4个字节组成，
		//所以我们不能简单的按照字节去遍历一个包含中文的字符串，否则就会出现上面输出中第一行的结果。
		// 字符串底层是一个byte数组，所以可以和[]byte类型相互转换
		//字符串是不能修改的 字符串是由byte字节组成，所以字符串的长度是byte字节的长度
		// rune类型用来表示utf8字符，一个rune字符由一个或多个byte组成。
	}
	fmt.Println()
	bb := []rune(aa) //把字符串强制转换成了一个rune切片
	bb[len(bb)-1] = '明'
	fmt.Printf(string(bb)) //把rune切片强制转换成字符串
	fmt.Println()
	/*
	   练习题
	   编写代码分别定义一个整型、浮点型、布尔型、字符串型变量，使用fmt.Printf()搭配%T分别打印出上述变量的值和类型。
	   编写代码统计出字符串"hello沙河小王子"中汉字的数量。
	   **/
	var (
		a = 1
		b = 1.1
		c = float64(2.1)
		d = true
		e = "哈哈"
	)
	fmt.Printf("a:%T b:%T c:%T d:%T e:%T", a, b, c, d, e)
	fmt.Println()

	sum := 0
	for _, char := range aa {
		if unicode.Is(unicode.Han, char) {
			sum++
		}
	}
	fmt.Println()
	fmt.Printf("中文字符数量是:%v", sum)

}
