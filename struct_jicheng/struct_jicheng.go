package main

import "fmt"

type dog struct {
	feet   int
	animal // 继承了 animal的方法

}
type animal struct {
	name string
}

func (d dog) wang() {
	fmt.Printf("%s:叫\n", d.name)
}

func (a animal) move() {
	fmt.Printf("%s:会动\n", a.name)
}

func main() {
	a := dog{
		animal: animal{"小红"},
		feet:   4,
	}
	a.wang()
	a.move()

}
