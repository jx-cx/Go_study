package main

import "fmt"

func assign(i interface{}) {
	fmt.Printf("%T\n", i)
	switch v := i.(type) {
	case string:
		fmt.Printf("x is a string，value is %v\n", v)
	case int:
		fmt.Printf("x is a int，value is %v\n", v)
	default:
		fmt.Printf("x isunsupport type！，value is %v\n", v)

	}
}
func main() {
	assign(12.22)
}
