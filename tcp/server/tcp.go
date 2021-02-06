package main

import (
	"fmt"
	"net"
)

// tcp server端
func perocessConn(conn net.Conn) {
	// 与客户端通信
	for {
		var tmp [128]byte
		n, err := conn.Read(tmp[:])
		if err != nil {
			fmt.Println("通信失败，err:", err)
			return
		}
		fmt.Println(string(tmp[:n]))
	}
}

func main() {
	//本地端口启动服务
	listener, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("打开本地端口失败，err:", err)
		return
	}
	// 等待别人建立连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("建立连接失败，err:", err)
			return
		}
		go perocessConn(conn)
	}

}
