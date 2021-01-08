package main

import "fmt"

/**
Go 语言中的25个关键字
    break        default      func         interface    select
    case         defer        go           map          struct
    chan         else         goto         package      switch
    const        fallthrough  if           range        type
	continue     for          import       return       var

Go语言中还有37个保留字
    Constants:    true  false  iota  nil

        Types:    int  int8  int16  int32  int64
                  uint  uint8  uint16  uint32  uint64  uintptr
                  float32  float64  complex128  complex64
                  bool  byte  rune  string  error

    Functions:   make  len  cap  new  append  copy  close  delete
                 complex  real  imag
                 panic  recover

*/

/*
声明变量
关键字  var
语法  var 变量名  变量类型  var studt string

批量声明变量

var (
	a int
	b string

)
**/
var (
	name string
	age  int
	isok bool
	//card
	// go语言中变量声明后必须使用，否则过不了编译
)

func main() {
	name = "小明"
	age = 18
	fmt.Print(isok)             //fmt 包是 打印  print 打印一句话，没有结束符
	fmt.Printf("name:%s", name) //printf 格式化输出   %s是占位符 使用name的值去替换%s
	fmt.Println(age)            // printlf 打印完指定的内容之后会在后面加一个换行符

}
