package main

import (
	"fmt"
	"net"
)

func main() {
	Conn, err := net.Dial("udp", "127.0.0.1:8001")
	if err != nil {
		fmt.Println("net.Dial", err)
		return
	}
	defer Conn.Close()

	fmt.Println("客服端以建立连接成功")

	_, err = Conn.Write([]byte("are you ok?"))

	if err != nil {
		fmt.Println("Conn.Write", err)
		return
	}
	buf := make([]byte, 1024*4)
	n, err := Conn.Read(buf)
	fmt.Println("客服端接收到服务端发送过来数据：", string(buf[:n]))

}
