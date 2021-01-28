package main

/*
每个接口由数个方法组成，接口的定义格式如下：

type 接口类型名 interface{
    方法名1( 参数列表1 ) (返回值列表1,返回值列表2....)
    方法名2( 参数列表2 ) 返回值列表2
    …
}

在Go语言中接口（interface）是一种类型，一种抽象的类型。

interface是一组method的集合，是duck-type programming的一种体现。接口做的事情就像是定义一个协议（规则），只要一台机器有洗衣服和甩干的功能，我就称它为洗衣机。不关心属性（数据），只关心行为（方法）。

为了保护你的Go语言职业生涯，请牢记接口（interface）是一种类型。


一个对象只要全部实现了接口中的方法，那么就实现了这个接口。换句话说，接口就是一个需要实现的方法列表。

我们来定义一个Sayer接口：

// Sayer 接口
type Sayer interface {
	say()
}
定义dog和cat两个结构体：

type dog struct {}

type cat struct {}
因为Sayer接口里只有一个say方法，所以我们只需要给dog和cat 分别实现say方法就可以实现Sayer接口了。

// dog实现了Sayer接口
func (d dog) say() {
	fmt.Println("汪汪汪")
}

// cat实现了Sayer接口
func (c cat) say() {
	fmt.Println("喵喵喵")
}

接口的实现就是这么简单，只要实现了接口中的所有方法，就实现了这个接口。

*/
import "fmt"

type cat struct {
}

type dog struct {
}

type speaker interface {
	speak()
}

func (c cat) speak() {
	fmt.Println("喵喵喵")
}

func (d dog) speak() {
	fmt.Println("汪汪汪")
}

func da(x speaker) {
	x.speak()
}

func main() {
	var (
		c1 dog
		d1 cat
	)
	da(c1)
	da(d1)
}
