package main

import (
	"fmt"
	"net"
)

//广播用户上线：
//        1.主go程中，创建监听套接字， defer  .close()
//        2.for 循环监听客服端连接请求accept()
//        3.有一个客户端连接，创建新Go程处理客户端数据handlerConnet(conn)  defer
//        4.定义全局结构体类型  C Name adder
//        5.创建全局map channel
//        6.实现HandlerConnet,获取客服端IP+ PORT  ---Remoteadder() 初始化新用户结构信息  name ==adder
//        7.创建Manager 实现管理GO程  在accept()之前
//        8.实现Manager 初始化在线用户map，循环读取channel 如果没有数据阻塞，如果有数据遍历在线用户map，讲数据写到用户C里
//        9.讲新的用户添加到在线用户map中。key(adder+port):value(新用户结构体)
//        10.创建WriteMsgToClient Go 程，专门给当前用户写数据  ，来源用用户自带的C中
//        11.实现WriteMsgToClient(clnt,conn)。遍历自带的C，读数据，conn.Write()到客服端
//        12.HandlerConnet中，结束位置，组织用户上线信息，将用户上线信息写到全局channel --Manager 的读就被激活，否则阻塞
//        13.HandlerConnet中加 结尾加 for{;}

type Client struct {
	C     chan string
	Name  string
	Adder string
}

// 创建全局map，存储在线用户
var onlineMap map[string]Client

var messages = make(chan string) //创建全局channel传递用户信息

func WriteMsgToClient(clnt Client, conn net.Conn) {
	for msg := range clnt.C {
		conn.Write([]byte(msg + "\n"))
	}

}
func handlerConnet(conn net.Conn) {
	defer conn.Close()

	netAddr := conn.RemoteAddr().String()
	fmt.Printf("%T%s\n", netAddr, netAddr)
	clnt := Client{make(chan string), netAddr, netAddr}
	onlineMap[netAddr] = clnt //将新连接的用户，添加到在线用户map中
	go WriteMsgToClient(clnt, conn)
	//发送消息到全局channel中
	messages <- "[" + netAddr + "]" + clnt.Name + "login"
	for {

	}

}
func Manager() {
	onlineMap = make(map[string]Client)

	for {
		msg := <-messages                //监听全局channel中是否有数据，讲数据存储到msg，无数据阻塞
		for _, clnt := range onlineMap { //循环发送消息给所有在线用户
			clnt.C <- msg
		}
	}
}

func main() {
	Listener, err := net.Listen("tcp", "127.0.0.1:8001")
	if err != nil {
		fmt.Println("net.Listen", err)
		return
	}
	defer Listener.Close()
	go Manager()

	for {
		Conn, err := Listener.Accept()
		if err != nil {
			fmt.Println("Listener.Accept", err)
			return
		}
		go handlerConnet(Conn)

	}

}
