package main

import "fmt"

type student struct {
	id   int64
	name string
}
type marmoset struct { //创建一个学生管理者，只有此管理者才能访问下面四个方法
	allstudent map[int64]student
}

func (m marmoset) students() {
	for _, stu := range m.allstudent {
		fmt.Printf("学号:%d  姓名:%s \n", stu.id, stu.name)
	}
}

func (m marmoset) addstudent() {
	var (
		stupid  int64
		stuname string
	)
	fmt.Println("学号")
	fmt.Scanln(&stupid)
	fmt.Println("姓名")
	fmt.Scanln(&stuname)
	newstu := student{
		id:   stupid,
		name: stuname,
	}
	m.allstudent[newstu.id] = newstu

}
func (m marmoset) delstudent() {
	var stupid int64
	fmt.Printf("学号")
	fmt.Scanln(&stupid)
	_, ok := m.allstudent[stupid]
	if !ok {
		fmt.Println("这个人不存在")
		return
	}
	delete(m.allstudent, stupid)
	fmt.Println("删除成功")
}
func (m marmoset) upstudents() {
	var stupid int64
	fmt.Println("学号")
	fmt.Scanln(&stupid)
	stdid, ok := m.allstudent[stupid]
	if !ok {
		fmt.Println("这个人不存在")
		return
	}
	fmt.Println("输入要修改的学生学号")
	var upuser string
	fmt.Scanln(&upuser)
	stdid.name = upuser
	m.allstudent[stupid] = stdid
	fmt.Println("修改成功")

}
