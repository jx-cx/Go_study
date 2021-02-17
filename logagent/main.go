package main

import (
	"Go_study/logagent/kafka"
	"Go_study/logagent/taillog"
	"fmt"
	"time"
)

// logAgent入口程序
func run() {
	// 读取日志
	select {
	//发送到kafka
	case line := <-taillog.ReadChan():
		kafka.SendTiKafka("web_log", line.Text)
	default:
		time.Sleep(time.Second)

	}

}
func main() {
	// 初始化kafka连接
	err := kafka.Init([]string{"127.0.0.1:9092"})
	if err != nil {
		fmt.Printf("init Kafka failed,err:%v\n", err)
		return
	}
	//打开日志文件准备收集日志
	err = taillog.Init("./my.log")
	if err != nil {
		fmt.Printf("Init taillog failed,err:%v\n", err)
		return
	}
	run()
}
