package main

import (
	"fmt"
	"net"
)

type Client2 struct {
	C     chan string
	Name  string
	Adder string
}

var messagesclientmap2 map[string]Client2

var messages2 = make(chan string)

func WriteMsgToClient2(conn net.Conn, clt Client2) {
	for msg := range clt.C {
		conn.Write([]byte(msg + "\n"))
	}
}
func MakeMsg2(clt Client2, data string) (buf string) {
	buf = "[" + clt.Adder + "] :" + clt.Name + " :" + data
	return
}
func handlerConnet2(conn net.Conn) {
	defer conn.Close()
	Addr := conn.RemoteAddr().String()
	clt := Client2{make(chan string), Addr, Addr}
	go WriteMsgToClient2(conn, clt)
	messagesclientmap2[Addr] = clt

	messages2 <- MakeMsg2(clt, "login")
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
			msg := string(buf[:n-1])
			if msg == "who" && len(msg) == 3 {
				for _, value := range messagesclientmap2 {
					msg_list := "用户列表[" + value.Name + value.Adder + "]\n"
					conn.Write([]byte(msg_list))
				}

			} else {
				messages2 <- MakeMsg2(clt, msg)
			}

		}

	}()
	for {

	}

}
func Manager2() {
	messagesclientmap2 = make(map[string]Client2)
	for {
		msg := <-messages2
		for _, clent := range messagesclientmap2 {
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
	go Manager2()
	for {
		Conn, err := Listener.Accept()
		if err != nil {
			fmt.Printf("Listener.Accept", err)
			return
		}
		go handlerConnet2(Conn)
	}
}
