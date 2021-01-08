package main

/*
Go语言中以bool类型进行声明布尔型数据，布尔型数据只有true（真）和false（假）两个值。

注意：

布尔类型变量的默认值为false。
Go 语言中不允许将整型强制转换为布尔型.
布尔型无法参与数值运算，也无法与其他类型进行转换。

**/

import "fmt"

func main() {
	var a bool
	b := true
	fmt.Printf("%T\n", a)
	fmt.Printf("%T value:%v\n", b, b)

	

}
