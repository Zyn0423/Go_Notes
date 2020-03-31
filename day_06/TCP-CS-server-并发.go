package main

import (
	"fmt"
	"net"
	"runtime"
	"strings"
)

func HandlerConnect(conn net.Conn) {
	defer conn.Close()
	addr := conn.RemoteAddr() // TODO 客户端的IP地址
	fmt.Println(addr, "服务器阻塞等待客服端连接")
	buf := make([]byte, 1024*4) //TODO 循环客服端数据
	for {
		n, err := conn.Read(buf)
		//fmt.Println("byte",buf[:n])  //TODO ASCLL 字符排查
		if n == 0 { //TODO 强制关闭会导致进入此判断
			fmt.Println("服务端检测到客服端已关闭，断开连接 ！！！")
			return
		}

		if "exit\n" == string(buf[:n]) || "exit\r\n" == string(buf[:n]) {
			fmt.Println("服务端检测到客户端以退出的请求，服务端关闭！！！")
			return
		}
		if err != nil {
			fmt.Println("conn.Read", err)
			//goto Block
			//return
			runtime.Goexit()
		}
		fmt.Println("服务端接收的数据", string(buf[:n]))

		conn.Write([]byte(strings.ToUpper(string(buf[:n])))) //回传大写数据

	}
	//Block:
}
func main() {
	Listener, err := net.Listen("tcp", "127.0.0.1:8001")
	if err != nil {
		fmt.Println("net.Listen 失败 ", err)
	}
	defer Listener.Close()
	for {
		conn, err := Listener.Accept()
		if err != nil {
			fmt.Println("Listener.Accept 失败", err)
			return
		}
		go HandlerConnect(conn)

	}

}
