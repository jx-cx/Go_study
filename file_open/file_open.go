package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

/*
打开文件
*/

func fileread1() {
	fileos, err := os.Open("./file_open.go")
	if err != nil {
		fmt.Printf("open file failed,err:%v \n", err)
		return
	}
	//关闭文件
	defer fileos.Close()
	var ss [128]byte // 设定长度
	for {
		n, err := fileos.Read(ss[:]) //读取数组长度的内容
		if err == io.EOF {
			fmt.Println("读完了")
			return
		}
		if err != nil {
			fmt.Printf("read from file failed,err %v \n", err)
			return
		}
		fmt.Printf("读取了%d个字节\n", n)
		fmt.Println(string(ss[:n]))
		if n < 128 {
			return
		}
	}
}

func fileread2() {
	fileos, err := os.Open("./file_open.go")
	if err != nil {
		fmt.Printf("open file failed,err:%v \n", err)
		return
	}
	//关闭文件
	defer fileos.Close()
	// 创建一个用来从文件中读取内容的对象
	reader := bufio.NewReader(fileos)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			return

		}
		if err != nil {
			fmt.Printf("read from file failed,err %v \n", err)
			return
		}
		fmt.Println(str)

	}
}
func fileread3() {

	str, err := ioutil.ReadFile("./file_open.go")
	if err != nil {
		fmt.Printf("open file failed,err:%v \n", err)
		return
	}
	fmt.Println(string(str))

}

func main() {
	//fileread1()
	//fileread2()
	fileread3()
}
