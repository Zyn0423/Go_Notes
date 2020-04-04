package main

import (
	"fmt"
	"net"
	"os"
)

func recvfile(conn net.Conn, file_name string) {
	f, err := os.Create(file_name)
	if err != nil {
		fmt.Println("os.Create", err)
		return
	}
	defer f.Close()
	buf := make([]byte, 1024*4)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("os.Create", err)
			return
		}
		if n == 0 {
			fmt.Println("接收完毕")
			return
		}
		_, err = f.Write(buf[:n])
		if err != nil {
			fmt.Println("f.Write", err)
			return
		}

	}

}

func main() {
	//		1.创建socket并建立监听 ,2.接收客服端发送的数据3.返回OK，4 创建一个新的文件 5.网络接收客服端发送的数据，并保存到创建的文件里
	listen, err := net.Listen("tcp", "127.0.0.1:8001")
	if err != nil {
		fmt.Println("net.Listen", err)
		return
	}
	defer listen.Close()

	Conn, err := listen.Accept()
	if err != nil {
		fmt.Println("listen.Accept", err)
		return
	}
	buf := make([]byte, 1024*4)
	n, err := Conn.Read(buf)
	if err != nil {
		fmt.Println("Conn.Read", err)
		return
	}
	Conn.Write([]byte("ok"))
	file_name := string(buf[:n])
	recvfile(Conn, file_name)

}

//Users/zxn/Downloads/index.html
