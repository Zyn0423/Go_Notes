package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	UDPAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8001") //TODO 获取UDP地址
	if err != nil {
		fmt.Println("net.ResolveUDPAddr错误", err)
		return
	}
	UDPConn, err := net.ListenUDP("udp", UDPAddr)
	if err != nil {
		fmt.Println("net.ListenUDP错误", err)
		return
	}
	fmt.Println("UDP服务端连接成功")
	defer UDPConn.Close() //TODO 关闭UDP连接
	buf := make([]byte, 1024*4)
	n, udpaddr, err := UDPConn.ReadFromUDP(buf) // TODO 接收数据该方法返回3个参数 1.字节  2. 客服端地址  3. 错误
	if err != nil {
		fmt.Println("UDPConn.ReadFromUDP", err)
		return

	}

	fmt.Printf("服务端收到客户端 %v\n 数据：%s\n", udpaddr, string(buf[:n]))
	time_now := time.Now().String()
	n, err = UDPConn.WriteToUDP([]byte(time_now), udpaddr) // TODO WriteToUDP acts like WriteTo but takes a UDPAddr.  int, error
	if err != nil {
		fmt.Println("UDPConn.WriteToUDP", err)
		return
	}

}
