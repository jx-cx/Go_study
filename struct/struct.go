package main

import "fmt"

// 定义结构体 person
type person struct {
	name, gender string
	age          int
	hobby        []string
}

func c(gender person) {
	gender.gender = "女" //函数调用都是copy，无法改变其根本的值
}

func c1(gender *person) {
	(*gender).gender = "女" // 根据内存地址找到原来那个变量，修改的就是原来的变量
}
func c2(age *person) {
	age.age = 20 // 语法糖简写
}

type student struct {
	name string
	age  int
}

func main() {
	var xm person
	//1、结构体初始化
	xm.name = "小明"
	xm.gender = "男"
	xm.age = 18
	xm.hobby = []string{"爱钱", "爱江山", "爱美人"}
	//2 结构体 key-value初始化
	var a1 = person{
		name: "小红",
	}

	//3 使用值列表来进行初始化
	a2 := person{
		"小爱",
		"xx",
		60,
		[]string{"ss"},
	}
	fmt.Println(xm)
	fmt.Printf("%T\n", xm)
	//匿名结构体
	var a struct {
		a int
		b string
	}
	a.a = 100
	fmt.Printf("%T %v\n", a, a)
	c(xm)
	fmt.Println(xm.gender)
	c1(&xm)
	fmt.Println(xm.gender)
	c2(&xm)
	fmt.Println(xm.age)
	fmt.Println(a1, a2)
	fmt.Printf("%p %p %p %p\n", &xm.age, &xm.hobby, &xm.gender, &xm.name)

	m := make(map[string]*student)
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}

	for _, stu := range stus {
		m[stu.name] = &stu
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}

}
