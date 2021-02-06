package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

//tcp client
func main() {
	// 与server端建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("与server端建立连接失败，err:", err)
		return
	}
	// 发送数据
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("发送：")
		msg, _ := reader.ReadString('\n')
		msg = strings.TrimSpace(msg)
		if msg == "exit" {
			break
		}
		conn.Write([]byte(msg))
	}
	conn.Close()
}
