package main

import (
	"fmt"
	"net"
)

type Client1 struct {
	C     chan string
	Name  string
	Adder string
}

var messagesclientmap map[string]Client1

var messages1 = make(chan string)

func WriteMsgToClient1(conn net.Conn, clt Client1) {
	for msg := range clt.C {
		conn.Write([]byte(msg + "\n"))
	}
}
func MakeMsg(clt Client1, data string) (buf string) {
	buf = "[" + clt.Adder + "] :" + clt.Name + " :" + data
	return
}
func handlerConnet1(conn net.Conn) {
	defer conn.Close()
	Addr := conn.RemoteAddr().String()
	clt := Client1{make(chan string), Addr, Addr}
	go WriteMsgToClient1(conn, clt)
	messagesclientmap[Addr] = clt

	messages1 <- MakeMsg(clt, "login")
	go func() {
		buf := make([]byte, 1024*4)
		for {
			n, err := conn.Read(buf)
			if err != nil {
				fmt.Println("conn.Write", err)
				return
			}
			if n == 0 {
				fmt.Printf("检测到用户端%s退出\n", clt.Name)
				return
			}
			//将读到用户消息保存到msg中
			msg := string(buf[:n])
			messages1 <- MakeMsg(clt, msg)

		}

	}()
	for {

	}

}
func Manager1() {
	messagesclientmap = make(map[string]Client1)
	for {
		msg := <-messages1
		for _, clent := range messagesclientmap {
			clent.C <- msg
		}
	}
}
func main() {
	Listener, err := net.Listen("tcp", "127.0.0.1:8001")
	if err != nil {
		fmt.Printf("net.Listen", err)
		return
	}

	defer Listener.Close()
	go Manager1()
	for {
		Conn, err := Listener.Accept()
		if err != nil {
			fmt.Printf("Listener.Accept", err)
			return
		}
		go handlerConnet1(Conn)
	}
}
