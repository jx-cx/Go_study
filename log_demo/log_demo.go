package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	fileobj, err := os.OpenFile("./xx.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open file failed ,err:%s", err)
		return
	}
	log.SetOutput(fileobj) //将终端的输出传到fileobj文件中
	for {
		log.Println("测试日志！") // 向终端输入日志
		time.Sleep(time.Second * 3)
	}
}
