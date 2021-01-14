package main

import "fmt"

func main() {
	a := 5
	if a == 1 {
		fmt.Println("大拇指")
	} else if a == 2 {
		fmt.Println("食指")
	} else if a == 3 {
		fmt.Println("中指")
	} else if a == 4 {
		fmt.Println("无名指")
	} else if a == 5 {
		fmt.Println("小拇指")
	} else {
		fmt.Println("你TM没手吗")
	}

	//使用switch语句可方便地对大量的值进行条件判断。
	switch a { // switch a:4; a {}   这种写法a只在这switch中生效
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("你TM没手吗")
	}
	//Go语言规定每个switch只能有一个default分支。

	//一个分支可以有多个值，多个case值中间使用英文逗号分隔。
	switch n := 7; n {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	default:
		fmt.Println(n)
	}

	//分支还可以使用表达式，这时候switch语句后面不需要再跟判断变量。例如：
	age := 30
	switch {
	case age < 25:
		fmt.Println("好好学习吧")
	case age > 25 && age < 35:
		fmt.Println("好好工作吧")
	case age > 60:
		fmt.Println("好好享受吧")
	default:
		fmt.Println("活着真好")
	}

	//fallthrough语法可以执行满足条件的case的下一个case，是为了兼容C语言中的case设计的。=
	s := "a"
	switch {
	case s == "a":
		fmt.Println("a")
		fallthrough
	case s == "b":
		fmt.Println("b")
	case s == "c":
		fmt.Println("c")
	default:
		fmt.Println("...")
	}

	/*
			goto(跳转到指定标签)
		goto语句通过标签进行代码间的无条件跳转。goto语句可以在快速跳出循环、避免重复退出上有一定的帮助。
		Go语言中使用goto语句能简化一些代码的实现过程。
		例如双层嵌套的for循环要退出时
	**/
	var breakFlag bool
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 设置退出标签
				breakFlag = true
				break
			}
			fmt.Printf("%v-%v\n", i, j)
		}
		// 外层for循环判断
		fmt.Println("结束for循环")

		if breakFlag { //if breakFlag=true 则跳出外层for循环
			break
		}

	}

	//使用goto语句能简化代码：
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 设置退出标签
				goto breakTag
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	return
	// 标签
breakTag:
	fmt.Println("结束for循环")

}
