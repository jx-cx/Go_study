package main

import (
	"fmt"
	"strings"
)

/*
Go语言中的字符串以原生数据类型出现，使用字符串就像使用其他原生数据类型（int、bool、float32、float64 等）一样。
Go 语言里的字符串的内部实现使用UTF-8编码。 字符串的值为双引号(")中的内容，可以在Go语言的源码中直接添加非ASCII码字符，例如：
s1 := "hello"
s2 := "你好"
Go 语言的字符串常见转义符包含回车、换行、单双引号、制表符等，如下表所示。

转义符	含义
\r	回车符（返回行首）
\n	换行符（直接跳到下一行的同列位置）
\t	制表符
\'	单引号
\"	双引号
\\	反斜杠

多行字符串
Go语言中要定义一个多行字符串时，就必须使用反引号字符：

s1 := `第一行
第二行
第三行
`
fmt.Println(s1)

字符串的常用操作
方法	介绍
len(str)	求长度
+或fmt.Sprintf	拼接字符串
strings.Split	分割
strings.contains	判断是否包含
strings.HasPrefix,strings.HasSuffix	前缀/后缀判断
strings.Index(),strings.LastIndex()	子串出现的位置
strings.Join(a[]string, sep string)	join操作
**/
func main() {
	str := "'D:\\Project\\Go_study\\string'"
	fmt.Println(str)
	str1 := "\"D:\\Project\\Go_study\\string\""
	fmt.Println(str1)
	s := `
	哈哈
	嘿嘿
	`
	fmt.Println(s)
	d := `D:\Project\Go_study\string\`
	fmt.Println(d)

	//字符串相关操作

	//字符串长度
	fmt.Println(len(d))

	name := "我"
	word := "sg"

	//字符串拼接
	s1 := name + word
	fmt.Printf(s1)
	n1 := fmt.Sprintf("%s%s", name, word) //fmt.printf("%s%s",name,word)
	fmt.Println(n1)

	//分割
	ret := strings.Split(d, "\\")
	fmt.Println(ret)

	//包含
	fmt.Println(strings.Contains(s1, "他"))
	fmt.Println(strings.Contains(s1, "我"))

	//前缀
	fmt.Println(strings.HasPrefix(s1, "我"))

	//后缀
	fmt.Println(strings.HasSuffix(s1, "我"))

	//子串出现的位置
	s2 := "abcdef"
	fmt.Println(strings.Index(s2, "c"))
	fmt.Println(strings.LastIndex(s2, "c"))

	//拼接
	fmt.Println(strings.Join(ret, "+"))

}
