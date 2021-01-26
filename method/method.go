package main

import "fmt"

//方法是作用于特定类型的函数
//方法接收者是调用该方法的具体类型变量，一般用类型首字母小写
type dog struct {
	name string
}

func newDog(name string) dog {
	return dog{
		name: name,
	}
}

//方法写在函数名前面
func (d dog) wang() {
	fmt.Printf("%s:----\n", d.name)
}
func main() {
	xx := newDog("yy")
	xx.wang()
}
