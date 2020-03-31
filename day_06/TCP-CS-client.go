package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Dial连接失败")
		return

	}
	defer conn.Close()
	//TODO 主动写给服务器
	_, err = conn.Write([]byte("are you ready?"))
	if err != nil {
		fmt.Println("conn.Write失败", err)
		return
	}

	buf := make([]byte, 1024*4)
	//TODO 接收服务器回发的数据
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read失败")
		return
	}
	fmt.Println("客服端接收到服务器发送过来的数据 ", string(buf[:n]))
}
