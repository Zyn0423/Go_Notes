package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func file_fun(conn net.Conn, file_name string) {
	f, err := os.Open(file_name)
	if err != nil {
		fmt.Println("os.Open", err)
		return
	}
	defer f.Close()

	buf := make([]byte, 1024*4)
	for {
		n, err := f.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客服端发送成功")
			} else {
				fmt.Println("os.Open", err)
			}
			return
		}
		_, err = conn.Write(buf[:n])
		if err != nil {
			fmt.Println("conn.Write", err)
			return
		}
	}
	//fmt.Println("客服端发送成功")

}

func main() {
	//	TODO 1.获取命令行字符串2.建立发送客服端 3.写入文件名给服务端。4.接收服务端发送OK  5.打开文件读，然后写入到缓冲区里发送给服务端
	file_name := os.Args
	if len(file_name) != 2 {
		fmt.Println("请输入要路径")
		return
	}

	FileInfo, err := os.Stat(file_name[1])
	if err != nil {
		fmt.Println("os.Stat", err)
		return
	}

	Conn, err := net.Dial("tcp", "127.0.0.1:8001")
	if err != nil {
		fmt.Println("net.Dial", err)
		return
	}
	defer Conn.Close()

	_, err = Conn.Write([]byte(FileInfo.Name()))

	if err != nil {
		fmt.Println("Conn.Write", err)
		return
	}
	fmt.Println("客服端发送成功")
	buf := make([]byte, 4096)
	n, err := Conn.Read(buf)
	if string(buf[:n]) == "ok" && err == nil {
		fmt.Println("")
		file_fun(Conn, file_name[1])
	} else {
		fmt.Println("接收失败")
	}

}
