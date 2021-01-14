package main

/*
if else(分支结构)
if条件判断基本写法
Go语言中if条件判断的格式如下：

if 表达式1 {
    分支1
} else if 表达式2 {
    分支2
} else{
    分支3
}
当表达式1的结果为true时，执行分支1，否则判断表达式2，如果满足则执行分支2，都不满足时，则执行分支3。 if判断中的else if和else都是可选的，可以根据实际需要进行选择。

Go语言规定与if匹配的左括号{必须与if和表达式放在同一行，{放在其他位置会触发编译错误。 同理，与else匹配的{也必须与else写在同一行，else也必须与上一个if或else if右边的大括号在同一行。
**/

import "fmt"

func main() {
	age := 60
	if age > 60 { // if age:=40;age>60  这种写法只在if条件作用域中生效
		fmt.Println("老骨头，蹦跶不动要死了")
	} else if age > 30 {
		fmt.Println("有钱的老男人")
	} else {
		fmt.Println("没钱的屌丝")
	}
}
