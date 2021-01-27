package main

import (
	"fmt"
	"os"
)

/*
   学生管理系统
*/

var srm marmoset

func menstudent() {
	fmt.Println("-----学生管理系统----")
	fmt.Println(`
	1、查看所有学生
	2、添加学生
	3、修改学习
	4、删除学生
	5、退出
		`)
}

func main() {
	srm = marmoset{
		allstudent: make(map[int64]student, 100),
	}
	for { //无限循环此函数
		menstudent()
		fmt.Print("请输入你的操作:")
		var choice int // 获取用户输入
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			srm.students()
		case 2:
			srm.addstudent()
		case 3:
			srm.upstudents()
		case 4:
			srm.delstudent()
		case 5:
			os.Exit(1)
		default:
			fmt.Println("无此操作选项，程序退出！")
		}
	}

}
