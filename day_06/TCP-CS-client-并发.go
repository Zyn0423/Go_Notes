package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	//客服端的创建
	Conn, err := net.Dial("tcp", "127.0.0.1:8001")
	if err != nil {
		fmt.Println("net.Dia 已关闭")
	}
	defer Conn.Close()
	go func() {
		str := make([]byte, 1024*4)
		for {
			//TODO 获取用户键盘输入将输入数据发给服务端
			n, err := os.Stdin.Read(str)
			if err != nil {
				fmt.Println("os.Stdin.Rea", err)
				continue
			}
			Conn.Write(str[:n]) //TODO 写给服务器，读多少写多少
		}
	}()
	//TODO 回显数据
	buf := make([]byte, 1024*4)
	for {
		n, err := Conn.Read(buf)
		if n == 0 {
			fmt.Println("服务端以关闭,请重新连接")
			return
		}
		if err != nil {
			fmt.Println("Conn.Read(buf)", err)
			return
		}
		fmt.Println("客服端读到服务器的回发", string(buf[:n]))
	}

}
