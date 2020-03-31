package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8000") //(Listener, error)
	if err != nil {
		fmt.Println("创建连接失败。。。。。", err)
		return
	}
	defer listener.Close()
	fmt.Println("阻塞等待连接")
	Conn, err := listener.Accept() //Conn, error
	if err != nil {
		fmt.Println("服务端和客服端连接失败。。。")
		return
	}
	defer Conn.Close()
	buf := make([]byte, 1024*4)
	n, err := Conn.Read(buf)
	if err != nil {
		fmt.Println("Conn.Read 接收失败.....")
		return
	}
	fmt.Println("服务器读到的数据", string(buf[:n]))
}
