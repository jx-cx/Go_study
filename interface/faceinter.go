package main

import "fmt"

type falali struct {
	name string
}
type baoshijie struct {
	name string
}

func (f falali) run() {
	fmt.Printf("%s垃圾,速度50\n", f.name)
}

func (b baoshijie) run() {
	fmt.Printf("%s我喜欢，速度200\n", b.name)
}

type car interface {
	run()
}

func drive(c car) {
	c.run()
}

func main() {
	var f1 = falali{
		name: "法拉利",
	}
	var b1 = baoshijie{
		name: "保时捷",
	}
	drive(f1)
	drive(b1)

}
