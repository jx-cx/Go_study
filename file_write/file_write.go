package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func filewrite1() {
	fileobj, err := os.OpenFile("./xxx.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		fmt.Printf("open file failed err:%v", err)
		return
	}
	fileobj.Write([]byte("hah"))
	fileobj.WriteString("嘿嘿")
	fileobj.Close()
}
func filewrite2() {
	fileobj, err := os.OpenFile("./xxx.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		fmt.Printf("open file failed err:%v", err)
		return
	}
	defer fileobj.Close()
	wr := bufio.NewWriter(fileobj)
	wr.WriteString("呵呵")
	wr.Flush()
}

func filewrite3() {
	strs := "啦啦"
	err := ioutil.WriteFile("./xx.txt", []byte(strs), 0666)
	if err != nil {
		fmt.Printf("open file failed err:%v", err)
		return
	}

}

func filewrite4() {
	// 打开文件
	fileobj, err := os.OpenFile("./sb.txt", os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf(" 文件打开错误，%v", err)
		return
	}

	//因为没办法直接在文件中间插入内容，所以要借助临时文件
	tmpobj, err := os.OpenFile("./hb.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf(" 文件打开错误，%v", err)
		return
	}
	defer tmpobj.Close()
	// 读取要变更的文件内容写入到临时文件
	var ret [1]byte
	n, err := fileobj.Read(ret[:]) // 取这个文件的定义义内容
	if err != nil {
		fmt.Printf(" 文件打开错误，%v", err)
		return
	}
	// 写入临时文件
	tmpobj.Write(ret[:n])

	// 再写入要插入的内容
	var s []byte
	s = []byte{'c'}
	tmpobj.Write(s)
	// 紧接着把源文件后续的内容写入临时文件
	var x [1024]byte
	for {
		n, err := fileobj.Read(x[:])
		if err == io.EOF {
			tmpobj.Write(x[:n])
			break
		}
		if err != nil {
			fmt.Printf(" 文件打开错误，%v\n", err)
			return
		}
		tmpobj.Write(x[:n])
	}
	// 原文件写入临时文件中
	fileobj.Close()
	tmpobj.Close()

	//重命名源文件
	os.Rename("./hb.txt", "./sb.txt")

}

func main() {
	//filewrite1()
	//filewrite2()
	filewrite4()
}
