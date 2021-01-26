package main

/*
recover()必须搭配defer使用。
defer一定要在可能引发panic的语句之前定义。
*/
import "fmt"

func a() {
	fmt.Println("a")
}

func b() {
	defer func() {
		err := recover()
		fmt.Println(err)
	}()
	panic("haha，你服务器挂啦！")
	fmt.Println("b")

}
func c() {
	fmt.Println("c")
}

func main() {
	a()
	b()
	c()
}
