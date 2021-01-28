package main

import "fmt"

type checking struct {
	feet int8
}

func (c checking) move() {
	fmt.Println("鸡动")
}

func (c checking) eat(food string) {
	fmt.Printf("鸡吃%s", food)
}

type cat1 struct {
	name string
	feet int8
}

func (c cat1) move() {
	fmt.Println("猫动")
}

func (c cat1) eat(food string) {
	fmt.Printf("猫吃%s \n", food)
}

type animal interface {
	eat(string)
	move()
}

func main() {
	var h1 animal
	fmt.Printf("%T\n", h1)
	bc := cat1{
		name: "淘气",
		feet: 4,
	}
	h1 = bc
	fmt.Println(h1)
	fmt.Printf("%T\n", h1)

	h1.eat("小黄鱼")
	ck := checking{
		feet: 4,
	}
	h1 = ck
	fmt.Printf("%T\n", h1)
	h1.eat("kfc")

}
